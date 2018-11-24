# 构建服务
1.打包(一定要改日志输出文件夹配置)
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o system_setup main.go
```
2.构建docker镜像（一定要改config.js中的数据库地址）
```
docker build -t system_setup:1.0 .
```
3.导出镜像(在`_docker`文件夹中)
```
docker save -o system_setup_1.0.tar system_setup:1.0
```
4.导入镜像
```docker
docker load -i system_setup_1.0.tar
```