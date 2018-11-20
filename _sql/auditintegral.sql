/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.5.53-log : Database - auditintegral
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`auditintegral` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `auditintegral`;

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
  `id` int(9) NOT NULL COMMENT '部门ID',
  `parent_id` int(9) DEFAULT NULL COMMENT '上级部门ID',
  `code` char(50) DEFAULT NULL COMMENT '部门编码',
  `level` char(20) DEFAULT NULL COMMENT '部门级别',
  `grade` int(3) DEFAULT NULL COMMENT '所在部门树的层级',
  `address` char(250) DEFAULT NULL COMMENT '地址',
  `phone` char(11) DEFAULT NULL COMMENT '手机号或者电话',
  `update_time` time DEFAULT NULL COMMENT '更新时间',
  `delete` tinyint(1) DEFAULT '0' COMMENT '是否删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `department` */

/*Table structure for table `department_user` */

DROP TABLE IF EXISTS `department_user`;

CREATE TABLE `department_user` (
  `department_id` int(11) NOT NULL COMMENT '部门id',
  `user_id` int(11) DEFAULT NULL COMMENT '人员id',
  `type` char(20) DEFAULT NULL COMMENT 'type',
  PRIMARY KEY (`department_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `department_user` */

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
) ENGINE=MyISAM AUTO_INCREMENT=90 DEFAULT CHARSET=utf8;

/*Data for the table `dictionary` */

insert  into `dictionary`(`id`,`type_id`,`key`,`value`,`order`,`describe`,`delete`) values (-1,-1,'other','其他',0,'系统分类（不允许删除）',0),(1,-1,'system','系统',0,NULL,0),(2,41,'test','测试',1,'这是测试',0),(3,42,'test1','测试1',1,'这是测试1',0),(4,42,'test2','测试2',2,'这是测试2',0),(21,58,'','',0,'',0),(36,63,'','',0,'',0),(35,62,'','',0,'',0),(34,62,'','',0,'',0),(33,62,'','',0,'',0),(32,61,'','',0,'',0),(31,61,'','',0,'',0),(30,61,'','',0,'',0),(29,60,'','',0,'',0),(28,60,'','',0,'',0),(27,60,'','',0,'',0),(26,59,'','',0,'',0),(25,59,'','',0,'',0),(24,59,'','',0,'',0),(23,58,'','',0,'',0),(22,58,'','',0,'',0),(37,63,'','',0,'',0),(38,63,'','',0,'',0),(39,64,'','',0,'',0),(40,64,'','',0,'',0),(41,64,'','',0,'',0),(42,65,'','',0,'',0),(43,65,'','',0,'',0),(44,65,'','',0,'',0),(45,66,'','',0,'',0),(46,66,'','',0,'',0),(47,66,'','',0,'',0),(48,67,'','',0,'',0),(49,67,'','',0,'',0),(50,67,'','',0,'',0),(51,68,'','',0,'',0),(52,68,'','',0,'',0),(53,68,'','',0,'',0),(54,69,'','',0,'',0),(55,69,'','',0,'',0),(56,69,'','',0,'',0),(57,70,'','',0,'',0),(58,70,'','',0,'',0),(59,70,'','',0,'',0),(60,71,'','',0,'',0),(61,71,'','',0,'',0),(62,71,'','',0,'',0),(63,72,'','',0,'',0),(64,72,'','',0,'',0),(65,72,'','',0,'',0),(66,73,'','',0,'',0),(67,73,'','',0,'',0),(68,73,'','',0,'',0),(69,74,'','',0,'',0),(70,74,'','',0,'',0),(71,74,'','',0,'',0),(72,75,'','',0,'',0),(73,75,'','',0,'',0),(74,75,'','',0,'',0),(75,76,'','',0,'',0),(76,76,'','',0,'',0),(77,76,'','',0,'',0),(78,77,'','',0,'',0),(79,77,'','',0,'',0),(80,77,'','',0,'',0),(81,78,'test1','测试1',1,'这是测试1',0),(82,78,'test2','测试2',2,'这是测试2',0),(83,78,'test3','测试3',3,'这是测试3',0),(84,79,'test1','测试1',1,'这是测试1',0),(85,79,'test2','测试2',2,'这是测试2',0),(86,79,'test3','测试3',3,'这是测试3',0),(87,80,'test1','测试1',1,'这是测试1',0),(88,80,'test2','测试2',2,'这是测试2',0),(89,80,'test3','测试3',3,'这是测试3',0);

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
) ENGINE=MyISAM AUTO_INCREMENT=81 DEFAULT CHARSET=utf8;

/*Data for the table `dictionary_type` */

insert  into `dictionary_type`(`id`,`type_id`,`key`,`title`,`is_use`,`update_time`,`user_id`,`describe`,`delete`) values (-1,-1,'dictionaryType','字典分类',1,'2018-11-12 15:37:43',0,'系统字典分类',0),(-2,-1,'jobType','人员岗位',1,'2018-11-12 15:58:30',1,'人员岗位分类',0),(1,1,'yes','666',1,'2018-11-15 16:03:16',1,'这是描述文字',0),(2,1,'yes','666',1,'2018-11-15 16:03:19',0,'这是描述文字',0),(3,1,'yes','666',1,'2018-11-15 16:03:22',0,'这是描述文字',1),(4,1,'yes','666',1,'2018-11-15 16:03:24',0,'这是描述文字',1),(5,1,'yes','666',1,'2018-11-15 16:03:27',0,'这是描述文字',1),(6,1,'yes','666',1,'2018-11-15 16:03:29',0,'这是描述文字',0),(7,1,'yes','666',1,'2018-11-14 14:12:48',0,'这是描述文字',0),(8,1,'yes','666',0,'2018-11-14 14:13:01',0,'这是描述文字',0),(9,1,'yes','666',0,'2018-11-14 15:56:19',0,'这是描述文字',0),(10,1,'yes','666',0,'2018-11-15 11:03:39',0,'这是描述文字',0),(11,1,'yes','666',0,'2018-11-15 13:31:25',0,'这是描述文字',0),(12,1,'yes','666',0,'2018-11-15 16:18:42',0,'这是描述文字',0),(13,1,'yes','666',0,'2018-11-15 16:18:42',0,'这是描述文字',0),(14,1,'yes','666',0,'2018-11-15 16:18:43',0,'这是描述文字',0),(15,1,'yes','666',0,'2018-11-15 16:18:49',0,'这是描述文字',0),(16,1,'yes','666',0,'2018-11-15 16:18:52',0,'这是描述文字',0),(17,1,'yes','666',0,'2018-11-15 16:18:56',0,'这是描述文字',0),(18,1,'yes','666',0,'2018-11-15 16:18:58',0,'这是描述文字',0),(19,1,'yes','666',0,'2018-11-15 16:19:00',0,'这是描述文字',0),(20,1,'yes','666',0,'2018-11-15 16:19:03',0,'这是描述文字',0),(21,1,'yes','666',0,'2018-11-15 16:19:04',0,'这是描述文字',0),(22,1,'yes','666',0,'2018-11-15 16:19:04',0,'这是描述文字',0),(23,1,'yes','666',0,'2018-11-15 16:19:05',0,'这是描述文字',0),(24,1,'yes','666',0,'2018-11-15 16:45:28',0,'这是描述文字',0),(25,1,'yes','666',0,'2018-11-15 16:46:27',0,'这是描述文字',0),(26,1,'yes','666',0,'2018-11-15 16:46:33',0,'这是描述文字',0),(27,1,'yes','666',0,'2018-11-15 16:54:29',0,'这是描述文字',0),(28,1,'yes','666',0,'2018-11-15 16:56:19',0,'这是描述文字',0),(29,1,'yes','666',0,'2018-11-15 17:01:45',0,'这是描述文字',0),(30,1,'yes','666',0,'2018-11-15 17:01:51',0,'这是描述文字',0),(31,1,'yes','666',0,'2018-11-15 17:01:52',0,'这是描述文字',0),(32,1,'yes','666',0,'2018-11-15 17:07:21',0,'这是描述文字',0),(33,1,'yes','666',0,'2018-11-15 17:13:47',0,'这是描述文字',0),(34,1,'yes','666',0,'2018-11-15 17:13:50',0,'这是描述文字',0),(35,1,'yes','666',0,'2018-11-15 17:13:51',0,'这是描述文字',0),(36,1,'yes','666',0,'2018-11-15 17:13:52',0,'这是描述文字',0),(37,1,'yes','666',0,'2018-11-15 17:13:57',0,'这是描述文字',0),(38,1,'yes','666',0,'2018-11-15 17:13:59',0,'这是描述文字',0),(39,1,'yes','666',0,'2018-11-15 17:14:00',1,'这是描述文字',0),(40,1,'yes','666',0,'2018-11-15 17:14:01',1,'这是描述文字',0),(41,1,'yes','666',0,'2018-11-16 08:56:42',1,'这是描述文字',0),(42,1,'yes','666',0,'2018-11-16 08:57:56',1,'这是描述文字',0),(43,3,'yes','666',0,'2018-11-20 16:52:23',NULL,'这是描述文字11',0),(44,3,'yes','666',0,'2018-11-20 16:52:46',NULL,'这是描述文字11',0),(45,3,'yes','666',0,'2018-11-20 16:52:46',NULL,'这是描述文字11',0),(46,3,'yes','666',0,'2018-11-20 16:52:47',NULL,'这是描述文字11',0),(47,3,'yes','666',0,'2018-11-20 16:52:49',NULL,'这是描述文字11',0),(48,3,'yes','666',0,'2018-11-20 16:52:49',NULL,'这是描述文字11',0),(49,3,'yes','666',0,'2018-11-20 16:52:49',NULL,'这是描述文字11',0),(50,3,'yes','666',0,'2018-11-20 17:05:45',NULL,'这是描述文字11',0),(51,3,'yes','666',0,'2018-11-20 17:09:22',NULL,'这是描述文字11',0),(52,3,'yes','666',0,'2018-11-20 17:11:36',NULL,'这是描述文字11',0),(53,3,'yes','666',0,'2018-11-20 17:11:36',NULL,'这是描述文字11',0),(54,3,'yes','666',0,'2018-11-20 17:11:37',NULL,'这是描述文字11',0),(55,3,'yes','666',0,'2018-11-20 17:11:37',NULL,'这是描述文字11',0),(56,3,'yes','666',0,'2018-11-20 17:11:38',NULL,'这是描述文字11',0),(57,3,'yes','666',0,'2018-11-20 17:11:39',NULL,'这是描述文字11',0),(58,3,'yes','666',0,'2018-11-20 17:20:06',NULL,'这是描述文字11',0),(59,3,'yes','666',0,'2018-11-20 17:20:35',NULL,'这是描述文字11',0),(60,3,'yes','666',0,'2018-11-20 17:21:48',NULL,'这是描述文字11',0),(61,3,'yes','666',0,'2018-11-20 17:21:49',NULL,'这是描述文字11',0),(62,3,'yes','666',0,'2018-11-20 17:22:39',NULL,'这是描述文字11',0),(63,3,'yes','666',0,'2018-11-20 17:22:40',NULL,'这是描述文字11',0),(64,3,'yes','666',0,'2018-11-20 17:32:23',NULL,'这是描述文字11',0),(65,3,'yes','666',0,'2018-11-20 17:32:27',NULL,'这是描述文字11',0),(66,3,'yes','666',0,'2018-11-20 17:36:18',NULL,'这是描述文字11',0),(67,3,'yes','666',0,'2018-11-20 17:37:08',NULL,'这是描述文字11',0),(68,3,'yes','666',0,'2018-11-20 17:37:10',NULL,'这是描述文字11',0),(69,3,'yes','666',0,'2018-11-20 17:37:11',NULL,'这是描述文字11',0),(70,3,'yes','666',0,'2018-11-20 17:37:12',NULL,'这是描述文字11',0),(71,3,'yes','666',0,'2018-11-20 17:38:53',NULL,'这是描述文字11',0),(72,3,'yes','666',0,'2018-11-20 17:39:13',NULL,'这是描述文字11',0),(73,3,'yes','666',0,'2018-11-20 17:40:07',NULL,'这是描述文字11',0),(74,3,'yes','666',0,'2018-11-20 17:41:01',NULL,'这是描述文字11',0),(75,3,'yes','666',0,'2018-11-20 17:41:41',NULL,'这是描述文字11',0),(76,3,'yes','666',0,'2018-11-20 17:42:16',NULL,'这是描述文字11',0),(77,3,'yes','666',0,'2018-11-20 17:43:15',NULL,'这是描述文字11',0),(78,3,'yes','666',0,'2018-11-20 17:44:02',NULL,'这是描述文字11',0),(79,3,'yes','666',0,'2018-11-20 17:45:22',NULL,'这是描述文字11',0),(80,3,'yes','666',0,'2018-11-20 17:45:24',NULL,'这是描述文字11',0);

/*Table structure for table `file` */

DROP TABLE IF EXISTS `file`;

CREATE TABLE `file` (
  `id` int(11) NOT NULL COMMENT '附件ID',
  `name` char(100) DEFAULT NULL COMMENT '附件名称',
  `suffix` char(10) DEFAULT NULL COMMENT '附件后缀',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上传时间',
  `path` char(250) DEFAULT NULL COMMENT '附件存放位置',
  `filename` char(100) DEFAULT NULL COMMENT '存放名称',
  `delete` tinyint(1) DEFAULT NULL COMMENT '是否删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `file` */

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
  `user_code` int(11) NOT NULL COMMENT '员工号',
  `password` char(50) DEFAULT NULL COMMENT '登录密码',
  `is_use` enum('1','0') DEFAULT '1' COMMENT '是否禁用（1启用0禁用）',
  `change_pd_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改密码时间',
  `login_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `author_id` int(11) DEFAULT NULL COMMENT '授权人',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`user_code`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `login` */

/*Table structure for table `logs` */

DROP TABLE IF EXISTS `logs`;

CREATE TABLE `logs` (
  `id` int(11) DEFAULT NULL COMMENT '日志ID',
  `type` char(50) DEFAULT NULL COMMENT '日志类型',
  `user_id` int(11) NOT NULL COMMENT '操作人ID',
  `msg` varchar(5000) DEFAULT NULL COMMENT '操作说明',
  `data` varchar(5000) DEFAULT NULL COMMENT '参数',
  `time` timestamp NULL DEFAULT NULL COMMENT '操作时间',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `logs` */

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
  `id` int(11) DEFAULT NULL COMMENT '公告ID',
  `department_id` int(11) NOT NULL COMMENT '机构ID',
  `title` char(100) DEFAULT NULL COMMENT '标题',
  `content` longtext COMMENT '公告内容',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '发布时间',
  `range` enum('1','2') DEFAULT '1' COMMENT '通知范围（1所有部门，2指定部门）',
  `inform_id` char(1) DEFAULT NULL COMMENT '指定部门时的部门ids',
  `state` enum('draft','publish') DEFAULT 'draft' COMMENT '状态（draft草稿publish发布）',
  `delete` int(11) DEFAULT '0' COMMENT '删除（0未删1已删）',
  PRIMARY KEY (`department_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `notice` */

/*Table structure for table `notice_file` */

DROP TABLE IF EXISTS `notice_file`;

CREATE TABLE `notice_file` (
  `file_id` int(11) NOT NULL COMMENT '通知附件ID',
  `notice_id` int(11) NOT NULL COMMENT '通知ID',
  PRIMARY KEY (`file_id`,`notice_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `notice_file` */

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
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '人员ID',
  `department_id` int(11) NOT NULL COMMENT '部门ID',
  `user_name` char(50) DEFAULT NULL COMMENT '姓名',
  `user_code` char(20) NOT NULL COMMENT '员工号',
  `class` char(20) DEFAULT NULL COMMENT '民族',
  `sex` enum('0','1','2') NOT NULL DEFAULT '0' COMMENT '性别（0保密1女2男）',
  `id_card` char(18) DEFAULT NULL COMMENT '身份证',
  `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `delete` tinyint(1) DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`,`department_id`,`user_code`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Data for the table `users` */

insert  into `users`(`id`,`department_id`,`user_name`,`user_code`,`class`,`sex`,`id_card`,`update_time`,`delete`) values (1,0,'小明','',NULL,'0',NULL,'0000-00-00 00:00:00',0);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
