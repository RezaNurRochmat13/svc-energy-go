-- MySQL dump 10.13  Distrib 5.7.27, for Linux (x86_64)
--
-- Host: localhost    Database: energy
-- ------------------------------------------------------
-- Server version	5.7.27-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `interval_block`
--

DROP TABLE IF EXISTS `interval_block`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `interval_block` (
  `interval_block_id` int(11) NOT NULL AUTO_INCREMENT,
  `interval_readin_id` varchar(45) DEFAULT NULL,
  `end_time` varchar(45) DEFAULT NULL,
  `interval_length` varchar(45) DEFAULT NULL,
  `interval_value` varchar(45) DEFAULT NULL,
  `interval_flags` varchar(45) DEFAULT NULL,
  `collection_time` varchar(45) DEFAULT NULL,
  `meter_code` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`interval_block_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `interval_block`
--

LOCK TABLES `interval_block` WRITE;
/*!40000 ALTER TABLE `interval_block` DISABLE KEYS */;
INSERT INTO `interval_block` VALUES (1,'P15MIN_1-0:32.24.0.255','2017-12-22T18:00:00-08:00','3600','5','0','2017-12-23T00:00:00-08:00','SAG654647'),(5,'P15MIN_1-0:32.24.0.256','2019-08-06','4','6','2','2019-08-06','SGV-187747892');
/*!40000 ALTER TABLE `interval_block` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meter`
--

DROP TABLE IF EXISTS `meter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meter` (
  `meter_id` int(11) NOT NULL AUTO_INCREMENT,
  `meter_code` varchar(45) DEFAULT NULL,
  `meter_type` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`meter_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meter`
--

LOCK TABLES `meter` WRITE;
/*!40000 ALTER TABLE `meter` DISABLE KEYS */;
INSERT INTO `meter` VALUES (1,'SAG654647','METER_X_UDC_ASSET_ID'),(6,'SGV-187747892','Meter SGV 187747892');
/*!40000 ALTER TABLE `meter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meter_information`
--

DROP TABLE IF EXISTS `meter_information`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meter_information` (
  `meter_information_id` int(11) NOT NULL AUTO_INCREMENT,
  `verb` varchar(45) DEFAULT NULL,
  `noun` varchar(45) DEFAULT NULL,
  `revision` varchar(45) DEFAULT NULL,
  `date_time` varchar(100) DEFAULT NULL,
  `source` varchar(45) DEFAULT NULL,
  `meter_code` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`meter_information_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meter_information`
--

LOCK TABLES `meter_information` WRITE;
/*!40000 ALTER TABLE `meter_information` DISABLE KEYS */;
INSERT INTO `meter_information` VALUES (1,'reply','MeterReads','1','2017-12-23T00:31:01-08:00','SICONIA','SAG654647'),(2,'reply','MeterReads','1','2017-12-23T00:31:01-08:00','SICONIA','SAG654647');
/*!40000 ALTER TABLE `meter_information` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-08-06  9:28:06
