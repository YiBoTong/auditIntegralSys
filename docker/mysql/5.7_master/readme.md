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
```
mkdir -p -m 777 /home/mysql57/{master_3310,slave_3311,slave_3312}/{data,cnof,logs}
# 创建文件夹并设置文件夹权限
# 关联文件
docker run --name mysql57_master_3310 -v $PWD/data:/var/lib/mysql -v $PWD/conf:/etc/mysql/conf.d -v $PWD/logs:/usr/local/mysql/logs --restart always -d -p 3310:3306 -e MYSQL_ROOT_PASSWORD=root mysql57_master:1.0
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
# 创建注册日志账号
create user 'slave'@'%' identified by 'slavemima';
# 设置注册日志账号权限
grant replication slave on *.* to 'slave'@'%';
# 创建开发者账号
CREATE USER 'dev'@'%' IDENTIFIED BY 'ybt0005';
# 设置开发者对数据库的操作权限
GRANT ALL ON *.* TO 'dev'@'%';
# 查看当前数据的状态信息
show master status;
# 刷新数据库
flush tables;
# 刷新权限
flush privileges;
```