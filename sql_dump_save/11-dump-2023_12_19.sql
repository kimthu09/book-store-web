-- MySQL dump 10.13  Distrib 8.0.34, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: bookstoremanagement
-- ------------------------------------------------------
-- Server version	8.0.33

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
-- Table structure for table `Author`
--

DROP TABLE IF EXISTS `Author`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Author` (
  `id` varchar(12) NOT NULL,
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Author`
--

LOCK TABLES `Author` WRITE;
INSERT INTO `Author` VALUES ('tgak','Adam Khoo','2023-12-02 01:51:49','2023-12-02 01:51:49',NULL,1),('tghat','Hồ Anh Thái','2023-12-19 00:26:41','2023-12-19 00:26:41',NULL,1),('tgic','Iris Cao','2023-12-19 00:26:41','2023-12-19 00:26:41',NULL,1),('tgnc','Nam Cao','2023-12-19 00:28:00','2023-12-19 00:28:00',NULL,1),('tgnna','Nguyễn Nhật Ánh','2023-12-02 01:51:49','2023-12-02 01:51:49',NULL,1),('tgnnt','Nguyễn Ngọc Tư','2023-12-19 00:28:00','2023-12-19 00:28:00',NULL,1),('tgnpv','Nguyễn Phong Việt','2023-12-19 00:26:41','2023-12-19 00:26:41',NULL,1),('tgth','Trang Hạ','2023-12-19 00:26:41','2023-12-19 00:26:41',NULL,1),('tgtn','Tuệ Nghi','2023-12-19 00:26:41','2023-12-19 00:26:41',NULL,1),('tgvef','Viktor E Frankl','2023-12-02 01:51:49','2023-12-02 01:51:49',NULL,1),('tgvtp','Vũ Trọng Phụng','2023-12-19 00:28:00','2023-12-19 00:28:00',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `Book`
--

DROP TABLE IF EXISTS `Book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Book` (
  `id` varchar(12) NOT NULL,
  `name` varchar(100) NOT NULL,
  `booktitleid` varchar(12) NOT NULL,
  `publisherid` varchar(12) NOT NULL,
  `edition` int NOT NULL DEFAULT '1',
  `quantity` int NOT NULL DEFAULT '0',
  `listedPrice` int DEFAULT NULL,
  `sellPrice` int DEFAULT NULL,
  `importPrice` int DEFAULT NULL,
  `imgUrl` text,
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`,`booktitleid`),
  KEY `Book_BookTitle_id_fk` (`booktitleid`),
  KEY `Book_Publisher_id_fk` (`publisherid`),
  CONSTRAINT `Book_BookTitle_id_fk` FOREIGN KEY (`booktitleid`) REFERENCES `BookTitle` (`id`),
  CONSTRAINT `Book_Publisher_id_fk` FOREIGN KEY (`publisherid`) REFERENCES `Publisher` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Book`
--

LOCK TABLES `Book` WRITE;
INSERT INTO `Book` VALUES ('smb','Mắt biếc','dsmb','nxbdk',1,20,85000,85000,85000,'https://salt.tikicdn.com/cache/w1200/ts/product/10/d1/35/b2098bf8884bb8a5fbcd42a978a6b601.jpg','2023-12-19 01:00:19','2023-12-19 01:03:23',NULL,1),('Z9SfNPvIR','Tôi là Bêtô','stlbt','nxbdk',1,95,100000,120000,60000,'https://www.nxbtre.com.vn/Images/Book/nxbtre_full_05112021_111104.jpg','2023-12-14 06:31:35','2023-12-19 00:48:08',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `BookTitle`
--

DROP TABLE IF EXISTS `BookTitle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `BookTitle` (
  `id` varchar(12) NOT NULL,
  `name` varchar(100) NOT NULL,
  `desc` text,
  `authorIds` text NOT NULL,
  `categoryIds` text NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `BookTitle`
--

LOCK TABLES `BookTitle` WRITE;
INSERT INTO `BookTitle` VALUES ('dsdtls','Đi Tìm Lẽ Sống','Cuốn sách giúp người ta tìm được ý nghĩa cuộc sống','tgvef','dmkns','2023-12-14 18:44:24','2023-12-19 00:57:08',NULL,1),('dsmb','Mắt biếc','Mắt biếc là tiểu thuyết của nhà văn Nguyễn Nhật Ánh trong loạt truyện viết về tình yêu thanh thiếu niên của tác giả này cùng với Thằng quỷ nhỏ, Cô gái đến từ hôm qua,...','tgnna','dmtt|dmtruyen','2023-12-19 00:59:30','2023-12-19 00:59:30',NULL,1),('stlbt','Tôi là Bêtô','Một tác phẩm của Nguyễn Nhật Ánh','tgnna','dmtt|dmtruyen','2023-12-09 20:41:28','2023-12-11 09:54:37',NULL,1),('sttgbct','Tôi tài giỏi, bạn cũng thế!','Tôi tài giỏi, bạn cũng thế! (nhan đề gốc tiếng Anh: I Am Gifted, So Are You!) là quyển sách bán chạy nhất của doanh nhân người Singapore Adam Khoo, viết về những phương pháp học tập tiên tiến. Quyển sách đã được dịch ra hàng chục thứ tiếng, trong đó Tôi tài giỏi, bạn cũng thế! là phiên bản tiếng Việt được dịch bởi hai dịch giả nổi tiếng Trần Đăng Khoa và Uông Xuân Vy của TGM Books. Tại Việt Nam, quyển sách đã trở thành một hiện tượng giáo dục trong những năm 2009-2011 và đạt được nhiều thành tựu trong lĩnh vực xuất bản, tạo ra kỷ lục mới cho ngành xuất bản Việt Nam với hơn 200.000 bản in được bán ra và hơn 400.000 e-book được phân phối.','tgak','dmkns','2023-12-10 16:07:24','2023-12-11 09:50:45',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `Category`
--

DROP TABLE IF EXISTS `Category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Category` (
  `id` varchar(12) NOT NULL,
  `name` varchar(50) NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Category`
--

LOCK TABLES `Category` WRITE;
INSERT INTO `Category` VALUES ('dmdna','Dạy nấu ăn','2023-12-19 00:54:38','2023-12-19 00:54:38',NULL,1),('dmkh','Khoa học','2023-12-19 00:54:38','2023-12-19 00:54:38',NULL,1),('dmkns','Kỹ năng sống','2023-12-02 01:52:21','2023-12-02 01:52:21',NULL,1),('dmls','Lịch sử','2023-12-19 00:54:38','2023-12-19 00:54:38',NULL,1),('dmsgk','Sách giáo khoa','2023-12-02 01:52:21','2023-12-02 01:52:21',NULL,1),('dmtch','Truyển cảm hứng','2023-12-19 00:54:38','2023-12-19 00:54:38',NULL,1),('dmtg','Tôn giáo','2023-12-19 00:54:38','2023-12-19 00:54:38',NULL,1),('dmtl','Tâm lý','2023-12-19 00:54:38','2023-12-19 00:54:38',NULL,1),('dmtn','Thiếu nhi','2023-12-19 00:54:38','2023-12-19 00:54:38',NULL,1),('dmtruyen','Truyện','2023-12-02 01:52:21','2023-12-02 01:52:21',NULL,1),('dmtt','Tiểu thuyết','2023-12-02 01:52:21','2023-12-02 01:52:21',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `Feature`
--

DROP TABLE IF EXISTS `Feature`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Feature` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` text,
  `groupName` text,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Feature`
--

LOCK TABLES `Feature` WRITE;
INSERT INTO `Feature` VALUES ('AUTHOR_CREATE','Tạo tác giả','Tác giả','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('AUTHOR_DELETE','Xóa tác giả','Tác giả','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('AUTHOR_UPDATE','Chỉnh sửa thông tin tác giả','Tác giả','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('AUTHOR_VIEW','Xem tác giả','Tác giả','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('BOOK_CREATE','Tạo sách','Sách','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('BOOK_DELETE','Xóa sách','Sách','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('BOOK_TITLE_CREATE','Tạo đầu sách','Đầu sách','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('BOOK_TITLE_DELETE','Xóa đầu sách','Đầu sách','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('BOOK_TITLE_UPDATE','Chỉnh sửa thông tin đầu sách','Đầu sách','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('BOOK_TITLE_VIEW','Xem đầu sách','Đầu sách','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('BOOK_UPDATE','Chỉnh sửa thông tin sách','Sách','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('BOOK_VIEW','Xem sách','Sách','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('CATEGORY_CREATE','Tạo danh mục','Danh mục','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('CATEGORY_DELETE','Xóa danh mục','Danh mục','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('CATEGORY_UPDATE','Chỉnh sửa thông tin danh mục','Danh mục','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('CATEGORY_VIEW','Xem danh mục','Danh mục','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('IMPORT_NOTE_CREATE','Tạo phiếu nhập','Phiếu nhập','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('IMPORT_NOTE_STATUS','Chỉnh sửa trạng thái phiếu nhập','Phiếu nhập','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('IMPORT_NOTE_VIEW','Xem phiếu nhập','Phiếu nhập','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('INVENTORY_NOTE_CREATE','Tạo phiếu kiểm kho','Phiếu kiểm kho','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('INVENTORY_NOTE_VIEW','Xem phiếu kiểm kho','Phiếu kiểm kho','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('INVOICE_CREATE','Bán hàng','Hóa đơn','2023-12-15 01:46:28','2023-12-15 01:46:28',NULL,1),('INVOICE_VIEW','Xem hóa đơn','Hóa đơn','2023-12-15 01:46:28','2023-12-15 01:46:28',NULL,1),('PUBLISHER_CREATE','Tạo nhà sản xuất','Nhà sản xuất','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('PUBLISHER_VIEW','Xem nhà sản xuất','Nhà sản xuất','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('REPORT_VIEW_SALE','Xem báo cáo doanh thu','Báo cáo','2023-12-15 07:34:11','2023-12-15 07:34:11',NULL,1),('REPORT_VIEW_STOCK','Xem báo cáo tồn kho','Báo cáo','2023-12-15 07:34:11','2023-12-15 07:34:11',NULL,1),('REPORT_VIEW_SUPPLIER','Xem báo cáo nợ','Báo cáo','2023-12-15 07:34:11','2023-12-15 07:34:11',NULL,1),('SUPPLIER_CREATE','Tạo nhà cung cấp','Nhà cung cấp','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('SUPPLIER_PAY','Trả nợ nhà cung cấp','Nhà cung cấp','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('SUPPLIER_UPDATE_INFO','Chỉnh sửa thông tin nhà cung cấp','Nhà cung cấp','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('SUPPLIER_VIEW','Xem nhà cung cấp','Nhà cung cấp','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('USER_UPDATE_INFO','Chỉnh sửa thông tin người dùng','Nhân viên','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('USER_UPDATE_STATE','Chỉnh sửa trạng thái','Nhân viên','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1),('USER_VIEW','Xem người dùng','Nhân viên','2023-12-13 08:54:39','2023-12-14 07:56:30',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `ImportNote`
--

DROP TABLE IF EXISTS `ImportNote`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ImportNote` (
  `id` varchar(12) NOT NULL,
  `supplierId` varchar(12) NOT NULL,
  `totalPrice` int DEFAULT '0',
  `status` enum('InProgress','Done','Cancel') DEFAULT 'InProgress',
  `closedBy` varchar(12) DEFAULT NULL,
  `closedAt` datetime DEFAULT NULL,
  `createdBy` varchar(12) NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `supplierId` (`supplierId`),
  KEY `closedBy` (`closedBy`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `ImportNote_ibfk_1` FOREIGN KEY (`supplierId`) REFERENCES `Supplier` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ImportNote_ibfk_2` FOREIGN KEY (`closedBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ImportNote_ibfk_3` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ImportNote`
--

LOCK TABLES `ImportNote` WRITE;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_AUTO_VALUE_ON_ZERO' */ ;
/*!50032 DROP TRIGGER IF EXISTS update_closedAt */;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER `update_closedAt` BEFORE UPDATE ON `ImportNote` FOR EACH ROW BEGIN
    IF NEW.status != 'InProgress' THEN
        SET NEW.closedAt = CURRENT_TIMESTAMP;
    END IF;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `ImportNoteDetail`
--

DROP TABLE IF EXISTS `ImportNoteDetail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ImportNoteDetail` (
  `importNoteId` varchar(12) NOT NULL,
  `bookId` varchar(12) NOT NULL,
  `price` int NOT NULL,
  `qtyImport` int DEFAULT '0',
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`importNoteId`,`bookId`),
  KEY `bookId` (`bookId`),
  CONSTRAINT `ImportNoteDetail_ibfk_1` FOREIGN KEY (`importNoteId`) REFERENCES `ImportNote` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ImportNoteDetail_ibfk_2` FOREIGN KEY (`bookId`) REFERENCES `Book` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ImportNoteDetail`
--

LOCK TABLES `ImportNoteDetail` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `InventoryCheckNote`
--

DROP TABLE IF EXISTS `InventoryCheckNote`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `InventoryCheckNote` (
  `id` varchar(12) NOT NULL,
  `qtyDifferent` int NOT NULL,
  `qtyAfterAdjust` int NOT NULL,
  `createdBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `InventoryCheckNote_ibfk_1` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `InventoryCheckNote`
--

LOCK TABLES `InventoryCheckNote` WRITE;
INSERT INTO `InventoryCheckNote` VALUES ('tXLuY3DSg',-10,95,'g3W21A7SR','2023-12-17 12:56:45','2023-12-17 12:56:45',NULL,1),('ye3wLqDSg',20,105,'g3W21A7SR','2023-12-17 12:55:49','2023-12-17 12:55:49',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `InventoryCheckNoteDetail`
--

DROP TABLE IF EXISTS `InventoryCheckNoteDetail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `InventoryCheckNoteDetail` (
  `inventoryCheckNoteId` varchar(12) NOT NULL,
  `bookId` varchar(12) NOT NULL,
  `initial` int NOT NULL,
  `difference` int NOT NULL,
  `final` int NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`inventoryCheckNoteId`,`bookId`),
  KEY `bookId` (`bookId`),
  CONSTRAINT `InventoryCheckNoteDetail_ibfk_1` FOREIGN KEY (`inventoryCheckNoteId`) REFERENCES `InventoryCheckNote` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `InventoryCheckNoteDetail_ibfk_2` FOREIGN KEY (`bookId`) REFERENCES `Book` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `InventoryCheckNoteDetail`
--

LOCK TABLES `InventoryCheckNoteDetail` WRITE;
INSERT INTO `InventoryCheckNoteDetail` VALUES ('tXLuY3DSg','Z9SfNPvIR',105,-10,95,'2023-12-17 12:56:45','2023-12-17 12:56:45',NULL,1),('ye3wLqDSg','Z9SfNPvIR',85,20,105,'2023-12-17 12:55:49','2023-12-17 12:55:49',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `Invoice`
--

DROP TABLE IF EXISTS `Invoice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Invoice` (
  `id` varchar(13) NOT NULL,
  `totalPrice` int NOT NULL,
  `createdBy` varchar(13) NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Invoice`
--

LOCK TABLES `Invoice` WRITE;
INSERT INTO `Invoice` VALUES ('4SFAYqvIg',240000,'g3W21A7SR','2023-12-17 12:50:33','2023-12-17 12:50:33',NULL,1),('RlFAL3vIR',240000,'g3W21A7SR','2023-12-17 12:50:34','2023-12-17 12:50:34',NULL,1),('XUI0Y3vIg',240000,'g3W21A7SR','2023-12-17 12:50:30','2023-12-17 12:50:30',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `InvoiceDetail`
--

DROP TABLE IF EXISTS `InvoiceDetail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `InvoiceDetail` (
  `invoiceId` varchar(13) NOT NULL,
  `bookId` varchar(13) NOT NULL,
  `bookName` text,
  `qty` int NOT NULL,
  `unitPrice` int NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`invoiceId`,`bookId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `InvoiceDetail`
--

LOCK TABLES `InvoiceDetail` WRITE;
INSERT INTO `InvoiceDetail` VALUES ('4SFAYqvIg','Z9SfNPvIR','Tôi là Bêtô',2,120000,'2023-12-17 12:50:33','2023-12-17 12:50:33',NULL,1),('RlFAL3vIR','Z9SfNPvIR','Tôi là Bêtô',2,120000,'2023-12-17 12:50:34','2023-12-17 12:50:34',NULL,1),('XUI0Y3vIg','Z9SfNPvIR','Tôi là Bêtô',2,120000,'2023-12-17 12:50:30','2023-12-17 12:50:30',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `MUser`
--

DROP TABLE IF EXISTS `MUser`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `MUser` (
  `id` varchar(12) NOT NULL,
  `name` text NOT NULL,
  `phone` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` text NOT NULL,
  `salt` text NOT NULL,
  `roleId` varchar(12) NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`) USING BTREE,
  KEY `roleId` (`roleId`),
  CONSTRAINT `MUser_ibfk_1` FOREIGN KEY (`roleId`) REFERENCES `Role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `MUser`
--

LOCK TABLES `MUser` WRITE;
INSERT INTO `MUser` VALUES ('bgIqwQSIg','user','','','user@gmail.com','0dd71ba5a82e98ccdc6f5edb6fb870a5','ByVwWucjSGZkozLFeQcopssBrHPbCHoqRuUCFUbpfIhhqGUujj','user','2023-12-02 01:52:32','2023-12-04 01:24:10',NULL,1),('g3W21A7SR','admin','1234567890','','admin@gmail.com','5e107317df151f6e8e0015c4f2ee7936','mVMxRDAHpAJfyzuiXWRELghNpynUqBKueSboGBcrwHUuzEWsms','admin','2023-12-02 01:52:32','2023-12-04 01:24:10',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `Publisher`
--

DROP TABLE IF EXISTS `Publisher`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Publisher` (
  `id` varchar(12) NOT NULL,
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Publisher`
--

LOCK TABLES `Publisher` WRITE;
INSERT INTO `Publisher` VALUES ('nxbdk','Kim Đồng','2023-12-02 01:52:21','2023-12-02 01:52:21',NULL,1),('nxbgd','Giáo dục','2023-12-02 01:52:21','2023-12-10 16:07:18',NULL,1),('nxbld','Lao động','2023-12-19 00:23:52','2023-12-19 00:23:52',NULL,1),('nxbnn','Nhã Nam','2023-12-19 00:23:52','2023-12-19 00:23:52',NULL,1),('nxbpn','Tri thức','2023-12-19 00:23:52','2023-12-19 00:23:52',NULL,1),('nxbtn','Thanh niên','2023-12-19 00:23:52','2023-12-19 00:23:52',NULL,1),('nxbtre','Trẻ','2023-12-19 00:23:52','2023-12-19 00:23:52',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `Role`
--

DROP TABLE IF EXISTS `Role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Role` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Role`
--

LOCK TABLES `Role` WRITE;
INSERT INTO `Role` VALUES ('admin','admin','2023-12-02 01:52:40','2023-12-17 12:49:31',NULL,1),('user','user','2023-12-02 01:52:40','2023-12-02 01:52:40',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `RoleFeature`
--

DROP TABLE IF EXISTS `RoleFeature`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `RoleFeature` (
  `roleId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `featureId` varchar(30) NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`roleId`,`featureId`),
  KEY `featureId` (`featureId`),
  CONSTRAINT `RoleFeature_ibfk_1` FOREIGN KEY (`roleId`) REFERENCES `Role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `RoleFeature_ibfk_2` FOREIGN KEY (`featureId`) REFERENCES `Feature` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RoleFeature`
--

LOCK TABLES `RoleFeature` WRITE;
INSERT INTO `RoleFeature` VALUES ('admin','AUTHOR_CREATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','AUTHOR_DELETE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','AUTHOR_UPDATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','AUTHOR_VIEW','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','BOOK_CREATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','BOOK_DELETE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','BOOK_TITLE_CREATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','BOOK_TITLE_DELETE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','BOOK_TITLE_UPDATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','BOOK_TITLE_VIEW','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','BOOK_UPDATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','BOOK_VIEW','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','CATEGORY_CREATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','CATEGORY_DELETE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','CATEGORY_UPDATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','CATEGORY_VIEW','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','IMPORT_NOTE_CREATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','IMPORT_NOTE_STATUS','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','IMPORT_NOTE_VIEW','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','INVENTORY_NOTE_CREATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','INVENTORY_NOTE_VIEW','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','INVOICE_CREATE','2023-12-15 01:47:00','2023-12-15 01:47:00',NULL,1),('admin','INVOICE_VIEW','2023-12-15 01:47:00','2023-12-15 01:47:00',NULL,1),('admin','PUBLISHER_CREATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','PUBLISHER_VIEW','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','REPORT_VIEW_SALE','2023-12-15 07:34:53','2023-12-15 07:34:53',NULL,1),('admin','REPORT_VIEW_STOCK','2023-12-15 07:34:53','2023-12-15 07:34:53',NULL,1),('admin','REPORT_VIEW_SUPPLIER','2023-12-15 07:34:53','2023-12-15 07:34:53',NULL,1),('admin','SUPPLIER_CREATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','SUPPLIER_PAY','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','SUPPLIER_UPDATE_INFO','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','SUPPLIER_VIEW','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','USER_UPDATE_INFO','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','USER_UPDATE_STATE','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('admin','USER_VIEW','2023-12-12 08:46:33','2023-12-12 08:46:33',NULL,1),('user','AUTHOR_CREATE','2023-12-12 08:48:06','2023-12-12 08:48:06',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `StockChangeHistory`
--

DROP TABLE IF EXISTS `StockChangeHistory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `StockChangeHistory` (
  `id` varchar(12) NOT NULL,
  `bookId` varchar(12) NOT NULL,
  `qty` int NOT NULL,
  `qtyLeft` int NOT NULL,
  `type` enum('Sell','Import','Modify') NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`,`bookId`),
  KEY `bookId` (`bookId`),
  CONSTRAINT `StockChangeHistory_ibfk_1` FOREIGN KEY (`bookId`) REFERENCES `Book` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `StockChangeHistory`
--

LOCK TABLES `StockChangeHistory` WRITE;
INSERT INTO `StockChangeHistory` VALUES ('4SFAYqvIg','Z9SfNPvIR',-2,86,'Sell','2023-12-17 12:50:33','2023-12-17 12:50:33',NULL,1),('73PYLqvIR','Z9SfNPvIR',1,85,'Import','2023-12-17 12:54:20','2023-12-17 12:54:20',NULL,1),('RlFAL3vIR','Z9SfNPvIR',-2,84,'Sell','2023-12-17 12:50:34','2023-12-17 12:50:34',NULL,1),('tXLuY3DSg','Z9SfNPvIR',-10,95,'Modify','2023-12-17 12:56:45','2023-12-17 12:56:45',NULL,1),('XUI0Y3vIg','Z9SfNPvIR',-2,88,'Sell','2023-12-17 12:50:30','2023-12-17 12:50:30',NULL,1),('ye3wLqDSg','Z9SfNPvIR',20,105,'Modify','2023-12-17 12:55:48','2023-12-17 12:55:48',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `StockReport`
--

DROP TABLE IF EXISTS `StockReport`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `StockReport` (
  `id` varchar(12) NOT NULL,
  `timeFrom` timestamp NOT NULL,
  `timeTo` timestamp NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `StockReport`
--

LOCK TABLES `StockReport` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `StockReportDetail`
--

DROP TABLE IF EXISTS `StockReportDetail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `StockReportDetail` (
  `reportId` varchar(12) NOT NULL,
  `bookId` varchar(12) NOT NULL,
  `initial` int NOT NULL,
  `sell` int NOT NULL,
  `import` int NOT NULL,
  `modify` int NOT NULL,
  `final` int NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`reportId`,`bookId`),
  KEY `bookId` (`bookId`),
  CONSTRAINT `StockReportDetail_ibfk_1` FOREIGN KEY (`reportId`) REFERENCES `StockReport` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `StockReportDetail_ibfk_2` FOREIGN KEY (`bookId`) REFERENCES `Book` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `StockReportDetail`
--

LOCK TABLES `StockReportDetail` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `Supplier`
--

DROP TABLE IF EXISTS `Supplier`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Supplier` (
  `id` varchar(12) NOT NULL,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `phone` varchar(11) NOT NULL,
  `debt` int DEFAULT '0',
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Supplier`
--

LOCK TABLES `Supplier` WRITE;
INSERT INTO `Supplier` VALUES ('123','Đông A','donga@gmail.com','0123456789',-6160000,'2023-12-14 06:31:50','2023-12-19 01:10:43',NULL,1),('ncc1980','1980 Books','1980books@gmail.com','0123456787',0,'2023-12-19 01:09:49','2023-12-19 01:09:49',NULL,1),('nccapb','Alpha Books','alphabooks@gmail.com','0123456784',0,'2023-12-19 01:08:18','2023-12-19 01:08:18',NULL,1),('nccfn','First News','firstnews@gmail.com','0123456785',0,'2023-12-19 01:08:18','2023-12-19 01:08:18',NULL,1),('ncchnb','Hanoi Books','hanoibooks@gmail.com','0123456782',0,'2023-12-19 01:06:25','2023-12-19 01:06:25',NULL,1),('ncchtt','Hoa học trò','hoahoctro@gmail.com','0123456788',0,'2023-12-19 01:10:28','2023-12-19 01:10:28',NULL,1),('ncckd','Kim Đồng','kimdong@gmail.com','0123456781',0,'2023-12-19 01:06:25','2023-12-19 01:06:25',NULL,1),('nccnn','Nhã Nam','nhanam@gmail.com','0123456780',0,'2023-12-19 01:05:13','2023-12-19 01:05:13',NULL,1),('nccpdb','PandaBooks','pandabooks@gmail.com','0123456783',0,'2023-12-19 01:08:18','2023-12-19 01:08:18',NULL,1),('nccttt','Tri Thức Trẻ','trithuctre@gmail.com','0123456786',0,'2023-12-19 01:09:49','2023-12-19 01:09:49',NULL,1);
UNLOCK TABLES;

--
-- Table structure for table `SupplierDebt`
--

DROP TABLE IF EXISTS `SupplierDebt`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `SupplierDebt` (
  `id` varchar(12) NOT NULL,
  `supplierId` varchar(12) NOT NULL,
  `qty` int NOT NULL,
  `qtyLeft` int NOT NULL,
  `type` enum('Debt','Pay') NOT NULL,
  `createdBy` varchar(9) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`,`supplierId`),
  KEY `createdBy` (`createdBy`),
  KEY `supplierId` (`supplierId`),
  CONSTRAINT `SupplierDebt_ibfk_1` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `SupplierDebt_ibfk_2` FOREIGN KEY (`supplierId`) REFERENCES `Supplier` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SupplierDebt`
--

LOCK TABLES `SupplierDebt` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `SupplierDebtReport`
--

DROP TABLE IF EXISTS `SupplierDebtReport`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `SupplierDebtReport` (
  `id` varchar(12) NOT NULL,
  `timeFrom` timestamp NOT NULL,
  `timeTo` timestamp NOT NULL,
  `initial` int NOT NULL,
  `debt` int NOT NULL,
  `pay` int NOT NULL,
  `final` int NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SupplierDebtReport`
--

LOCK TABLES `SupplierDebtReport` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `SupplierDebtReportDetail`
--

DROP TABLE IF EXISTS `SupplierDebtReportDetail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `SupplierDebtReportDetail` (
  `reportId` varchar(12) NOT NULL,
  `supplierId` varchar(12) NOT NULL,
  `initial` int NOT NULL,
  `debt` int NOT NULL,
  `pay` int NOT NULL,
  `final` int NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`reportId`,`supplierId`),
  KEY `supplierId` (`supplierId`),
  CONSTRAINT `SupplierDebtReportDetail_ibfk_1` FOREIGN KEY (`supplierId`) REFERENCES `Supplier` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `SupplierDebtReportDetail_ibfk_2` FOREIGN KEY (`reportId`) REFERENCES `SupplierDebtReport` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SupplierDebtReportDetail`
--

LOCK TABLES `SupplierDebtReportDetail` WRITE;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-19  8:12:31
