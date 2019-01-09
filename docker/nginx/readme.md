# 使用 nginx 镜像
下载nginx镜像
```docker
docker pull nginx
```
构建master镜像
```docker
docker build -t ai_nginx:1.0 .
```
运行镜像（内置配置）
```docker
docker run --name ai_nginx --restart always -d -p 80:80 ai_nginx:1.0
docker run -p 80:80 --name nginx -v $PWD/www:/www -v $PWD/conf:/etc/nginx/conf.d -d nginx
```
运行镜像（外置配置）
```docker
cd /home # 进入home文件夹
docker run -p 80:80 --name ai_nginx -v $PWD/www:/www -v $PWD/conf/nginx.conf:/etc/nginx/nginx.conf -v $PWD/conf/audit_integral.conf:/etc/nginx/conf.d/audit_integral.conf -d nginx
# -p 80:80：将容器的80端口映射到主机的80端口
# --name ai_nginx：将容器命名为ai_nginx
# -v $PWD/www:/www：将主机中当前目录下的www挂载到容器的/www
# -v $PWD/conf/nginx.conf:/etc/nginx/nginx.conf：将主机中当前目录下的nginx.conf挂载到容器的/etc/nginx/nginx.conf
# $PWD/home/conf/audit_integral.conf:/etc/nginx/conf.d/audit_integral.conf：将主机中当前目录下的logs挂载到容器的/wwwlogs
# -v $PWD/logs:/wwwlogs：将主机中当前目录下的logs挂载到容器的/wwwlogs
```
hmoe文件夹结构（外置运行镜像时使用）
```
images  # docker镜像文件夹
conf    # 配置文件夹
www     # 前端项目文件夹
```
导出镜像
```docker
docker save -o ai_nginx_1.0.tar ai_nginx:1.0 
```
导入镜像
```docker
docker load -i ai_nginx_1.0.tar
```
静态服务
```docker
docker run -p 8094:80 --name ai_static_server --restart always -v $PWD/static:/usr/share/nginx/html -d nginx
```