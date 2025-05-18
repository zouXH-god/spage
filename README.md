# liteyuki-pages
Self-hosted static pages hosting and managing

## 开发规范
- JSON类响应统一使用resps库的方法进行

## 手动构建
1. 先构建前端，前端产物文件夹dist内的东西全部放到./web/dist/目录下
2. 按照正常的go项目构建流程即可`go build github.com/LiteyukiStudio/spage/cmd/server`

## 常见问题
- 跨域问题：开发模式正常情况下不会遇到跨域问题，开发模式下允许的域为`http://localhost:5173`(Vite开发服务器默认地址)，如果需要其他域名请配置`frontend-url`配置项