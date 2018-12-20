# docker run redis
构建master镜像
```docker
docker build -t redis50:1.0 .
```
运行镜像
```docker
docker run -p 6379:6379 -d redis:4.0 redis-server --appendonly yes
# -p 6379:6379 : 将容器的6379端口映射到主机的6379端口
# -v $PWD/data:/data : 将主机中当前目录下的data挂载到容器的/data
# redis-server --appendonly yes : 在容器执行redis-server启动命令，并打开redis持久化配置
# docker run -p 6379:6379 --name redis --restart always -v $PWD/data:/data -d redis:5.0 redis-server --appendonly yes
docker run -p 6379:6379 --name redis --restart always -v $PWD/data:/data -v $PWD/conf/redis.conf:/etc/redis/redis.conf -d redis:5.0 redis-server /etc/redis/redis.conf --appendonly yes
```
导出镜像
```docker
docker save -o redis50_1.0.tar redis50:1.0 
```
导入镜像
```docker
docker load -i redis50_1.0.tar
```