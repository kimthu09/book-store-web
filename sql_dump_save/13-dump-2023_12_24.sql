-- MySQL dump 10.13  Distrib 8.0.34, for Linux (x86_64)
--
-- Host: 103.57.221.113    Database: bookstoremanagement
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

DROP TABLE IF EXISTS `Book`;
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

DROP TABLE IF EXISTS `BookTitle`;
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

DROP TABLE IF EXISTS `Category`;
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

DROP TABLE IF EXISTS `Feature`;
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

DROP TABLE IF EXISTS `ImportNote`;
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

DROP TABLE IF EXISTS `ImportNoteDetail`;
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

DROP TABLE IF EXISTS `InventoryCheckNote`;
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

DROP TABLE IF EXISTS `InventoryCheckNoteDetail`;
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

DROP TABLE IF EXISTS `Invoice`;
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

DROP TABLE IF EXISTS `InvoiceDetail`;
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

DROP TABLE IF EXISTS `MUser`;
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
  `imgUrl` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `isActive` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`) USING BTREE,
  KEY `roleId` (`roleId`),
  CONSTRAINT `MUser_ibfk_1` FOREIGN KEY (`roleId`) REFERENCES `Role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Publisher`;
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

DROP TABLE IF EXISTS `Role`;
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

DROP TABLE IF EXISTS `RoleFeature`;
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

DROP TABLE IF EXISTS `StockChangeHistory`;
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

DROP TABLE IF EXISTS `StockReport`;
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

DROP TABLE IF EXISTS `StockReportDetail`;
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

DROP TABLE IF EXISTS `Supplier`;
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

DROP TABLE IF EXISTS `SupplierDebt`;
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

DROP TABLE IF EXISTS `SupplierDebtReport`;
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

DROP TABLE IF EXISTS `SupplierDebtReportDetail`;
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

INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tgak', 'Adam Khoo', '2023-12-02 01:51:49', '2023-12-02 01:51:49', NULL, 1);
INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tgFFujio', 'Fujiko Fujio', '2023-12-19 01:50:09', '2023-12-19 01:50:09', NULL, 1);
INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tghat', 'Hồ Anh Thái', '2023-12-19 00:26:41', '2023-12-19 00:26:41', NULL, 1);
INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tgic', 'Iris Cao', '2023-12-19 00:26:41', '2023-12-19 00:26:41', NULL, 1),
('tgnc', 'Nam Cao', '2023-12-19 00:28:00', '2023-12-19 00:28:00', NULL, 1),
('tgnna', 'Nguyễn Nhật Ánh', '2023-12-02 01:51:49', '2023-12-02 01:51:49', NULL, 1),
('tgnnt', 'Nguyễn Ngọc Tư', '2023-12-19 00:28:00', '2023-12-19 00:28:00', NULL, 1),
('tgnpv', 'Nguyễn Phong Việt', '2023-12-19 00:26:41', '2023-12-19 00:26:41', NULL, 1),
('tgntp', 'Nguyễn Thị Phụng', '2023-12-19 02:05:31', '2023-12-19 02:05:31', NULL, 1),
('tgnull', 'Không/chưa tác giả', '2023-12-19 01:41:51', '2023-12-19 01:41:51', NULL, 1),
('tgth', 'Trang Hạ', '2023-12-19 00:26:41', '2023-12-19 00:26:41', NULL, 1),
('tgtn', 'Tuệ Nghi', '2023-12-19 00:26:41', '2023-12-19 00:26:41', NULL, 1),
('tgvef', 'Viktor E Frankl', '2023-12-02 01:51:49', '2023-12-02 01:51:49', NULL, 1),
('tgvtp', 'Vũ Trọng Phụng', '2023-12-19 00:28:00', '2023-12-19 00:28:00', NULL, 1),
('tgynh', 'Yuval Noah Harari', '2023-12-19 02:16:00', '2023-12-19 02:16:00', NULL, 1);

INSERT INTO `Book` (`id`, `name`, `booktitleid`, `publisherid`, `edition`, `quantity`, `listedPrice`, `sellPrice`, `importPrice`, `imgUrl`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dsslsln', 'Sapiens Lược Sử Loài Người', 'dsslsln', 'nxbls', 1, 5, 299000, 299001, 299000, 'https://cdn0.fahasa.com/media/catalog/product/8/9/8935270703554.jpg', '2023-12-19 02:17:08', '2023-12-19 02:17:45', NULL, 1);
INSERT INTO `Book` (`id`, `name`, `booktitleid`, `publisherid`, `edition`, `quantity`, `listedPrice`, `sellPrice`, `importPrice`, `imgUrl`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('s100ma', '100 Món Ăn Ngày Thường', 'ds100ma', 'nxbtn', 1, 12, 46000, 46000, 46000, 'https://cdn0.fahasa.com/media/catalog/product/1/1/1118020260362_1.jpg', '2023-12-19 02:07:56', '2023-12-19 02:23:21', NULL, 1);
INSERT INTO `Book` (`id`, `name`, `booktitleid`, `publisherid`, `edition`, `quantity`, `listedPrice`, `sellPrice`, `importPrice`, `imgUrl`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('sdoraemont12', 'Doraemon - Tập 12', 'dsdoraemon', 'nxbdk', 1, 29, 30000, 30000, 30000, 'https://momotaro.vn/upload/images/12_50.jpg', '2023-12-19 01:51:57', '2023-12-19 02:09:26', NULL, 1);
INSERT INTO `Book` (`id`, `name`, `booktitleid`, `publisherid`, `edition`, `quantity`, `listedPrice`, `sellPrice`, `importPrice`, `imgUrl`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('sdoraemonv23', 'Doraemon vol23. Nobita và những pháp sư gió bí ẩn', 'dsdoraemon', 'nxbdk', 1, 40, 35000, 35000, 35000, 'https://bizweb.dktcdn.net/thumb/1024x1024/100/299/021/products/8935244814316.jpg?v=1679371436627', '2023-12-19 01:56:51', '2023-12-19 01:57:10', NULL, 1),
('sdtls', 'Đi Tìm Lẽ Sống', 'dsdtls', 'nxbnn', 1, 96, 80000, 80000, 80000, 'https://salt.tikicdn.com/ts/product/80/14/8b/61fb657f347d14d9d7bf6fe901001a8e.jpg', '2023-12-19 01:47:37', '2023-12-19 02:23:21', NULL, 1),
('sgktoan5', 'Sách giáo khoa Toán lớp 5', 'dsgktoan', 'nxbgd', 1, 1000, 18000, 18000, 18000, 'https://hieusach24h.com/wp-content/uploads/2021/09/Toan-5-1.jpg', '2023-12-19 02:12:33', '2023-12-19 02:12:56', NULL, 1),
('sgktoan7', 'Sách giáo khoa Toán lớp 7', 'dsgktoan', 'nxbgd', 1, 1000, 18000, 18000, 18000, 'https://bizweb.dktcdn.net/100/397/635/products/giai-bai-tap-sgk-toan-lop-7-tap-1.png?v=1620215042633', '2023-12-19 02:12:33', '2023-12-19 02:12:56', NULL, 1),
('sipm2', 'Official IELTS Practice Materials 2 with DVD', 'dsipm2', 'nxbgd', 1, 10, 500000, 500000, 500000, 'https://cdn0.fahasa.com/media/catalog/product/i/m/image_195509_1_25616.jpg', '2023-12-19 01:43:25', '2023-12-19 01:43:48', NULL, 1),
('smb', 'Mắt biếc', 'dsmb', 'nxbdk', 1, 20, 85000, 85000, 85000, 'https://salt.tikicdn.com/cache/w1200/ts/product/10/d1/35/b2098bf8884bb8a5fbcd42a978a6b601.jpg', '2023-12-19 01:00:19', '2023-12-19 01:03:23', NULL, 1),
('stlbt', 'Tôi là Bêtô', 'stlbt', 'nxbdk', 1, 95, 100000, 120000, 60000, 'https://www.nxbtre.com.vn/Images/Book/nxbtre_full_05112021_111104.jpg', '2023-12-14 06:31:35', '2023-12-19 00:48:08', NULL, 1),
('sttgbct', 'Tôi tài giỏi, bạn cũng thế!', 'sttgbct', 'nxbtn', 1, 99, 150000, 150000, 150000, 'https://metaisach.com/wp-content/uploads/2019/05/toi-tai-gioi-ban-cung-the.jpg', '2023-12-19 01:35:19', '2023-12-19 02:23:21', NULL, 1);

INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('ds100ma', '100 Món Ăn Ngày Thường', 'Quyển sách 100 Món Ăn Ngày Thường cung cấp cho bạn những công thức nấu ăn món ăn ngày thường thông dụng nhất, dễ thực hiện với những kỹ thuật không quá cao.', 'tgntp', 'dmdna', '2023-12-19 02:06:50', '2023-12-19 02:06:50', NULL, 1);
INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dsdoraemon', 'Doraemon', 'Doraemon', 'tgFFujio', 'dmtruyen|dmtn', '2023-12-19 01:50:40', '2023-12-19 01:56:39', NULL, 1);
INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dsdtls', 'Đi Tìm Lẽ Sống', 'Cuốn sách giúp người ta tìm được ý nghĩa cuộc sống', 'tgvef', 'dmkns', '2023-12-14 18:44:24', '2023-12-19 00:57:08', NULL, 1);
INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dsgktoan', 'SGK Toán', 'Sách giáo khoa toán', 'tgnull', 'dmgd|dmsgk', '2023-12-19 02:11:35', '2023-12-19 02:11:35', NULL, 1),
('dsipm2', 'Official IELTS Practice Materials 2', 'Official IELTS Practice Materials 2', 'tgnull', 'dmtk|dmgd', '2023-12-19 01:41:59', '2023-12-19 01:41:59', NULL, 1),
('dsmb', 'Mắt biếc', 'Mắt biếc là tiểu thuyết của nhà văn Nguyễn Nhật Ánh trong loạt truyện viết về tình yêu thanh thiếu niên của tác giả này cùng với Thằng quỷ nhỏ, Cô gái đến từ hôm qua,...', 'tgnna', 'dmtt|dmtruyen', '2023-12-19 00:59:30', '2023-12-19 00:59:30', NULL, 1),
('dsslsln', 'Sapiens Lược Sử Loài Người', 'Sapiens Lược Sử Loài Người', 'tgynh', 'dmls', '2023-12-19 02:16:07', '2023-12-19 02:16:07', NULL, 1),
('stlbt', 'Tôi là Bêtô', 'Một tác phẩm của Nguyễn Nhật Ánh', 'tgnna', 'dmtt|dmtruyen', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('sttgbct', 'Tôi tài giỏi, bạn cũng thế!', 'Tôi tài giỏi, bạn cũng thế! (nhan đề gốc tiếng Anh: I Am Gifted, So Are You!) là quyển sách bán chạy nhất của doanh nhân người Singapore Adam Khoo, viết về những phương pháp học tập tiên tiến. Quyển sách đã được dịch ra hàng chục thứ tiếng, trong đó Tôi tài giỏi, bạn cũng thế! là phiên bản tiếng Việt được dịch bởi hai dịch giả nổi tiếng Trần Đăng Khoa và Uông Xuân Vy của TGM Books. Tại Việt Nam, quyển sách đã trở thành một hiện tượng giáo dục trong những năm 2009-2011 và đạt được nhiều thành tựu trong lĩnh vực xuất bản, tạo ra kỷ lục mới cho ngành xuất bản Việt Nam với hơn 200.000 bản in được bán ra và hơn 400.000 e-book được phân phối.', 'tgak', 'dmkns', '2023-12-10 16:07:24', '2023-12-11 09:50:45', NULL, 1);

INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmdna', 'Dạy nấu ăn', '2023-12-19 00:54:38', '2023-12-19 00:54:38', NULL, 1);
INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmgd', 'Giáo dục ', '2023-12-19 01:40:41', '2023-12-19 01:40:41', NULL, 1);
INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmkh', 'Khoa học', '2023-12-19 00:54:38', '2023-12-19 00:54:38', NULL, 1);
INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmkns', 'Kỹ năng sống', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('dmls', 'Lịch sử', '2023-12-19 00:54:38', '2023-12-19 00:54:38', NULL, 1),
('dmsgk', 'Sách giáo khoa', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('dmtch', 'Truyển cảm hứng', '2023-12-19 00:54:38', '2023-12-19 00:54:38', NULL, 1),
('dmtg', 'Tôn giáo', '2023-12-19 00:54:38', '2023-12-19 00:54:38', NULL, 1),
('dmtk', 'Tham khảo', '2023-12-19 01:40:29', '2023-12-19 01:40:29', NULL, 1),
('dmtl', 'Tâm lý', '2023-12-19 00:54:38', '2023-12-19 00:54:38', NULL, 1),
('dmtn', 'Thiếu nhi', '2023-12-19 00:54:38', '2023-12-19 00:54:38', NULL, 1),
('dmtruyen', 'Truyện', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('dmtt', 'Tiểu thuyết', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('SzG-PmOIg', 'sách nấu ăn', '2023-12-19 02:26:31', '2023-12-19 02:26:31', NULL, 1);

INSERT INTO `Feature` (`id`, `description`, `groupName`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_CREATE', 'Tạo tác giả', 'Tác giả', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1);
INSERT INTO `Feature` (`id`, `description`, `groupName`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_DELETE', 'Xóa tác giả', 'Tác giả', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1);
INSERT INTO `Feature` (`id`, `description`, `groupName`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_UPDATE', 'Chỉnh sửa thông tin tác giả', 'Tác giả', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1);
INSERT INTO `Feature` (`id`, `description`, `groupName`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_VIEW', 'Xem tác giả', 'Tác giả', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_CREATE', 'Tạo sách', 'Sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_DELETE', 'Xóa sách', 'Sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_TITLE_CREATE', 'Tạo đầu sách', 'Đầu sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_TITLE_DELETE', 'Xóa đầu sách', 'Đầu sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_TITLE_UPDATE', 'Chỉnh sửa thông tin đầu sách', 'Đầu sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_TITLE_VIEW', 'Xem đầu sách', 'Đầu sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_UPDATE', 'Chỉnh sửa thông tin sách', 'Sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_VIEW', 'Xem sách', 'Sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('CATEGORY_CREATE', 'Tạo danh mục', 'Danh mục', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('CATEGORY_DELETE', 'Xóa danh mục', 'Danh mục', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('CATEGORY_UPDATE', 'Chỉnh sửa thông tin danh mục', 'Danh mục', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('CATEGORY_VIEW', 'Xem danh mục', 'Danh mục', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('IMPORT_NOTE_CREATE', 'Tạo phiếu nhập', 'Phiếu nhập', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('IMPORT_NOTE_STATUS', 'Chỉnh sửa trạng thái phiếu nhập', 'Phiếu nhập', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('IMPORT_NOTE_VIEW', 'Xem phiếu nhập', 'Phiếu nhập', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('INVENTORY_NOTE_CREATE', 'Tạo phiếu kiểm kho', 'Phiếu kiểm kho', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('INVENTORY_NOTE_VIEW', 'Xem phiếu kiểm kho', 'Phiếu kiểm kho', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('INVOICE_CREATE', 'Bán hàng', 'Hóa đơn', '2023-12-15 01:46:28', '2023-12-15 01:46:28', NULL, 1),
('INVOICE_VIEW', 'Xem hóa đơn', 'Hóa đơn', '2023-12-15 01:46:28', '2023-12-15 01:46:28', NULL, 1),
('PUBLISHER_CREATE', 'Tạo nhà sản xuất', 'Nhà sản xuất', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('PUBLISHER_VIEW', 'Xem nhà sản xuất', 'Nhà sản xuất', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('REPORT_VIEW_SALE', 'Xem báo cáo doanh thu', 'Báo cáo', '2023-12-15 07:34:11', '2023-12-15 07:34:11', NULL, 1),
('REPORT_VIEW_STOCK', 'Xem báo cáo tồn kho', 'Báo cáo', '2023-12-15 07:34:11', '2023-12-15 07:34:11', NULL, 1),
('REPORT_VIEW_SUPPLIER', 'Xem báo cáo nợ', 'Báo cáo', '2023-12-15 07:34:11', '2023-12-15 07:34:11', NULL, 1),
('SUPPLIER_CREATE', 'Tạo nhà cung cấp', 'Nhà cung cấp', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('SUPPLIER_PAY', 'Trả nợ nhà cung cấp', 'Nhà cung cấp', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('SUPPLIER_UPDATE_INFO', 'Chỉnh sửa thông tin nhà cung cấp', 'Nhà cung cấp', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('SUPPLIER_VIEW', 'Xem nhà cung cấp', 'Nhà cung cấp', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('USER_UPDATE_INFO', 'Chỉnh sửa thông tin người dùng', 'Nhân viên', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('USER_UPDATE_STATE', 'Chỉnh sửa trạng thái', 'Nhân viên', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('USER_VIEW', 'Xem người dùng', 'Nhân viên', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1);









INSERT INTO `Invoice` (`id`, `totalPrice`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('oVFpPmOSR', 482000, 'g3W21A7SR', '2023-12-19 02:23:21', '2023-12-19 02:23:21', NULL, 1);
INSERT INTO `Invoice` (`id`, `totalPrice`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('Wzu-YiOIR', 156000, 'g3W21A7SR', '2023-12-19 02:09:26', '2023-12-19 02:09:26', NULL, 1);


INSERT INTO `InvoiceDetail` (`invoiceId`, `bookId`, `bookName`, `qty`, `unitPrice`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('oVFpPmOSR', 's100ma', '100 Món Ăn Ngày Thường', 2, 46000, '2023-12-19 02:23:21', '2023-12-19 02:23:21', NULL, 1);
INSERT INTO `InvoiceDetail` (`invoiceId`, `bookId`, `bookName`, `qty`, `unitPrice`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('oVFpPmOSR', 'sdtls', 'Đi Tìm Lẽ Sống', 3, 80000, '2023-12-19 02:23:21', '2023-12-19 02:23:21', NULL, 1);
INSERT INTO `InvoiceDetail` (`invoiceId`, `bookId`, `bookName`, `qty`, `unitPrice`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('oVFpPmOSR', 'sttgbct', 'Tôi tài giỏi, bạn cũng thế!', 1, 150000, '2023-12-19 02:23:21', '2023-12-19 02:23:21', NULL, 1);
INSERT INTO `InvoiceDetail` (`invoiceId`, `bookId`, `bookName`, `qty`, `unitPrice`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('Wzu-YiOIR', 's100ma', '100 Món Ăn Ngày Thường', 1, 46000, '2023-12-19 02:09:26', '2023-12-19 02:09:26', NULL, 1),
('Wzu-YiOIR', 'sdoraemont12', 'Doraemon - Tập 12', 1, 30000, '2023-12-19 02:09:26', '2023-12-19 02:09:26', NULL, 1),
('Wzu-YiOIR', 'sdtls', 'Đi Tìm Lẽ Sống', 1, 80000, '2023-12-19 02:09:26', '2023-12-19 02:09:26', NULL, 1);

INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `createdAt`, `updatedAt`, `deletedAt`, `imgUrl`, `isActive`) VALUES
('bgIqwQSIg', 'user', '', '', 'user@gmail.com', '0dd71ba5a82e98ccdc6f5edb6fb870a5', 'ByVwWucjSGZkozLFeQcopssBrHPbCHoqRuUCFUbpfIhhqGUujj', 'user', '2023-12-02 01:52:32', '2023-12-24 09:56:50', NULL, 'https://cdn-icons-png.flaticon.com/512/149/149071.png', 1);
INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `createdAt`, `updatedAt`, `deletedAt`, `imgUrl`, `isActive`) VALUES
('g3W21A7SR', 'admin', '1234567890', '', 'admin@gmail.com', '5e107317df151f6e8e0015c4f2ee7936', 'mVMxRDAHpAJfyzuiXWRELghNpynUqBKueSboGBcrwHUuzEWsms', 'admin', '2023-12-02 01:52:32', '2023-12-24 09:56:50', NULL, 'https://cdn-icons-png.flaticon.com/512/149/149071.png', 1);
INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `createdAt`, `updatedAt`, `deletedAt`, `imgUrl`, `isActive`) VALUES
('kJ68EidIg', 'Thu Nguyen', '0987654321', 'None', 'nguyenkimanhthu25092003@gmail.com', 'd06c7bc74ad018e3ade1fdd536b1ec96', 'AXUGqeiDINLbHeaGmxKauptFKnwJgOUxZaMGaCBuJncsphurtf', 'user', '2023-12-19 02:29:43', '2023-12-24 09:56:50', NULL, 'https://cdn-icons-png.flaticon.com/512/149/149071.png', 1);

INSERT INTO `Publisher` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('nxbdk', 'Kim Đồng', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1);
INSERT INTO `Publisher` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('nxbgd', 'Giáo dục', '2023-12-02 01:52:21', '2023-12-10 16:07:18', NULL, 1);
INSERT INTO `Publisher` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('nxbld', 'Lao động', '2023-12-19 00:23:52', '2023-12-19 00:23:52', NULL, 1);
INSERT INTO `Publisher` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('nxbls', 'Lịch sử', '2023-12-19 02:17:06', '2023-12-20 18:12:07', NULL, 1),
('nxbnn', 'Nhã Nam', '2023-12-19 00:23:52', '2023-12-19 00:23:52', NULL, 1),
('nxbpn', 'Tri thức', '2023-12-19 00:23:52', '2023-12-19 00:23:52', NULL, 1),
('nxbtn', 'Thanh niên', '2023-12-19 00:23:52', '2023-12-19 00:23:52', NULL, 1),
('nxbtre', 'Trẻ', '2023-12-19 00:23:52', '2023-12-19 00:23:52', NULL, 1);

INSERT INTO `Role` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('admin', 'admin', '2023-12-02 01:52:40', '2023-12-17 12:49:31', NULL, 1);
INSERT INTO `Role` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('user', 'user', '2023-12-02 01:52:40', '2023-12-02 01:52:40', NULL, 1);
INSERT INTO `Role` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('WEyQEidSg', 'thu ngân', '2023-12-19 02:30:07', '2023-12-19 02:30:23', NULL, 1);

INSERT INTO `RoleFeature` (`roleId`, `featureId`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('admin', 'AUTHOR_CREATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1);
INSERT INTO `RoleFeature` (`roleId`, `featureId`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('admin', 'AUTHOR_DELETE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1);
INSERT INTO `RoleFeature` (`roleId`, `featureId`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('admin', 'AUTHOR_UPDATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1);
INSERT INTO `RoleFeature` (`roleId`, `featureId`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('admin', 'AUTHOR_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'BOOK_CREATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'BOOK_DELETE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'BOOK_TITLE_CREATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'BOOK_TITLE_DELETE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'BOOK_TITLE_UPDATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'BOOK_TITLE_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'BOOK_UPDATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'BOOK_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'CATEGORY_CREATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'CATEGORY_DELETE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'CATEGORY_UPDATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'CATEGORY_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'IMPORT_NOTE_CREATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'IMPORT_NOTE_STATUS', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'IMPORT_NOTE_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'INVENTORY_NOTE_CREATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'INVENTORY_NOTE_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'INVOICE_CREATE', '2023-12-15 01:47:00', '2023-12-15 01:47:00', NULL, 1),
('admin', 'INVOICE_VIEW', '2023-12-15 01:47:00', '2023-12-15 01:47:00', NULL, 1),
('admin', 'PUBLISHER_CREATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'PUBLISHER_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'REPORT_VIEW_SALE', '2023-12-15 07:34:53', '2023-12-15 07:34:53', NULL, 1),
('admin', 'REPORT_VIEW_STOCK', '2023-12-15 07:34:53', '2023-12-15 07:34:53', NULL, 1),
('admin', 'REPORT_VIEW_SUPPLIER', '2023-12-15 07:34:53', '2023-12-15 07:34:53', NULL, 1),
('admin', 'SUPPLIER_CREATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'SUPPLIER_PAY', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'SUPPLIER_UPDATE_INFO', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'SUPPLIER_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'USER_UPDATE_INFO', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'USER_UPDATE_STATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'USER_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('user', 'AUTHOR_CREATE', '2023-12-12 08:48:06', '2023-12-12 08:48:06', NULL, 1),
('WEyQEidSg', 'AUTHOR_CREATE', '2023-12-19 02:30:07', '2023-12-19 02:30:07', NULL, 1),
('WEyQEidSg', 'AUTHOR_DELETE', '2023-12-19 02:30:07', '2023-12-19 02:30:07', NULL, 1),
('WEyQEidSg', 'AUTHOR_UPDATE', '2023-12-19 02:30:07', '2023-12-19 02:30:07', NULL, 1),
('WEyQEidSg', 'AUTHOR_VIEW', '2023-12-19 02:30:07', '2023-12-19 02:30:07', NULL, 1),
('WEyQEidSg', 'CATEGORY_CREATE', '2023-12-19 02:30:07', '2023-12-19 02:30:07', NULL, 1),
('WEyQEidSg', 'CATEGORY_DELETE', '2023-12-19 02:30:07', '2023-12-19 02:30:07', NULL, 1),
('WEyQEidSg', 'CATEGORY_UPDATE', '2023-12-19 02:30:07', '2023-12-19 02:30:07', NULL, 1),
('WEyQEidSg', 'CATEGORY_VIEW', '2023-12-19 02:30:07', '2023-12-19 02:30:07', NULL, 1);







INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('ncc1980', '1980 Book', '1980books@gmail.com', '0345689012', 0, '2023-12-19 01:09:49', '2023-12-19 02:28:19', NULL, 1);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('nccapb', 'Alpha Books', 'alphabooks@gmail.com', '0123456784', 0, '2023-12-19 01:08:18', '2023-12-19 01:08:18', NULL, 1);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('nccfn', 'First News', 'firstnews@gmail.com', '0123456785', 0, '2023-12-19 01:08:18', '2023-12-19 01:08:18', NULL, 1);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('ncchnb', 'Hanoi Books', 'hanoibooks@gmail.com', '0123456782', 0, '2023-12-19 01:06:25', '2023-12-19 01:06:25', NULL, 1),
('ncchtt', 'Hoa học trò', 'hoahoctro@gmail.com', '0123456788', 0, '2023-12-19 01:10:28', '2023-12-19 01:10:28', NULL, 1),
('ncckd', 'Kim Đồng', 'kimdong@gmail.com', '0123456781', 0, '2023-12-19 01:06:25', '2023-12-19 01:06:25', NULL, 1),
('nccnn', 'Nhã Nam', 'nhanam@gmail.com', '0123456780', 0, '2023-12-19 01:05:13', '2023-12-19 01:05:13', NULL, 1),
('nccpdb', 'PandaBooks', 'pandabooks@gmail.com', '0123456783', 0, '2023-12-19 01:08:18', '2023-12-19 01:08:18', NULL, 1),
('nccttt', 'Tri Thức Trẻ', 'trithuctre@gmail.com', '0123456786', 0, '2023-12-19 01:09:49', '2023-12-19 01:09:49', NULL, 1);

DELIMITER //
CREATE TRIGGER update_closedAt
BEFORE UPDATE ON ImportNote
FOR EACH ROW
BEGIN
    IF NEW.status != 'InProgress' THEN
        SET NEW.closedAt = CURRENT_TIMESTAMP;
    END IF;
END;
//
DELIMITER ;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

