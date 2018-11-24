# yapi
构建镜像
```docker
docker build -t yapi:1.0 .
```
导出镜像
```docker
docker save -o yapi_1.0.tar yapi:1.0
```
运行
```
cd vendors
npm install --production --registry https://registry.npm.taobao.org
npm run install-server //安装程序会初始化数据库索引和管理员账号，管理员账号名可在 config.json 配置
node server/app.js //启动服务器后，请访问 127.0.0.1:{config.json配置的端口}，初次运行会有个编译的过程，请耐心等候
```