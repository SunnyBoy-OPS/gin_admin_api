-- MySQL dump 10.13  Distrib 8.0.44, for Linux (x86_64)
--
-- Host: 10.0.0.105    Database: ginblog
-- ------------------------------------------------------
-- Server version	8.0.44

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `ginblog`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `ginblog` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `ginblog`;

--
-- Table structure for table `Department`
--

DROP TABLE IF EXISTS `Department`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Department` (
  `DepartmentId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `DepartmentName` varchar(10) NOT NULL COMMENT '部门名称',
  `CreateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`DepartmentId`),
  UNIQUE KEY `DepartmentName` (`DepartmentName`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='部门表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Department`
--

LOCK TABLES `Department` WRITE;
/*!40000 ALTER TABLE `Department` DISABLE KEYS */;
INSERT INTO `Department` VALUES (1,'运维部','2026-01-03 20:23:51');
/*!40000 ALTER TABLE `Department` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `User` (
  `Id` int unsigned NOT NULL AUTO_INCREMENT,
  `UserName` varchar(20) NOT NULL,
  `UserPassword` varchar(20) NOT NULL,
  `CreateTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `Name` varchar(20) NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `User_pk` (`UserName`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (1,'admin','admin','2026-01-01 20:09:50',''),(2,'wjs','wjs','2026-01-01 20:09:50','');
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys.UserInfo`
--

DROP TABLE IF EXISTS `sys.UserInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys.UserInfo` (
  `Id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `UserName` varchar(20) NOT NULL COMMENT '账号',
  `UserPassword` varchar(20) NOT NULL COMMENT '密码',
  `NickNane` varchar(20) NOT NULL COMMENT '姓名',
  `Icon` varchar(40) DEFAULT NULL COMMENT '头像',
  `Sex` int NOT NULL DEFAULT '1' COMMENT '性别(1->男,n->女)',
  `Email` varchar(30) DEFAULT NULL COMMENT '邮箱',
  `Status` int NOT NULL DEFAULT '1' COMMENT '账号启用状态(1->启用,n->停止)',
  `DepartmentId` int NOT NULL COMMENT '部门ID',
  `CreateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UserName` (`UserName`),
  UNIQUE KEY `DepartmentId` (`DepartmentId`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys.UserInfo`
--

LOCK TABLES `sys.UserInfo` WRITE;
/*!40000 ALTER TABLE `sys.UserInfo` DISABLE KEYS */;
INSERT INTO `sys.UserInfo` VALUES (1,'admin','admin','wjs',NULL,1,'17671050860@163.com',1,1,'2026-01-03 20:24:31');
/*!40000 ALTER TABLE `sys.UserInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_test_models`
--

DROP TABLE IF EXISTS `user_test_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_test_models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(10) DEFAULT NULL,
  `user_pass` longtext,
  `creat_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_test_models_user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_test_models`
--

LOCK TABLES `user_test_models` WRITE;
/*!40000 ALTER TABLE `user_test_models` DISABLE KEYS */;
INSERT INTO `user_test_models` VALUES (4,'wjs','123456',NULL,'吴佳顺','2026-01-21 23:31:42.343',NULL),(8,'root','root',NULL,'管理员','2026-01-21 23:37:40.851',NULL),(21,'admin','admin',NULL,'1','2026-01-24 00:33:56.000',NULL),(22,'admin1','admin1',NULL,'2','2026-01-24 02:10:57.000',NULL),(23,'admin2','admin2',NULL,'3','2026-01-24 02:15:29.000',NULL);
/*!40000 ALTER TABLE `user_test_models` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-03-09 23:11:46
