/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.5.53-log : Database - audit_integral
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`audit_integral` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `audit_integral`;

/*Table structure for table `clause` */

DROP TABLE IF EXISTS `clause`;

CREATE TABLE `clause` (
  `id` int(11) NOT NULL COMMENT '管理办法及条款ID',
  `department_id` int(11) DEFAULT NULL COMMENT '所属部门ID',
  `title` char(200) DEFAULT NULL COMMENT '管理办法及条款标题',
  `author_id` int(11) DEFAULT NULL COMMENT '发布人ID',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `state` enum('draft','publish') DEFAULT 'draft' COMMENT '状态（draft草稿publish发布）',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `clause` */

/*Table structure for table `clause_content` */

DROP TABLE IF EXISTS `clause_content`;

CREATE TABLE `clause_content` (
  `id` int(11) NOT NULL COMMENT '管理办法内容ID',
  `is_title` tinyint(1) DEFAULT '0' COMMENT '是否为标题（0内容1标题）',
  `title_level` char(20) DEFAULT NULL COMMENT '标题级别',
  `content` varchar(1000) DEFAULT NULL COMMENT '内容',
  `order` int(4) DEFAULT NULL COMMENT '顺序',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `clause_content` */

/*Table structure for table `clause_file` */

DROP TABLE IF EXISTS `clause_file`;

CREATE TABLE `clause_file` (
  `clause_id` int(11) DEFAULT NULL COMMENT '管理办法ID',
  `file_id` int(11) DEFAULT NULL COMMENT '附件ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `clause_file` */

/*Table structure for table `confirmation` */

DROP TABLE IF EXISTS `confirmation`;

CREATE TABLE `confirmation` (
  `id` int(11) NOT NULL COMMENT '确认书ID',
  `manuscript_id` int(11) NOT NULL COMMENT '工作底稿ID',
  `confirmation_receipt_id` int(11) DEFAULT NULL COMMENT '确认书回执ID',
  `file_id` int(11) DEFAULT NULL COMMENT '图片ID',
  `proposal` varchar(2000) DEFAULT NULL COMMENT '单位意见或者建议',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`,`manuscript_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `confirmation` */

/*Table structure for table `confirmation_receipt` */

DROP TABLE IF EXISTS `confirmation_receipt`;

CREATE TABLE `confirmation_receipt` (
  `id` int(11) NOT NULL COMMENT '确认书回执ID',
  `content` varchar(5000) DEFAULT NULL COMMENT '回复内容',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '回执时间',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `confirmation_receipt` */

/*Table structure for table `confirmation_receipt_content` */

DROP TABLE IF EXISTS `confirmation_receipt_content`;

CREATE TABLE `confirmation_receipt_content` (
  `id` int(11) NOT NULL COMMENT '整改项ID',
  `confirmation_receipt` int(11) DEFAULT NULL COMMENT '确认书回执ID',
  `behavior_id` int(11) DEFAULT NULL COMMENT '违规行为ID',
  `result` char(250) DEFAULT NULL COMMENT '整改结果',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '整改时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `confirmation_receipt_content` */

/*Table structure for table `department` */

DROP TABLE IF EXISTS `department`;

CREATE TABLE `department` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `name` char(50) DEFAULT NULL COMMENT '部门名称',
  `parent_id` int(11) DEFAULT NULL COMMENT '上级部门ID',
  `code` char(50) DEFAULT NULL COMMENT '部门编码',
  `level` char(20) DEFAULT NULL COMMENT '部门级别',
  `grade` int(3) DEFAULT NULL COMMENT '所在部门树的层级',
  `address` char(250) DEFAULT NULL COMMENT '地址',
  `phone` char(11) DEFAULT NULL COMMENT '手机号或者电话',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete` tinyint(1) DEFAULT '0' COMMENT '是否删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;

/*Data for the table `department` */

insert  into `department`(`id`,`name`,`parent_id`,`code`,`level`,`grade`,`address`,`phone`,`update_time`,`delete`) values (1,'部门1',-1,'d1','1',NULL,'0','','0000-00-00 00:00:00',0),(2,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(3,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',1),(4,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(5,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(6,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(7,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(8,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(9,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(10,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(11,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(12,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(13,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(14,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(15,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(16,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(17,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',1),(18,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(19,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(20,'部门1',-1,'d12','1',NULL,'贵州贵阳1','','2018-11-28 17:19:21',0),(21,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(22,'部门1',-1,'d1','1',NULL,'贵州贵阳','','0000-00-00 00:00:00',0),(23,'部门1',-1,'d1','1',NULL,'贵州贵阳','1008611','2018-11-26 15:53:55',0),(24,'部门1',-1,'d1','1',NULL,'贵州贵阳','','2018-11-23 16:53:39',0),(25,'部门1',-1,'d1','1',NULL,'贵州贵阳','','2018-11-23 16:54:15',0),(26,'部门1',-1,'d1','1',NULL,'贵州贵阳','','2018-11-26 10:28:01',0),(27,'部门26-1',26,'d1','1',NULL,'贵州贵阳','','2018-11-26 10:28:14',0),(28,'部门26-2',26,'d1','1',NULL,'贵州贵阳','','2018-11-26 10:28:18',0),(29,'部门26-3',26,'d1','1',NULL,'贵州贵阳','','2018-11-26 10:28:21',0),(30,'部门26-4',26,'d1','1',NULL,'贵州贵阳','','2018-11-26 10:28:25',0),(31,'',0,'','0',NULL,'','','2018-11-26 11:47:10',0),(32,'',0,'','0',NULL,'','','2018-11-26 11:47:18',0),(33,'',26,'','0',NULL,'','','2018-11-26 11:51:01',1),(34,'',26,'','0',NULL,'','','2018-11-27 11:50:28',0),(35,'部门1',-1,'d1','1',NULL,'贵州贵阳','','2018-11-28 09:49:52',0),(36,'',26,'','0',NULL,'','','2018-11-28 14:25:59',0),(37,'部门1',-1,'d1','1',NULL,'贵州贵阳','','2018-11-28 17:18:37',1),(38,'部门1',-1,'d1','1',NULL,'贵州贵阳','','2018-11-28 17:19:21',1),(39,'测试部门',0,'10023','2',NULL,'休息休息','1815310','2018-11-29 15:37:50',0);

/*Table structure for table `department_user` */

DROP TABLE IF EXISTS `department_user`;

CREATE TABLE `department_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `department_id` int(11) DEFAULT NULL COMMENT '部门id',
  `user_id` int(11) DEFAULT NULL COMMENT '人员id',
  `type` char(20) DEFAULT NULL COMMENT '角色类型',
  `delete` tinyint(1) DEFAULT '0' COMMENT '软删除',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

/*Data for the table `department_user` */

insert  into `department_user`(`id`,`department_id`,`user_id`,`type`,`delete`) values (1,18,1,'admin',0),(2,19,1,'admin',0),(3,20,1,'admins',0),(4,21,1,'admin',0),(5,22,1,'admin',0),(6,23,1,'admin',0),(7,20,2,'admins',1),(8,20,2,'admins',1),(9,20,2,'admins',1),(10,20,2,'admins',1),(11,20,2,'admins',1),(12,31,0,'',0),(13,32,0,'',0),(14,33,0,'',0),(15,34,-2,'',0),(16,20,2,'admins',1),(17,36,-2,'',0),(18,20,2,'admins',1),(19,20,2,'admins',0),(20,39,-2,'admin',0);

/*Table structure for table `dictionary` */

DROP TABLE IF EXISTS `dictionary`;

CREATE TABLE `dictionary` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '字典ID',
  `type_id` int(11) DEFAULT NULL COMMENT '字典类型ID',
  `key` char(50) DEFAULT NULL COMMENT '字典值',
  `value` char(50) DEFAULT NULL COMMENT '字典名称',
  `order` int(4) DEFAULT '0' COMMENT '排序顺序',
  `describe` char(250) DEFAULT NULL COMMENT '字典描述',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=179 DEFAULT CHARSET=utf8;

/*Data for the table `dictionary` */

insert  into `dictionary`(`id`,`type_id`,`key`,`value`,`order`,`describe`,`delete`) values (-102,-1,'other','其他',0,'系统分类（不允许删除）',0),(-101,-1,'system','系统',0,'系统分类（不允许删除）',0),(2,41,'test','测试',1,'这是测试',0),(3,42,'test1','测试1',1,'这是测试1',0),(4,42,'test2','测试2',2,'这是测试2',0),(-301,-3,'list','获取列表',1,'系统分类（不允许删除）',0),(36,63,'','',0,'',0),(35,62,'','',0,'',0),(34,62,'','',0,'',0),(33,62,'','',0,'',0),(32,61,'','',0,'',0),(31,61,'','',0,'',0),(30,61,'','',0,'',0),(29,60,'','',0,'',0),(-308,-3,'login','登录',2,'系统分类（不允许删除）',0),(-307,-3,'is-use','变更',3,'系统分类（不允许删除）',0),(-306,-3,'tree','获取',4,'系统分类（不允许删除）',0),(-305,-3,'delete','删除',5,'系统分类（不允许删除）',0),(-304,-3,'edit','编辑',6,'系统分类（不允许删除）',0),(-303,-3,'get','获取',7,'系统分类（不允许删除）',0),(-302,-3,'add','添加',8,'系统分类（不允许删除）',0),(37,63,'','',0,'',0),(38,63,'','',0,'',0),(39,64,'','',0,'',0),(40,64,'','',0,'',0),(41,64,'','',0,'',0),(42,65,'','',0,'',0),(43,65,'','',0,'',0),(44,65,'','',0,'',0),(45,66,'','',0,'',0),(46,66,'','',0,'',0),(47,66,'','',0,'',0),(48,67,'','',0,'',0),(49,67,'','',0,'',0),(50,67,'','',0,'',0),(51,68,'','',0,'',0),(52,68,'','',0,'',0),(53,68,'','',0,'',0),(54,69,'','',0,'',0),(55,69,'','',0,'',0),(56,69,'','',0,'',0),(57,70,'','',0,'',0),(58,70,'','',0,'',0),(59,70,'','',0,'',0),(60,71,'','',0,'',0),(61,71,'','',0,'',0),(62,71,'','',0,'',0),(63,72,'','',0,'',0),(64,72,'','',0,'',0),(65,72,'','',0,'',0),(66,73,'','',0,'',0),(67,73,'','',0,'',0),(68,73,'','',0,'',0),(69,74,'','',0,'',0),(70,74,'','',0,'',0),(71,74,'','',0,'',0),(72,75,'','',0,'',0),(73,75,'','',0,'',0),(74,75,'','',0,'',0),(75,76,'','',0,'',0),(76,76,'','',0,'',0),(77,76,'','',0,'',0),(78,77,'','',0,'',0),(79,77,'','',0,'',0),(80,77,'','',0,'',0),(81,78,'test1','测试1',1,'这是测试1',0),(82,78,'test2','测试2',2,'这是测试2',0),(83,78,'test3','测试3',3,'这是测试3',0),(84,79,'test1','测试1',1,'这是测试1',0),(85,79,'test2','测试2',2,'这是测试2',0),(86,79,'test3','测试3',3,'这是测试3',0),(87,80,'test1','测试1',1,'这是测试1',0),(88,80,'test2','测试2',2,'这是测试2',0),(89,80,'test3','测试3',3,'这是测试3',0),(90,81,'test1','测试10',1,'这是测试10',0),(91,81,'test2','测试2',2,'这是测试2',1),(92,81,'test3','测试30',3,'这是测试30',0),(93,82,'test1','测试1',1,'这是测试1',0),(94,82,'test2','测试2',2,'这是测试2',0),(95,82,'test3','测试3',3,'这是测试3',0),(96,83,'test1','测试1',1,'这是测试1',0),(97,83,'test2','测试2',2,'这是测试2',0),(98,83,'test3','测试3',3,'这是测试3',0),(99,84,'test1','测试1',1,'这是测试1',0),(100,84,'test2','测试2',2,'这是测试2',0),(101,84,'test3','测试3',3,'这是测试3',0),(102,81,'test2','测试2',2,'这是测试2',1),(103,81,'test2','测试20',2,'这是测试20',1),(104,81,'test2','测试20',2,'这是测试20',1),(105,81,'test2','测试20',2,'这是测试20',1),(106,85,'test1','测试1',1,'这是测试1',0),(107,85,'test2','测试2',2,'这是测试2',0),(108,85,'test3','测试3',3,'这是测试3',0),(109,86,'test1','测试1',1,'这是测试1',0),(110,86,'test2','测试2',2,'这是测试2',0),(111,86,'test3','测试3',3,'这是测试3',0),(112,81,'test2','测试20',2,'这是测试20',1),(113,87,'test1','测试1',1,'这是测试1',0),(114,87,'test2','测试2',2,'这是测试2',0),(115,87,'test3','测试3',3,'这是测试3',0),(116,88,'test1','测试1',1,'这是测试1',0),(117,88,'test2','测试2',2,'这是测试2',0),(118,88,'test3','测试3',3,'这是测试3',0),(119,81,'test2','测试20',2,'这是测试20',1),(120,89,'test1','测试1',1,'这是测试1',0),(121,89,'test2','测试2',2,'这是测试2',0),(122,89,'test3','测试3',3,'这是测试3',0),(123,81,'test2','测试20',2,'这是测试20',1),(124,81,'test2','测试20',2,'这是测试20',1),(125,81,'test2','测试20',2,'这是测试20',1),(126,90,'test1','测试1',1,'这是测试1',0),(127,90,'test2','测试2',2,'这是测试2',0),(128,90,'test3','测试3',3,'这是测试3',0),(129,81,'test2','测试20',2,'这是测试20',1),(130,91,'','',0,'',0),(131,92,'key','value',0,'sdasda',0),(132,93,'a','1',0,'3',0),(133,93,'b','2',0,'4',0),(134,94,'3','2',1,'6',0),(135,94,'65','56',2,'56',0),(136,95,'key1','value1',1,'这是备注1',0),(137,95,'key2','value2',2,'这是备注2',0),(138,96,'test1','测试1',1,'这是测试1',0),(139,96,'test2','测试2',2,'这是测试2',0),(140,96,'test3','测试3',3,'这是测试3',0),(141,81,'test2','测试20',2,'这是测试20',1),(142,97,'test1','测试1',1,'这是测试1',0),(143,97,'test2','测试2',2,'这是测试2',0),(144,97,'test3','测试3',3,'这是测试3',0),(145,81,'test2','测试20',2,'这是测试20',1),(146,98,'test1','测试1',1,'这是测试1',0),(147,98,'test2','测试2',2,'这是测试2',0),(148,98,'test3','测试3',3,'这是测试3',0),(149,81,'test2','测试20',2,'这是测试20',0),(150,112,'t)x7','tYtU^I',1,'I5ts',0),(151,112,'eF8^J','8hSD8*Y',2,'#t[LW',0),(152,112,'Ymsw9','CHlosJ',3,'vvK*i',0),(-201,-2,'admin','管理员',1,'系统管理员',0),(-202,-2,'management','负责人',2,'部门负责人',0),(-203,-2,'staff','业务员',3,'部门业务员',0),(156,114,'test1','测试1',1,'这是测试1',0),(157,114,'test2','测试2',2,'这是测试2',0),(158,114,'test3','测试3',3,'这是测试3',0),(159,115,'test1','测试1',1,'这是测试1',0),(160,115,'test2','测试2',2,'这是测试2',0),(161,115,'test3','测试3',3,'这是测试3',0),(162,116,'test1','测试1',1,'这是测试1',0),(163,116,'test2','测试2',2,'这是测试2',0),(164,116,'test3','测试3',3,'这是测试3',0),(165,117,'test1','测试1',1,'这是测试1',0),(166,117,'test2','测试2',2,'这是测试2',0),(167,117,'test3','测试3',3,'这是测试3',0),(168,118,'test1','测试1',1,'这是测试1',0),(169,118,'test2','测试2',2,'这是测试2',0),(170,118,'test3','测试3',3,'这是测试3',0),(171,119,'test1','测试1',1,'这是测试1',0),(172,119,'test2','测试2',2,'这是测试2',0),(173,119,'test3','测试3',3,'这是测试3',0),(174,120,'test1','测试1',1,'这是测试1',0),(175,120,'test2','测试2',2,'这是测试2',0),(176,120,'test3','测试3',3,'这是测试3',0),(-309,-3,'password','编辑',9,'系统分类（不允许删除）',0),(-310,-3,'upload','文件',10,'系统分类（不允许删除）',0);

/*Table structure for table `dictionary_type` */

DROP TABLE IF EXISTS `dictionary_type`;

CREATE TABLE `dictionary_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '字典类型ID',
  `type_id` int(11) NOT NULL COMMENT '字典分类ID',
  `key` char(20) NOT NULL COMMENT '字典分类键',
  `title` char(50) NOT NULL COMMENT '字典类型名称',
  `is_use` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `user_id` int(11) DEFAULT NULL COMMENT '创建人ID',
  `describe` char(250) DEFAULT NULL COMMENT '字典类型描述',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '软删除',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=121 DEFAULT CHARSET=utf8;

/*Data for the table `dictionary_type` */

insert  into `dictionary_type`(`id`,`type_id`,`key`,`title`,`is_use`,`update_time`,`user_id`,`describe`,`delete`) values (-1,-1,'system','字典分类',1,'2018-11-12 15:37:43',0,'系统字典分类',0),(-3,-1,'system','日志分类',1,'2018-11-29 18:15:50',0,'系统日志分类',0),(119,3,'yes','666',0,'2018-11-28 17:18:03',0,'这是描述文字11',0),(118,3,'yes','666',0,'2018-11-28 17:14:02',0,'这是描述文字11',0),(117,3,'yes','666',0,'2018-11-28 16:43:54',0,'这是描述文字11',0),(116,3,'yes','666',0,'2018-11-28 16:37:25',0,'这是描述文字11',0),(115,3,'yes','666',0,'2018-11-28 15:40:49',0,'这是描述文字11',0),(114,3,'yes','666',0,'2018-11-28 09:49:40',0,'这是描述文字11',0),(-2,-1,'system','人员角色',1,'2018-11-26 12:08:34',0,'部门人员角色字典',0);

/*Table structure for table `files` */

DROP TABLE IF EXISTS `files`;

CREATE TABLE `files` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '附件ID',
  `name` char(100) DEFAULT NULL COMMENT '附件名称',
  `suffix` char(10) DEFAULT NULL COMMENT '附件后缀',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上传时间',
  `path` char(250) DEFAULT NULL COMMENT '附件存放位置',
  `size` int(10) DEFAULT '0' COMMENT '文件大小（单位B）',
  `file_name` char(100) DEFAULT NULL COMMENT '存放名称',
  `form_id` int(11) DEFAULT '0' COMMENT '外部ID',
  `from` char(10) NOT NULL DEFAULT '-' COMMENT '来自那张表',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=55 DEFAULT CHARSET=utf8;

/*Data for the table `files` */

insert  into `files`(`id`,`name`,`suffix`,`time`,`path`,`size`,`file_name`,`form_id`,`from`,`delete`) values (17,'系统设置','json','2018-11-29 16:23:36','201811/',0,'1543479816_系统设置',0,'-',0),(3,'系统设置','json','2018-11-29 16:05:49','201811/',0,'1543478749_系统设置',0,'-',0),(4,'系统设置','json','2018-11-29 16:05:51','201811/',0,'1543478751_系统设置',0,'-',0),(5,'系统设置','json','2018-11-29 16:05:52','201811/',0,'1543478752_系统设置',0,'-',0),(6,'系统设置','json','2018-11-29 16:05:53','201811/',0,'1543478753_系统设置',0,'-',0),(7,'系统设置','json','2018-11-29 16:05:54','201811/',0,'1543478754_系统设置',0,'-',0),(8,'系统设置','json','2018-11-29 16:10:00','201811/',0,'1543479000_系统设置',0,'-',0),(9,'系统设置','json','2018-11-29 16:10:01','201811/',0,'1543479001_系统设置',0,'-',0),(10,'系统设置','json','2018-11-29 16:10:02','201811/',0,'1543479002_系统设置',0,'-',0),(11,'系统设置','json','2018-11-29 16:10:03','201811/',0,'1543479003_系统设置',0,'-',0),(12,'系统设置','json','2018-11-29 16:10:03','201811/',0,'1543479003_系统设置',0,'-',0),(13,'系统设置','json','2018-11-29 16:10:04','201811/',0,'1543479004_系统设置',0,'-',0),(14,'系统设置','json','2018-11-29 16:10:05','201811/',0,'1543479005_系统设置',0,'-',0),(15,'系统设置','json','2018-11-29 16:10:05','201811/',0,'1543479005_系统设置',0,'-',0),(16,'系统设置','json','2018-11-29 16:10:06','201811/',0,'1543479006_系统设置',0,'-',0),(18,'系统设置','json','2018-11-29 16:32:19','201811/',11672,'1543480339_系统设置',0,'-',0),(19,'vscode-setting','txt','2018-11-29 18:22:25','201811/',1902,'1543486945_vscode-setting',0,'-',0),(20,'vscode-setting','txt','2018-11-29 18:23:07','201811/',1902,'1543486987_vscode-setting',0,'-',0),(21,'刘亦菲','gif','2018-11-29 18:23:55','201811/',5116027,'1543487035_刘亦菲',0,'-',0),(22,'微信图片_20181122135157','jpg','2018-11-29 18:29:10','201811/',11887,'1543487350_微信图片_20181122135157',0,'-',0),(23,'微信图片_20181122135157','jpg','2018-11-29 18:29:25','201811/',11887,'1543487365_微信图片_20181122135157',0,'-',0),(24,'微信图片_20181122135157','jpg','2018-11-29 18:29:28','201811/',11887,'1543487368_微信图片_20181122135157',0,'-',0),(25,'微信图片_20181122135157','jpg','2018-11-29 18:29:33','201811/',11887,'1543487373_微信图片_20181122135157',0,'-',0),(26,'微信图片_20181122135157','jpg','2018-11-29 18:32:49','201811/',11887,'1543487569_微信图片_20181122135157',0,'-',0),(27,'微信图片_20181122135157','jpg','2018-11-29 18:36:27','201811/',11887,'1543487787_微信图片_20181122135157',0,'-',0),(28,'微信图片_20181122135157','jpg','2018-11-29 18:36:32','201811/',11887,'1543487792_微信图片_20181122135157',0,'-',0),(29,'微信图片_20181122135157','jpg','2018-11-29 18:36:35','201811/',11887,'1543487795_微信图片_20181122135157',0,'-',0),(30,'微信图片_20181122135157','jpg','2018-11-29 18:36:58','201811/',11887,'1543487818_微信图片_20181122135157',0,'-',0),(31,'微信图片_20181122135157','jpg','2018-11-29 18:37:10','201811/',11887,'1543487830_微信图片_20181122135157',0,'-',0),(32,'微信图片_20181122135157','jpg','2018-11-29 18:37:35','201811/',11887,'1543487855_微信图片_20181122135157',0,'-',0),(33,'微信图片_20181122135157','jpg','2018-11-29 18:37:38','201811/',11887,'1543487858_微信图片_20181122135157',0,'-',0),(34,'微信图片_20181122135157','jpg','2018-11-29 18:38:41','201811/',11887,'1543487921_微信图片_20181122135157',0,'-',0),(35,'微信图片_20181122135157','jpg','2018-11-29 18:39:03','201811/',11887,'1543487943_微信图片_20181122135157',0,'-',0),(36,'微信图片_20181122135157','jpg','2018-11-29 18:39:07','201811/',11887,'1543487947_微信图片_20181122135157',0,'-',0),(37,'微信图片_20181122135157','jpg','2018-11-29 18:39:10','201811/',11887,'1543487950_微信图片_20181122135157',0,'-',0),(38,'微信图片_20181122135157','jpg','2018-11-29 18:39:57','201811/',11887,'1543487997_微信图片_20181122135157',0,'-',0),(39,'微信图片_20181122135157','jpg','2018-11-29 18:40:01','201811/',11887,'1543488001_微信图片_20181122135157',0,'-',0),(40,'微信图片_20181122135157','jpg','2018-11-29 18:40:03','201811/',11887,'1543488003_微信图片_20181122135157',0,'-',0),(41,'微信图片_20181122135157','jpg','2018-11-29 18:41:39','201811/',11887,'1543488099_微信图片_20181122135157',0,'-',0),(42,'微信图片_20181122135157','jpg','2018-11-29 18:41:42','201811/',11887,'1543488102_微信图片_20181122135157',0,'-',0),(43,'微信图片_20181122135157','jpg','2018-11-29 18:41:45','201811/',11887,'1543488105_微信图片_20181122135157',0,'-',0),(44,'刘亦菲','gif','2018-11-29 18:42:04','201811/',5116027,'1543488124_刘亦菲',0,'-',0),(45,'刘亦菲','gif','2018-11-29 18:42:07','201811/',5116027,'1543488127_刘亦菲',0,'-',0),(46,'js运行机制','png','2018-11-29 18:44:55','201811/',187655,'1543488295_js运行机制',0,'-',0),(47,'刘亦菲','gif','2018-11-29 18:46:48','201811/',5116027,'1543488408_刘亦菲',0,'-',0),(48,'刘亦菲','gif','2018-11-29 18:46:53','201811/',5116027,'1543488413_刘亦菲',0,'-',0),(49,'刘亦菲','gif','2018-11-29 18:47:00','201811/',5116027,'1543488420_刘亦菲',0,'-',0),(50,'刘亦菲','gif','2018-11-29 18:52:41','201811/',5116027,'1543488761_刘亦菲',0,'-',0),(51,'刘亦菲','gif','2018-11-29 18:53:02','201811/',5116027,'1543488782_刘亦菲',0,'-',0),(52,'刘亦菲','gif','2018-11-29 18:53:15','201811/',5116027,'1543488795_刘亦菲',0,'-',0),(53,'刘亦菲','gif','2018-11-29 18:53:18','201811/',5116027,'1543488798_刘亦菲',0,'-',0),(54,'刘亦菲','gif','2018-11-29 18:53:22','201811/',5116027,'1543488802_刘亦菲',0,'-',0);

/*Table structure for table `integral` */

DROP TABLE IF EXISTS `integral`;

CREATE TABLE `integral` (
  `id` int(11) NOT NULL COMMENT '积分ID',
  `cognizance_user_id` int(11) DEFAULT NULL COMMENT '认定人ID',
  `responsibility_user_id` int(11) DEFAULT NULL COMMENT '责任人ID',
  `manuscript_id` int(11) DEFAULT NULL COMMENT '工作底稿ID',
  `receipt_id` int(11) DEFAULT NULL COMMENT '回执ID',
  `score` int(11) DEFAULT NULL COMMENT '分值（除以100显示）',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '日期',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删|删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `integral` */

/*Table structure for table `integral_edit` */

DROP TABLE IF EXISTS `integral_edit`;

CREATE TABLE `integral_edit` (
  `id` int(11) NOT NULL COMMENT '积分修改ID',
  `integral_id` int(11) DEFAULT NULL COMMENT '积分表ID',
  `score` int(11) DEFAULT NULL COMMENT '积分（除以100显示）',
  `user_id` int(11) DEFAULT NULL COMMENT '发起人修改ID',
  `describe` varchar(5000) DEFAULT NULL COMMENT '描述',
  `state` enum('draft','report','reject','adopt') DEFAULT 'draft' COMMENT '状态（draft草稿|report上报|reject驳回|adopt通过）',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `integral_edit` */

/*Table structure for table `login` */

DROP TABLE IF EXISTS `login`;

CREATE TABLE `login` (
  `login_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '登录id',
  `user_code` int(11) NOT NULL COMMENT '员工号',
  `password` char(50) DEFAULT NULL COMMENT '登录密码',
  `is_use` tinyint(1) DEFAULT '0' COMMENT '是否禁用（1启用0禁用）',
  `change_pd_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后修改密码时间',
  `login_num` int(6) DEFAULT '0' COMMENT '登录次数',
  `login_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `author_id` int(11) DEFAULT NULL COMMENT '授权人',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`user_code`),
  KEY `login_id` (`login_id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

/*Data for the table `login` */

insert  into `login`(`login_id`,`user_code`,`password`,`is_use`,`change_pd_time`,`login_num`,`login_time`,`author_id`,`delete`) values (3,10001,'3b1ac6e9a16ff76879e5888483e118a8',1,'2018-11-29 18:10:57',0,'0000-00-00 00:00:00',2,0),(4,10002,'78dc8bbc86eb472b3db1d0b025714ec1',1,'0000-00-00 00:00:00',0,'2018-11-29 18:28:33',2,0);

/*Table structure for table `logs` */

DROP TABLE IF EXISTS `logs`;

CREATE TABLE `logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `url` char(200) DEFAULT NULL COMMENT 'URL',
  `user_id` int(11) DEFAULT NULL COMMENT '操作人ID',
  `msg` varchar(5000) DEFAULT NULL COMMENT '操作说明',
  `method` char(10) DEFAULT NULL COMMENT '请求方法',
  `data` varchar(5000) DEFAULT NULL COMMENT '参数',
  `time` timestamp NULL DEFAULT NULL COMMENT '操作时间',
  `server` char(20) DEFAULT NULL COMMENT '服务',
  `ip` char(15) DEFAULT NULL COMMENT 'IP地址',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=1942 DEFAULT CHARSET=utf8;

/*Data for the table `logs` */

insert  into `logs`(`id`,`url`,`user_id`,`msg`,`method`,`data`,`time`,`server`,`ip`,`delete`) values (1857,'/api/systemSetup/log/list',2,'获取日志列表','POST','-','2018-11-29 18:46:28','systemSetup','192.168.1.20',0),(1858,'/api/worker/file/upload',2,'上传文件','POST','-','2018-11-29 18:46:48','worker','192.168.1.13',0),(1859,'/api/worker/file/upload',2,'上传文件','POST','-','2018-11-29 18:46:53','worker','192.168.1.13',0),(1860,'/api/worker/file/upload',2,'上传文件','POST','-','2018-11-29 18:47:00','worker','192.168.1.13',0),(1861,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:50:05','systemSetup','192.168.1.20',0),(1862,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:50:05','systemSetup','192.168.1.20',0),(1863,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:50:09','systemSetup','192.168.1.20',0),(1864,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:50:09','systemSetup','192.168.1.20',0),(1865,'/api/systemSetup/dictionaries/get',2,'获取字典','GET','id=-1&rnd=0.26184496348662334','2018-11-29 18:50:09','systemSetup','192.168.1.20',0),(1866,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:50:11','systemSetup','192.168.1.20',0),(1867,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:50:11','systemSetup','192.168.1.20',0),(1868,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:50:22','systemSetup','192.168.1.20',0),(1869,'/api/org/clause/list',2,'获取列表','POST','-','2018-11-29 18:50:24','org','192.168.1.20',0),(1870,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.07988402732560806','2018-11-29 18:50:24','org','192.168.1.20',0),(1871,'/api/org/user/list',2,'获取人员列表','POST','-','2018-11-29 18:50:27','org','192.168.1.20',0),(1872,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.8881055103830275','2018-11-29 18:50:27','org','192.168.1.20',0),(1873,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.07862746411675436','2018-11-29 18:50:29','org','192.168.1.20',0),(1874,'/api/org/department/list',2,'获取部门/结构/网点列表','POST','-','2018-11-29 18:50:30','org','192.168.1.20',0),(1875,'/api/org/notice/list',2,'获取通知公告列表','POST','-','2018-11-29 18:50:32','org','192.168.1.20',0),(1876,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.27292329057203135','2018-11-29 18:50:32','org','192.168.1.20',0),(1877,'/api/systemSetup/log/list',2,'获取日志列表','POST','-','2018-11-29 18:50:39','systemSetup','192.168.1.20',0),(1878,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:50:41','systemSetup','192.168.1.20',0),(1879,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:50:41','systemSetup','192.168.1.20',0),(1880,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:50:42','systemSetup','192.168.1.20',0),(1881,'/api/systemSetup/dictionaries/get',2,'获取字典','GET','id=-1&rnd=0.567638259881063','2018-11-29 18:50:43','systemSetup','192.168.1.20',0),(1882,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:50:43','systemSetup','192.168.1.20',0),(1883,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:50:48','systemSetup','192.168.1.20',0),(1884,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:50:50','systemSetup','192.168.1.20',0),(1885,'/api/systemSetup/log/list',2,'获取日志列表','POST','-','2018-11-29 18:50:52','systemSetup','192.168.1.20',0),(1886,'/api/systemSetup/log/list',2,'获取日志列表','POST','-','2018-11-29 18:50:55','systemSetup','192.168.1.20',0),(1887,'/api/systemSetup/log/list',2,'获取日志列表','POST','-','2018-11-29 18:50:57','systemSetup','192.168.1.20',0),(1888,'/api/systemSetup/log/list',2,'获取日志列表','POST','-','2018-11-29 18:51:01','systemSetup','192.168.1.20',0),(1889,'/api/systemSetup/log/list',2,'获取日志列表','POST','-','2018-11-29 18:51:03','systemSetup','192.168.1.20',0),(1890,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:51:31','systemSetup','192.168.1.20',0),(1891,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:51:31','systemSetup','192.168.1.20',0),(1892,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:51:34','systemSetup','192.168.1.20',0),(1893,'/api/systemSetup/dictionaries/get',2,'获取字典','GET','id=-1&rnd=0.9079222434261094','2018-11-29 18:51:34','systemSetup','192.168.1.20',0),(1894,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:51:34','systemSetup','192.168.1.20',0),(1895,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.9267410050029019','2018-11-29 18:51:35','org','192.168.1.20',0),(1896,'/api/org/clause/list',2,'获取列表','POST','-','2018-11-29 18:51:35','org','192.168.1.20',0),(1897,'/api/org/user/list',2,'获取人员列表','POST','-','2018-11-29 18:51:37','org','192.168.1.20',0),(1898,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.5111624977617799','2018-11-29 18:51:37','org','192.168.1.20',0),(1899,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.02544853091701782','2018-11-29 18:51:39','org','192.168.1.20',0),(1900,'/api/org/department/list',2,'获取部门/结构/网点列表','POST','-','2018-11-29 18:51:39','org','192.168.1.20',0),(1901,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.5867528098919121','2018-11-29 18:51:39','org','192.168.1.20',0),(1902,'/api/org/notice/list',2,'获取通知公告列表','POST','-','2018-11-29 18:51:39','org','192.168.1.20',0),(1903,'/api/worker/user/get',2,'获取','PUT','-','2018-11-29 18:51:43','worker','192.168.1.20',0),(1904,'/api/worker/file/upload',2,'上传文件','POST','-','2018-11-29 18:52:41','worker','192.168.1.13',0),(1905,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.532120962181972','2018-11-29 18:52:42','org','192.168.1.20',0),(1906,'/api/org/department/list',2,'获取部门/结构/网点列表','POST','-','2018-11-29 18:52:42','org','192.168.1.20',0),(1907,'/api/org/user/list',2,'获取人员列表','POST','-','2018-11-29 18:52:42','org','192.168.1.20',0),(1908,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.05552409353569532','2018-11-29 18:52:42','org','192.168.1.20',0),(1909,'/api/org/clause/list',2,'获取列表','POST','-','2018-11-29 18:52:46','org','192.168.1.20',0),(1910,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.7987151497120664','2018-11-29 18:52:47','org','192.168.1.20',0),(1911,'/api/worker/file/upload',2,'上传文件','POST','-','2018-11-29 18:53:02','worker','192.168.1.13',0),(1912,'/api/worker/file/upload',2,'上传文件','POST','-','2018-11-29 18:53:14','worker','192.168.1.13',0),(1913,'/api/worker/file/upload',2,'上传文件','POST','-','2018-11-29 18:53:18','worker','192.168.1.13',0),(1914,'/api/worker/file/upload',2,'上传文件','POST','-','2018-11-29 18:53:21','worker','192.168.1.13',0),(1915,'/api/org/department/list',2,'获取部门/结构/网点列表','POST','-','2018-11-29 18:55:05','org','192.168.1.13',0),(1916,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.910453360797578','2018-11-29 18:55:05','org','192.168.1.13',0),(1917,'/api/org/user/list',2,'获取人员列表','POST','-','2018-11-29 18:55:06','org','192.168.1.13',0),(1918,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.034586995842187385','2018-11-29 18:55:06','org','192.168.1.13',0),(1919,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.6113692258213184','2018-11-29 18:55:10','org','192.168.1.13',0),(1920,'/api/org/department/list',2,'获取部门/结构/网点列表','POST','-','2018-11-29 18:55:10','org','192.168.1.13',0),(1921,'/api/org/user/list',2,'获取人员列表','POST','-','2018-11-29 18:55:27','org','192.168.1.13',0),(1922,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.9405075362103852','2018-11-29 18:55:27','org','192.168.1.13',0),(1923,'/api/org/user/get',2,'获取人员','GET','rnd=0.06280703302793311','2018-11-29 18:55:34','org','192.168.1.13',0),(1924,'/api/org/user/get',2,'获取人员','GET','id=13&rnd=0.588791268375664','2018-11-29 18:55:45','org','192.168.1.13',0),(1925,'/api/org/user/get',2,'获取人员','GET','rnd=0.37036793179845184','2018-11-29 18:55:49','org','192.168.1.13',0),(1926,'/api/org/user/get',2,'获取人员','GET','id=13&rnd=0.8480723584702499','2018-11-29 18:55:57','org','192.168.1.13',0),(1927,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.8068928929452794','2018-11-29 18:56:13','org','192.168.1.13',0),(1928,'/api/org/department/list',2,'获取部门/结构/网点列表','POST','-','2018-11-29 18:56:13','org','192.168.1.13',0),(1929,'/api/org/department/get',2,'获取部门/结构/网点','GET','id=35&rnd=0.8647393915751471','2018-11-29 18:56:14','org','192.168.1.13',0),(1930,'/api/org/user/list',2,'获取人员列表','POST','-','2018-11-29 18:56:18','org','192.168.1.13',0),(1931,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.8841741596384249','2018-11-29 18:56:18','org','192.168.1.13',0),(1932,'/api/org/user/get',2,'获取人员','GET','id=13&rnd=0.8876736414181217','2018-11-29 18:56:22','org','192.168.1.13',0),(1933,'/api/systemSetup/dictionaries/get',2,'获取字典','GET','id=-1&rnd=0.46708861349672914','2018-11-29 18:56:33','systemSetup','192.168.1.13',0),(1934,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:56:33','systemSetup','192.168.1.13',0),(1935,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:56:33','systemSetup','192.168.1.13',0),(1936,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:56:37','systemSetup','192.168.1.13',0),(1937,'/api/systemSetup/login/list',2,'获取人员列表','POST','-','2018-11-29 18:56:37','systemSetup','192.168.1.13',0),(1938,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.13210680792287488','2018-11-29 18:57:06','org','192.168.1.13',0),(1939,'/api/org/department/list',2,'获取部门/结构/网点列表','POST','-','2018-11-29 18:57:06','org','192.168.1.13',0),(1940,'/api/org/notice/list',2,'获取通知公告列表','POST','-','2018-11-29 18:57:08','org','192.168.1.13',0),(1941,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.4043837256712792','2018-11-29 18:57:08','org','192.168.1.13',0),(1856,'/api/worker/file/upload',2,'上传文件','POST','-','2018-11-29 18:44:55','worker','192.168.1.13',0),(1848,'/api/systemSetup/log/list',2,'获取日志列表','POST','-','2018-11-29 18:43:15','systemSetup','192.168.1.20',0),(1849,'/api/systemSetup/log/list',2,'获取日志列表','POST','-','2018-11-29 18:43:19','systemSetup','192.168.1.20',0),(1850,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:43:29','systemSetup','192.168.1.20',0),(1851,'/api/systemSetup/dictionaries/get',2,'获取字典','GET','id=-1&rnd=0.36223512106415856','2018-11-29 18:43:29','systemSetup','192.168.1.20',0),(1852,'/api/systemSetup/dictionaries/list',2,'获取字典列表','POST','-','2018-11-29 18:43:29','systemSetup','192.168.1.20',0),(1853,'/api/org/notice/list',2,'获取通知公告列表','POST','-','2018-11-29 18:43:33','org','192.168.1.20',0),(1854,'/api/org/department/tree',2,'获取部门/结构/网点','GET','parentId=-1&rnd=0.8575095303974929','2018-11-29 18:43:33','org','192.168.1.20',0),(1855,'/api/org/notice/get',2,'获取通知公告','GET','id=39&rnd=0.32166533458646684','2018-11-29 18:43:41','org','192.168.1.20',0);

/*Table structure for table `manuscript` */

DROP TABLE IF EXISTS `manuscript`;

CREATE TABLE `manuscript` (
  `id` int(11) NOT NULL COMMENT '工作底稿ID',
  `query_department_id` int(11) NOT NULL COMMENT '被查询机构ID',
  `department_id` int(11) NOT NULL COMMENT '检查机构ID',
  `number` char(100) DEFAULT NULL COMMENT '编号',
  `project_name` char(250) DEFAULT NULL COMMENT '项目名称',
  `state` enum('draft','publish') DEFAULT NULL COMMENT '底稿状态（draft草稿|publish发布）',
  PRIMARY KEY (`id`,`query_department_id`,`department_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `manuscript` */

/*Table structure for table `manuscript_admin_user` */

DROP TABLE IF EXISTS `manuscript_admin_user`;

CREATE TABLE `manuscript_admin_user` (
  `manuscript_id` int(11) DEFAULT NULL COMMENT '底稿ID',
  `user_id` int(11) DEFAULT NULL COMMENT '复查人员ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `manuscript_admin_user` */

/*Table structure for table `manuscript_content` */

DROP TABLE IF EXISTS `manuscript_content`;

CREATE TABLE `manuscript_content` (
  `id` int(11) NOT NULL COMMENT '底稿内容ID',
  `manuscript_id` int(11) NOT NULL COMMENT '底稿ID',
  `type` enum('behavior','other') DEFAULT 'other' COMMENT '内容类型（behavior违规行为|other其他）',
  `behavior_id` int(11) DEFAULT NULL COMMENT '违规行为ID',
  `behavior_content` varchar(2000) DEFAULT NULL COMMENT '违规行为描述或者其他',
  PRIMARY KEY (`id`,`manuscript_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `manuscript_content` */

/*Table structure for table `manuscript_file` */

DROP TABLE IF EXISTS `manuscript_file`;

CREATE TABLE `manuscript_file` (
  `manuscript_id` int(11) NOT NULL COMMENT '底稿ID',
  `file_id` int(11) NOT NULL COMMENT '附件ID',
  PRIMARY KEY (`manuscript_id`,`file_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `manuscript_file` */

/*Table structure for table `manuscript_inspect_user` */

DROP TABLE IF EXISTS `manuscript_inspect_user`;

CREATE TABLE `manuscript_inspect_user` (
  `manuscript_id` int(11) DEFAULT NULL COMMENT '底稿ID',
  `user_id` int(11) DEFAULT NULL COMMENT '被检查人员ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `manuscript_inspect_user` */

/*Table structure for table `manuscript_query_user` */

DROP TABLE IF EXISTS `manuscript_query_user`;

CREATE TABLE `manuscript_query_user` (
  `manuscript_id` int(11) DEFAULT NULL COMMENT '底稿ID',
  `user_id` int(11) DEFAULT NULL COMMENT '检查人员ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `manuscript_query_user` */

/*Table structure for table `manuscript_review_user` */

DROP TABLE IF EXISTS `manuscript_review_user`;

CREATE TABLE `manuscript_review_user` (
  `manuscript_id` int(11) DEFAULT NULL COMMENT '底稿ID',
  `user_id` int(11) DEFAULT NULL COMMENT '负责人员ID'
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `manuscript_review_user` */

/*Table structure for table `notice` */

DROP TABLE IF EXISTS `notice`;

CREATE TABLE `notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `department_id` int(11) NOT NULL COMMENT '机构ID',
  `title` char(100) DEFAULT NULL COMMENT '标题',
  `content` longtext COMMENT '公告内容',
  `time` timestamp NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '发布时间',
  `range` enum('1','2') DEFAULT '1' COMMENT '通知范围（1所有部门，2指定部门）',
  `state` enum('draft','publish') DEFAULT 'draft' COMMENT '状态（draft草稿publish发布）',
  `delete` int(11) DEFAULT '0' COMMENT '删除（0未删1已删）',
  PRIMARY KEY (`id`,`department_id`)
) ENGINE=MyISAM AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;

/*Data for the table `notice` */

insert  into `notice`(`id`,`department_id`,`title`,`content`,`time`,`range`,`state`,`delete`) values (1,1,'公告1','公告内容1','2018-11-26 15:20:19','2','',1),(2,1,'公告1','公告内容1','2018-11-26 15:22:20','2','',1),(3,1,'公告1','公告内容1','2018-11-29 16:51:15','2','publish',0),(4,1,'公告1','公告内容1','2018-11-29 16:51:17','2','publish',0),(5,1,'公告1','公告内容1','2018-11-26 16:10:55','2','',1),(6,1,'公告1','公告内容1','2018-11-26 16:12:04','2','',1),(7,1,'公告1','公告内容1','2018-11-29 16:51:19','2','publish',0),(8,1,'公告1','公告内容1','2018-11-29 16:51:21','2','publish',0),(9,1,'公告1','公告内容1','2018-11-29 16:51:22','2','publish',0),(10,1,'公告1','公告内容1','2018-11-29 16:51:24','2','publish',0),(11,1,'公告1','公告内容1','2018-11-29 16:51:26','2','publish',0),(12,1,'公告1','公告内容1','2018-11-29 16:51:28','2','publish',0),(13,1,'公告1','公告内容1','2018-11-29 16:51:31','2','publish',0),(14,1,'公告1','公告内容1','2018-11-29 16:51:33','2','publish',0),(15,1,'公告1','公告内容1','2018-11-26 16:18:31','2','',1),(16,1,'公告1','公告内容1','2018-11-26 17:15:51','1','',1),(17,1,'公告1','公告内容1','2018-11-26 16:19:38','2','',1),(18,1,'公告18','公告内容18','2018-11-28 12:12:38','2','draft',0),(19,1,'公告1','公告内容1','2018-11-27 09:44:51','2','draft',0),(20,1,'公告1','公告内容1','2018-11-29 16:51:10','2','publish',0),(21,1,'公告1','公告内容1','2018-11-29 16:51:12','2','publish',0),(22,1,'公告1','公告内容1','2018-11-29 16:50:37','2','publish',0),(23,1,'公告1','公告内容1','2018-11-29 16:50:36','2','publish',0),(24,1,'公告1','公告内容1','2018-11-29 16:50:32','2','publish',0),(25,1,'公告1','公告内容1','2018-11-28 10:09:58','2','draft',1),(26,1,'公告1','<p>公告内容1</p>','2018-11-28 10:19:04','2','draft',1),(27,1,'公告1','<p>公告内容234567</p>','2018-11-28 10:12:35','1','draft',0),(28,1,'公告1','公告内容1','2018-11-28 10:09:44','2','draft',1),(29,1,'公告1','公告内容1','2018-11-29 16:50:18','2','publish',0),(30,1,'','','2018-11-29 17:49:36','2','draft',1),(31,1,'','','2018-11-29 17:49:33','','draft',1),(32,1,'','','2018-11-29 18:10:48','','draft',1),(33,1,'','','2018-11-29 18:10:42','','draft',1),(34,1,'','','2018-11-29 18:10:47','','draft',1),(35,1,'','','2018-11-29 18:10:44','','draft',1),(36,1,'','','2018-11-29 18:11:15','','draft',0),(37,1,'','','2018-11-29 18:12:01','','draft',0),(38,1,'','','2018-11-29 18:20:14','','draft',0),(39,1,'','','2018-11-29 18:22:34','','draft',0);

/*Table structure for table `notice_file` */

DROP TABLE IF EXISTS `notice_file`;

CREATE TABLE `notice_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `file_id` int(11) NOT NULL COMMENT '通知附件ID',
  `notice_id` int(11) NOT NULL COMMENT '通知ID',
  PRIMARY KEY (`id`,`notice_id`)
) ENGINE=MyISAM AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

/*Data for the table `notice_file` */

insert  into `notice_file`(`id`,`file_id`,`notice_id`) values (1,1,3),(2,2,3),(3,3,3),(4,1,4),(5,2,4),(6,3,4),(7,1,13),(8,2,13),(9,3,13),(10,1,14),(11,2,14),(12,3,14),(13,1,16),(14,2,16);

/*Table structure for table `notice_inform` */

DROP TABLE IF EXISTS `notice_inform`;

CREATE TABLE `notice_inform` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `department_id` int(11) NOT NULL COMMENT '机构ID',
  `notice_id` int(11) NOT NULL COMMENT '公告ID',
  PRIMARY KEY (`id`,`notice_id`)
) ENGINE=MyISAM AUTO_INCREMENT=39 DEFAULT CHARSET=utf8;

/*Data for the table `notice_inform` */

insert  into `notice_inform`(`id`,`department_id`,`notice_id`) values (1,1,3),(2,2,3),(3,3,3),(4,4,3),(5,1,4),(6,2,4),(7,3,4),(8,4,4),(9,2,13),(10,23,13),(11,5,13),(12,5,14),(36,5,18),(35,2,18),(34,1,18),(18,1,19),(22,1,20),(23,1,21),(24,1,22),(25,1,23),(26,1,24),(38,1,30),(37,1,29),(29,1,27);

/*Table structure for table `punish_notice` */

DROP TABLE IF EXISTS `punish_notice`;

CREATE TABLE `punish_notice` (
  `id` int(11) NOT NULL COMMENT '处罚通知ID',
  `confirmation_id` int(11) DEFAULT NULL COMMENT '确认书ID',
  `manuscript_id` int(11) DEFAULT NULL COMMENT '工作底稿ID',
  `clause_id` int(11) DEFAULT NULL COMMENT '管理办法ID',
  `integral_id` int(11) DEFAULT NULL COMMENT '积分表ID',
  `number` char(50) DEFAULT NULL COMMENT '文件号',
  `state` enum('draft','publish') DEFAULT NULL COMMENT '状态（draft草稿|publish发布）',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `punish_notice` */

/*Table structure for table `user_job` */

DROP TABLE IF EXISTS `user_job`;

CREATE TABLE `user_job` (
  `user_id` int(11) NOT NULL COMMENT '人员ID',
  `job` char(20) NOT NULL COMMENT '岗位',
  PRIMARY KEY (`user_id`,`job`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `user_job` */

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '人员ID',
  `department_id` int(11) NOT NULL COMMENT '部门ID',
  `user_name` char(50) DEFAULT NULL COMMENT '姓名',
  `user_code` char(20) NOT NULL COMMENT '员工号',
  `class` char(20) DEFAULT NULL COMMENT '民族',
  `sex` enum('0','1','2') NOT NULL DEFAULT '0' COMMENT '性别（0保密1女2男）',
  `id_card` char(18) DEFAULT NULL COMMENT '身份证',
  `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `delete` tinyint(1) DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`user_id`,`department_id`,`user_code`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

/*Data for the table `users` */

insert  into `users`(`user_id`,`department_id`,`user_name`,`user_code`,`class`,`sex`,`id_card`,`update_time`,`delete`) values (1,1,'小明','10001','','0','522623-----------','2018-11-26 17:04:36',0),(2,0,'小王','10002',NULL,'0',NULL,'2018-11-27 14:06:43',0),(11,1,'测试','10086','苗','','56541321564','2018-11-26 17:06:32',0),(13,1,'张国荣','10234','汉','0','5649825134546','2018-11-29 15:53:22',0),(12,1,'1','21','汉','0','123456578','2018-11-28 09:49:53',0);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
