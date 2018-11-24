# docker 镜像制作
镜像名称只能是小写字母及下划线
```docker
docker build -t demo:latest .
```
构建master镜像
```docker
docker build -t mysql57_master:1.0 .
```
运行镜像
```docker
docker run --name mysql57_master_3310 --restart always -d -p 3310:3306 -e MYSQL_ROOT_PASSWORD=root mysql57_master:1.0
```
导出镜像
```docker
docker save -o mysql57_master_1.0.tar mysql57_master:1.0 
```
导入镜像
```docker
docker load -i mysql57_master_1.0.tar
```
MySQL Master
```mysql
create user 'slave'@'%' identified by 'slavemima';
grant replication slave on *.* to 'slave'@'%';
show master status;
flush tables;
flush privileges;
```