package utils

import (
	"fmt"
	"github.com/LiteyukiStudio/spage/pkg/constants"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RunWithMode(h *server.Hertz, mode string) error {
	if mode == constants.ModeDev {
		err := h.Run()
		if err != nil {
			return err
		}
	} else if mode == constants.ModeProd {
		h.Spin()
	} else {
		return fmt.Errorf("unsupported mode: %s", mode)
	}
	return nil
}
