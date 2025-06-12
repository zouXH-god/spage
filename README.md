

<center><img src="https://socialify.git.ci/LiteyukiStudio/spage/image?description=1&font=Bitter&forks=1&issues=1&logo=https%3A%2F%2Fcdn.liteyuki.org%2Flogos%2Fapage.png&name=1&pattern=Overlapping+Hexagons&pulls=1&stargazers=1&theme=Auto" alt="spage" width="780" height="320" />
    <!-- <a href="./README.md">简体中文</a>
    |
    <a href="./README/en.md">English</a>
    |
    <a href="./README/zh.md">繁體中文</a>
    |
    <a href="./README/ja.md">日本語</a>
    |
    <a href="./README/ko.md">한국어</a> -->
</center>

# Spage - 自托管静态页面托管服务

> 项目处于开发阶段，暂未发布第一个版本，感兴趣的可以先订阅一下

## 简介

一个基于Go语言开发的, 开源自托管静态页面托管服务, 使用Caddy作为web服务器

类似于`Vercel Pages`, `GitHub Pages`, `Cloudflare Pages`等PaaS服务, 但它是一个开源的, 可自托管的**平替**

## 快速安装

### 容器化部署

```bash
# Docker Hub源
liteyukistudio/spage:latest

# GitHub Container Registry源
ghcr.io/liteyukistudio/spage:latest

# Liteyuki Container Registry源
reg.liteyuki.org/spage/spage:latest

# 推荐优先使用公共源，即Docker Hub源和GHCR源，若访问不了再使用Liteyuki Container Registry源
# 以减轻Liteyuki Container Registry的压力
```

你可使用docker，podman等工具部署，也可以将其部署到Kubernetes集群中

默认容器内服务端口是`8888`，你可以按需暴露，并挂载`./config.yaml`到容器内部

若你需要使用最新提交，可以将`latest`替换为`nightly`标签

### 二进制部署

如果不想容器化，也可以直接跑二进制，支持，可以在[Release](./releases)界面找到大部分平台和架构的二进制文件
支持Linux，macOS(Darwin)，Windows，FreeBSD等操作系统
AMD64兼容性：v1支持所有AMD64架构的CPU，v3支持2013年及以后的AMD64架构CPU
如果找不到你目标平台的二进制文件，可以尝试从源代码构建
    `go build ./cmd/server`

## 技术栈

自托管静态页面托管服务, 基于以下技术栈构建:
- 后端: `Golang`, `Hertz`框架, `Kitex`(RPC框架), `GORM`(ORM框架)
- 前端: `Next.js(React.js)`, `TypeScript`
- 数据库: `SQLite3`, `PostgreSQL`
- Web服务器: `Caddy`(仅提供接口)
- CLI自动化: 任意, 但尽可能和现有的前端工具链集成
- 容器化: `Docker`

理念:
- 尽可能的在保证完整功能的同时在开发和使用过程中不引入过多的依赖
- 考虑更具有优势和简洁完善的技术栈

## 基础功能

- **用户管理, 组织管理, 项目和站点管理**

- **支持OAuth2以便团队使用和管理**

- **可配置后台域名, 支持自定义域名**
如 `panel.liteyuki.org`

- **支持全局域名, 可配置自定义域名, 用于站点自动生成使用**
使用`{version.replace(".", "_")}-{prefix}.{pages-domain}`
前缀`prefix`随机生成, 可后期指定, 不能包含"`.`"
版本号`version`默认为`latest`时可不加前缀

- **TLS配置自定义**
内网部署时, 如果外网存在一个统一网关, 可以不配置内网TLS
只需要将泛域名`CNAME`反向代理到内网Pages服务Http端点即可

- **动态管理站点主机路由**
WEB服务器, 我们优先选用`Caddy`作为Web服务器, 因为它的`REATful API`可以用来让我们的Pages服务与`Caddy`交互, 动态地管理站点主机名路由
例如: 对CDN或者客户端(直连时)声明缓存时间以减少源服务器流量开销

- **ClearURLs支持**
支持使用 example.com/about 访问 example.com/about.html/htm
省略文件扩展名, 减少冗余

## CI/CD集成

- **CLI**
使用`TypeScript`编写, 可以使用`pnpm`或`npm`等包管理器进行安装依赖, 部署, 构建等操作
部署时需要指定**输出文件夹**(包含index.html的目录), **站点ID**以及**站点Token**, 还有版本号Tag(可选)
CLI将输出文件夹压缩后使用指定接口上传到服务器
(可以考虑使用其他语言编写, 但是在前端构建工作流容器中使用已有前端工具链更方便)

- **Git平台自动化集成**
可以无缝衔接**GitHub**, **Gitea**等平台的工作流
例如: 一个项目有`Release`/`Nightly`两个稳定站点, 在CLI中推送时就需要指定站点ID, 优先返回第一个配置的自定义域名, 在PR预览模式下, 无需指定站点URL, 自动创建新站点并返回随机前缀的URL

## 开发

在不修改默认配置的情况下，开发模式有预设配置，可以直接上手开发

### 主控后端和agent部分

- 安装go工具链
- 安装依赖`go mod tidy`
- 启动主控后端`go run ./cmd/server`

### 主控前端

- 安装pnpm和node(或其他运行时，例如bun，deno)
- 在项目根目录下切换到前端源码目录：`cd web-src`
- 使用`pnpm install`安装依赖
- 使用`pnpm dev` `bun dev`等方式启动前端开发服务器