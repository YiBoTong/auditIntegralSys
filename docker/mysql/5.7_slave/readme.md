# docker 镜像制作
镜像名称只能是小写字母及下划线
```docker
docker build -t demo:latest .
```
构建master镜像
```docker
docker build -t mysql57_slave:1.0 .
```
运行镜像
```docker
docker run --name mysql57_slave_3311 --restart always -d -p 3311:3306 -e MYSQL_ROOT_PASSWORD=root mysql57_slave:1.0
```
导出镜像
```docker
docker save -o mysql57_slave_1.0.tar mysql57_slave:1.0 
```
导入镜像
```docker
docker load -i mysql57_slave_1.0.tar
```
MySQL Slave
```mysql
stop slave;
change master to master_host='192.168.10.107', master_port=3310, master_user='slave', master_password='slavemima', master_log_file='mysql-bin.000005', master_log_pos=30511;
start slave;
show slave status;
```