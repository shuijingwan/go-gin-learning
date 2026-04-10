go-gin-learning
# go-gin-learning
一个基于Go和Gin构建RESTful API的渐进式学习项目，包含数据库集成、身份认证及Docker部署。
## 项目介绍
基于Go和Gin框架的渐进式学习项目，帮助学习者掌握Go Web开发全流程，包含RESTful API实战和Go基础学习内容，适合初学者及入门开发者。
## 目录结构
```plain text
go-gin-learning/
├── .devcontainer/    # 容器开发环境配置
├── go-tour/          # Go官方教程（A Tour of Go）示例代码
├── .gitignore        # Git忽略文件
├── README.md         # 项目说明文档
├── docker-compose.yml# Docker部署配置文件
├── go.mod            # Go模块依赖配置
├── go.sum            # 依赖校验文件
└── main.go           # Gin项目入口文件
```
## 核心目录及文件
### 1. go-tour/
包含Go官方教程（A Tour of Go）完整示例代码，可直接运行，用于巩固Go基础。
### 2. 核心文件
- main.go：Gin项目入口文件，实现基础API功能。
- docker-compose.yml：Docker部署配置，支持容器化开发部署。
- go.mod & go.sum：Go模块依赖管理文件。
- .devcontainer/：集成容器开发环境配置，保障开发环境一致。
## 快速开始
### 前置要求
- Go 1.18+
- Git
- Docker（可选，用于容器化部署）
### 运行项目
1. 克隆仓库（注：结合系统“link dead”报错，需检查链接正确性）
```bash
git clone https://github.com/shuijingwan/go-gin-learning.git
cd go-gin-learning
```
2. 运行Gin主项目
```bash
go run main.go
```
3. 运行go-tour示例代码
```bash
cd go-tour/[目录名]
go run main.go
```
## Git提交备注规范（适配仓库实际提交）
规范写法：`docs: 完善 README 文档`（与仓库最新提交备注一致，规避此前doc:前缀错误）
补充：结合“link dead”报错，可补充为：`docs: 完善 README 文档（修正链接提示，适配GitHub渲染）`
## 贡献者
- shuijingwan (wangqiang)
## 开发语言
Go 100.0%