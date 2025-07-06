package caddy

import (
	"bufio"
	"context"
	"github.com/LiteyukiStudio/spage/agent/env"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"resty.dev/v3"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	caddyProcess *os.Process
	caddyMutex   sync.Mutex
	daemonActive bool
)

// runCaddy 启动Caddy服务器
func runCaddy() (*os.Process, error) {
	cmd := exec.Command(env.CaddyBin, env.CaddyCommandArgs)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	caddyLogs(stdout, stderr)
	logrus.Infof("Caddy Server started, PID: %d", cmd.Process.Pid)
	return cmd.Process, nil
}

// healthCheck 用于检查caddy服务的健康状态
func healthCheck() (bool, error) {
	client := resty.New()
	defer client.Close()
	res, err := client.R().Get(env.CaddyApiEndpoint + "/config/")
	if err != nil {
		logrus.Errorf("Caddy health check failed: %v", err)
		return false, err
	}
	if res.IsError() {
		logrus.Errorf("Caddy health check returned error: %s", res.String())
		return false, nil
	}
	logrus.Debug("Caddy health check passed")
	return true, nil
}

// stopDaemon 停止守护进程（私有方法）
func stopDaemon() {
	caddyMutex.Lock()
	defer caddyMutex.Unlock()

	if !daemonActive {
		logrus.Info("Caddy daemon not running")
		return
	}

	daemonActive = false
	logrus.Info("Stopping Caddy daemon")

	if caddyProcess != nil {
		if err := caddyProcess.Signal(syscall.SIGTERM); err != nil {
			logrus.Errorf("Failed to terminate Caddy: %v", err)
			_ = caddyProcess.Kill()
		}
		caddyProcess = nil
	}
}

// StartDaemon 启动守护进程，定期检查Caddy状态并在需要时重启
func StartDaemon(ctx context.Context) {
	checkInterval := time.Duration(env.CaddyCheckInterval) * time.Second
	caddyMutex.Lock()
	if daemonActive {
		caddyMutex.Unlock()
		logrus.Info("Caddy daemon already running")
		return
	}
	daemonActive = true
	caddyMutex.Unlock()

	go func() {
		// 确保在函数退出时停止守护进程
		defer func() {
			stopDaemon()
			logrus.Info("Caddy daemon goroutine exited")
		}()

		logrus.Info("Starting Caddy daemon")

		// 首次启动
		process, err := runCaddy()
		if err != nil {
			logrus.Errorf("Failed to start Caddy: %v", err)
		} else {
			caddyMutex.Lock()
			caddyProcess = process
			caddyMutex.Unlock()
		}

		ticker := time.NewTicker(checkInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				// 上下文被取消，停止守护进程
				logrus.Info("Context canceled, stopping Caddy daemon")
				return
			case <-ticker.C:
				caddyMutex.Lock()
				isActive := daemonActive
				caddyMutex.Unlock()

				if !isActive {
					logrus.Info("Caddy daemon marked for stopping")
					return
				}

				// 检查Caddy健康状态
				healthy, _ := healthCheck()
				if !healthy {
					logrus.Warn("Caddy is not healthy, attempting to restart")

					// 尝试关闭现有进程
					caddyMutex.Lock()
					if caddyProcess != nil {
						_ = caddyProcess.Signal(syscall.SIGTERM)
						time.Sleep(2 * time.Second) // 给进程一些时间来优雅关闭
						_ = caddyProcess.Kill()     // 如果还在运行，强制终止
					}

					// 重启Caddy
					process, err := runCaddy()
					if err != nil {
						logrus.Errorf("Failed to restart Caddy: %v", err)
					} else {
						caddyProcess = process
						logrus.Info("Caddy restarted successfully")
					}
					caddyMutex.Unlock()
				}
			}
		}
	}()
}

// caddyLogs 用于获取Caddy日志
func caddyLogs(stdout, stderr io.Reader) {
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			text := scanner.Text()
			logrus.Info("[Caddy] " + text)
		}
		if err := scanner.Err(); err != nil {
			logrus.Errorf("Error reading Caddy stdout logs: %v", err)
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			text := scanner.Text()
			if strings.Contains(text, `"level":"debug"`) {
				logrus.Debug("[Caddy] " + text)
			} else if strings.Contains(text, `"level":"info"`) {
				logrus.Info("[Caddy] " + text)
			} else if strings.Contains(text, `"level":"warn"`) || strings.Contains(text, `"level":"warning"`) {
				logrus.Warn("[Caddy] " + text)
			} else if strings.Contains(text, `"level":"error"`) {
				logrus.Error("[Caddy] " + text)
			} else if strings.Contains(text, `"level":"fatal"`) {
				logrus.Fatal("[Caddy] " + text)
			} else {
				logrus.Info("[Caddy] " + text)
			}
		}
		if err := scanner.Err(); err != nil {
			logrus.Errorf("Error reading Caddy stderr logs: %v", err)
		}
	}()
}
