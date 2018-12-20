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
# 停止slave服务
stop slave;
# 主从配置
change master to master_host='192.168.10.107', master_port=3310, master_user='slave', master_password='slavemima', master_log_file='mysql-bin.000005', master_log_pos=30511;
# 启动服务
start slave;
# 查看slave状态信息
show slave status;
# 创建开发者账号
CREATE USER 'dev'@'%' IDENTIFIED BY 'ybt0005';
# 设置开发者账号对所有数据库只读权限
GRANT SELECT ON *.* TO 'dev'@'%';
# 刷新权限
flush privileges;
# 查看当前数据库的只读配置
show global variables like "%read_only%";
# 设置当前数据为只读
set global read_only=1;
```