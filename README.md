# Spage - 自托管静态页面托管服务

<center><img src="https://socialify.git.ci/Nanaloveyuki/spage/image?description=1&font=Bitter&forks=1&issues=1&logo=https%3A%2F%2Fcdn.liteyuki.icu%2Flogos%2Fapage.png&name=1&pattern=Overlapping+Hexagons&pulls=1&stargazers=1&theme=Auto" alt="spage" width="780" height="320" />
    <a href="./README.md">简体中文</a>
    |
    <a href="./README/en.md">English</a>
    |
    <a href="./README/zh.md">繁體中文</a>
    |
    <a href="./README/ja.md">日本語</a>
    |
    <a href="./README/ko.md">한국어</a>
</center>

---

## 快速开始 

容器化一键部署

```bash
# Docker Hub源
liteyukistudio/spage:latest

# GitHub Container Registry源
ghcr.io/liteyukistudio/spage:latest

# Liteyuki Container Registry源
reg.liteyuki.org/spage/spage:latest
```

你可使用docker，podman等工具部署，也可以将其部署到Kubernetes集群中

如果不想容器化，也可以直接跑二进制

## 简介

一个基于Go语言开发的, 开源自托管静态页面托管服务

## 初衷

### 为什么要开发这个平台

常规的前端应用管理的流程是
1. 在构建平台构建好
2. 上传输出目录到服务器的某处
3. 在Nginx, Apache, Caddy等这种Web服务器中配置静态资源目录
4. 用户访问

通常需要配置一大堆东西, 部署过程还不太好**全自动化**
我们希望能够让前端应用的部署过程更加**简单**, 更**快速**

为了让现有和未来的前后端分离架构的应用前端部分能够快速上线, 与现有的CI/CD集成, 通过RESTful API部署和管理项目, 且能够托管到自己的服务器, 于是我们打算开发这个平台

### 为什么不使用其他同类项目

我们不希望这个平台是一个SaaS服务, 而是一个开源的, 自托管的平台, 我们希望这个平台能够被更多的人使用, 并且能够被更多的人贡献代码

它是`Vercel Pages`, `GitHub Pages`, `Cloudflare Pages`这种PaaS服务的开源, 可自托管的**平替**

也许有人会问:
> 为什么我不用公共的SaaS服务, 要使用这个呢？

如果是个人小静态站, 使用公共免费SaaS完全没有问题！

本项目面向的用户更多倾向于使用自托管, 例如企业使用一个服务专门来托管团队/企业中静态页面项目

我们的设计是在前端构建完成后, 只需要**一条命令**就可以将构建产物部署到平台

在此之前我们参考过非常多的同类项目, 例如: `getmeli`/`meli`, 然而其已经在几年前停止更新了, `coolify`和`dokploy`功能过于臃肿和难以部署, 给每个前端项目起一个容器又过于浪费资源

## 技术栈

自托管静态页面托管服务, 基于以下技术栈构建:
- 后端: `Golang`, `Hertz`框架, `Kitex`(RPC框架), `GORM`(ORM框架)
- 前端: `Vue3`, `Element`, `TypeScript`
- 数据库: `SQLite3`, `PostgreSQL`
- Web服务器: `Caddy`, `Nginx`(仅提供接口)
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

- **站点版本管理**
每个站点最新版本默认标签为latest, 当部署站点时, 指定相同的标签会被覆盖, 不指定默认latest
正常情况下, 这样的站点域名为`{version.replace(".", "_")}-{prefix}.{pages-domain}`

---

更多细节请查看[文档](https://docs.apage.dev/)

---

## 快速开始

从[Github Release](https://github.com/LiteyukiStudio/spage/releases)下载对应版本的二进制文件, 运行即可

**Todo: 编写快速开始部分的文档**

## 手动构建

**Todo: 完善构建文档**

1. **首先Clone本仓库**
```bash
git clone https://github.com/LiteyukiStudio/spage ./spage && cd spage
```


2. **随后构建您的前端**


3. **按照正常的Go语言项目构建server二进制文件即可**
(从github获取源代码构建)
```bash
go build github.com/LiteyukiStudio/spage/cmd/server
```

(从本地获取源代码构建)
```bash
go build ./cmd/server
```

## 常见问题
- 跨域问题：开发模式正常情况下不会遇到跨域问题, 开发模式下允许的域为`http://localhost:5173`(Vite开发服务器默认地址), 如果需要其他域名请配置`frontend-url`配置项