# docker run redis
```docker
docker run -p 6379:6379 -v $PWD/data:/data  -d redis:3.2 redis-server --appendonly yes
# -p 6379:6379 : 将容器的6379端口映射到主机的6379端口
# -v $PWD/data:/data : 将主机中当前目录下的data挂载到容器的/data
# redis-server --appendonly yes : 在容器执行redis-server启动命令，并打开redis持久化配置
```