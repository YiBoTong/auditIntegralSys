# Golang包管理工具  
安装
```cmd
go get -u -v github.com/kardianos/govendor
```
初始化vendor目录  
```cmd
govendor init
```
将GOPATH中本工程使用到的依赖包自动移动到vendor目录中  
说明：如果本地GOPATH没有依赖包，先go get相应的依赖包
```cmd
govendor add +e
```