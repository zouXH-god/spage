# liteyuki-pages
Self-hosted static pages hosting and managing

# 手动构建文档
1. 先构建前端，前端产物文件夹dist内的东西全部放到./web/dist/目录下
2. 按照正常的go项目构建流程即可`go build github.com/LiteyukiStudio/spage/cmd/server`