# MySQL主从同步

环境  
master：10.83.3.102  
slave: 10.83.3.103  
MySQL：5.6
---
## 在master上操作
### 更改配置文件
```cmd
vim /etc/mysql/mysql.conf.d/mysqld.cnf
#在配置文件里添加：
server-id = 1 #在master-slave架构中，每台机器节点都需要有唯一的server-id
log_bin = /var/log/mysql/mysql-bin.log #开启binlog
systemctl restart mysql
```
### 创建主从同步的mysql user
```cmd
root@newbie-unknown85882:~# mysql
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 2
Server version: 5.7.21-log MySQL Community Server (GPL)

Copyright (c) 2000, 2018, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> create user 'slave'@'10.83.3.103' identified by 'slavemima';
Query OK, 0 rows affected (0.01 sec)
#创建slave用户，并指定该用户只能在10.83.3.103上登录。

mysql> grant replication slave on *.* to 'slave'@'10.83.3.103';
Query OK, 0 rows affected (0.01 sec)
#为slave赋予replication slave权限。
```
### 为MySQL加读锁
为了使主库与从库的数据保持一致，先为MySQL加入读锁，使其变为只读。
```cmd
mysql> flush tables with read lock;
Query OK, 0 rows affected (0.00 sec)
```
记录master replication log的位置，稍后会用到
```cmd
mysql> show master status;
+------------------+----------+--------------+------------------+------------------------------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set                        |
+------------------+----------+--------------+------------------+------------------------------------------+
| mysql-bin.000001 |      617 |              |                  | 43425f23-2760-11e8-a564-52540035ed32:1-2 |
+------------------+----------+--------------+------------------+------------------------------------------+
1 row in set (0.00 sec)
```
将master DB中现有的数据导出
```cmd
root@newbie-unknown85882:~# mysqldump -u root -p --all-databases --master-data > dbdump.sql
Enter password: 
Warning: A partial dump from a server that has GTIDs will by default include the GTIDs of all transactions, even those that changed suppressed parts of the database. If you don't want to restore GTIDs, pass --set-gtid-purged=OFF. To make a complete dump, pass --all-databases --triggers --routines --events. 
root@newbie-unknown85882:~# ls
adduser.sh   backend2.py         cloudinit.sh  __pycache__  公共  视频  文档  音乐
backend1.py  cloudinit-Linux.sh  dbdump.sql    sarfile      模板  图片  下载  桌面
```
解开master DB的读锁
```cmd
mysql> unlock tables;
Query OK, 0 rows affected (0.00 sec)
```
将dbdump.sql文件复制到slave

---
## 在slave上操作
更改配置文件
```cmd
root@newbie-unknown85883:~# vim /etc/mysql/mysql.conf.d/mysqld.cnf
#在配置文件里添加：
bind-address = 10.83.3.103 #slave ip
server-id = 2 #在master-slave架构中，每台机器节点都需要有唯一的server-id
log_bin = /var/log/mysql/mysql-bin.log #开启binlog
root@newbie-unknown85883:~# systemctl restart mysql
```
导入master DB的dbdump.sql文件，使master-slave数据一致
```cmd
root@newbie-unknown85883:~# systemctl restart mysql
root@newbie-unknown85883:~# mysql -u root -p < /tmp/dbdump.sql 
Enter password: 
```
使slave与master建立连接，从而实现MySQL主从同步
```cmd
root@newbie-unknown85883:~# mysql
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 3
Server version: 5.7.21-log MySQL Community Server (GPL)

Copyright (c) 2000, 2018, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> stop slave;
Query OK, 0 rows affected, 1 warning (0.01 sec)
mysql> change master to
    -> master_host='10.83.3.102',
    -> master_port=3306,
    -> master_user='slave',
    -> master_password='slavemima',
    -> master_log_file='mysql-bin.000001',
    -> master_log_pos=617;
Query OK, 0 rows affected, 2 warnings (0.04 sec)

mysql> start slave;
Query OK, 0 rows affected (0.01 sec)

mysql> 
#master_log_file='mysql-bin.000001'和master_log_pos=617的值是从master上 show master status得到的。
```
设置从库为只读模式
```mysql
mysql> show global variables like "%read_only%";
+-----------------------+-------+
| Variable_name         | Value |
+-----------------------+-------+
| innodb_read_only      | OFF   |
| read_only             | OFF   |
| super_read_only       | OFF   |
| transaction_read_only | OFF   |
| tx_read_only          | OFF   |
+-----------------------+-------+
5 rows in set (0.03 sec)

mysql> set global read_only=1;
Query OK, 0 rows affected (0.00 sec)

mysql> show global variables like "%read_only%";
+-----------------------+-------+
| Variable_name         | Value |
+-----------------------+-------+
| innodb_read_only      | OFF   |
| read_only             | ON    |
| super_read_only       | OFF   |
| transaction_read_only | OFF   |
| tx_read_only          | OFF   |
+-----------------------+-------+
5 rows in set (0.00 sec)
```
分别在master和slave上放通3306端口
```cmd
#先在master上操作
root@newbie-unknown85882:~# vim /etc/network/iptables.up.rules 
# Generated by iptables-save v1.6.0 on Thu Feb  8 09:48:16 2018
*filter
:INPUT ACCEPT [0:0]
:FORWARD ACCEPT [0:0]
:OUTPUT ACCEPT [31:3088]
-A INPUT -s 10.83.3.104/32 -p tcp -m tcp --dport 3000 -j ACCEPT
-A INPUT -s 10.83.3.103/32 -p tcp -m tcp --dport 3000 -j ACCEPT
-A INPUT -s 10.83.3.11/32 -p tcp -j ACCEPT
-A INPUT -s 192.168.0.0/16 -p tcp -m multiport --dports 8888,8889,8890,9500,9999,35000 -j ACCEPT
-A INPUT -s 172.16.0.0/12 -p tcp -m multiport --dports 8888,8889,8890,9500,9999,35000 -j ACCEPT
-A INPUT -s 10.0.0.0/8 -p tcp -m multiport --dports 8888,8889,8890,9500,9999,35000,3306 -j ACCEPT #在这里添加3306
-A INPUT -s 127.0.0.0/8 -p tcp -m multiport --dports 8888,8889,8890,9500,9999,35000 -j ACCEPT
-A INPUT -m state --state RELATED,ESTABLISHED -j ACCEPT
-A INPUT -s 127.0.0.1/32 -j ACCEPT
-A INPUT -p icmp -j ACCEPT
-A INPUT -p tcp -m tcp --dport 32200 -j ACCEPT
-A INPUT -j DROP
COMMIT
# Completed on Thu Feb  8 09:48:16 2018
root@newbie-unknown85882:~# iptables-restore</etc/network/iptables.up.rules
#然后在slave上也执行同样的操作
```
```cmd
root@newbie-unknown85883:~# telnet 10.83.3.102 3306
Trying 10.83.3.102...
Connected to 10.83.3.102.
Escape character is '^]'.
N
5.7.21-log
          X?CuSD\�<m   yDi9 mysql_native_password
^CConnection closed by foreign host.
#在slave上测试，能连通master的3306端口。
```
确保master和slave都开启GTID事务
```mysql
root@newbie-unknown85883:~# vim /etc/mysql/mysql.conf.d/mysqld.cnf
#在配置文件中添加：
gtid_mode = on
enforce_gtid_consistency = 1


mysql> show variables like '%gtid_mode%';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| gtid_mode     | ON    |
+---------------+-------+
1 row in set (0.00 sec)
```
---
## 测试：
现在MySQL主从同步已经配置好了，但是测试的过程中可能会出现各种问题导致不能同步数据。这时候需要查看配置文件是否正确，查看错误日志获取关键信息以便解决问题。
分别在master和slave上执行`show master status;`和`show slave status\G;`
```mysql
mysql> show master status;
+------------------+----------+--------------+------------------+------------------------------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set                        |
+------------------+----------+--------------+------------------+------------------------------------------+
| mysql-bin.000007 |      194 |              |                  | 43425f23-2760-11e8-a564-52540035ed32:1-4 |
+------------------+----------+--------------+------------------+------------------------------------------+
```
```mysql
mysql> show slave status\G;
*************************** 1. row ***************************
               Slave_IO_State: Waiting for master to send event
                  Master_Host: 10.83.3.102
                  Master_User: slave
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File: mysql-bin.000007
          Read_Master_Log_Pos: 672
               Relay_Log_File: newbie-unknown85883-relay-bin.000002
                Relay_Log_Pos: 633
        Relay_Master_Log_File: mysql-bin.000007
             Slave_IO_Running: No
            Slave_SQL_Running: Yes
```
当Slave_IO_Running或者Slave_SQL_Running显示非“Yes”状态的时，说明配置有错误，需要查看配置或者错误日志。
```cmd
cat /var/log/mysql/error.log
2018-03-16T11:20:00.929581Z 7 [ERROR] Slave I/O for channel '': The replication receiver thread cannot start because the master has GTID_MODE = ON and this server has GTID_MODE = OFF. Error_code: 1593
```
我是在slave上查看错误日志的，由this server has GTID_MODE = OFF可得知slave的GTID没有开启，这时候需要更改配置文件开启GTID。

什么是GTID？  
GTID(Global Transaction ID)是对于一个已提交事务的编号，并且是一个全局唯一的编号。它的官方定义如下：GTID = source_id ：transaction_id每一个 GTID 代表一个数据库事务。在上面的定义中，source_id 表示执行事务的主库uuid（server_uuid），transaction_id 是一个从 1 开始的自增计数，表示在这个主库上执行的第 n 个事务。MySQL 只要保证每台数据库的 server_uuid 全局唯一，以及每台数据库生成的 transaction_id 自身唯一，就能保证 GTID 的全局唯一性。

纠正了错误后继续测试：  
在master上操作
```mysql
mysql> show master status;
+------------------+----------+--------------+------------------+------------------------------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set                        |
+------------------+----------+--------------+------------------+------------------------------------------+
| mysql-bin.000007 |      359 |              |                  | 43425f23-2760-11e8-a564-52540035ed32:1-5 |
+------------------+----------+--------------+------------------+------------------------------------------+
1 row in set (0.00 sec)
```
由master上查出的信息得知，master replication log的位置和Position和我们之前查的不一样了，那么要以现在的为准。  

在slave上操作
```mysql
mysql> stop slave;
Query OK, 0 rows affected (0.01 sec)

mysql> change master to
    -> master_host='10.83.3.102',
    -> master_user='slave',
    -> master_password='slavemima',
    -> master_log_file='mysql-bin.000007',
    -> master_log_pos=359;
Query OK, 0 rows affected, 2 warnings (0.04 sec)

mysql> start slave;
Query OK, 0 rows affected (0.01 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| newbie             |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)
```
现在所有的配置都已完成且正确，接下来就是最关键的测试，同步测试。  
在master上操作
```mysql
mysql> create database test;
Query OK, 1 row affected (0.01 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| newbie             |
| performance_schema |
| sys                |
| test               |
+--------------------+
6 rows in set (0.01 sec)
```
在slave上操作
```mysql
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| newbie             |
| performance_schema |
| sys                |
| test               |
+--------------------+
6 rows in set (0.01 sec)
```
由此可知，在master上创建一个test库，slave会同步创建相应的test库！MySQL主从同步配置成功！