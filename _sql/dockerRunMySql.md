# docker run mySql  
1.拉取mySql 5.6  
  ```docker
  docker pull mysql:5.6
  ```
  
2.查看镜像  
  ```docker
  docker images
  ```

3.运行一个mysql容器
  ```docker
  docker run --name mySql3320 --restart always -p 3320:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.6
  
  
  ```
  ```
  run               运行一个容器
  --name            后面是这个镜像的名称
  --restart always  docker服务挂了后容器也挂了，加上这个表示跟随容器启动
  -p 3320:3306      表示在这个容器中使用3306端口(第二个)映射到本机的端口号也为3320(第一个)
  -d                表示使用守护进程运行，即服务挂在后台
  ```
4.检查镜像运行状态
  ```docker
  docker ps
  ```
5.删除容器
  ```docker
  docker rm [容器名称]
  ```
6.删除镜像
  ```docker
  docker rmi [镜像id]
  ```
7.镜像迁移  
`save`和`load`参数迁移镜像  
`export`和`import`参数迁移容器
  - 导出容器（export）
  ```docker
  docker export -o savePath/文件名称.tar <容器名称>
  ```
  - 导入容器（import）
  ```docker
  
  ```
  
 ## 分析 Docker 空间分布
 Docker 的内置 CLI 指令docker system df，可用于查询镜像（Images）、容器（Containers）和本地卷（Local Volumes）等空间使用大户的空间占用情况。 
 ```docker
docker system df
```
## 查看空间占用细节
```docker
docker system df -v
```