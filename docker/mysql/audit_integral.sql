/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.7.24-log : Database - audit_integral
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

/*Table structure for table `audit_notice` */

DROP TABLE IF EXISTS `audit_notice`;

CREATE TABLE `audit_notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '审计通知书ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿ID',
  `year` year(4) DEFAULT NULL COMMENT '年份',
  `number` char(100) DEFAULT NULL COMMENT '编号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `audit_notice` */

/*Table structure for table `audit_report` */

DROP TABLE IF EXISTS `audit_report`;

CREATE TABLE `audit_report` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '审计报告ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '方案ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿ID',
  `confirmation_id` int(11) NOT NULL DEFAULT '0' COMMENT '事实确认书ID',
  `rectify_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '整改报告ID',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|publish发布）',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `audit_report` */

/*Table structure for table `audit_report_basic_info` */

DROP TABLE IF EXISTS `audit_report_basic_info`;

CREATE TABLE `audit_report_basic_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '审计报告基本信息ID',
  `audit_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '审计报告ID',
  `content` varchar(2500) DEFAULT NULL COMMENT '基本信息',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `audit_report_basic_info` */

/*Table structure for table `audit_report_plan` */

DROP TABLE IF EXISTS `audit_report_plan`;

CREATE TABLE `audit_report_plan` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '审计报告基本信息ID',
  `audit_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '审计报告ID',
  `content` varchar(2500) DEFAULT NULL COMMENT '基本信息',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `audit_report_plan` */

/*Table structure for table `audit_report_reason` */

DROP TABLE IF EXISTS `audit_report_reason`;

CREATE TABLE `audit_report_reason` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '审计报告基本信息ID',
  `audit_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '审计报告ID',
  `content` varchar(2500) DEFAULT NULL COMMENT '基本信息',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `audit_report_reason` */

/*Table structure for table `clause` */

DROP TABLE IF EXISTS `clause`;

CREATE TABLE `clause` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '管理办法及条款ID',
  `department_id` int(11) NOT NULL DEFAULT '-1' COMMENT '所属部门ID（为-1时所有部门通用）',
  `title` char(200) DEFAULT NULL COMMENT '管理办法及条款标题',
  `from` char(200) DEFAULT NULL COMMENT '来源',
  `type` char(80) DEFAULT NULL COMMENT '分类（字典）',
  `author_id` int(11) DEFAULT NULL COMMENT '发布人ID',
  `number` char(100) DEFAULT NULL COMMENT '文件号',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿publish发布）',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `clause` */

/*Table structure for table `clause_content` */

DROP TABLE IF EXISTS `clause_content`;

CREATE TABLE `clause_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '管理办法内容ID',
  `clause_id` int(11) DEFAULT NULL COMMENT '管理办法ID',
  `is_title` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为标题（0内容1标题）',
  `title_level` char(80) NOT NULL DEFAULT '-' COMMENT '标题级别',
  `content` varchar(2000) DEFAULT NULL COMMENT '内容',
  `order` int(4) NOT NULL DEFAULT '0' COMMENT '顺序',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `clause_content` */

/*Table structure for table `clause_file` */

DROP TABLE IF EXISTS `clause_file`;

CREATE TABLE `clause_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `clause_id` int(11) DEFAULT NULL COMMENT '管理办法ID',
  `file_id` int(11) DEFAULT NULL COMMENT '附件ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;

/*Data for the table `clause_file` */

/*Table structure for table `confirmation` */

DROP TABLE IF EXISTS `confirmation`;

CREATE TABLE `confirmation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '确认书ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿ID',
  `confirmation_receipt_id` int(11) NOT NULL DEFAULT '0' COMMENT '确认书回执ID',
  `file_id` int(11) DEFAULT '0' COMMENT '图片ID',
  `has_read` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已读（0未读|1已读）',
  `has_read_time` timestamp NULL DEFAULT NULL COMMENT '已读时间',
  `proposal` varchar(2000) DEFAULT NULL COMMENT '单位意见或者建议',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|publish发布）',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `year` year(4) DEFAULT NULL COMMENT '年份',
  `number` int(11) DEFAULT NULL COMMENT '编号',
  `author_id` int(11) DEFAULT NULL COMMENT '创建人id',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`),
  KEY `工作底稿` (`draft_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `confirmation` */

/*Table structure for table `confirmation_basis` */

DROP TABLE IF EXISTS `confirmation_basis`;

CREATE TABLE `confirmation_basis` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '依据主键ID',
  `confirmation_id` int(11) NOT NULL DEFAULT '0' COMMENT '事实确认书ID',
  `basis_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿依据ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `confirmation_basis` */

/*Table structure for table `confirmation_content` */

DROP TABLE IF EXISTS `confirmation_content`;

CREATE TABLE `confirmation_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '发现问题内容ID',
  `confirmation_id` int(11) DEFAULT NULL COMMENT '事实确认书ID',
  `order` int(4) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `type` enum('title','behavior','other') NOT NULL DEFAULT 'other' COMMENT '内容类型（behavior违规行为|title标题|other其他）',
  `behavior_id` int(11) NOT NULL DEFAULT '0' COMMENT '违规行为ID',
  `behavior_content` varchar(2000) DEFAULT NULL COMMENT '违规行为描述或者其他',
  `delete` int(11) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `confirmation_content` */

/*Table structure for table `confirmation_receipt` */

DROP TABLE IF EXISTS `confirmation_receipt`;

CREATE TABLE `confirmation_receipt` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '确认书回执ID',
  `content` varchar(5000) DEFAULT NULL COMMENT '回复内容',
  `time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '回执时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Data for the table `confirmation_receipt` */

/*Table structure for table `confirmation_receipt_content` */

DROP TABLE IF EXISTS `confirmation_receipt_content`;

CREATE TABLE `confirmation_receipt_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改项ID',
  `confirmation_receipt` int(11) DEFAULT NULL COMMENT '确认书回执ID',
  `behavior_id` int(11) DEFAULT NULL COMMENT '违规行为ID',
  `result` varchar(500) DEFAULT NULL COMMENT '整改结果',
  `time` timestamp NULL DEFAULT NULL COMMENT '整改时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `confirmation_receipt_content` */

/*Table structure for table `confirmation_user` */

DROP TABLE IF EXISTS `confirmation_user`;

CREATE TABLE `confirmation_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `confirmation_id` int(11) DEFAULT NULL COMMENT '事实确认书ID',
  `user_id` int(11) DEFAULT NULL COMMENT '违规人员ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `confirmation_user` */

/*Table structure for table `department` */

DROP TABLE IF EXISTS `department`;

CREATE TABLE `department` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `name` char(50) DEFAULT NULL COMMENT '部门名称',
  `parent_id` int(11) DEFAULT NULL COMMENT '上级部门ID',
  `code` char(150) DEFAULT NULL COMMENT '部门编码',
  `has_child` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有子部门（0没有1有）',
  `level` char(20) DEFAULT NULL COMMENT '部门级别',
  `type` char(80) DEFAULT NULL COMMENT '类型（字典key）',
  `grade` int(3) DEFAULT NULL COMMENT '所在部门树的层级',
  `address` char(250) DEFAULT NULL COMMENT '地址',
  `phone` char(11) DEFAULT NULL COMMENT '手机号或者电话',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete` tinyint(1) DEFAULT '0' COMMENT '是否删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

/*Data for the table `department` */

insert  into `department`(`id`,`name`,`parent_id`,`code`,`has_child`,`level`,`type`,`grade`,`address`,`phone`,`update_time`,`delete`) values (1,'农商行',-1,'bank',1,'0',NULL,NULL,'','','2019-01-15 14:22:31',0),(2,'小微企业金融服务中心',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:12:07',0),(3,'业务授权中心',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:12:14',0),(4,'工会工作部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:12:21',0),(5,'事后监督部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:12:26',0),(6,'农村业务部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:12:31',0),(7,'电子银行部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:12:36',0),(8,'资金运营部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:12:43',0),(9,'后勤服务部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:12:49',0),(10,'安全保卫部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:12:55',0),(11,'稽核审计部',1,'',0,'0',NULL,NULL,'','','2019-01-08 10:02:22',0),(12,'财务统计部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:13:08',0),(13,'合规风险管理部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:13:13',0),(14,'业务发展部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:13:21',0),(15,'督导室',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:13:27',0),(16,'纪检监察部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:13:35',0),(17,'党群工作部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:13:40',0),(18,'人力资源部',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:13:46',0),(19,'办公室',1,'',0,'0',NULL,NULL,'','','2019-01-08 09:55:35',0),(20,'董事会办公室',1,'',0,'0',NULL,NULL,'','','2019-01-06 22:13:59',0);

/*Table structure for table `department_user` */

DROP TABLE IF EXISTS `department_user`;

CREATE TABLE `department_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `department_id` int(11) DEFAULT NULL COMMENT '部门id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '人员id',
  `type` char(80) DEFAULT NULL COMMENT '角色类型',
  `delete` tinyint(1) DEFAULT '0' COMMENT '软删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Data for the table `department_user` */

insert  into `department_user`(`id`,`department_id`,`user_id`,`type`,`delete`) values (1,1,-1,'admin',0);

/*Table structure for table `dictionary` */

DROP TABLE IF EXISTS `dictionary`;

CREATE TABLE `dictionary` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '字典ID',
  `type_id` int(11) DEFAULT NULL COMMENT '字典类型ID',
  `key` char(80) DEFAULT NULL COMMENT '字典值',
  `value` char(100) DEFAULT NULL COMMENT '字典名称',
  `order` int(4) DEFAULT '0' COMMENT '排序顺序',
  `describe` char(250) DEFAULT NULL COMMENT '字典描述',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=199 DEFAULT CHARSET=utf8;

/*Data for the table `dictionary` */

insert  into `dictionary`(`id`,`type_id`,`key`,`value`,`order`,`describe`,`delete`) values (-1002,-10,'other','其他文件',2,'',0),(-1001,-10,'managementMeasures','管理办法',1,'管理办法分类',0),(-405,-4,'h5','五级标题',5,'系统分类（不允许删除）',0),(-404,-4,'h4','四级标题',4,'系统分类（不允许删除）',0),(-403,-4,'h3','三级标题',3,'系统分类（不允许删除）',0),(-402,-4,'h2','二级标题',2,'系统分类（不允许删除）',0),(-401,-4,'h1','一级标题',1,'系统分类（不允许删除）',0),(-303,-3,'address','贵州省',3,'单位地址',0),(-302,-3,'name','兴仁农村商业银行',1,'单位简称',0),(-301,-3,'com','贵州兴仁农村商业银行股份有限公司',2,'单位名称',0),(-221,-2,'fgld','分管领导',21,'分管领导',0),(-220,-2,'zcn','总出纳',20,'总出纳',0),(-219,-2,'zrzl','主任助理',19,'主任助理',0),(-218,-2,'ghzx','工会主席',18,'工会主席',0),(-217,-2,'fxzj','风险总监',17,'风险总监',0),(-216,-2,'jsz','监事长',3,'监事长',0),(-215,-2,'lsz','理事长',5,'理事长',0),(-214,-2,'tw','团委',16,'团委',0),(-213,-2,'zr','主任',6,'主任',0),(-212,-2,'fjl','副经理',15,'副经理',0),(-211,-2,'xdnq','信贷内勤',14,'信贷内勤',0),(-210,-2,'gzy','工作员',13,'工作员',0),(-209,-2,'zdy','指导员',12,'指导员',0),(-208,-2,'jl','经理',11,'经理',0),(-207,-2,'xdy','信贷员',10,'信贷员',0),(-206,-2,'zhgy','综合柜员',9,'综合柜员',0),(-205,-2,'fzr','副主任',8,'副主任',0),(-204,-2,'zbkj','主办会计',7,'主办会计',0),(-203,-2,'staff','业务员',4,'业务员',0),(-202,-2,'management','部门负责人',2,'部门负责人',0),(-201,-2,'admin','超级管理员',1,'超级管理员',0),(-102,-1,'other','其他',0,'系统分类（不允许删除）',0),(-101,-1,'system','系统',0,'系统分类（不允许删除）',0),(184,-5,'auditKey1','方案类型1',1,'',0),(185,-5,'auditKey2','方案类型2',2,'',0),(186,-5,'auditKey3','方案类型3',3,'',0),(187,-6,'auditType1','审计方式1',1,'',0),(188,-6,'auditType2','审计方式2',2,'',0),(189,-6,'auditType3','审计方式3',3,'',0),(190,-8,'userJob1','人员职务1',1,'',0),(191,-8,'userJob2','人员职务2',2,'',0),(192,-8,'userJob3','人员职务3',3,'',0),(193,-7,'userTitle1','人员技术职称1',1,'',0),(194,-7,'userTitle2','人员技术职称2',2,'',0),(195,-7,'userTitle3','人员技术职称3',3,'',0),(196,-9,'auditTask1','员工分工1',1,'',0),(197,-9,'auditTask2','员工分工2',2,'',0),(198,-9,'auditTask3','员工分工3',3,'',0);

/*Table structure for table `dictionary_type` */

DROP TABLE IF EXISTS `dictionary_type`;

CREATE TABLE `dictionary_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '字典类型ID',
  `type_id` int(11) NOT NULL COMMENT '字典分类ID',
  `key` char(80) NOT NULL COMMENT '字典分类键',
  `title` char(50) NOT NULL COMMENT '字典类型名称',
  `is_use` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `user_id` int(11) DEFAULT NULL COMMENT '创建人ID',
  `describe` char(250) DEFAULT NULL COMMENT '字典类型描述',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '软删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `dictionary_type` */

insert  into `dictionary_type`(`id`,`type_id`,`key`,`title`,`is_use`,`update_time`,`user_id`,`describe`,`delete`) values (-10,-1,'system','管理办法分类',1,'2019-01-12 16:25:09',0,'机构管理下的文件分类',0),(-9,-1,'system','员工分工',1,'2018-12-19 16:51:04',0,'稽核员工分工',0),(-8,-1,'system','人员技术职称',1,'2018-12-19 16:49:39',0,'稽核人员技术职称',0),(-7,-1,'system','人员职务',1,'2018-12-19 16:48:09',0,'稽核人员职务',0),(-6,-1,'system','审计方式',1,'2018-12-19 18:07:13',0,'审计方式',0),(-5,-1,'system','方案类型',1,'2018-12-19 18:06:38',0,'方案类型',0),(-4,-1,'system','文件正文标题级别',0,'2019-01-12 17:30:06',0,'文件正文标题级别选项',0),(-3,-1,'system','单位信息',1,'2019-01-13 11:23:42',0,'系统单位信息',0),(-2,-1,'system','人员角色',1,'2019-01-14 14:39:21',0,'部门人员角色字典',0),(-1,-1,'system','字典分类',1,'2018-11-12 15:37:43',0,'系统字典分类',0);

/*Table structure for table `draft` */

DROP TABLE IF EXISTS `draft`;

CREATE TABLE `draft` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '工作底稿ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '方案ID',
  `query_department_id` int(11) NOT NULL DEFAULT '0' COMMENT '被查询机构ID',
  `department_id` int(11) NOT NULL DEFAULT '0' COMMENT '检查机构ID',
  `number` int(11) DEFAULT NULL COMMENT '编号',
  `year` year(4) DEFAULT NULL COMMENT '年份',
  `project_name` char(250) DEFAULT NULL COMMENT '项目名称',
  `public` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否公开（0私有|1公开）【被检查人是否能查看】',
  `query_start_time` date DEFAULT NULL COMMENT '检查开始日期',
  `query_end_time` date DEFAULT NULL COMMENT '检查结束日期',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '底稿状态（draft草稿|publish发布）',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建人id',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `draft` */

/*Table structure for table `draft_admin_user` */

DROP TABLE IF EXISTS `draft_admin_user`;

CREATE TABLE `draft_admin_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '复查ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '底稿ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '复查人员ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `draft_admin_user` */

/*Table structure for table `draft_content` */

DROP TABLE IF EXISTS `draft_content`;

CREATE TABLE `draft_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '底稿内容ID',
  `draft_id` int(11) DEFAULT NULL COMMENT '底稿ID',
  `order` int(4) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `type` enum('title','behavior','other') NOT NULL DEFAULT 'other' COMMENT '内容类型（behavior违规行为|title标题|other其他）',
  `behavior_id` int(11) NOT NULL DEFAULT '0' COMMENT '违规行为ID',
  `behavior_content` varchar(2000) DEFAULT NULL COMMENT '违规行为描述或者其他',
  `delete` int(11) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `draft_content` */

/*Table structure for table `draft_file` */

DROP TABLE IF EXISTS `draft_file`;

CREATE TABLE `draft_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `draft_id` int(11) DEFAULT NULL COMMENT '底稿ID',
  `file_id` int(11) DEFAULT NULL COMMENT '附件ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`),
  KEY `工作底稿1` (`draft_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `draft_file` */

/*Table structure for table `draft_inspect_user` */

DROP TABLE IF EXISTS `draft_inspect_user`;

CREATE TABLE `draft_inspect_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `draft_id` int(11) DEFAULT NULL COMMENT '底稿ID',
  `user_id` int(11) DEFAULT NULL COMMENT '被检查人员ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`),
  KEY `draft_id` (`draft_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `draft_inspect_user` */

/*Table structure for table `draft_query_user` */

DROP TABLE IF EXISTS `draft_query_user`;

CREATE TABLE `draft_query_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `draft_id` int(11) DEFAULT NULL COMMENT '底稿ID',
  `user_id` int(11) DEFAULT NULL COMMENT '检查人员ID',
  `is_leader` tinyint(1) NOT NULL DEFAULT '0' COMMENT '组长（0否|1是）',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `draft_query_user` */

/*Table structure for table `draft_review_user` */

DROP TABLE IF EXISTS `draft_review_user`;

CREATE TABLE `draft_review_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `draft_id` int(11) DEFAULT NULL COMMENT '底稿ID',
  `user_id` int(11) DEFAULT NULL COMMENT '负责人员ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `draft_review_user` */

/*Table structure for table `files` */

DROP TABLE IF EXISTS `files`;

CREATE TABLE `files` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '附件ID',
  `name` char(200) DEFAULT NULL COMMENT '附件名称',
  `suffix` char(10) DEFAULT NULL COMMENT '附件后缀',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上传时间',
  `path` char(250) DEFAULT NULL COMMENT '附件存放位置',
  `size` int(10) DEFAULT '0' COMMENT '文件大小（单位B）',
  `file_name` char(200) DEFAULT NULL COMMENT '存放名称',
  `form_id` int(11) NOT NULL DEFAULT '0' COMMENT '外部ID',
  `form` char(50) NOT NULL DEFAULT '-' COMMENT '来自那张表',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `files` */

/*Table structure for table `integral` */

DROP TABLE IF EXISTS `integral`;

CREATE TABLE `integral` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '积分ID',
  `cognizance_user_id` int(11) NOT NULL DEFAULT '0' COMMENT '认定人ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '责任人ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿ID',
  `punish_notice_id` int(11) NOT NULL DEFAULT '0' COMMENT '处罚通知ID',
  `receipt_id` int(11) NOT NULL DEFAULT '0' COMMENT '回执ID',
  `score` int(5) NOT NULL DEFAULT '0' COMMENT '分值（除以1000显示）',
  `money` int(10) NOT NULL DEFAULT '0' COMMENT '罚款（除以1000显示）',
  `time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '日期',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `integral` */

/*Table structure for table `integral_edit` */

DROP TABLE IF EXISTS `integral_edit`;

CREATE TABLE `integral_edit` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '积分修改ID',
  `integral_id` int(11) NOT NULL DEFAULT '0' COMMENT '积分表ID',
  `money` int(11) NOT NULL DEFAULT '0' COMMENT '罚款（除以1000显示）',
  `score` int(5) NOT NULL DEFAULT '0' COMMENT '积分（除以1000显示）',
  `user_id` int(11) DEFAULT NULL COMMENT '发起人修改ID',
  `describe` varchar(500) DEFAULT NULL COMMENT '修改分数原因',
  `suggestion` char(250) DEFAULT NULL COMMENT '审核意见',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `state` enum('draft','report','reject','adopt') DEFAULT 'draft' COMMENT '状态（draft草稿|report上报|reject驳回|adopt通过）',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `integral_edit` */

/*Table structure for table `introduction` */

DROP TABLE IF EXISTS `introduction`;

CREATE TABLE `introduction` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '查库介绍信id',
  `draft_id` int(11) DEFAULT NULL COMMENT '工作底稿id',
  `number` int(11) DEFAULT NULL COMMENT '编号',
  `year` year(4) DEFAULT NULL COMMENT '年份',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `introduction` */

/*Table structure for table `login` */

DROP TABLE IF EXISTS `login`;

CREATE TABLE `login` (
  `login_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '登录id',
  `user_code` char(30) NOT NULL DEFAULT '-' COMMENT '员工号',
  `password` char(50) DEFAULT NULL COMMENT '登录密码',
  `is_use` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否禁用（1启用0禁用）',
  `change_pd_time` timestamp NULL DEFAULT NULL COMMENT '最后修改密码时间',
  `login_num` int(6) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `author_id` int(11) DEFAULT NULL COMMENT '授权人',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`login_id`),
  KEY `login_id` (`login_id`)
) ENGINE=MyISAM AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

/*Data for the table `login` */

insert  into `login`(`login_id`,`user_code`,`password`,`is_use`,`change_pd_time`,`login_num`,`login_time`,`author_id`,`delete`) values (-1,'admin','95efad0975e613c0609fcbeb1b23cb5d',1,NULL,16,'2019-01-15 14:27:58',-1,0);

/*Table structure for table `logs` */

DROP TABLE IF EXISTS `logs`;

CREATE TABLE `logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `url` char(200) DEFAULT NULL COMMENT 'URL',
  `user_id` int(11) DEFAULT NULL COMMENT '操作人ID',
  `msg` char(100) DEFAULT NULL COMMENT '操作说明',
  `method` char(10) DEFAULT NULL COMMENT '请求方法',
  `data` varchar(5000) DEFAULT NULL COMMENT '参数',
  `time` timestamp NULL DEFAULT NULL COMMENT '操作时间',
  `server` char(20) DEFAULT NULL COMMENT '服务',
  `ip` char(15) DEFAULT NULL COMMENT 'IP地址',
  `delete` tinyint(1) DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=77538 DEFAULT CHARSET=utf8;

/*Data for the table `logs` */

/*Table structure for table `menu` */

DROP TABLE IF EXISTS `menu`;

CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `path` char(250) DEFAULT NULL COMMENT '路径',
  `name` char(250) DEFAULT NULL COMMENT '路由名称',
  `title` char(250) DEFAULT NULL COMMENT '名称',
  `icon` char(50) DEFAULT NULL COMMENT '图标',
  `no_cache` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否缓存',
  `parent_id` int(11) NOT NULL DEFAULT '-1' COMMENT '父级菜单',
  `has_child` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有子菜单',
  `order` int(2) NOT NULL DEFAULT '0' COMMENT '排序',
  `is_use` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否启用（0禁用1启用）',
  `time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;

/*Data for the table `menu` */

insert  into `menu`(`id`,`path`,`name`,`title`,`icon`,`no_cache`,`parent_id`,`has_child`,`order`,`is_use`,`time`,`delete`) values (1,'','','','',1,-1,1,1,1,'2018-12-06 09:15:55',0),(2,'dashboard','Dashboard','dashboard','dashboard',1,1,0,1,1,'2019-01-06 16:37:51',0),(3,'/guide','','','',1,-1,1,2,1,'2018-12-06 09:17:06',0),(4,'index','Guide','guide','guide',1,3,0,1,1,'2019-01-06 16:54:08',0),(5,'/personal','','','',1,-1,1,3,1,'2018-12-06 09:18:09',0),(6,'index','Personal','personal','user',1,5,0,0,1,'2019-01-04 15:34:18',0),(7,'/organization','organization','organization','component',1,-1,1,4,1,'2018-12-06 09:20:43',0),(8,'notice','notice','notice','',1,7,0,1,1,'2019-01-06 16:41:06',0),(9,'departmentManagement','departmentManagement','departmentManagement','',1,7,0,2,1,'2018-12-06 09:21:31',0),(10,'personnelManagement','personnelManagement','personnelManagement','',1,7,0,3,1,'2019-01-05 10:26:15',0),(11,'managementMethods','managementMethods','managementMethods','',1,7,0,4,1,'2018-12-06 09:21:47',0),(12,'/audit','audit','audit','international',1,-1,1,5,1,'2018-12-06 09:22:10',0),(13,'workManuscript','workManuscript','workManuscript','',1,12,0,2,1,'2019-01-04 15:30:54',0),(14,'confirmation','confirmation','confirmation','',1,12,0,6,1,'2019-01-14 15:05:49',0),(15,'punishNotice','punishNotice','punishNotice','',1,12,0,7,1,'2019-01-14 15:05:50',0),(16,'integralTable','integralTable','integralTable','',1,12,0,8,1,'2019-01-14 15:05:51',0),(17,'statisticalAnalysis','statisticalAnalysis','statisticalAnalysis','',1,12,0,13,1,'2019-01-14 15:05:59',0),(18,'/system','system','system','example',1,-1,1,6,1,'2018-12-06 09:23:47',0),(19,'dictionaryManagement','dictionaryManagement','dictionaryManagement','',1,18,0,1,1,'2018-12-06 09:25:58',0),(20,'loginManagement','loginManagement','loginManagement','',1,18,0,2,1,'2018-12-06 09:26:06',0),(21,'menusManagement','menusManagement','menusManagement','',1,18,0,3,1,'2018-12-06 09:26:19',0),(22,'powerManagement','powerManagement','powerManagement','',1,18,0,4,1,'2018-12-06 09:26:27',0),(23,'systemLog','systemLog','systemLog','',1,18,0,5,1,'2018-12-06 09:26:33',0),(24,'auditPlan','auditPlan','auditPlan','',1,12,0,1,1,'2019-01-04 15:34:12',0),(25,'rectifyNotice','rectifyNotice','rectifyNotice','',1,12,0,9,1,'2019-01-14 15:05:52',0),(26,'rectifyReport','rectifyReport','rectifyReport','',1,12,0,10,1,'2019-01-14 15:05:54',0),(27,'auditReport','auditReport','auditReport','',1,12,0,11,1,'2019-01-14 15:05:57',0),(28,'introduction','introduction','introduction','',1,12,0,3,1,'2019-01-14 15:05:30',0),(29,'track','track','track','',1,12,0,12,1,'2019-01-14 15:05:58',0),(30,'auditNotice','auditNotice','auditNotice','',1,12,0,4,1,'2019-01-14 15:07:01',0);

/*Table structure for table `notice` */

DROP TABLE IF EXISTS `notice`;

CREATE TABLE `notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `department_id` int(11) NOT NULL COMMENT '机构ID',
  `title` char(200) DEFAULT NULL COMMENT '标题',
  `content` longtext COMMENT '公告内容',
  `time` timestamp NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '发布时间',
  `range` enum('1','2') DEFAULT '1' COMMENT '通知范围（1所有部门，2指定部门）',
  `author_id` int(11) NOT NULL DEFAULT '0' COMMENT '作者ID',
  `state` enum('draft','publish') DEFAULT 'draft' COMMENT '状态（draft草稿publish发布）',
  `delete` int(11) DEFAULT '0' COMMENT '删除（0未删1已删）',
  PRIMARY KEY (`id`,`department_id`)
) ENGINE=MyISAM AUTO_INCREMENT=222 DEFAULT CHARSET=utf8;

/*Data for the table `notice` */

/*Table structure for table `notice_file` */

DROP TABLE IF EXISTS `notice_file`;

CREATE TABLE `notice_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `file_id` int(11) DEFAULT NULL COMMENT '通知附件ID',
  `notice_id` int(11) NOT NULL COMMENT '通知ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`,`notice_id`)
) ENGINE=MyISAM AUTO_INCREMENT=56 DEFAULT CHARSET=utf8;

/*Data for the table `notice_file` */

/*Table structure for table `notice_inform` */

DROP TABLE IF EXISTS `notice_inform`;

CREATE TABLE `notice_inform` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `department_id` int(11) NOT NULL COMMENT '机构ID',
  `notice_id` int(11) NOT NULL COMMENT '公告ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`,`notice_id`)
) ENGINE=MyISAM AUTO_INCREMENT=527 DEFAULT CHARSET=utf8;

/*Data for the table `notice_inform` */

/*Table structure for table `programme` */

DROP TABLE IF EXISTS `programme`;

CREATE TABLE `programme` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '方案ID',
  `title` char(200) DEFAULT NULL COMMENT '方案名称',
  `key` char(80) DEFAULT NULL COMMENT '方案分类',
  `query_department_id` int(11) NOT NULL DEFAULT '0' COMMENT '被稽核审计机构ID',
  `user_id` int(11) DEFAULT '0' COMMENT '创建方案的人员ID',
  `query_point_id` int(11) NOT NULL DEFAULT '0' COMMENT '被稽核审计网点ID',
  `purpose` varchar(500) DEFAULT NULL COMMENT '稽核目的',
  `type` char(80) DEFAULT NULL COMMENT '稽核审计方式',
  `start_time` date NOT NULL COMMENT '稽核审计开始时间',
  `end_time` date NOT NULL COMMENT '稽核审计结束时间',
  `plan_start_time` date DEFAULT NULL COMMENT '稽核业务开始时间',
  `plan_end_time` date DEFAULT NULL COMMENT '稽核业务结束时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `number` int(11) DEFAULT NULL COMMENT '编号',
  `year` year(4) DEFAULT NULL COMMENT '年份',
  `state` enum('draft','report','dep_reject','dep_adopt','admin_reject','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|report上报|publish发布|adopt通过|reject驳回【dep_部门负责人前缀|admin_分管领导前缀】）',
  `author_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建人id',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `programme` */

/*Table structure for table `programme_basis` */

DROP TABLE IF EXISTS `programme_basis`;

CREATE TABLE `programme_basis` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '实施稽核的依据ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '稽核方案ID',
  `clause_id` int(11) NOT NULL DEFAULT '0' COMMENT '管理办法ID',
  `content` char(250) DEFAULT NULL COMMENT '依据内容',
  `order` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `programme_basis` */

/*Table structure for table `programme_business` */

DROP TABLE IF EXISTS `programme_business`;

CREATE TABLE `programme_business` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '工作方案业务范围ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '稽核方案ID',
  `content` char(250) DEFAULT NULL COMMENT '内容',
  `order` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `programme_business` */

/*Table structure for table `programme_content` */

DROP TABLE IF EXISTS `programme_content`;

CREATE TABLE `programme_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '工作方案主要内容ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '稽核方案ID',
  `content` varchar(500) DEFAULT NULL COMMENT '主要内容',
  `order` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未除|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=168 DEFAULT CHARSET=utf8;

/*Data for the table `programme_content` */

insert  into `programme_content`(`id`,`programme_id`,`content`,`order`,`delete`) values (129,81,'这是方案主要内容1',1,0),(130,81,'这是方案主要内容2',2,0),(131,82,'这是方案主要内容1',1,0),(132,82,'这是方案主要内容2',2,0),(134,88,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。',1,0),(135,90,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。',1,0),(136,91,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。',1,0),(137,92,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。',1,0),(138,93,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。 ',1,0),(139,94,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。',1,0),(140,95,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。',1,0),(141,96,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。',1,0),(142,99,'qwe',1,0),(143,100,'达到',1,0),(144,101,'胜多负少的GV',1,1),(145,102,'大师as',1,1),(146,103,'大师as',1,1),(147,104,'胜多负少的GV',1,1),(148,105,'达到',1,0),(149,106,'达到',1,0),(150,107,'达到',1,1),(151,108,'达到',1,1),(152,109,'达到',1,1),(153,110,'达到',1,1),(154,111,'qwe',1,1),(155,112,'qwe',1,1),(156,113,'qwe',1,1),(157,114,'qwe',1,1),(158,115,'qwe',1,1),(159,116,'qwe',1,1),(160,117,'qwe',1,1),(161,118,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。',1,0),(162,119,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。1',1,0),(163,119,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。2',2,0),(164,120,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。1',1,0),(165,120,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。2',2,0),(166,121,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。1',1,0),(167,121,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。2',2,0);

/*Table structure for table `programme_emphases` */

DROP TABLE IF EXISTS `programme_emphases`;

CREATE TABLE `programme_emphases` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '工作方案重点ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '稽核方案ID',
  `content` char(250) DEFAULT NULL COMMENT '内容',
  `order` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=126 DEFAULT CHARSET=utf8;

/*Data for the table `programme_emphases` */

insert  into `programme_emphases`(`id`,`programme_id`,`content`,`order`,`delete`) values (87,81,'方案重点1',1,0),(88,81,'方案重点2',2,0),(89,82,'方案重点1',1,0),(90,82,'方案重点2',2,0),(91,88,'即本次审计重要性标准为：人民币9000万元。采用就地审计的各分支行重要性标准根据准则，要求以上述重要性标准1/3-1/6把握，采用抽查与送达审计相结合的分支行重要性标准不得高于1500万元，全部调整事项应予汇总后编制未调整不符事项汇总表。对于采用抽查与送达审计相结合的分支行最终误差超过1500万元的要开展进一步审计程度，进行详细审计。',1,0),(93,90,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。',1,0),(94,91,' 本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。',1,0),(95,92,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。',1,0),(96,93,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。',1,0),(97,94,'\n本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。',1,0),(98,95,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。',1,0),(99,96,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。',1,0),(100,99,'qwe',1,0),(101,100,'设定',1,0),(102,101,'防守打法好吧',1,1),(103,102,'阿萨斯的a',1,1),(104,103,'阿萨斯的a',1,1),(105,104,'防守打法好吧',1,1),(106,105,'设定',1,0),(107,106,'设定',1,0),(108,107,'设定',1,1),(109,108,'设定',1,1),(110,109,'设定',1,1),(111,110,'设定',1,1),(112,111,'qwe',1,1),(113,112,'qwe',1,1),(114,113,'qwe',1,1),(115,114,'qwe',1,1),(116,115,'qwe',1,1),(117,116,'qwe',1,1),(118,117,'qwe',1,1),(119,118,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。',1,0),(120,119,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。1',1,0),(121,119,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。2',2,0),(122,120,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。1',1,0),(123,120,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。2',2,0),(124,121,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。1',1,0),(125,121,'本次审计根据以往我们对次级信贷业务内部控制了解及其对次级信贷业务的初步审阅，我们认为本次审计相对于上一年而言,信贷余额没有发生重大变化。因此，本次审计的重要性标准依据次级信贷的0.5%计算约为0.9亿元（即9000万元）作为本次审计的重要性标准,本次所有单独或汇总起来后对本行信贷业务产生重大影响的调整或重分类不超过上述标准。2',2,0);

/*Table structure for table `programme_examine_admin` */

DROP TABLE IF EXISTS `programme_examine_admin`;

CREATE TABLE `programme_examine_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '部门负责人审核ID',
  `programme_id` int(11) DEFAULT NULL COMMENT '方案ID',
  `user_id` int(11) DEFAULT NULL COMMENT '审核人ID',
  `content` varchar(500) DEFAULT NULL COMMENT '审核意见内容',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '审核时间',
  `state` char(20) NOT NULL DEFAULT 'reject' COMMENT '审核状态（reject驳回|adopt通过）',
  `file_id` int(11) DEFAULT NULL COMMENT '签名附件ID(未启用)',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `programme_examine_admin` */

/*Table structure for table `programme_examine_dep` */

DROP TABLE IF EXISTS `programme_examine_dep`;

CREATE TABLE `programme_examine_dep` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '部门负责人审核ID',
  `programme_id` int(11) DEFAULT NULL COMMENT '方案ID',
  `user_id` int(11) DEFAULT NULL COMMENT '审核人ID',
  `content` varchar(500) DEFAULT NULL COMMENT '审核意见内容',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '审核时间',
  `state` char(20) NOT NULL DEFAULT 'reject' COMMENT '审核状态（reject驳回|adopt通过）',
  `file_id` int(11) DEFAULT NULL COMMENT '签名附件ID(未启用)',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `programme_examine_dep` */

/*Table structure for table `programme_step` */

DROP TABLE IF EXISTS `programme_step`;

CREATE TABLE `programme_step` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '审查方案实施步骤ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '稽核方案ID',
  `type` enum('step','title','content') NOT NULL DEFAULT 'content' COMMENT '类型（step步骤|title标题|content内容）',
  `content` varchar(500) DEFAULT NULL COMMENT '内容',
  `order` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `programme_step` */

/*Table structure for table `programme_user` */

DROP TABLE IF EXISTS `programme_user`;

CREATE TABLE `programme_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '工作方案主要内容ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '稽核方案ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '员工ID',
  `job` char(100) NOT NULL DEFAULT '-' COMMENT '员工行政职务',
  `title` char(100) NOT NULL DEFAULT '-' COMMENT '员工技术职称',
  `task` varchar(500) DEFAULT NULL COMMENT '员工分工',
  `order` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未除|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `programme_user` */

/*Table structure for table `punish_notice` */

DROP TABLE IF EXISTS `punish_notice`;

CREATE TABLE `punish_notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '处罚通知ID',
  `confirmation_id` int(11) NOT NULL DEFAULT '0' COMMENT '确认书ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '被惩罚人ID',
  `number` char(50) DEFAULT NULL COMMENT '文件号',
  `time` timestamp NULL DEFAULT NULL COMMENT '生成时间',
  `basis_clause_id` int(11) DEFAULT NULL COMMENT '依据文件id',
  `state` enum('draft','jh_draft','jh_publish','ld_draft','ld_publish','bgs_draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|publish发布）',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `punish_notice` */

/*Table structure for table `punish_notice_behavior` */

DROP TABLE IF EXISTS `punish_notice_behavior`;

CREATE TABLE `punish_notice_behavior` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '违规行为ID',
  `punish_notice_id` int(11) NOT NULL DEFAULT '0' COMMENT '惩罚通知ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '记录人ID',
  `behavior_id` int(11) DEFAULT NULL COMMENT '违规行为ID',
  `content` char(250) DEFAULT NULL COMMENT '违规行为内容',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '填写时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `punish_notice_behavior` */

/*Table structure for table `punish_notice_score` */

DROP TABLE IF EXISTS `punish_notice_score`;

CREATE TABLE `punish_notice_score` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '处罚分数ID',
  `cognizance_user_id` int(11) NOT NULL DEFAULT '0' COMMENT '处罚人员ID',
  `punish_notice_id` int(11) NOT NULL DEFAULT '0' COMMENT '处罚通知ID',
  `score` int(5) NOT NULL DEFAULT '0' COMMENT '处罚分数(填写的数*1000保存)',
  `money` int(11) DEFAULT NULL COMMENT '处罚金额(填写的数*1000保存)',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '处罚时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `punish_notice_score` */

/*Table structure for table `rbac` */

DROP TABLE IF EXISTS `rbac`;

CREATE TABLE `rbac` (
  `rid` int(11) NOT NULL AUTO_INCREMENT COMMENT '权限主键ID',
  `key` char(80) NOT NULL DEFAULT '-' COMMENT '角色key（字典）',
  `menu_id` int(11) NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `is_read` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否可读',
  `is_write` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否可写',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`rid`)
) ENGINE=MyISAM AUTO_INCREMENT=1263 DEFAULT CHARSET=utf8;

/*Data for the table `rbac` */

insert  into `rbac`(`rid`,`key`,`menu_id`,`is_read`,`is_write`,`delete`) values (763,'staff',23,0,0,0),(762,'staff',22,0,0,0),(761,'staff',21,0,0,0),(760,'staff',20,0,0,0),(759,'staff',19,0,0,0),(758,'staff',18,0,0,0),(757,'staff',27,1,1,0),(756,'staff',26,1,1,0),(755,'staff',25,1,1,0),(754,'staff',24,1,1,0),(753,'staff',17,1,1,0),(752,'staff',16,1,1,0),(751,'staff',15,1,1,0),(750,'staff',14,1,1,0),(1153,'management',22,1,1,0),(1152,'management',21,1,1,0),(1151,'management',20,1,1,0),(1150,'management',19,1,1,0),(1149,'management',18,1,1,0),(1148,'management',17,1,1,0),(1147,'management',29,1,1,0),(1146,'management',27,1,1,0),(1145,'management',26,1,1,0),(1144,'management',25,1,1,0),(1143,'management',16,1,1,0),(1142,'management',15,1,1,0),(1141,'management',14,1,1,0),(1140,'management',0,1,1,0),(1139,'management',28,1,1,0),(1138,'management',13,1,1,0),(1137,'management',24,1,1,0),(1136,'management',12,1,1,0),(1135,'management',11,1,0,0),(1134,'management',10,1,0,0),(1133,'management',9,1,0,0),(1132,'management',8,1,0,0),(1131,'management',7,1,0,0),(1045,'admin',22,1,1,0),(1044,'admin',21,1,1,0),(1043,'admin',20,1,1,0),(1042,'admin',19,1,1,0),(1041,'admin',18,1,1,0),(1040,'admin',30,1,1,0),(1039,'admin',29,1,1,0),(1038,'admin',28,1,1,0),(1037,'admin',27,1,1,0),(1036,'admin',26,1,1,0),(1035,'admin',25,1,1,0),(1034,'admin',24,1,1,0),(1033,'admin',17,1,1,0),(1032,'admin',16,1,1,0),(1031,'admin',15,1,1,0),(1030,'admin',14,1,1,0),(1029,'admin',13,1,1,0),(1028,'admin',12,1,1,0),(1027,'admin',11,1,1,0),(1026,'admin',10,1,1,0),(1025,'admin',9,1,1,0),(1024,'admin',8,1,1,0),(1023,'admin',7,1,1,0),(749,'staff',13,1,1,0),(748,'staff',12,1,1,0),(747,'staff',11,0,0,0),(746,'staff',10,0,0,0),(745,'staff',9,0,0,0),(744,'staff',8,0,0,0),(743,'staff',7,0,0,0),(742,'staff',5,1,0,0),(741,'staff',3,1,0,0),(1022,'admin',5,1,1,0),(1130,'management',5,1,1,0),(740,'staff',1,1,0,0),(1124,'zr',20,0,0,0),(1123,'zr',19,0,0,0),(1122,'zr',18,0,0,0),(1121,'zr',17,1,1,0),(1120,'zr',0,0,0,0),(1119,'zr',27,1,1,0),(1118,'zr',26,1,1,0),(1117,'zr',25,1,1,0),(1116,'zr',16,1,1,0),(1115,'zr',15,1,1,0),(1114,'zr',14,1,1,0),(1113,'zr',0,1,1,0),(1112,'zr',0,0,0,0),(1111,'zr',13,1,1,0),(1110,'zr',24,1,1,0),(1109,'zr',12,1,1,0),(1108,'zr',11,1,0,0),(1107,'zr',10,1,0,0),(1106,'zr',9,1,0,0),(1105,'zr',8,1,0,0),(1104,'zr',7,1,0,0),(1103,'zr',5,1,1,0),(1102,'zr',3,0,0,0),(1101,'zr',1,1,1,0),(1097,'fzr',20,0,0,0),(1096,'fzr',19,0,0,0),(1095,'fzr',18,0,0,0),(1094,'fzr',17,1,1,0),(1093,'fzr',0,0,0,0),(1092,'fzr',27,1,1,0),(1091,'fzr',26,1,1,0),(1090,'fzr',25,1,1,0),(1089,'fzr',16,1,1,0),(1088,'fzr',15,1,1,0),(1087,'fzr',14,1,1,0),(1086,'fzr',0,0,0,0),(1085,'fzr',0,0,0,0),(1084,'fzr',13,1,1,0),(1083,'fzr',24,1,1,0),(1082,'fzr',12,1,1,0),(1081,'fzr',11,1,1,0),(1080,'fzr',10,1,1,0),(1079,'fzr',9,1,1,0),(1078,'fzr',8,1,1,0),(1077,'fzr',7,1,1,0),(1076,'fzr',5,1,0,0),(1075,'fzr',3,0,0,0),(1074,'fzr',1,1,0,0),(1021,'admin',3,1,1,0),(1020,'admin',1,1,1,0),(1129,'management',3,1,1,0),(1128,'management',1,1,1,0),(1046,'admin',23,1,1,0),(1262,'fgld',23,1,1,0),(1261,'fgld',22,1,1,0),(1260,'fgld',21,1,1,0),(1259,'fgld',20,1,1,0),(1258,'fgld',19,1,1,0),(1257,'fgld',18,1,1,0),(1256,'fgld',17,1,1,0),(1255,'fgld',29,1,1,0),(1254,'fgld',27,1,1,0),(1253,'fgld',26,1,1,0),(1252,'fgld',25,1,1,0),(1251,'fgld',16,1,1,0),(1250,'fgld',15,1,1,0),(1249,'fgld',14,1,1,0),(1248,'fgld',30,1,1,0),(1247,'fgld',28,1,1,0),(1246,'fgld',13,1,1,0),(1245,'fgld',24,1,1,0),(1244,'fgld',12,1,1,0),(1243,'fgld',11,1,1,0),(1242,'fgld',10,1,1,0),(1241,'fgld',9,1,1,0),(1240,'fgld',8,1,1,0),(1239,'fgld',7,1,1,0),(1238,'fgld',6,1,1,0),(1237,'fgld',4,1,1,0),(1236,'fgld',2,1,1,0),(1098,'fzr',21,0,0,0),(1099,'fzr',22,0,0,0),(1100,'fzr',23,0,0,0),(1125,'zr',21,0,0,0),(1126,'zr',22,0,0,0),(1127,'zr',23,0,0,0),(1154,'management',23,1,1,0),(1235,'jsz',23,0,0,0),(1234,'jsz',22,0,0,0),(1233,'jsz',21,0,0,0),(1232,'jsz',20,0,0,0),(1231,'jsz',19,0,0,0),(1230,'jsz',18,0,0,0),(1229,'jsz',17,1,1,0),(1228,'jsz',29,1,1,0),(1227,'jsz',27,1,1,0),(1226,'jsz',26,1,1,0),(1225,'jsz',25,1,1,0),(1224,'jsz',16,1,1,0),(1223,'jsz',15,1,1,0),(1222,'jsz',14,1,1,0),(1221,'jsz',30,1,1,0),(1220,'jsz',28,1,1,0),(1219,'jsz',13,1,1,0),(1218,'jsz',24,1,1,0),(1217,'jsz',12,1,1,0),(1216,'jsz',11,1,0,0),(1215,'jsz',10,1,0,0),(1214,'jsz',9,1,0,0),(1213,'jsz',8,1,0,0),(1212,'jsz',7,1,0,0),(1211,'jsz',6,0,0,0),(1210,'jsz',4,0,0,0),(1209,'jsz',2,0,0,0),(1182,'xdnq',0,0,0,0),(1183,'xdnq',0,1,0,0),(1184,'xdnq',0,0,0,0),(1185,'xdnq',0,1,0,0),(1186,'xdnq',0,1,0,0),(1187,'xdnq',0,1,0,0),(1188,'xdnq',0,1,0,0),(1189,'xdnq',0,1,0,0),(1190,'xdnq',0,1,1,0),(1191,'xdnq',0,1,1,0),(1192,'xdnq',0,1,1,0),(1193,'xdnq',0,1,1,0),(1194,'xdnq',0,1,1,0),(1195,'xdnq',0,1,1,0),(1196,'xdnq',0,1,1,0),(1197,'xdnq',0,1,1,0),(1198,'xdnq',0,1,1,0),(1199,'xdnq',0,1,1,0),(1200,'xdnq',0,1,1,0),(1201,'xdnq',0,1,1,0),(1202,'xdnq',0,1,1,0),(1203,'xdnq',0,0,0,0),(1204,'xdnq',0,0,0,0),(1205,'xdnq',0,0,0,0),(1206,'xdnq',0,0,0,0),(1207,'xdnq',0,0,0,0),(1208,'xdnq',0,0,0,0);

/*Table structure for table `rectify` */

DROP TABLE IF EXISTS `rectify`;

CREATE TABLE `rectify` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改通知主键ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿ID',
  `confirmation_id` int(11) NOT NULL DEFAULT '0' COMMENT '事实确认书ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '签署人ID',
  `last_time` date DEFAULT NULL COMMENT '整改报告提交时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新日期',
  `year` year(4) DEFAULT NULL COMMENT '年份',
  `number` int(11) DEFAULT '0' COMMENT '编号',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|publish发布）',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `rectify` */

/*Table structure for table `rectify_demand` */

DROP TABLE IF EXISTS `rectify_demand`;

CREATE TABLE `rectify_demand` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改要求ID',
  `rectify_id` int(11) DEFAULT NULL COMMENT '整改通知书id',
  `content` varchar(10000) DEFAULT NULL COMMENT '整改要求',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `rectify_demand` */

/*Table structure for table `rectify_report` */

DROP TABLE IF EXISTS `rectify_report`;

CREATE TABLE `rectify_report` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改报告ID',
  `rectify_id` int(11) NOT NULL DEFAULT '0' COMMENT '整改通知ID',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|publish发布）',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `rectify_report` */

/*Table structure for table `rectify_report_content` */

DROP TABLE IF EXISTS `rectify_report_content`;

CREATE TABLE `rectify_report_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改详情ID',
  `rectify_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '整改报告ID',
  `draft_content_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿违规内容ID',
  `content` varchar(500) DEFAULT NULL COMMENT '整改详情',
  `time` date DEFAULT NULL COMMENT '整改时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `rectify_report_content` */

/*Table structure for table `rectify_report_content_user` */

DROP TABLE IF EXISTS `rectify_report_content_user`;

CREATE TABLE `rectify_report_content_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改人信息ID',
  `rectify_report_id` int(11) DEFAULT NULL COMMENT '整改报告ID',
  `rectify_report_content_id` int(11) NOT NULL DEFAULT '0' COMMENT '整改报告内容ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '整改人ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `rectify_report_content_user` */

/*Table structure for table `rectify_report_file` */

DROP TABLE IF EXISTS `rectify_report_file`;

CREATE TABLE `rectify_report_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改附件ID',
  `file_id` int(11) NOT NULL DEFAULT '0' COMMENT '附件ID',
  `rectify_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '整改报告ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `rectify_report_file` */

/*Table structure for table `rectify_suggest` */

DROP TABLE IF EXISTS `rectify_suggest`;

CREATE TABLE `rectify_suggest` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改要求ID',
  `rectify_id` int(11) DEFAULT NULL COMMENT '整改通知书id',
  `content` varchar(10000) DEFAULT NULL COMMENT '整改建议',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `rectify_suggest` */

/*Table structure for table `user_job` */

DROP TABLE IF EXISTS `user_job`;

CREATE TABLE `user_job` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` int(11) DEFAULT NULL COMMENT '人员ID',
  `job` char(20) DEFAULT NULL COMMENT '岗位',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

/*Data for the table `user_job` */

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '人员ID',
  `department_id` int(11) NOT NULL DEFAULT '0' COMMENT '部门ID',
  `user_name` char(50) DEFAULT NULL COMMENT '姓名',
  `user_code` char(50) NOT NULL COMMENT '员工号',
  `class` char(20) DEFAULT NULL COMMENT '民族',
  `sex` enum('0','1','2') NOT NULL DEFAULT '0' COMMENT '性别（0保密1女2男）',
  `phone` char(20) DEFAULT NULL COMMENT '联系方式',
  `id_card` char(18) DEFAULT NULL COMMENT '身份证',
  `portrait_id` int(11) NOT NULL DEFAULT '0' COMMENT '头像ID',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `users` */

insert  into `users`(`user_id`,`department_id`,`user_name`,`user_code`,`class`,`sex`,`phone`,`id_card`,`portrait_id`,`update_time`,`delete`) values (-1,1,'超级管理员','admin','汉','0','18153100614','522623199706141234',0,'2019-01-07 18:38:53',0);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
