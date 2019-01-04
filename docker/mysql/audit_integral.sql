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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Data for the table `audit_report` */

insert  into `audit_report`(`id`,`programme_id`,`draft_id`,`confirmation_id`,`rectify_report_id`,`state`,`update_time`,`delete`) values (1,78,70,7,41,'publish','2019-01-04 11:33:30',0);

/*Table structure for table `audit_report_basic_info` */

DROP TABLE IF EXISTS `audit_report_basic_info`;

CREATE TABLE `audit_report_basic_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '审计报告基本信息ID',
  `audit_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '审计报告ID',
  `content` varchar(2500) DEFAULT NULL COMMENT '基本信息',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

/*Data for the table `audit_report_basic_info` */

insert  into `audit_report_basic_info`(`id`,`audit_report_id`,`content`,`delete`) values (5,1,'<p>这是基本情况</p>',0);

/*Table structure for table `audit_report_plan` */

DROP TABLE IF EXISTS `audit_report_plan`;

CREATE TABLE `audit_report_plan` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '审计报告基本信息ID',
  `audit_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '审计报告ID',
  `content` varchar(2500) DEFAULT NULL COMMENT '基本信息',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

/*Data for the table `audit_report_plan` */

insert  into `audit_report_plan`(`id`,`audit_report_id`,`content`,`delete`) values (5,1,'<p>这是下一步工资措施</p>',0);

/*Table structure for table `audit_report_reason` */

DROP TABLE IF EXISTS `audit_report_reason`;

CREATE TABLE `audit_report_reason` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '审计报告基本信息ID',
  `audit_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '审计报告ID',
  `content` varchar(2500) DEFAULT NULL COMMENT '基本信息',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

/*Data for the table `audit_report_reason` */

insert  into `audit_report_reason`(`id`,`audit_report_id`,`content`,`delete`) values (5,1,'<p>这是问题形成的原因</p>',0);

/*Table structure for table `clause` */

DROP TABLE IF EXISTS `clause`;

CREATE TABLE `clause` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '管理办法及条款ID',
  `department_id` int(11) NOT NULL DEFAULT '-1' COMMENT '所属部门ID（为-1时所有部门通用）',
  `title` char(200) DEFAULT NULL COMMENT '管理办法及条款标题',
  `author_id` int(11) DEFAULT NULL COMMENT '发布人ID',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿publish发布）',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

/*Data for the table `clause` */

insert  into `clause`(`id`,`department_id`,`title`,`author_id`,`update_time`,`state`,`delete`) values (2,1,'公告1',1,'2018-11-30 15:05:50','publish',0),(6,-1,'普定县农村信用合作联社员工 违规行为积分管理实施细则',2,'2018-12-27 18:20:53','publish',0),(7,-1,'安顺市信用合作联社员工 违规行为积分管理实施细则',2,'2018-12-27 18:20:59','publish',0),(8,-1,'普定县农村信用合作联社员工 违规行为积分管理实施细则2',2,'2018-12-03 21:28:05','draft',0);

/*Table structure for table `clause_content` */

DROP TABLE IF EXISTS `clause_content`;

CREATE TABLE `clause_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '管理办法内容ID',
  `clause_id` int(11) DEFAULT NULL COMMENT '管理办法ID',
  `is_title` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为标题（0内容1标题）',
  `title_level` char(20) NOT NULL DEFAULT '-' COMMENT '标题级别',
  `content` varchar(1000) DEFAULT NULL COMMENT '内容',
  `order` int(4) NOT NULL DEFAULT '0' COMMENT '顺序',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;

/*Data for the table `clause_content` */

insert  into `clause_content`(`id`,`clause_id`,`is_title`,`title_level`,`content`,`order`,`delete`) values (1,4,1,'h1','标题测试',1,0),(2,5,1,'h1','标题测试1',1,0),(8,5,1,'h2','标题测试3',2,0),(9,6,1,'h1','第一章 总　则',1,0),(10,6,0,'h2','第一条 为加强内控管理，切实解决有章不循、违规操作问题，有效防范和控制操作风险，进一步促进我县农村信用社安全稳健运行，根据国家有关法律、金融法规和省联社有关内控管理制度规定，制定本办法。',2,0),(11,6,0,'h2','第二条 本办法所称的违规行为是指不按制度或流程办理、审批有关事项，未造成不良后果或情节轻微，不足以按照有关规定实施处理的行为。',3,0),(12,6,0,'h2','第三条 本办法所称的积分管理是对员工违规行为进行年度累计积分，并将累计积分分别与行政处罚、绩效薪酬、评先选优、职务晋升等挂钩进行处理的管理方法。',4,0),(13,6,0,'h2','第四条 本办法所称部门是指县联社所辖各部室、各乡镇信用社、分社、小微企业金融服务中心等；所称人员是指普定联社所有在岗人员（含高管人员、劳务派遣员工）。',5,0),(14,6,0,'h2','第五条 违规积分分为直接违规积分和连带责任积分。直接违规积分是对违规当事人进行的积分，连带责任积分是对违规行为负有管理或监督责任的人员进行的积分。',6,0),(28,7,1,'h1','第一章 总　则',1,0),(29,7,0,'h1','第一条 为加强内控管理，切实解决有章不循、违规操作问题，有效防范和控制操作风险，进一步促进我县农村信用社安全稳健运行，根据国家有关法律、金融法规和省联社有关内控管理制度规定，制定本办法。',2,0),(30,7,0,'h1','第二条 本办法所称的违规行为是指不按制度或流程办理、审批有关事项，未造成不良后果或情节轻微，不足以按照有关规定实施处理的行为。',3,0),(31,7,0,'h1','第三条 本办法所称的积分管理是对员工违规行为进行年度累计积分，并将累计积分分别与行政处罚、绩效薪酬、评先选优、职务晋升等挂钩进行处理的管理方法。',4,0),(32,8,1,'','updateContentArr',1,0),(39,9,0,'','1',1,0);

/*Table structure for table `clause_file` */

DROP TABLE IF EXISTS `clause_file`;

CREATE TABLE `clause_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `clause_id` int(11) DEFAULT NULL COMMENT '管理办法ID',
  `file_id` int(11) DEFAULT NULL COMMENT '附件ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=34 DEFAULT CHARSET=utf8;

/*Data for the table `clause_file` */

insert  into `clause_file`(`id`,`clause_id`,`file_id`,`delete`) values (1,2,3,0),(2,2,4,0),(3,2,5,0),(31,3,4,0),(30,3,3,0),(33,4,4,0),(32,4,3,0),(29,5,8,0),(28,5,5,0);

/*Table structure for table `confirmation` */

DROP TABLE IF EXISTS `confirmation`;

CREATE TABLE `confirmation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '确认书ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿ID',
  `confirmation_receipt_id` int(11) NOT NULL DEFAULT '0' COMMENT '确认书回执ID',
  `file_id` int(11) DEFAULT '0' COMMENT '图片ID',
  `has_read` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已读（0未读|1已读）',
  `has_read_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '已读时间',
  `proposal` varchar(2000) DEFAULT NULL COMMENT '单位意见或者建议',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|publish发布）',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`),
  KEY `工作底稿` (`draft_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

/*Data for the table `confirmation` */

insert  into `confirmation`(`id`,`draft_id`,`confirmation_receipt_id`,`file_id`,`has_read`,`has_read_time`,`proposal`,`state`,`update_time`,`delete`) values (7,70,0,0,1,'2018-12-28 11:48:20',NULL,'publish','2018-12-28 19:57:21',0);

/*Table structure for table `confirmation_basis` */

DROP TABLE IF EXISTS `confirmation_basis`;

CREATE TABLE `confirmation_basis` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '依据主键ID',
  `confirmation_id` int(11) NOT NULL DEFAULT '0' COMMENT '事实确认书ID',
  `basis_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿依据ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

/*Data for the table `confirmation_basis` */

insert  into `confirmation_basis`(`id`,`confirmation_id`,`basis_id`,`delete`) values (15,7,138,0),(16,7,139,0);

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

/*Table structure for table `department` */

DROP TABLE IF EXISTS `department`;

CREATE TABLE `department` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `name` char(50) DEFAULT NULL COMMENT '部门名称',
  `parent_id` int(11) DEFAULT NULL COMMENT '上级部门ID',
  `code` char(50) DEFAULT NULL COMMENT '部门编码',
  `has_child` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有子部门（0没有1有）',
  `level` char(20) DEFAULT NULL COMMENT '部门级别',
  `type` char(20) DEFAULT NULL COMMENT '类型（字典key）',
  `grade` int(3) DEFAULT NULL COMMENT '所在部门树的层级',
  `address` char(250) DEFAULT NULL COMMENT '地址',
  `phone` char(11) DEFAULT NULL COMMENT '手机号或者电话',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete` tinyint(1) DEFAULT '0' COMMENT '是否删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8;

/*Data for the table `department` */

insert  into `department`(`id`,`name`,`parent_id`,`code`,`has_child`,`level`,`type`,`grade`,`address`,`phone`,`update_time`,`delete`) values (48,'部门2',47,'d12',0,'1',NULL,NULL,'贵州贵阳1','','2018-12-10 15:10:17',0),(55,'部门3',47,'',0,'0',NULL,NULL,'','','2018-12-10 12:20:24',0),(57,'xx农商行',-1,'',1,'0',NULL,NULL,'','','2018-12-20 17:08:27',0),(58,'公司部',57,'',0,'0',NULL,NULL,'','','2018-12-20 17:09:03',0),(59,'行社领导',57,'',0,'0',NULL,NULL,'','','2018-12-20 17:09:25',0),(60,'保卫部',57,'',0,'0',NULL,NULL,'','','2018-12-20 17:11:30',0),(61,'党建办',57,'',0,'0',NULL,NULL,'','','2018-12-20 17:11:39',0),(62,'电子银行部',57,'',0,'0',NULL,NULL,'','','2018-12-20 17:11:53',0),(63,'合规风险部',57,'',0,'0',NULL,NULL,'','','2018-12-20 17:12:05',0),(64,'基建办公室',57,'',0,'0',NULL,NULL,'','','2018-12-20 17:12:28',0),(65,'稽核审计部',57,'',0,'0',NULL,NULL,'','','2018-12-20 17:12:35',0),(66,'纪检监察室',57,'',0,'0',NULL,NULL,'','','2018-12-20 17:12:44',0),(67,'监控中心',57,'',0,'0',NULL,NULL,'','','2018-12-20 17:12:52',0);

/*Table structure for table `department_user` */

DROP TABLE IF EXISTS `department_user`;

CREATE TABLE `department_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `department_id` int(11) DEFAULT NULL COMMENT '部门id',
  `user_id` int(11) DEFAULT NULL COMMENT '人员id',
  `type` char(20) DEFAULT NULL COMMENT '角色类型',
  `delete` tinyint(1) DEFAULT '0' COMMENT '软删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8;

/*Data for the table `department_user` */

insert  into `department_user`(`id`,`department_id`,`user_id`,`type`,`delete`) values (1,18,1,'admin',0),(2,19,1,'admin',0),(3,48,1,'admin',0),(4,21,1,'admin',0),(5,22,1,'admin',0),(6,23,1,'admin',0),(12,31,0,'',0),(13,32,0,'',0),(14,33,0,'',0),(15,34,-2,'',0),(17,36,-2,'',0),(19,20,2,'admins',0),(20,39,-2,'admin',0),(21,40,-2,'',0),(26,48,2,'admin',0),(27,55,-2,'',0),(28,56,-2,'',0),(29,57,-2,'',0),(30,58,-2,'',0),(31,59,-2,'',0),(32,60,-2,'',0),(33,61,-2,'',0),(34,62,-2,'',0),(35,63,-2,'',0),(36,64,-2,'',0),(37,65,-2,'',0),(38,66,-2,'',0),(39,67,-2,'',0),(40,68,0,'',0),(41,68,0,'',0),(42,69,0,'',0),(43,70,16,'',0),(44,70,1,'',0);

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
) ENGINE=InnoDB AUTO_INCREMENT=199 DEFAULT CHARSET=utf8;

/*Data for the table `dictionary` */

insert  into `dictionary`(`id`,`type_id`,`key`,`value`,`order`,`describe`,`delete`) values (-405,-4,'h5','无机标题',5,'系统分类（不允许删除）',0),(-404,-4,'h4','四级标题',4,'系统分类（不允许删除）',0),(-403,-4,'h3','三级标题',3,'系统分类（不允许删除）',0),(-402,-4,'h2','二级标题',2,'系统分类（不允许删除）',0),(-401,-4,'h1','一级标题',1,'系统分类（不允许删除）',0),(-310,-3,'upload','文件',10,'系统分类（不允许删除）',0),(-309,-3,'password','编辑',9,'系统分类（不允许删除）',0),(-308,-3,'login','登录',2,'系统分类（不允许删除）',0),(-307,-3,'is-use','变更',3,'系统分类（不允许删除）',0),(-306,-3,'tree','获取',4,'系统分类（不允许删除）',0),(-305,-3,'delete','删除',5,'系统分类（不允许删除）',0),(-304,-3,'edit','编辑',6,'系统分类（不允许删除）',0),(-303,-3,'get','获取',7,'系统分类（不允许删除）',0),(-302,-3,'add','添加',8,'系统分类（不允许删除）',0),(-301,-3,'list','获取列表',1,'系统分类（不允许删除）',0),(-203,-2,'staff','业务员',3,'部门业务员',0),(-202,-2,'management','负责人',2,'部门负责人',0),(-201,-2,'admin','管理员',1,'系统管理员',0),(-102,-1,'other','其他',0,'系统分类（不允许删除）',0),(-101,-1,'system','系统',0,'系统分类（不允许删除）',0),(184,-5,'auditKey1','方案类型1',1,'',0),(185,-5,'auditKey2','方案类型2',2,'',0),(186,-5,'auditKey3','方案类型3',3,'',0),(187,-6,'auditType1','审计方式1',1,'',0),(188,-6,'auditType2','审计方式2',2,'',0),(189,-6,'auditType3','审计方式3',3,'',0),(190,-8,'userJob1','人员职务1',1,'',0),(191,-8,'userJob2','人员职务2',2,'',0),(192,-8,'userJob3','人员职务3',3,'',0),(193,-7,'userTitle1','人员技术职称1',1,'',0),(194,-7,'userTitle2','人员技术职称2',2,'',0),(195,-7,'userTitle3','人员技术职称3',3,'',0),(196,-9,'auditTask1','员工分工1',1,'',0),(197,-9,'auditTask2','员工分工2',2,'',0),(198,-9,'auditTask3','员工分工3',3,'',0);

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `dictionary_type` */

insert  into `dictionary_type`(`id`,`type_id`,`key`,`title`,`is_use`,`update_time`,`user_id`,`describe`,`delete`) values (-9,-1,'system','员工分工',1,'2018-12-19 16:51:04',0,'稽核员工分工',0),(-8,-1,'system','人员技术职称',1,'2018-12-19 16:49:39',0,'稽核人员技术职称',0),(-7,-1,'system','人员职务',1,'2018-12-19 16:48:09',0,'稽核人员职务',0),(-6,-1,'system','审计方式',1,'2018-12-19 18:07:13',0,'审计方式',0),(-5,-1,'system','方案类型',1,'2018-12-19 18:06:38',0,'方案类型',0),(-4,-1,'system','管理办法标题级别',0,'2018-12-03 20:31:18',0,'管理办法标题级别选项',0),(-3,-1,'system','日志分类',1,'2018-11-29 18:15:50',0,'系统日志分类',0),(-2,-1,'system','人员角色',1,'2018-11-26 12:08:34',0,'部门人员角色字典',0),(-1,-1,'system','字典分类',1,'2018-11-12 15:37:43',0,'系统字典分类',0);

/*Table structure for table `draft` */

DROP TABLE IF EXISTS `draft`;

CREATE TABLE `draft` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '工作底稿ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '方案ID',
  `query_department_id` int(11) NOT NULL DEFAULT '0' COMMENT '被查询机构ID',
  `department_id` int(11) NOT NULL DEFAULT '0' COMMENT '检查机构ID',
  `number` char(100) DEFAULT NULL COMMENT '编号',
  `project_name` char(250) DEFAULT NULL COMMENT '项目名称',
  `public` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否公开（0私有|1公开）【被检查人是否能查看】',
  `time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '检查日期',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '底稿状态（draft草稿|publish发布）',
  `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8;

/*Data for the table `draft` */

insert  into `draft`(`id`,`programme_id`,`query_department_id`,`department_id`,`number`,`project_name`,`public`,`time`,`state`,`update_time`,`delete`) values (70,78,64,62,'jh123','测试三',1,'2018-12-20 00:00:00','publish','2018-12-28 11:18:13',0),(71,78,0,0,'1234','55',1,'2019-01-02 14:13:30','draft','2019-01-02 14:13:37',0);

/*Table structure for table `draft_admin_user` */

DROP TABLE IF EXISTS `draft_admin_user`;

CREATE TABLE `draft_admin_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '复查ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '底稿ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '复查人员ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

/*Data for the table `draft_admin_user` */

insert  into `draft_admin_user`(`id`,`draft_id`,`user_id`,`delete`) values (21,70,1,0);

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
) ENGINE=InnoDB AUTO_INCREMENT=136 DEFAULT CHARSET=utf8;

/*Data for the table `draft_content` */

insert  into `draft_content`(`id`,`draft_id`,`order`,`type`,`behavior_id`,`behavior_content`,`delete`) values (129,70,1,'title',0,'办公用房整改不及时性、不到位。',0),(130,70,2,'other',0,'审计核查发现，××年×月开展办公用房清退工作以来，上级职能部门先后×次对你单位办公用房清理腾退情况进行督促检查，检查中发现你单位领导干部办公用房超标，没有 按要求真腾实退。上级职能部门专门向你单位制发整改通知， 要求限期整改完毕。',0),(131,70,3,'other',0,'第一条 为加强内控管理，切实解决有章不循、违规操作问题，有效防范和控制操作风险，进一步促进我县农村信用社安全稳健运行，根据国家有关法律、金融法规和省联社有关内控管理制度规定，制定本办法。',0),(132,70,4,'other',0,'第二条 本办法所称的违规行为是指不按制度或流程办理、审批有关事项，未造成不良后果或情节轻微，不足以按照有关规定实施处理的行为。',0),(133,70,5,'title',0,'办理固定资产报销结算手续不规范。',0),(134,70,6,'other',0,'第三条 本办法所称的积分管理是对员工违规行为进行年度累计积分，并将累计积分分别与行政处罚、绩效薪酬、评先选优、职务晋升等挂钩进行处理的管理方法。',0),(135,70,7,'other',0,'第四条 本办法所称部门是指县联社所辖各部室、各乡镇信用社、分社、小微企业金融服务中心等；所称人员是指普定联社所有在岗人员（含高管人员、劳务派遣员工）。',0);

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
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8;

/*Data for the table `draft_inspect_user` */

insert  into `draft_inspect_user`(`id`,`draft_id`,`user_id`,`delete`) values (56,70,13,0),(57,70,14,0);

/*Table structure for table `draft_query_user` */

DROP TABLE IF EXISTS `draft_query_user`;

CREATE TABLE `draft_query_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `draft_id` int(11) DEFAULT NULL COMMENT '底稿ID',
  `user_id` int(11) DEFAULT NULL COMMENT '检查人员ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8;

/*Data for the table `draft_query_user` */

insert  into `draft_query_user`(`id`,`draft_id`,`user_id`,`delete`) values (57,70,16,0),(58,70,15,0);

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
  `name` char(100) DEFAULT NULL COMMENT '附件名称',
  `suffix` char(10) DEFAULT NULL COMMENT '附件后缀',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上传时间',
  `path` char(250) DEFAULT NULL COMMENT '附件存放位置',
  `size` int(10) DEFAULT '0' COMMENT '文件大小（单位B）',
  `file_name` char(100) DEFAULT NULL COMMENT '存放名称',
  `form_id` int(11) NOT NULL DEFAULT '0' COMMENT '外部ID',
  `form` char(30) NOT NULL DEFAULT '-' COMMENT '来自那张表',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除（0未删1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=474 DEFAULT CHARSET=utf8;

/*Data for the table `files` */

insert  into `files`(`id`,`name`,`suffix`,`time`,`path`,`size`,`file_name`,`form_id`,`form`,`delete`) values (3,'系统设置','json','2018-12-03 21:14:15','201811/',0,'1543478749_系统设置',4,'clause',0),(4,'系统设置','json','2018-12-03 21:14:15','201811/',0,'1543478751_系统设置',4,'clause',0),(5,'系统设置','json','2018-12-03 21:13:35','201811/',0,'1543478752_系统设置',5,'clause',0),(8,'系统设置','json','2018-12-03 21:13:35','201811/',0,'1543479000_系统设置',5,'clause',0),(9,'系统设置','json','2018-11-29 16:10:01','201811/',0,'1543479001_系统设置',0,'-',0),(10,'系统设置','json','2018-11-29 16:10:02','201811/',0,'1543479002_系统设置',0,'-',0),(11,'系统设置','json','2018-11-29 16:10:03','201811/',0,'1543479003_系统设置',0,'-',0),(12,'系统设置','json','2018-11-29 16:10:03','201811/',0,'1543479003_系统设置',0,'-',0),(13,'系统设置','json','2018-11-29 16:10:04','201811/',0,'1543479004_系统设置',0,'-',0),(14,'系统设置','json','2018-11-29 16:10:05','201811/',0,'1543479005_系统设置',0,'-',0),(15,'系统设置','json','2018-11-29 16:10:05','201811/',0,'1543479005_系统设置',0,'-',0),(16,'系统设置','json','2018-11-29 16:10:06','201811/',0,'1543479006_系统设置',0,'-',0),(17,'系统设置','json','2018-11-29 16:23:36','201811/',0,'1543479816_系统设置',0,'-',0),(18,'系统设置','json','2018-11-29 16:32:19','201811/',11672,'1543480339_系统设置',0,'-',0),(19,'vscode-setting','txt','2018-11-29 18:22:25','201811/',1902,'1543486945_vscode-setting',0,'-',0),(20,'vscode-setting','txt','2018-11-29 18:23:07','201811/',1902,'1543486987_vscode-setting',0,'-',0),(21,'刘亦菲','gif','2018-11-29 18:23:55','201811/',5116027,'1543487035_刘亦菲',0,'-',0),(22,'微信图片_20181122135157','jpg','2018-11-29 18:29:10','201811/',11887,'1543487350_微信图片_20181122135157',0,'-',0),(23,'微信图片_20181122135157','jpg','2018-11-29 18:29:25','201811/',11887,'1543487365_微信图片_20181122135157',0,'-',0),(24,'微信图片_20181122135157','jpg','2018-11-29 18:29:28','201811/',11887,'1543487368_微信图片_20181122135157',0,'-',0),(25,'微信图片_20181122135157','jpg','2018-11-29 18:29:33','201811/',11887,'1543487373_微信图片_20181122135157',0,'-',0),(26,'微信图片_20181122135157','jpg','2018-11-29 18:32:49','201811/',11887,'1543487569_微信图片_20181122135157',0,'-',0),(27,'微信图片_20181122135157','jpg','2018-11-29 18:36:27','201811/',11887,'1543487787_微信图片_20181122135157',0,'-',0),(28,'微信图片_20181122135157','jpg','2018-11-29 18:36:32','201811/',11887,'1543487792_微信图片_20181122135157',0,'-',0),(29,'微信图片_20181122135157','jpg','2018-11-29 18:36:35','201811/',11887,'1543487795_微信图片_20181122135157',0,'-',0),(31,'微信图片_20181122135157','jpg','2018-11-29 18:37:10','201811/',11887,'1543487830_微信图片_20181122135157',0,'-',0),(32,'微信图片_20181122135157','jpg','2018-11-29 18:37:35','201811/',11887,'1543487855_微信图片_20181122135157',0,'-',0),(33,'微信图片_20181122135157','jpg','2018-11-29 18:37:38','201811/',11887,'1543487858_微信图片_20181122135157',0,'-',0),(34,'微信图片_20181122135157','jpg','2018-11-29 18:38:41','201811/',11887,'1543487921_微信图片_20181122135157',0,'-',0),(35,'微信图片_20181122135157','jpg','2018-11-29 18:39:03','201811/',11887,'1543487943_微信图片_20181122135157',0,'-',0),(36,'微信图片_20181122135157','jpg','2018-11-29 18:39:07','201811/',11887,'1543487947_微信图片_20181122135157',0,'-',0),(37,'微信图片_20181122135157','jpg','2018-11-29 18:39:10','201811/',11887,'1543487950_微信图片_20181122135157',0,'-',0),(38,'微信图片_20181122135157','jpg','2018-11-29 18:39:57','201811/',11887,'1543487997_微信图片_20181122135157',0,'-',0),(41,'微信图片_20181122135157','jpg','2018-11-29 18:41:39','201811/',11887,'1543488099_微信图片_20181122135157',0,'-',0),(42,'微信图片_20181122135157','jpg','2018-11-29 18:41:42','201811/',11887,'1543488102_微信图片_20181122135157',0,'-',0),(43,'微信图片_20181122135157','jpg','2018-11-29 18:41:45','201811/',11887,'1543488105_微信图片_20181122135157',0,'-',0),(44,'刘亦菲','gif','2018-11-29 18:42:04','201811/',5116027,'1543488124_刘亦菲',0,'-',0),(45,'刘亦菲','gif','2018-11-29 18:42:07','201811/',5116027,'1543488127_刘亦菲',0,'-',0),(46,'js运行机制','png','2018-11-29 18:44:55','201811/',187655,'1543488295_js运行机制',0,'-',0),(47,'刘亦菲','gif','2018-11-29 18:46:48','201811/',5116027,'1543488408_刘亦菲',0,'-',0),(48,'刘亦菲','gif','2018-11-29 18:46:53','201811/',5116027,'1543488413_刘亦菲',0,'-',0),(49,'刘亦菲','gif','2018-11-29 18:47:00','201811/',5116027,'1543488420_刘亦菲',0,'-',0),(50,'刘亦菲','gif','2018-11-29 18:52:41','201811/',5116027,'1543488761_刘亦菲',0,'-',0),(51,'刘亦菲','gif','2018-11-29 18:53:02','201811/',5116027,'1543488782_刘亦菲',0,'-',0),(52,'刘亦菲','gif','2018-11-29 18:53:15','201811/',5116027,'1543488795_刘亦菲',0,'-',0),(53,'刘亦菲','gif','2018-11-29 18:53:18','201811/',5116027,'1543488798_刘亦菲',0,'-',0),(54,'刘亦菲','gif','2018-11-29 18:53:22','201811/',5116027,'1543488802_刘亦菲',0,'-',0),(55,'刘亦菲','gif','2018-11-29 19:00:26','201811/',5116027,'1543489226_刘亦菲',0,'-',0),(56,NULL,NULL,'2018-11-30 09:38:25',NULL,0,NULL,42,'notice',0),(57,NULL,NULL,'2018-11-30 09:38:25',NULL,0,NULL,42,'notice',0),(58,NULL,NULL,'2018-11-30 09:38:25',NULL,0,NULL,42,'notice',0),(59,'系统设置','json','2018-11-30 09:42:08','201811/',11672,'1543541593_系统设置',43,'notice',0),(61,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-11-30 09:47:43','201811/',403120,'1543542093_a532d333c895d14321efbc1e71f082025baf0783',18,'notice',0),(62,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-11-30 09:43:25','201811/',403120,'1543542205_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(63,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-11-30 09:43:54','201811/',403120,'1543542234_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(64,'6767178-a13d5324f8d8cd98','png','2018-11-30 09:43:57','201811/',229979,'1543542237_6767178-a13d5324f8d8cd98',0,'-',0),(65,'6767178-a13d5324f8d8cd98','png','2018-11-30 09:45:12','201811/',229979,'1543542312_6767178-a13d5324f8d8cd98',0,'-',0),(66,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-11-30 09:46:01','201811/',403120,'1543542361_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(67,'6767178-a13d5324f8d8cd98','png','2018-11-30 09:46:08','201811/',229979,'1543542368_6767178-a13d5324f8d8cd98',0,'-',0),(68,'6767178-a13d5324f8d8cd98','png','2018-11-30 09:51:17','201811/',229979,'1543542677_6767178-a13d5324f8d8cd98',0,'-',0),(69,'6767178-a13d5324f8d8cd98','png','2018-11-30 09:51:20','201811/',229979,'1543542680_6767178-a13d5324f8d8cd98',0,'-',0),(70,'6767178-a13d5324f8d8cd98','png','2018-11-30 09:51:22','201811/',229979,'1543542682_6767178-a13d5324f8d8cd98',0,'-',0),(71,'6767178-a13d5324f8d8cd98','png','2018-11-30 09:59:30','201811/',229979,'1543543161_6767178-a13d5324f8d8cd98',45,'notice',0),(72,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-11-30 09:59:30','201811/',403120,'1543543166_a532d333c895d14321efbc1e71f082025baf0783',45,'notice',0),(73,'6767178-a13d5324f8d8cd98','png','2018-11-30 10:03:15','201811/',229979,'1543543395_6767178-a13d5324f8d8cd98',0,'-',0),(74,'6767178-a13d5324f8d8cd98','png','2018-11-30 10:03:33','201811/',229979,'1543543413_6767178-a13d5324f8d8cd98',0,'-',0),(75,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-11-30 10:04:19','201811/',403120,'1543543459_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(76,'6767178-a13d5324f8d8cd98','png','2018-11-30 10:04:22','201811/',229979,'1543543462_6767178-a13d5324f8d8cd98',0,'-',0),(77,'6767178-a13d5324f8d8cd98','png','2018-12-03 13:40:10','201812/',229979,'1543815610_6767178-a13d5324f8d8cd98',0,'-',0),(78,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:03:12','201812/',229979,'1543816992_6767178-a13d5324f8d8cd98',0,'-',0),(79,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:26:39','201812/',229979,'1543818399_6767178-a13d5324f8d8cd98',0,'-',0),(80,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 14:35:18','201812/',403120,'1543818918_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(81,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:38:24','201812/',229979,'1543819104_6767178-a13d5324f8d8cd98',0,'-',0),(82,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:38:26','201812/',229979,'1543819106_6767178-a13d5324f8d8cd98',0,'-',0),(83,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 14:38:29','201812/',403120,'1543819109_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(84,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 14:40:34','201812/',403120,'1543819234_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(85,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:40:37','201812/',229979,'1543819237_6767178-a13d5324f8d8cd98',0,'-',0),(86,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:40:42','201812/',229979,'1543819242_6767178-a13d5324f8d8cd98',0,'-',0),(87,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:45:58','201812/',229979,'1543819558_6767178-a13d5324f8d8cd98',0,'-',0),(88,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 14:46:04','201812/',403120,'1543819564_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(89,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 14:48:22','201812/',403120,'1543819702_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(90,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:48:24','201812/',229979,'1543819704_6767178-a13d5324f8d8cd98',0,'-',0),(91,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 14:48:28','201812/',403120,'1543819708_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(92,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:49:11','201812/',229979,'1543819751_6767178-a13d5324f8d8cd98',0,'-',0),(93,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:49:14','201812/',229979,'1543819754_6767178-a13d5324f8d8cd98',0,'-',0),(94,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:49:16','201812/',229979,'1543819756_6767178-a13d5324f8d8cd98',0,'-',0),(95,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:51:30','201812/',229979,'1543819890_6767178-a13d5324f8d8cd98',0,'-',0),(96,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:51:32','201812/',229979,'1543819892_6767178-a13d5324f8d8cd98',0,'-',0),(97,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:51:35','201812/',229979,'1543819895_6767178-a13d5324f8d8cd98',0,'-',0),(98,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:52:37','201812/',229979,'1543819957_6767178-a13d5324f8d8cd98',0,'-',0),(99,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:52:39','201812/',229979,'1543819959_6767178-a13d5324f8d8cd98',0,'-',0),(100,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:52:41','201812/',229979,'1543819961_6767178-a13d5324f8d8cd98',0,'-',0),(101,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 14:55:26','201812/',403120,'1543820126_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(102,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:55:29','201812/',229979,'1543820129_6767178-a13d5324f8d8cd98',0,'-',0),(103,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 14:56:01','201812/',403120,'1543820161_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(104,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:56:04','201812/',229979,'1543820164_6767178-a13d5324f8d8cd98',0,'-',0),(105,'6767178-a13d5324f8d8cd98','png','2018-12-03 14:59:23','201812/',229979,'1543820363_6767178-a13d5324f8d8cd98',0,'-',0),(106,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 14:59:25','201812/',403120,'1543820365_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(107,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:00:03','201812/',229979,'1543820403_6767178-a13d5324f8d8cd98',0,'-',0),(108,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:00:06','201812/',403120,'1543820406_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(109,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:00:15','201812/',229979,'1543820415_6767178-a13d5324f8d8cd98',0,'-',0),(110,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:03:38','201812/',403120,'1543820618_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(111,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:03:41','201812/',229979,'1543820621_6767178-a13d5324f8d8cd98',0,'-',0),(112,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:03:44','201812/',403120,'1543820624_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(113,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:04:21','201812/',403120,'1543820661_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(114,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:04:24','201812/',229979,'1543820664_6767178-a13d5324f8d8cd98',0,'-',0),(115,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:04:26','201812/',403120,'1543820666_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(116,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:06:51','201812/',403120,'1543820811_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(117,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:06:53','201812/',229979,'1543820813_6767178-a13d5324f8d8cd98',0,'-',0),(118,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:06:56','201812/',403120,'1543820816_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(119,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:09:16','201812/',403120,'1543820956_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(120,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:09:19','201812/',403120,'1543820959_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(121,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:09:22','201812/',229979,'1543820962_6767178-a13d5324f8d8cd98',0,'-',0),(122,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:14:11','201812/',403120,'1543821251_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(123,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:14:13','201812/',229979,'1543821253_6767178-a13d5324f8d8cd98',0,'-',0),(124,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:14:33','201812/',229979,'1543821273_6767178-a13d5324f8d8cd98',0,'-',0),(125,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:14:35','201812/',403120,'1543821275_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(126,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:14:38','201812/',229979,'1543821278_6767178-a13d5324f8d8cd98',0,'-',0),(127,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:19:07','201812/',403120,'1543821547_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(128,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:19:09','201812/',229979,'1543821549_6767178-a13d5324f8d8cd98',0,'-',0),(129,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:19:11','201812/',403120,'1543821551_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(130,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:20:54','201812/',229979,'1543821654_6767178-a13d5324f8d8cd98',0,'-',0),(131,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:20:57','201812/',403120,'1543821657_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(132,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:20:59','201812/',229979,'1543821659_6767178-a13d5324f8d8cd98',0,'-',0),(133,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:26:48','201812/',229979,'1543822008_6767178-a13d5324f8d8cd98',0,'-',0),(134,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:26:49','201812/',403120,'1543822009_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(135,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:26:52','201812/',229979,'1543822012_6767178-a13d5324f8d8cd98',0,'-',0),(136,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:26:56','201812/',403120,'1543822016_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(137,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:27:26','201812/',229979,'1543822046_6767178-a13d5324f8d8cd98',0,'-',0),(138,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:27:28','201812/',403120,'1543822048_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(139,'6767178-a13d5324f8d8cd98','png','2018-12-03 15:27:30','201812/',229979,'1543822050_6767178-a13d5324f8d8cd98',0,'-',0),(140,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:27:33','201812/',403120,'1543822053_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(141,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:28:23','201812/',403120,'1543822103_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(142,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:28:26','201812/',403120,'1543822106_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(143,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 15:29:00','201812/',403120,'1543822140_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(145,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 17:33:59','201812/',403120,'1543822146_a532d333c895d14321efbc1e71f082025baf0783',47,'notice',0),(146,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 17:33:59','201812/',403120,'1543822154_a532d333c895d14321efbc1e71f082025baf0783',47,'notice',0),(147,'6767178-a13d5324f8d8cd98','png','2018-12-03 16:17:29','201812/',229979,'1543825049_6767178-a13d5324f8d8cd98',0,'-',0),(148,'6767178-a13d5324f8d8cd98','png','2018-12-03 16:18:03','201812/',229979,'1543825083_6767178-a13d5324f8d8cd98',0,'-',0),(149,'6767178-a13d5324f8d8cd98','png','2018-12-03 16:26:16','201812/',229979,'1543825576_6767178-a13d5324f8d8cd98',0,'-',0),(150,'6767178-a13d5324f8d8cd98','png','2018-12-03 16:30:59','201812/',229979,'1543825859_6767178-a13d5324f8d8cd98',0,'-',0),(151,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 16:32:08','201812/',403120,'1543825928_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(152,'a532d333c895d14321efbc1e71f082025baf0783','jpg','2018-12-03 16:32:22','201812/',403120,'1543825942_a532d333c895d14321efbc1e71f082025baf0783',0,'-',0),(153,'6767178-a13d5324f8d8cd98','png','2018-12-03 16:32:26','201812/',229979,'1543825946_6767178-a13d5324f8d8cd98',0,'-',0),(154,'6767178-a13d5324f8d8cd98','png','2018-12-03 16:32:37','201812/',229979,'1543825957_6767178-a13d5324f8d8cd98',0,'-',0),(155,'6767178-a13d5324f8d8cd98','png','2018-12-03 16:34:40','201812/',229979,'1543826080_6767178-a13d5324f8d8cd98',0,'-',0),(156,'6767178-a13d5324f8d8cd98','png','2018-12-03 16:35:19','201812/',229979,'1543826119_6767178-a13d5324f8d8cd98',0,'-',0),(157,'6767178-a13d5324f8d8cd98','png','2018-12-03 16:37:36','201812/',229979,'1543826256_6767178-a13d5324f8d8cd98',0,'-',0),(158,'6767178-a13d5324f8d8cd98','png','2018-12-03 17:33:59','201812/',229979,'1543826369_6767178-a13d5324f8d8cd98',47,'notice',0),(159,'Docker for Windows Installer','exe','2018-12-07 10:05:24','201812/',546318376,'1544148321_Docker for Windows Installer',0,'-',0),(401,'v2-91eeac568c7ff3891d8a5f7753e4e01c_hd','jpg','2018-12-24 18:28:07','201812/',27515,'1545647287_v2-91eeac568c7ff3891d8a5f7753e4e01c_hd',0,'-',0),(402,'v2-69638139a50869f6f9e457ce0d672203_hd','jpg','2018-12-24 23:05:06','201812/',23783,'1545663906_v2-69638139a50869f6f9e457ce0d672203_hd',0,'-',0),(403,'v2-69638139a50869f6f9e457ce0d672203_hd','jpg','2018-12-24 23:05:16','201812/',23783,'1545663916_v2-69638139a50869f6f9e457ce0d672203_hd',0,'-',0),(404,'v2-69638139a50869f6f9e457ce0d672203_hd','jpg','2018-12-25 17:07:32','201812/',23783,'1545728852_v2-69638139a50869f6f9e457ce0d672203_hd',0,'-',0),(405,'v2-69638139a50869f6f9e457ce0d672203_hd','jpg','2018-12-25 17:07:36','201812/',23783,'1545728856_v2-69638139a50869f6f9e457ce0d672203_hd',0,'-',0),(406,'v2-69638139a50869f6f9e457ce0d672203_hd','jpg','2018-12-25 18:05:08','201812/',23783,'1545732308_v2-69638139a50869f6f9e457ce0d672203_hd',0,'-',0),(407,'v2-91eeac568c7ff3891d8a5f7753e4e01c_hd','jpg','2018-12-25 18:05:13','201812/',27515,'1545732313_v2-91eeac568c7ff3891d8a5f7753e4e01c_hd',0,'-',0),(408,'v2-91eeac568c7ff3891d8a5f7753e4e01c_hd','jpg','2018-12-25 18:15:06','201812/',27515,'1545732906_v2-91eeac568c7ff3891d8a5f7753e4e01c_hd',0,'-',0),(409,'v2-91eeac568c7ff3891d8a5f7753e4e01c_hd','jpg','2018-12-25 18:15:18','201812/',27515,'1545732918_v2-91eeac568c7ff3891d8a5f7753e4e01c_hd',0,'-',0),(410,'v2-69638139a50869f6f9e457ce0d672203_hd','jpg','2018-12-25 18:15:21','201812/',23783,'1545732921_v2-69638139a50869f6f9e457ce0d672203_hd',0,'-',0),(411,'v2-91eeac568c7ff3891d8a5f7753e4e01c_hd','jpg','2018-12-25 18:18:43','201812/',27515,'1545733123_v2-91eeac568c7ff3891d8a5f7753e4e01c_hd',0,'-',0),(412,'v2-69638139a50869f6f9e457ce0d672203_hd','jpg','2018-12-25 18:18:56','201812/',23783,'1545733136_v2-69638139a50869f6f9e457ce0d672203_hd',0,'-',0),(413,'v2-69638139a50869f6f9e457ce0d672203_hd','jpg','2018-12-25 18:20:14','201812/',23783,'1545733214_v2-69638139a50869f6f9e457ce0d672203_hd',0,'-',0),(414,'v2-91eeac568c7ff3891d8a5f7753e4e01c_hd','jpg','2018-12-25 18:20:46','201812/',27515,'1545733246_v2-91eeac568c7ff3891d8a5f7753e4e01c_hd',0,'-',0),(415,'v2-91eeac568c7ff3891d8a5f7753e4e01c_hd','jpg','2018-12-26 10:56:01','201812/',27515,'1545792961_v2-91eeac568c7ff3891d8a5f7753e4e01c_hd',0,'-',0),(416,'v2-69638139a50869f6f9e457ce0d672203_hd','jpg','2018-12-26 10:56:07','201812/',23783,'1545792967_v2-69638139a50869f6f9e457ce0d672203_hd',0,'-',0),(417,'刘亦菲','gif','2018-12-26 10:57:05','201812/',5116027,'1545793025_刘亦菲',0,'-',0),(418,'刘亦菲','gif','2018-12-26 10:57:10','201812/',5116027,'1545793030_刘亦菲',0,'-',0),(419,'刘亦菲','gif','2018-12-26 11:00:13','201812/',5116027,'1545793213_刘亦菲',0,'-',0),(420,'刘亦菲','gif','2018-12-26 11:02:32','201812/',5116027,'1545793352_刘亦菲',0,'-',0),(421,'刘亦菲','gif','2018-12-26 11:04:35','201812/',5116027,'1545793475_刘亦菲',0,'-',0),(422,'刘亦菲','gif','2018-12-26 11:11:26','201812/',5116027,'1545793886_刘亦菲',0,'-',0),(423,'刘亦菲','gif','2018-12-26 11:12:31','201812/',5116027,'1545793951_刘亦菲',0,'-',0),(424,'刘亦菲','gif','2018-12-26 11:12:40','201812/',5116027,'1545793960_刘亦菲',0,'-',0),(425,'刘亦菲','gif','2018-12-26 11:13:41','201812/',5116027,'1545794021_刘亦菲',0,'-',0),(426,'刘亦菲','gif','2018-12-26 11:14:56','201812/',5116027,'1545794096_刘亦菲',0,'-',0),(427,'刘亦菲','gif','2018-12-26 11:15:08','201812/',5116027,'1545794108_刘亦菲',0,'-',0),(428,'刘亦菲','gif','2018-12-26 11:16:22','201812/',5116027,'1545794182_刘亦菲',0,'-',0),(429,'刘亦菲','gif','2018-12-26 11:16:37','201812/',5116027,'1545794197_刘亦菲',0,'-',0),(430,'刘亦菲','gif','2018-12-26 11:16:41','201812/',5116027,'1545794201_刘亦菲',0,'-',0),(431,'刘亦菲','gif','2018-12-26 11:18:23','201812/',5116027,'1545794303_刘亦菲',0,'-',0),(432,'刘亦菲','gif','2018-12-26 11:19:11','201812/',5116027,'1545794351_刘亦菲',0,'-',0),(433,'刘亦菲','gif','2018-12-26 11:19:59','201812/',5116027,'1545794399_刘亦菲',0,'-',0),(434,'刘亦菲','gif','2018-12-26 11:21:29','201812/',5116027,'1545794489_刘亦菲',0,'-',0),(435,'刘亦菲','gif','2018-12-26 11:24:23','201812/',5116027,'1545794663_刘亦菲',0,'-',0),(436,'刘亦菲','gif','2018-12-26 11:24:57','201812/',5116027,'1545794697_刘亦菲',0,'-',0),(437,'刘亦菲','gif','2018-12-26 11:25:36','201812/',5116027,'1545794736_刘亦菲',0,'-',0),(438,'刘亦菲','gif','2018-12-26 11:25:55','201812/',5116027,'1545794755_刘亦菲',0,'-',0),(439,'刘亦菲','gif','2018-12-26 11:26:30','201812/',5116027,'1545794790_刘亦菲',0,'-',0),(440,'刘亦菲','gif','2018-12-26 11:27:50','201812/',5116027,'1545794870_刘亦菲',0,'-',0),(441,'刘亦菲','gif','2018-12-26 11:28:31','201812/',5116027,'1545794911_刘亦菲',0,'-',0),(442,'刘亦菲','gif','2018-12-26 11:28:39','201812/',5116027,'1545794919_刘亦菲',0,'-',0),(443,'刘亦菲','gif','2018-12-26 11:28:52','201812/',5116027,'1545794932_刘亦菲',0,'-',0),(444,'刘亦菲','gif','2018-12-26 11:29:13','201812/',5116027,'1545794953_刘亦菲',0,'-',0),(445,'刘亦菲','gif','2018-12-26 11:29:22','201812/',5116027,'1545794962_刘亦菲',0,'-',0),(446,'刘亦菲','gif','2018-12-26 11:29:37','201812/',5116027,'1545794977_刘亦菲',0,'-',0),(447,'刘亦菲','gif','2018-12-26 11:30:53','201812/',5116027,'1545795053_刘亦菲',0,'-',0),(448,'刘亦菲','gif','2018-12-26 11:34:30','201812/',5116027,'1545795270_刘亦菲',0,'-',0),(449,'刘亦菲','gif','2018-12-26 11:36:02','201812/',5116027,'1545795362_刘亦菲',0,'-',0),(450,'刘亦菲','gif','2018-12-26 11:36:05','201812/',5116027,'1545795365_刘亦菲',0,'-',0),(451,'刘亦菲','gif','2018-12-26 11:36:31','201812/',5116027,'1545795391_刘亦菲',0,'-',0),(452,'刘亦菲','gif','2018-12-26 11:36:45','201812/',5116027,'1545795405_刘亦菲',0,'-',0),(453,'刘亦菲','gif','2018-12-26 11:38:55','201812/',5116027,'1545795535_刘亦菲',0,'-',0),(459,'刘亦菲','gif','2018-12-26 11:54:46','201812/',5116027,'1545796486_刘亦菲',0,'-',0),(460,'刘亦菲','gif','2018-12-26 11:56:12','201812/',5116027,'1545796572_刘亦菲',0,'-',0),(461,'刘亦菲','gif','2018-12-26 11:56:18','201812/',5116027,'1545796578_刘亦菲',0,'-',0),(466,'刘亦菲','gif','2018-12-26 12:01:13','201812/',5116027,'1545796873_刘亦菲',0,'-',0);

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
  `time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '日期',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

/*Data for the table `integral` */

insert  into `integral`(`id`,`cognizance_user_id`,`user_id`,`draft_id`,`punish_notice_id`,`receipt_id`,`score`,`time`,`delete`) values (3,2,14,70,11,0,1000,'2018-01-01 14:09:47',0),(4,2,14,70,9,0,2200,'2018-12-31 15:02:33',0);

/*Table structure for table `integral_edit` */

DROP TABLE IF EXISTS `integral_edit`;

CREATE TABLE `integral_edit` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '积分修改ID',
  `integral_id` int(11) NOT NULL DEFAULT '0' COMMENT '积分表ID',
  `score` int(11) DEFAULT NULL COMMENT '积分（除以100显示）',
  `user_id` int(11) DEFAULT NULL COMMENT '发起人修改ID',
  `describe` varchar(500) DEFAULT NULL COMMENT '修改分数原因',
  `suggestion` char(250) DEFAULT NULL COMMENT '审核意见',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `state` enum('draft','report','reject','adopt') DEFAULT 'draft' COMMENT '状态（draft草稿|report上报|reject驳回|adopt通过）',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

/*Data for the table `integral_edit` */

insert  into `integral_edit`(`id`,`integral_id`,`score`,`user_id`,`describe`,`suggestion`,`update_time`,`state`,`delete`) values (8,3,1100,NULL,'叮叮当当000',NULL,'2019-01-03 22:27:10','draft',0),(9,4,1300,NULL,'222222222',NULL,'2019-01-03 22:27:35','draft',0),(15,5,1100,NULL,'这是建议或者意见的内容1',NULL,'2019-01-03 22:51:51','draft',0);

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

insert  into `login`(`login_id`,`user_code`,`password`,`is_use`,`change_pd_time`,`login_num`,`login_time`,`author_id`,`delete`) values (3,10001,'3b1ac6e9a16ff76879e5888483e118a8',1,'2018-11-29 18:10:57',0,'2018-12-03 10:48:57',2,0),(4,10002,'78dc8bbc86eb472b3db1d0b025714ec1',1,'0000-00-00 00:00:00',111,'2019-01-04 11:57:45',2,0);

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
) ENGINE=MyISAM AUTO_INCREMENT=32465 DEFAULT CHARSET=utf8;

/*Data for the table `logs` */

insert  into `logs`(`id`,`url`,`user_id`,`msg`,`method`,`data`,`time`,`server`,`ip`,`delete`) values (32464,'/api/audit/punishNotice/get_accountability',2,'惩罚通知书问责情况','GET','confirmationId=7&rnd=0.6631875192082495','2019-01-04 15:06:21','audit','192.168.1.13',0),(32463,'/api/audit/rectifyReport/get',2,'获取整改报告','GET','id=41&rnd=0.9571857059574536','2019-01-04 15:06:21','audit','192.168.1.13',0),(32462,'/api/audit/confirmation/get',2,'获取事实确认书','GET','id=7&rnd=0.18815893848750687','2019-01-04 15:06:21','audit','192.168.1.13',0),(32461,'/api/audit/draft/get',2,'获取工作底稿','GET','id=70&rnd=0.16362381923326907','2019-01-04 15:06:21','audit','192.168.1.13',0),(32459,'/api/audit/auditReport/get',2,'获取','GET','id=1&rnd=0.10045631409195432','2019-01-04 15:06:21','audit','192.168.1.13',0),(32460,'/api/audit/programme/get',2,'获取审核方案','GET','id=78&rnd=0.07456154245646052','2019-01-04 15:06:21','audit','192.168.1.13',0);

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
  `time` timestamp NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新日期',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;

/*Data for the table `menu` */

insert  into `menu`(`id`,`path`,`name`,`title`,`icon`,`no_cache`,`parent_id`,`has_child`,`order`,`is_use`,`time`,`delete`) values (1,'','','','',1,-1,1,1,1,'2018-12-06 09:15:55',0),(2,'dashboard','Dashboard','dashboard','dashboard',1,1,0,0,1,'2018-12-06 09:16:41',0),(3,'/guide','','','',1,-1,1,2,1,'2018-12-06 09:17:06',0),(4,'index','Guide','guide','guide',1,3,0,0,0,'2018-12-06 09:17:44',0),(5,'/personal','','','',1,-1,1,3,1,'2018-12-06 09:18:09',0),(6,'index','Personal','personal','user',1,5,0,0,0,'2018-12-06 09:19:57',0),(7,'/organization','organization','organization','component',1,-1,1,4,1,'2018-12-06 09:20:43',0),(8,'notice','notice','notice','notice',1,7,0,1,1,'2018-12-06 09:21:10',0),(9,'departmentManagement','departmentManagement','departmentManagement','',1,7,0,2,1,'2018-12-06 09:21:31',0),(10,'personnelManagement','personnelManagement','personnelManagement','',1,7,0,3,1,'2018-12-06 09:21:41',0),(11,'managementMethods','managementMethods','managementMethods','',1,7,0,4,1,'2018-12-06 09:21:47',0),(12,'/audit','audit','audit','international',1,-1,1,5,1,'2018-12-06 09:22:10',0),(13,'workManuscript','workManuscript','workManuscript','',1,12,0,1,1,'2018-12-06 09:23:00',0),(14,'confirmation','confirmation','confirmation','',1,12,0,2,1,'2018-12-06 09:23:08',0),(15,'punishNotice','punishNotice','punishNotice','',1,12,0,3,1,'2018-12-06 09:23:16',0),(16,'integralTable','integralTable','integralTable','',1,12,0,4,1,'2018-12-06 09:23:25',0),(17,'statisticalAnalysis','statisticalAnalysis','statisticalAnalysis','',1,12,0,5,1,'2018-12-06 09:23:31',0),(18,'/system','system','system','example',1,-1,1,6,1,'2018-12-06 09:23:47',0),(19,'dictionaryManagement','dictionaryManagement','dictionaryManagement','',1,18,0,1,1,'2018-12-06 09:25:58',0),(20,'loginManagement','loginManagement','loginManagement','',1,18,0,2,1,'2018-12-06 09:26:06',0),(21,'menusManagement','menusManagement','menusManagement','',1,18,0,3,1,'2018-12-06 09:26:19',0),(22,'powerManagement','powerManagement','powerManagement','',1,18,0,4,1,'2018-12-06 09:26:27',0),(23,'systemLog','systemLog','systemLog','',1,18,0,5,1,'2018-12-06 09:26:33',0);

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
) ENGINE=MyISAM AUTO_INCREMENT=206 DEFAULT CHARSET=utf8;

/*Data for the table `notice` */

insert  into `notice`(`id`,`department_id`,`title`,`content`,`time`,`range`,`state`,`delete`) values (18,1,'公告18','公告内容18','2018-12-25 14:02:02','2','publish',0),(45,1,'测试标题','<p>测试内容</p>','2018-11-30 09:59:30','1','draft',0),(47,1,'测试标题1','<p>这是测试内容</p>','2018-12-20 10:33:39','1','publish',0);

/*Table structure for table `notice_file` */

DROP TABLE IF EXISTS `notice_file`;

CREATE TABLE `notice_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `file_id` int(11) DEFAULT NULL COMMENT '通知附件ID',
  `notice_id` int(11) NOT NULL COMMENT '通知ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`,`notice_id`)
) ENGINE=MyISAM AUTO_INCREMENT=46 DEFAULT CHARSET=utf8;

/*Data for the table `notice_file` */

insert  into `notice_file`(`id`,`file_id`,`notice_id`,`delete`) values (1,1,3,0),(2,2,3,0),(3,3,3,0),(4,1,4,0),(5,2,4,0),(6,3,4,0),(7,1,13,0),(8,2,13,0),(9,3,13,0),(10,1,14,0),(11,2,14,0),(12,3,14,0),(13,1,16,0),(14,2,16,0),(16,56,40,0),(17,57,40,0),(18,58,40,0),(19,56,41,0),(20,57,41,0),(21,58,41,0),(22,56,42,0),(23,57,42,0),(24,58,42,0),(25,59,43,0),(26,60,43,0),(29,61,18,0),(30,71,45,0),(31,72,45,0),(45,158,47,0),(44,146,47,0),(43,145,47,0);

/*Table structure for table `notice_inform` */

DROP TABLE IF EXISTS `notice_inform`;

CREATE TABLE `notice_inform` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `department_id` int(11) NOT NULL COMMENT '机构ID',
  `notice_id` int(11) NOT NULL COMMENT '公告ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`,`notice_id`)
) ENGINE=MyISAM AUTO_INCREMENT=513 DEFAULT CHARSET=utf8;

/*Data for the table `notice_inform` */

insert  into `notice_inform`(`id`,`department_id`,`notice_id`,`delete`) values (1,1,3,0),(2,2,3,0),(3,3,3,0),(4,4,3,0),(5,1,4,0),(6,2,4,0),(7,3,4,0),(8,4,4,0),(9,2,13,0),(10,23,13,0),(11,5,13,0),(12,5,14,0),(512,5,18,0),(511,2,18,0),(510,1,18,0),(18,1,19,0),(22,1,20,0),(23,1,21,0),(24,1,22,0),(25,1,23,0),(26,1,24,0),(38,1,30,0),(37,1,29,0),(29,1,27,0),(39,1,40,0),(40,1,41,0),(41,1,42,0),(42,1,43,0),(52,1,50,0),(57,59,55,0),(58,62,56,0),(59,61,56,0),(60,60,56,0),(61,57,57,0),(62,67,57,0),(63,66,57,0),(64,65,57,0),(65,64,57,0),(66,63,57,0),(67,62,57,0),(68,61,57,0),(69,60,57,0),(70,59,57,0),(71,58,57,0);

/*Table structure for table `programme` */

DROP TABLE IF EXISTS `programme`;

CREATE TABLE `programme` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '方案ID',
  `title` char(200) DEFAULT NULL COMMENT '方案名称',
  `key` char(20) DEFAULT NULL COMMENT '方案分类',
  `query_department_id` int(11) NOT NULL DEFAULT '0' COMMENT '被稽核审计机构ID',
  `user_id` int(11) DEFAULT '0' COMMENT '创建方案的人员ID',
  `query_point_id` int(11) NOT NULL DEFAULT '0' COMMENT '被稽核审计网点ID',
  `purpose` varchar(500) DEFAULT NULL COMMENT '稽核目的',
  `type` char(20) DEFAULT NULL COMMENT '稽核审计方式',
  `start_time` timestamp NULL DEFAULT NULL COMMENT '稽核审计开始时间',
  `end_time` timestamp NULL DEFAULT NULL COMMENT '稽核审计结束时间',
  `plan_start_time` timestamp NULL DEFAULT NULL COMMENT '方案计划工作开始时间',
  `plan_end_time` timestamp NULL DEFAULT NULL COMMENT '方案计划工作结束时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `state` enum('draft','report','dep_reject','dep_adopt','admin_reject','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|report上报|publish发布|adopt通过|reject驳回【dep_部门负责人前缀|admin_分管领导前缀】）',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8;

/*Data for the table `programme` */

insert  into `programme`(`id`,`title`,`key`,`query_department_id`,`user_id`,`query_point_id`,`purpose`,`type`,`start_time`,`end_time`,`plan_start_time`,`plan_end_time`,`update_time`,`state`,`delete`) values (78,'测试审计方案一','auditKey3',1,2,3,'测试审计方案一目的','auditType1','2018-10-10 10:10:10','2018-10-10 11:10:10','2018-10-12 10:10:10','2018-10-13 10:10:10','2018-12-28 10:31:03','publish',0),(79,'测试审计方案二','auditKey1',1,2,3,'测试审计方案一目的','auditType1','2018-10-10 10:10:10','2018-10-10 11:10:10','2018-10-12 10:10:10','2018-10-13 10:10:10','2018-12-29 11:42:41','draft',0),(80,'测试审计方案三️','auditKey3',1,2,3,'测试审计方案一目的','auditType3','2018-10-10 10:10:10','2018-10-10 11:10:10','2018-10-12 10:10:10','2018-10-13 10:10:10','2018-12-29 11:42:34','draft',0),(81,'测试审计方案四️','auditKey2',1,2,3,'测试审计方案一目的','auditType2','2018-10-10 10:10:10','2018-10-10 11:10:10','2018-10-12 10:10:10','2018-10-13 10:10:10','2018-12-29 11:42:26','draft',0),(82,'测试审计方案五️','auditKey2',1,2,3,'测试审计方案一目的','auditType2','2018-10-10 10:10:10','2018-10-10 11:10:10','2018-10-12 10:10:10','2018-10-13 10:10:10','2018-12-29 11:42:16','draft',0),(88,'测试审计方案五️','auditKey2',1,2,3,'测试审计方案一目的','auditType3','2018-10-10 10:10:10','2018-10-10 11:10:10','2018-10-12 10:10:10','2018-10-13 10:10:10','2019-01-04 11:07:16','draft',0),(89,'测试标题1-1-1','',0,0,0,'','','0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00','0000-00-00 00:00:00','2019-01-03 15:10:13','publish',0);

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
) ENGINE=InnoDB AUTO_INCREMENT=155 DEFAULT CHARSET=utf8;

/*Data for the table `programme_basis` */

insert  into `programme_basis`(`id`,`programme_id`,`clause_id`,`content`,`order`,`delete`) values (137,78,1,'这是依据内容，自己填写的',1,0),(138,78,2,'安顺市信用合作联社员工 违规行为积分管理实施细则',2,0),(139,78,0,'普定县农村信用合作联社员工 违规行为积分管理实施细则',3,0),(140,79,1,'这是依据内容，自己填写的',1,0),(141,79,2,'安顺市信用合作联社员工 违规行为积分管理实施细则',2,0),(142,79,0,'普定县农村信用合作联社员工 违规行为积分管理实施细则',3,0),(143,80,1,'这是依据内容，自己填写的',1,0),(144,80,2,'安顺市信用合作联社员工 违规行为积分管理实施细则',2,0),(145,80,0,'普定县农村信用合作联社员工 违规行为积分管理实施细则',3,0),(146,81,1,'这是依据内容，自己填写的',1,0),(147,81,2,'安顺市信用合作联社员工 违规行为积分管理实施细则',2,0),(148,81,0,'普定县农村信用合作联社员工 违规行为积分管理实施细则',3,0),(149,82,1,'这是依据内容，自己填写的',1,0),(150,82,2,'安顺市信用合作联社员工 违规行为积分管理实施细则',2,0),(151,82,0,'普定县农村信用合作联社员工 违规行为积分管理实施细则',3,0),(152,88,1,'这是依据内容，自己填写的',1,0),(153,88,2,'安顺市信用合作联社员工 违规行为积分管理实施细则',2,0),(154,88,0,'普定县农村信用合作联社员工 违规行为积分管理实施细则',3,0);

/*Table structure for table `programme_business` */

DROP TABLE IF EXISTS `programme_business`;

CREATE TABLE `programme_business` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '工作方案业务范围ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '稽核方案ID',
  `content` char(250) DEFAULT NULL COMMENT '内容',
  `order` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=123 DEFAULT CHARSET=utf8;

/*Data for the table `programme_business` */

insert  into `programme_business`(`id`,`programme_id`,`content`,`order`,`delete`) values (111,78,'方案业务范围1',1,0),(112,78,'方案业务范围2',2,0),(113,79,'方案业务范围1',1,0),(114,79,'方案业务范围2',2,0),(115,80,'方案业务范围1',1,0),(116,80,'方案业务范围2',2,0),(117,81,'方案业务范围1',1,0),(118,81,'方案业务范围2',2,0),(119,82,'方案业务范围1',1,0),(120,82,'方案业务范围2',2,0),(121,88,'方案业务范围1',1,0),(122,88,'方案业务范围2',2,0);

/*Table structure for table `programme_content` */

DROP TABLE IF EXISTS `programme_content`;

CREATE TABLE `programme_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '工作方案主要内容ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '稽核方案ID',
  `content` varchar(500) DEFAULT NULL COMMENT '主要内容',
  `order` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未除|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=135 DEFAULT CHARSET=utf8;

/*Data for the table `programme_content` */

insert  into `programme_content`(`id`,`programme_id`,`content`,`order`,`delete`) values (123,78,'这是方案主要内容1',1,0),(124,78,'这是方案主要内容2',2,0),(125,79,'这是方案主要内容1',1,0),(126,79,'这是方案主要内容2',2,0),(127,80,'这是方案主要内容1',1,0),(128,80,'这是方案主要内容2',2,0),(129,81,'这是方案主要内容1',1,0),(130,81,'这是方案主要内容2',2,0),(131,82,'这是方案主要内容1',1,0),(132,82,'这是方案主要内容2',2,0),(133,88,'这是方案主要内容1',1,0),(134,88,'这是方案主要内容2',2,0);

/*Table structure for table `programme_emphases` */

DROP TABLE IF EXISTS `programme_emphases`;

CREATE TABLE `programme_emphases` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '工作方案重点ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '稽核方案ID',
  `content` char(250) DEFAULT NULL COMMENT '内容',
  `order` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=93 DEFAULT CHARSET=utf8;

/*Data for the table `programme_emphases` */

insert  into `programme_emphases`(`id`,`programme_id`,`content`,`order`,`delete`) values (81,78,'方案重点1',1,0),(82,78,'方案重点2',2,0),(83,79,'方案重点1',1,0),(84,79,'方案重点2',2,0),(85,80,'方案重点1',1,0),(86,80,'方案重点2',2,0),(87,81,'方案重点1',1,0),(88,81,'方案重点2',2,0),(89,82,'方案重点1',1,0),(90,82,'方案重点2',2,0),(91,88,'方案重点1',1,0),(92,88,'方案重点2',2,0);

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

/*Data for the table `programme_examine_admin` */

insert  into `programme_examine_admin`(`id`,`programme_id`,`user_id`,`content`,`time`,`state`,`file_id`,`delete`) values (1,78,NULL,'没问题的','2018-12-28 10:31:32','adopt',NULL,0),(2,89,NULL,'','2019-01-03 15:10:32','adopt',NULL,0);

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

/*Data for the table `programme_examine_dep` */

insert  into `programme_examine_dep`(`id`,`programme_id`,`user_id`,`content`,`time`,`state`,`file_id`,`delete`) values (1,78,NULL,'好样的','2018-12-28 10:31:19','adopt',NULL,0),(2,89,NULL,'','2019-01-03 15:10:21','adopt',NULL,0);

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
) ENGINE=InnoDB AUTO_INCREMENT=172 DEFAULT CHARSET=utf8;

/*Data for the table `programme_step` */

insert  into `programme_step`(`id`,`programme_id`,`type`,`content`,`order`,`delete`) values (129,78,'title','这是方案实施步骤标题1',1,0),(130,78,'content','这是方案实施步骤内容1',2,0),(131,78,'content','这是方案实施步骤内容2',3,0),(132,79,'title','这是方案实施步骤标题1',1,0),(133,79,'content','这是方案实施步骤内容1',2,0),(134,79,'content','这是方案实施步骤内容2',3,0),(135,80,'title','这是方案实施步骤标题1',1,0),(136,80,'content','这是方案实施步骤内容1',2,0),(137,80,'content','这是方案实施步骤内容2',3,0),(138,81,'title','这是方案实施步骤标题1',1,0),(139,81,'content','这是方案实施步骤内容1',2,0),(140,81,'content','这是方案实施步骤内容2',3,0),(141,82,'title','这是方案实施步骤标题1',1,0),(142,82,'content','这是方案实施步骤内容1',2,0),(143,82,'content','这是方案实施步骤内容2',3,0),(149,89,'step','这是测试步骤1-1-1',3,0),(151,89,'step','这是测试步骤内容2-1',5,0),(152,89,'step','这是测试步骤内容2-2',6,0),(153,89,'title','这是测试步骤标题1',1,0),(154,89,'content','这是测试步骤内容1-1',2,0),(155,89,'content','这是测试步骤内容2',4,0),(156,89,'title','这是测试步骤标题2',7,0),(157,89,'content','这是测试步骤内容2-1',8,0),(158,89,'step','这是测试步骤2-1-1',9,0),(159,89,'step','这是测试步骤2-2-1',10,0),(160,89,'content','这是测试步骤内容2-2',11,0),(161,89,'content','这是测试步骤内容2-3',12,0),(162,89,'step','这是测试步骤2-3-1',13,0),(169,88,'title','这是方案实施步骤标题1',1,0),(170,88,'content','这是方案实施步骤内容1',2,0),(171,88,'content','这是方案实施步骤内容2',3,0);

/*Table structure for table `programme_user` */

DROP TABLE IF EXISTS `programme_user`;

CREATE TABLE `programme_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '工作方案主要内容ID',
  `programme_id` int(11) NOT NULL DEFAULT '0' COMMENT '稽核方案ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '员工ID',
  `job` char(20) NOT NULL DEFAULT '-' COMMENT '员工行政职务',
  `title` char(20) NOT NULL DEFAULT '-' COMMENT '员工技术职称',
  `task` varchar(500) DEFAULT NULL COMMENT '员工分工',
  `order` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未除|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=134 DEFAULT CHARSET=utf8;

/*Data for the table `programme_user` */

insert  into `programme_user`(`id`,`programme_id`,`user_id`,`job`,`title`,`task`,`order`,`delete`) values (116,78,1,'人员职务1','人员技术职称1','员工分工1',1,0),(117,78,2,'人员职务2','人员技术职称2','员工分工2',2,0),(118,79,1,'人员职务1','人员技术职称1','员工分工1',1,0),(119,79,2,'人员职务2','人员技术职称2','员工分工2',2,0),(120,80,1,'人员职务1','人员技术职称1','员工分工1',1,0),(121,80,2,'人员职务2','人员技术职称2','员工分工2',2,0),(122,81,1,'人员职务1','人员技术职称1','员工分工1',1,0),(123,81,2,'人员职务2','人员技术职称2','员工分工2',2,0),(124,82,1,'人员职务1','人员技术职称1','员工分工1',1,0),(125,82,2,'人员职务2','人员技术职称2','员工分工2',2,0),(131,88,1,'人员职务1','人员技术职称1','员工分工1',1,0),(132,88,2,'人员职务2','人员技术职称2','员工分工2',2,0),(133,89,0,'','','',1,0);

/*Table structure for table `punish_notice` */

DROP TABLE IF EXISTS `punish_notice`;

CREATE TABLE `punish_notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '处罚通知ID',
  `confirmation_id` int(11) NOT NULL DEFAULT '0' COMMENT '确认书ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '被惩罚人ID',
  `number` char(50) DEFAULT NULL COMMENT '文件号',
  `time` timestamp NULL DEFAULT NULL COMMENT '生成时间',
  `state` enum('draft','jh_draft','jh_publish','ld_draft','ld_publish','bgs_draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|publish发布）',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

/*Data for the table `punish_notice` */

insert  into `punish_notice`(`id`,`confirmation_id`,`draft_id`,`user_id`,`number`,`time`,`state`,`delete`) values (8,7,70,15,NULL,'2018-12-29 09:32:38','draft',0),(10,7,70,13,NULL,'2018-12-29 09:32:40','jh_draft',0),(11,7,70,14,'acx123245','2018-12-29 09:32:44','publish',0);

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
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

/*Data for the table `punish_notice_behavior` */

insert  into `punish_notice_behavior`(`id`,`punish_notice_id`,`user_id`,`behavior_id`,`content`,`update_time`,`delete`) values (3,11,2,0,'违规行为1-0','2018-12-29 13:09:54',0),(9,11,2,0,'违规行为2','2018-12-29 13:09:54',0),(10,10,2,0,'','2019-01-01 22:49:29',0);

/*Table structure for table `punish_notice_score` */

DROP TABLE IF EXISTS `punish_notice_score`;

CREATE TABLE `punish_notice_score` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '处罚分数ID',
  `cognizance_user_id` int(11) NOT NULL DEFAULT '0' COMMENT '处罚人员ID',
  `punish_notice_id` int(11) NOT NULL DEFAULT '0' COMMENT '处罚通知ID',
  `score` int(5) NOT NULL DEFAULT '0' COMMENT '处罚分数(填写的数*1000保存)',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '处罚时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

/*Data for the table `punish_notice_score` */

insert  into `punish_notice_score`(`id`,`cognizance_user_id`,`punish_notice_id`,`score`,`update_time`,`delete`) values (4,2,8,1000,'2018-12-29 09:37:38',0),(12,2,11,1000,'2018-12-29 14:09:47',0);

/*Table structure for table `rbac` */

DROP TABLE IF EXISTS `rbac`;

CREATE TABLE `rbac` (
  `rid` int(11) NOT NULL AUTO_INCREMENT COMMENT '权限主键ID',
  `key` char(20) NOT NULL DEFAULT '-' COMMENT '角色key（字典）',
  `menu_id` int(11) NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `is_read` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否可读',
  `is_write` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否可写',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`rid`)
) ENGINE=MyISAM AUTO_INCREMENT=130 DEFAULT CHARSET=utf8;

/*Data for the table `rbac` */

insert  into `rbac`(`rid`,`key`,`menu_id`,`is_read`,`is_write`,`delete`) values (109,'admin',23,0,0,0),(89,'management',23,0,0,0),(88,'management',22,0,0,0),(87,'management',21,0,0,0),(108,'admin',22,0,0,0),(107,'admin',21,0,0,0),(106,'admin',20,0,0,0),(105,'admin',19,0,0,0),(104,'admin',18,0,0,0),(103,'admin',17,0,0,0),(102,'admin',16,0,0,0),(101,'admin',15,0,0,0),(100,'admin',14,0,0,0),(99,'admin',13,0,0,0),(98,'admin',12,0,0,0),(97,'admin',11,1,0,0),(96,'admin',10,1,0,0),(95,'admin',9,1,0,0),(94,'admin',8,1,0,0),(93,'admin',7,1,0,0),(92,'admin',5,0,0,0),(91,'admin',3,0,0,0),(90,'admin',1,1,1,0),(86,'management',20,0,0,0),(85,'management',19,0,0,0),(84,'management',18,0,0,0),(83,'management',17,0,0,0),(82,'management',16,0,0,0),(81,'management',15,0,0,0),(80,'management',14,0,0,0),(79,'management',13,0,0,0),(78,'management',12,0,0,0),(77,'management',11,0,0,0),(76,'management',10,0,0,0),(75,'management',9,0,0,0),(74,'management',8,0,0,0),(73,'management',7,0,0,0),(72,'management',5,0,0,0),(71,'management',3,1,1,0),(70,'management',1,1,1,0),(110,'staff',1,0,0,0),(111,'staff',3,1,1,0),(112,'staff',5,1,1,0),(113,'staff',7,1,1,0),(114,'staff',8,1,1,0),(115,'staff',9,1,1,0),(116,'staff',10,1,1,0),(117,'staff',11,1,1,0),(118,'staff',12,0,0,0),(119,'staff',13,0,0,0),(120,'staff',14,0,0,0),(121,'staff',15,0,0,0),(122,'staff',16,0,0,0),(123,'staff',17,0,0,0),(124,'staff',18,0,0,0),(125,'staff',19,0,0,0),(126,'staff',20,0,0,0),(127,'staff',21,0,0,0),(128,'staff',22,0,0,0),(129,'staff',23,0,0,0);

/*Table structure for table `rectify` */

DROP TABLE IF EXISTS `rectify`;

CREATE TABLE `rectify` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改通知主键ID',
  `draft_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿ID',
  `confirmation_id` int(11) NOT NULL DEFAULT '0' COMMENT '事实确认书ID',
  `suggest` varchar(10000) DEFAULT NULL COMMENT '整改建议或意见',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '签署人ID',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新日期',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|publish发布）',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1删除）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

/*Data for the table `rectify` */

insert  into `rectify`(`id`,`draft_id`,`confirmation_id`,`suggest`,`user_id`,`update_time`,`state`,`delete`) values (5,70,7,'这是建议或者意见的内容1',2,'2018-12-29 18:10:52','publish',0);

/*Table structure for table `rectify_report` */

DROP TABLE IF EXISTS `rectify_report`;

CREATE TABLE `rectify_report` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改报告ID',
  `rectify_id` int(11) NOT NULL DEFAULT '0' COMMENT '整改通知ID',
  `state` enum('draft','publish') NOT NULL DEFAULT 'draft' COMMENT '状态（draft草稿|publish发布）',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8;

/*Data for the table `rectify_report` */

insert  into `rectify_report`(`id`,`rectify_id`,`state`,`update_time`,`delete`) values (41,5,'publish','2019-01-04 09:16:30',0);

/*Table structure for table `rectify_report_content` */

DROP TABLE IF EXISTS `rectify_report_content`;

CREATE TABLE `rectify_report_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改详情ID',
  `rectify_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '整改报告ID',
  `draft_content_id` int(11) NOT NULL DEFAULT '0' COMMENT '工作底稿违规内容ID',
  `content` varchar(500) DEFAULT NULL COMMENT '整改详情',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8;

/*Data for the table `rectify_report_content` */

insert  into `rectify_report_content`(`id`,`rectify_report_id`,`draft_content_id`,`content`,`delete`) values (4,5,1,'整改内容1',1),(5,6,1,'整改内容1',1),(6,7,130,'整改内容1',1),(7,7,131,'整改内容2',1),(8,8,130,'整改内容1',1),(9,8,131,'整改内容2',1),(10,9,130,'整改内容1',1),(11,9,131,'整改内容2',1),(12,10,130,'整改内容1',1),(13,10,131,'整改内容2',1),(14,11,130,'整改内容1',1),(15,11,131,'整改内容2',1),(16,12,130,'整改内容1',1),(17,12,131,'整改内容2',1),(18,13,130,'整改内容1',1),(19,13,131,'整改内容2',1),(20,14,130,'整改内容1',1),(21,14,131,'整改内容2',1),(22,15,130,'整改内容1',1),(23,15,131,'整改内容2',1),(24,16,130,'整改内容1',1),(25,16,131,'整改内容2',1),(26,17,130,'整改内容1',1),(27,17,131,'整改内容2',1),(28,18,130,'整改内容1',1),(29,18,131,'整改内容2',1),(32,20,130,'整改内容1',1),(33,20,131,'整改内容2',1),(34,21,130,'整改内容1',1),(35,21,131,'整改内容2',1),(36,22,130,'整改内容1',1),(37,22,131,'整改内容2',1),(38,23,130,'整改内容1',1),(39,23,131,'整改内容2',1),(40,24,130,'整改内容1',1),(41,24,131,'整改内容2',1),(42,25,130,'整改内容1',1),(43,25,131,'整改内容2',1),(44,26,130,'整改内容1',1),(45,26,131,'整改内容2',1),(46,27,130,'整改内容1',1),(47,27,131,'整改内容2',1),(48,28,130,'整改内容1',1),(49,28,131,'整改内容2',1),(50,29,0,'1',1),(51,29,0,'2',1),(52,29,0,'3',1),(53,29,0,'4',1),(54,29,0,'5',1),(55,30,0,'1',1),(56,30,0,'2',1),(57,30,0,'3',1),(58,30,0,'4',1),(59,30,0,'5',1),(60,31,130,'整改内容1',1),(61,31,131,'整改内容2',1),(62,34,130,'整改内容1',1),(63,34,131,'整改内容2',1),(64,35,130,'整改内容1',1),(65,35,131,'整改内容2',1),(66,36,130,'整改内容1',1),(67,36,131,'整改内容2',1),(68,36,132,'3',1),(69,37,130,'整改内容1',1),(70,37,131,'整改内容2',1),(71,37,134,'4',1),(72,38,130,'整改内容1',1),(73,38,131,'整改内容2',1),(74,39,130,'整改内容1',1),(75,39,131,'整改内容2',1),(78,41,130,'整改内容1',0),(79,41,131,'整改内容2',0);

/*Table structure for table `rectify_report_file` */

DROP TABLE IF EXISTS `rectify_report_file`;

CREATE TABLE `rectify_report_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '整改附件ID',
  `file_id` int(11) NOT NULL DEFAULT '0' COMMENT '附件ID',
  `rectify_report_id` int(11) NOT NULL DEFAULT '0' COMMENT '整改报告ID',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '删除（0未删|1已删）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8;

/*Data for the table `rectify_report_file` */

insert  into `rectify_report_file`(`id`,`file_id`,`rectify_report_id`,`delete`) values (26,473,26,1),(27,472,26,1),(28,473,27,1),(29,472,27,1),(30,473,28,1),(31,472,28,1),(32,473,31,1),(33,472,31,1),(34,473,34,1),(35,472,34,1),(36,473,39,1),(37,472,39,1),(40,473,41,0),(41,472,41,0);

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
  `user_code` char(20) NOT NULL COMMENT '员工号',
  `class` char(20) DEFAULT NULL COMMENT '民族',
  `sex` enum('0','1','2') NOT NULL DEFAULT '0' COMMENT '性别（0保密1女2男）',
  `id_card` char(18) DEFAULT NULL COMMENT '身份证',
  `update_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`user_id`)
) ENGINE=MyISAM AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

/*Data for the table `users` */

insert  into `users`(`user_id`,`department_id`,`user_name`,`user_code`,`class`,`sex`,`id_card`,`update_time`,`delete`) values (1,66,'小明','10001','汉','0','522623-----------','2018-12-28 10:47:44',0),(2,65,'小王','10002','汉','0','123456789013554','2018-12-28 10:47:01',0),(11,1,'测试','10086','苗','0','56541321564','2018-11-26 17:06:32',0),(15,1,'刘德华','561641','汉','0','41461386995252','2018-11-30 10:41:37',0),(14,1,'梁朝伟','10023146','汉','0','564841318411','2018-11-30 10:40:41',0),(13,1,'张国荣','10234','汉','0','5649825134546','2018-11-29 15:53:22',0),(16,58,'张学友','2131446','汉','0','318656461315941','2018-12-21 11:17:43',0);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
