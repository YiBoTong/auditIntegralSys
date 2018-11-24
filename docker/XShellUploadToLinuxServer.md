# XShell upload or download to linux server
## Upload
1. 检查是否安装了`lrzsz`
```
rz
```
2.安装lrzsz
```
yum -y install lrzsz
```
3.检查是否安装成功
```
rpm -qa lrzsz
```
4.使用命令上传文件
```
rz -y
```
5.文件上传成功后查看文件
```
ls
```

## download
```
sz {文件}
```