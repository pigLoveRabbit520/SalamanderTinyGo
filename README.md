# SalamanderTinyGo
基于go的微小框架

# 安装依赖
利用[glide](https://glide.sh/)管理依赖，借助docker容器安装依赖，修改docker-compose.yml中command
```yaml
command: ["glide", "install"]
```
然后运行容器``docker-compose up`

# 编译
修改docker-compose.yml中command
```yaml
command: ["go", "build", "."]
```

