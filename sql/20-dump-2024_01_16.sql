/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

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

DROP TABLE IF EXISTS `Customer`;
CREATE TABLE `Customer` (
  `id` varchar(12) NOT NULL,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `phone` varchar(11) NOT NULL,
  `point` int DEFAULT '0',
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
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
  `customerId` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `totalPrice` int NOT NULL,
  `totalImportPrice` int NOT NULL,
  `amountReceived` int NOT NULL,
  `amountPriceUsePoint` int NOT NULL,
  `pointUse` int NOT NULL,
  `pointReceive` int NOT NULL,
  `createdBy` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `customerId` (`customerId`),
  CONSTRAINT `Invoice_ibfk_1` FOREIGN KEY (`customerId`) REFERENCES `Customer` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
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
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
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

DROP TABLE IF EXISTS `ShopGeneral`;
CREATE TABLE `ShopGeneral` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` text NOT NULL,
  `address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `wifiPass` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `accumulatePointPercent` float NOT NULL,
  `usePointPercent` float NOT NULL,
  PRIMARY KEY (`id`)
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
  `initial` int NOT NULL,
  `sell` int NOT NULL,
  `import` int NOT NULL,
  `modify` int NOT NULL,
  `final` int NOT NULL,
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
  `createdBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
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
('dsslsln', 'Sapiens Lược Sử Loài Người', 'dsslsln', 'nxbls', 1, 9, 299000, 299000, 290000, 'https://cdn0.fahasa.com/media/catalog/product/8/9/8935270703554.jpg', '2023-12-19 02:17:08', '2024-01-08 02:12:23', NULL, 1);
INSERT INTO `Book` (`id`, `name`, `booktitleid`, `publisherid`, `edition`, `quantity`, `listedPrice`, `sellPrice`, `importPrice`, `imgUrl`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('s100ma', '100 Món Ăn Ngày Thường', 'ds100ma', 'nxbtn', 1, 29, 46000, 46000, 40000, 'https://cdn0.fahasa.com/media/catalog/product/1/1/1118020260362_1.jpg', '2023-12-19 02:07:56', '2024-01-16 04:47:49', NULL, 1);
INSERT INTO `Book` (`id`, `name`, `booktitleid`, `publisherid`, `edition`, `quantity`, `listedPrice`, `sellPrice`, `importPrice`, `imgUrl`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('sdoraemont12', 'Doraemon - Tập 12', 'dsdoraemon', 'nxbdk', 1, 15, 30000, 30000, 25000, 'https://momotaro.vn/upload/images/12_50.jpg', '2023-12-19 01:51:57', '2024-01-16 04:47:44', NULL, 1);
INSERT INTO `Book` (`id`, `name`, `booktitleid`, `publisherid`, `edition`, `quantity`, `listedPrice`, `sellPrice`, `importPrice`, `imgUrl`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('sdoraemonv23', 'Doraemon vol23. Nobita và những pháp sư gió bí ẩn', 'dsdoraemon', 'nxbdk', 1, 18, 35000, 35000, 30000, 'https://bizweb.dktcdn.net/thumb/1024x1024/100/299/021/products/8935244814316.jpg?v=1679371436627', '2023-12-19 01:56:51', '2024-01-07 03:12:23', NULL, 1),
('sdtls', 'Đi Tìm Lẽ Sống', 'dsdtls', 'nxbnn', 1, 11, 80000, 80000, 70000, 'https://salt.tikicdn.com/ts/product/80/14/8b/61fb657f347d14d9d7bf6fe901001a8e.jpg', '2023-12-19 01:47:37', '2024-01-16 04:47:49', NULL, 1),
('sgktoan5', 'Sách giáo khoa Toán lớp 5', 'dsgktoan', 'nxbgd', 1, 49, 18000, 18000, 15000, 'https://hieusach24h.com/wp-content/uploads/2021/09/Toan-5-1.jpg', '2023-12-19 02:12:33', '2024-01-03 06:33:23', NULL, 1),
('sgktoan7', 'Sách giáo khoa Toán lớp 7', 'dsgktoan', 'nxbgd', 1, 49, 18000, 18000, 15000, 'https://bizweb.dktcdn.net/100/397/635/products/giai-bai-tap-sgk-toan-lop-7-tap-1.png?v=1620215042633', '2023-12-19 02:12:33', '2024-01-03 06:33:23', NULL, 1),
('sipm2', 'Official IELTS Practice Materials 2 with DVD', 'dsipm2', 'nxbgd', 1, 6, 500000, 500000, 470000, 'https://cdn0.fahasa.com/media/catalog/product/i/m/image_195509_1_25616.jpg', '2023-12-19 01:43:25', '2024-01-16 04:47:10', NULL, 1),
('smb', 'Mắt biếc', 'dsmb', 'nxbdk', 1, 13, 85000, 85000, 70000, 'https://salt.tikicdn.com/cache/w1200/ts/product/10/d1/35/b2098bf8884bb8a5fbcd42a978a6b601.jpg', '2023-12-19 01:00:19', '2024-01-16 04:47:54', NULL, 1),
('stlbt', 'Tôi là Bêtô', 'stlbt', 'nxbdk', 1, 15, 100000, 120000, 90000, 'https://www.nxbtre.com.vn/Images/Book/nxbtre_full_05112021_111104.jpg', '2023-12-14 06:31:35', '2024-01-16 04:47:54', NULL, 1),
('sttgbct', 'Tôi tài giỏi, bạn cũng thế!', 'sttgbct', 'nxbtn', 1, 39, 150000, 150000, 130000, 'https://metaisach.com/wp-content/uploads/2019/05/toi-tai-gioi-ban-cung-the.jpg', '2023-12-19 01:35:19', '2024-01-16 04:45:30', NULL, 1);

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

INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('1utsgF5SR', 'Nguyễn Lê Ngọc Mai', 'mai@gmail.com', '0902845188', 2042, '2024-01-16 04:28:57', '2024-01-16 04:47:49', NULL, 1);
INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('2DiuRF5IR', 'Bùi Vĩ Quốc', 'quoc@gmail.com', '0905341132', 1885, '2024-01-16 04:30:56', '2024-01-16 04:47:44', NULL, 1);
INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('ZqeugFcSR', 'Vũ Hoàng', 'hoang@gmail.com', '0983985681', 1948, '2024-01-16 04:31:26', '2024-01-16 04:47:54', NULL, 1);

INSERT INTO `Feature` (`id`, `description`, `groupName`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_CREATE', 'Tạo tác giả', 'Tác giả', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1);
INSERT INTO `Feature` (`id`, `description`, `groupName`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_UPDATE', 'Chỉnh sửa thông tin tác giả', 'Tác giả', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1);
INSERT INTO `Feature` (`id`, `description`, `groupName`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_VIEW', 'Xem tác giả', 'Tác giả', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1);
INSERT INTO `Feature` (`id`, `description`, `groupName`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('BOOK_CREATE', 'Tạo sách', 'Sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_TITLE_CREATE', 'Tạo đầu sách', 'Đầu sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_TITLE_UPDATE', 'Chỉnh sửa thông tin đầu sách', 'Đầu sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_TITLE_VIEW', 'Xem đầu sách', 'Đầu sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_UPDATE', 'Chỉnh sửa thông tin sách', 'Sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('BOOK_VIEW', 'Xem sách', 'Sách', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('CATEGORY_CREATE', 'Tạo danh mục', 'Danh mục', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('CATEGORY_UPDATE', 'Chỉnh sửa thông tin danh mục', 'Danh mục', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('CATEGORY_VIEW', 'Xem danh mục', 'Danh mục', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('CUSTOMER_CREATE', 'Tạo khách hàng', 'Khách hàng', '2024-01-13 15:18:19', '2024-01-13 15:18:19', NULL, 1),
('CUSTOMER_UPDATE_INFO', 'Chỉnh sửa thông tin khách hàng', 'Khách hàng', '2024-01-13 15:18:19', '2024-01-13 15:18:19', NULL, 1),
('CUSTOMER_VIEW', 'Xem khách hàng', 'Khách hàng', '2024-01-13 15:18:19', '2024-01-13 15:18:19', NULL, 1),
('IMPORT_NOTE_CREATE', 'Tạo phiếu nhập', 'Phiếu nhập', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('IMPORT_NOTE_STATUS', 'Chỉnh sửa trạng thái phiếu nhập', 'Phiếu nhập', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('IMPORT_NOTE_VIEW', 'Xem phiếu nhập', 'Phiếu nhập', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('INVENTORY_NOTE_CREATE', 'Tạo phiếu kiểm kho', 'Phiếu kiểm kho', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('INVENTORY_NOTE_VIEW', 'Xem phiếu kiểm kho', 'Phiếu kiểm kho', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('INVOICE_CREATE', 'Bán hàng', 'Hóa đơn', '2023-12-15 01:46:28', '2023-12-15 01:46:28', NULL, 1),
('INVOICE_VIEW', 'Xem hóa đơn', 'Hóa đơn', '2023-12-15 01:46:28', '2023-12-15 01:46:28', NULL, 1),
('PUBLISHER_CREATE', 'Tạo nhà xuất bản', 'Nhà xuất bản', '2023-12-13 08:54:39', '2024-01-02 16:56:12', NULL, 1),
('PUBLISHER_UPDATE', 'Chỉnh sửa thông tin nhà xuất bản', 'Nhà xuất bản', '2023-12-31 09:59:38', '2024-01-02 16:56:12', NULL, 1),
('PUBLISHER_VIEW', 'Xem nhà xuất bản', 'Nhà xuất bản', '2023-12-13 08:54:39', '2024-01-02 16:56:12', NULL, 1),
('REPORT_VIEW_SALE', 'Xem báo cáo doanh thu', 'Báo cáo', '2023-12-15 07:34:11', '2023-12-15 07:34:11', NULL, 1),
('REPORT_VIEW_STOCK', 'Xem báo cáo tồn kho', 'Báo cáo', '2023-12-15 07:34:11', '2023-12-15 07:34:11', NULL, 1),
('REPORT_VIEW_SUPPLIER', 'Xem báo cáo nợ', 'Báo cáo', '2023-12-15 07:34:11', '2023-12-15 07:34:11', NULL, 1),
('SUPPLIER_CREATE', 'Tạo nhà cung cấp', 'Nhà cung cấp', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('SUPPLIER_PAY', 'Trả nợ nhà cung cấp', 'Nhà cung cấp', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('SUPPLIER_UPDATE_INFO', 'Chỉnh sửa thông tin nhà cung cấp', 'Nhà cung cấp', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('SUPPLIER_VIEW', 'Xem nhà cung cấp', 'Nhà cung cấp', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('USER_UPDATE_INFO', 'Chỉnh sửa thông tin người dùng', 'Nhân viên', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1),
('USER_VIEW', 'Xem người dùng', 'Nhân viên', '2023-12-13 08:54:39', '2023-12-14 07:56:30', NULL, 1);

INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `closedBy`, `closedAt`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('4b3WjO5Ig', 'nccapb', 400000, 'Done', 'g3W21A7SR', '2023-12-01 23:55:23', 'g3W21A7SR', '2023-12-01 22:55:23', '2023-12-01 23:55:23', NULL, 1);
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `closedBy`, `closedAt`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('9MhHCd5Ig', 'ncc1980', 3700000, 'Done', 'g3W21A7SR', '2023-12-02 17:10:22', 'g3W21A7SR', '2023-12-02 06:10:22', '2023-12-02 17:10:22', NULL, 1);
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `closedBy`, `closedAt`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('DeT7Cd5Ig', 'nccttt', 700000, 'Done', 'g3W21A7SR', '2023-12-07 13:22:12', 'g3W21A7SR', '2023-12-02 18:15:22', '2023-12-07 13:22:12', NULL, 1);
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `closedBy`, `closedAt`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('FkAEzKcSR', 'nccapb', 550000, 'InProgress', NULL, NULL, 'g3W21A7SR', '2024-01-16 04:45:55', '2024-01-16 04:45:55', NULL, 1),
('hAgSCd5Sg', 'nccttt', 3850000, 'Done', 'g3W21A7SR', '2023-12-04 08:34:23', 'g3W21A7SR', '2023-12-03 16:34:23', '2023-12-04 08:34:23', NULL, 1),
('HmzEzF5IR', 'nccfn', 70000, 'Cancel', 'g3W21A7SR', '2024-01-16 04:45:43', 'g3W21A7SR', '2024-01-16 04:45:39', '2024-01-16 04:52:09', NULL, 1),
('hNMVjdcIg', 'nccpdb', 1850000, 'Cancel', 'g3W21A7SR', '2023-12-04 16:12:23', 'g3W21A7SR', '2023-12-04 12:53:23', '2023-12-04 16:12:23', NULL, 1),
('iAJGCOcIR', 'ncckd', 1800000, 'Done', 'g3W21A7SR', '2023-12-02 19:22:14', 'g3W21A7SR', '2023-12-07 03:26:12', '2023-12-02 19:22:14', NULL, 1),
('KMNnjdcSR', 'ncckd', 4800000, 'Done', 'g3W21A7SR', '2023-12-09 18:26:12', 'g3W21A7SR', '2023-12-08 08:26:12', '2023-12-09 18:26:12', NULL, 1),
('p6-LkKcIR', 'nccfn', 5100000, 'Done', 'g3W21A7SR', '2024-01-16 04:45:30', 'g3W21A7SR', '2024-01-16 04:45:26', '2024-01-16 04:51:44', NULL, 1),
('QFmvjd5Ig', 'nccapb', 700000, 'Done', 'g3W21A7SR', '2023-12-12 22:26:12', 'g3W21A7SR', '2023-12-10 10:26:12', '2023-12-12 22:26:12', NULL, 1),
('vI8PkF5SR', 'nccnn', 1410000, 'InProgress', NULL, NULL, 'g3W21A7SR', '2024-01-16 04:46:03', '2024-01-16 04:46:03', NULL, 1);

INSERT INTO `ImportNoteDetail` (`importNoteId`, `bookId`, `price`, `qtyImport`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('4b3WjO5Ig', 's100ma', 40000, 10, '2023-12-01 22:55:23', '2023-12-01 22:55:23', NULL, 1);
INSERT INTO `ImportNoteDetail` (`importNoteId`, `bookId`, `price`, `qtyImport`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('9MhHCd5Ig', 'dsslsln', 290000, 10, '2023-12-02 06:10:22', '2023-12-02 06:10:22', NULL, 1);
INSERT INTO `ImportNoteDetail` (`importNoteId`, `bookId`, `price`, `qtyImport`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('9MhHCd5Ig', 's100ma', 40000, 20, '2023-12-02 06:10:22', '2023-12-02 06:10:22', NULL, 1);
INSERT INTO `ImportNoteDetail` (`importNoteId`, `bookId`, `price`, `qtyImport`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('DeT7Cd5Ig', 'smb', 70000, 10, '2023-12-02 18:15:22', '2023-12-02 18:15:22', NULL, 1),
('FkAEzKcSR', 's100ma', 40000, 5, '2024-01-16 04:45:55', '2024-01-16 04:45:55', NULL, 1),
('FkAEzKcSR', 'sdtls', 70000, 5, '2024-01-16 04:45:55', '2024-01-16 04:45:55', NULL, 1),
('hAgSCd5Sg', 'sgktoan5', 15000, 50, '2023-12-03 16:34:23', '2023-12-03 16:34:23', NULL, 1),
('hAgSCd5Sg', 'sgktoan7', 15000, 50, '2023-12-03 16:34:23', '2023-12-03 16:34:23', NULL, 1),
('hAgSCd5Sg', 'sipm2', 470000, 5, '2023-12-03 16:34:23', '2023-12-03 16:34:23', NULL, 1),
('HmzEzF5IR', 'sdtls', 70000, 1, '2024-01-16 04:45:39', '2024-01-16 04:45:39', NULL, 1),
('hNMVjdcIg', 'dsslsln', 290000, 5, '2023-12-04 12:53:23', '2023-12-04 12:53:23', NULL, 1),
('hNMVjdcIg', 's100ma', 40000, 10, '2023-12-04 12:53:23', '2023-12-04 12:53:23', NULL, 1),
('iAJGCOcIR', 'sdoraemont12', 25000, 20, '2023-12-07 03:26:12', '2023-12-07 03:26:12', NULL, 1),
('iAJGCOcIR', 'sdoraemonv23', 30000, 20, '2023-12-07 03:26:12', '2023-12-07 03:26:12', NULL, 1),
('iAJGCOcIR', 'smb', 70000, 10, '2023-12-07 03:26:12', '2023-12-07 03:26:12', NULL, 1),
('KMNnjdcSR', 'stlbt', 90000, 10, '2023-12-08 08:26:12', '2023-12-08 08:26:12', NULL, 1),
('KMNnjdcSR', 'sttgbct', 130000, 30, '2023-12-08 08:26:12', '2023-12-08 08:26:12', NULL, 1),
('p6-LkKcIR', 's100ma', 40000, 5, '2024-01-16 04:45:26', '2024-01-16 04:45:26', NULL, 1),
('p6-LkKcIR', 'sdtls', 70000, 5, '2024-01-16 04:45:26', '2024-01-16 04:45:26', NULL, 1),
('p6-LkKcIR', 'sipm2', 470000, 5, '2024-01-16 04:45:26', '2024-01-16 04:45:26', NULL, 1),
('p6-LkKcIR', 'stlbt', 90000, 10, '2024-01-16 04:45:26', '2024-01-16 04:45:26', NULL, 1),
('p6-LkKcIR', 'sttgbct', 130000, 10, '2024-01-16 04:45:26', '2024-01-16 04:45:26', NULL, 1),
('QFmvjd5Ig', 'sdtls', 70000, 10, '2023-12-10 10:26:12', '2023-12-10 10:26:12', NULL, 1),
('vI8PkF5SR', 'sipm2', 470000, 3, '2024-01-16 04:46:03', '2024-01-16 04:46:03', NULL, 1);

INSERT INTO `InventoryCheckNote` (`id`, `qtyDifferent`, `qtyAfterAdjust`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('BYwUzF5IR', -4, 36, 'g3W21A7SR', '2024-01-16 04:47:10', '2024-01-16 04:47:10', NULL, 1);
INSERT INTO `InventoryCheckNote` (`id`, `qtyDifferent`, `qtyAfterAdjust`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('eFsOCOcIg', 1, 5, 'g3W21A7SR', '2023-12-30 12:10:12', '2023-12-30 12:10:12', NULL, 1);
INSERT INTO `InventoryCheckNote` (`id`, `qtyDifferent`, `qtyAfterAdjust`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('PXPDCOcIg', 1, 11, 'g3W21A7SR', '2023-12-15 15:10:12', '2023-12-15 15:10:12', NULL, 1);
INSERT INTO `InventoryCheckNote` (`id`, `qtyDifferent`, `qtyAfterAdjust`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('QkaOCd5SR', 1, 19, 'g3W21A7SR', '2023-12-25 10:10:12', '2023-12-25 10:10:12', NULL, 1),
('tSpOCd5Sg', -4, 41, 'g3W21A7SR', '2023-12-20 09:05:23', '2023-12-20 09:05:23', NULL, 1);

INSERT INTO `InventoryCheckNoteDetail` (`inventoryCheckNoteId`, `bookId`, `initial`, `difference`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('BYwUzF5IR', 's100ma', 33, -3, 30, '2024-01-16 04:47:10', '2024-01-16 04:47:10', NULL, 1);
INSERT INTO `InventoryCheckNoteDetail` (`inventoryCheckNoteId`, `bookId`, `initial`, `difference`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('BYwUzF5IR', 'sipm2', 7, -1, 6, '2024-01-16 04:47:10', '2024-01-16 04:47:10', NULL, 1);
INSERT INTO `InventoryCheckNoteDetail` (`inventoryCheckNoteId`, `bookId`, `initial`, `difference`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('eFsOCOcIg', 'sipm2', 4, 1, 5, '2023-12-30 12:10:12', '2023-12-30 12:10:12', NULL, 1);
INSERT INTO `InventoryCheckNoteDetail` (`inventoryCheckNoteId`, `bookId`, `initial`, `difference`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('PXPDCOcIg', 'sdtls', 10, 1, 11, '2023-12-15 15:10:12', '2023-12-15 15:10:12', NULL, 1),
('QkaOCd5SR', 'sdoraemont12', 18, 1, 19, '2023-12-25 10:10:12', '2023-12-25 10:10:12', NULL, 1),
('tSpOCd5Sg', 'sdoraemont12', 20, -2, 18, '2023-12-20 09:05:23', '2023-12-20 09:05:23', NULL, 1),
('tSpOCd5Sg', 'sdoraemonv23', 20, -1, 19, '2023-12-20 09:05:23', '2023-12-20 09:05:23', NULL, 1),
('tSpOCd5Sg', 'sipm2', 5, -1, 4, '2023-12-20 09:05:23', '2023-12-20 09:05:23', NULL, 1);

INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `totalImportPrice`, `amountReceived`, `amountPriceUsePoint`, `pointUse`, `pointReceive`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('5k3RkFcSR', '2DiuRF5IR', 500000, 470000, 500000, 0, 0, 5000, 'g3W21A7SR', '2024-01-09 04:12:23', '2024-01-09 04:12:23', NULL, 1);
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `totalImportPrice`, `amountReceived`, `amountPriceUsePoint`, `pointUse`, `pointReceive`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('bcSkzFcIg', NULL, 85000, 70000, 85000, 0, 0, 0, 'g3W21A7SR', '2024-01-10 10:12:23', '2024-01-10 10:12:23', NULL, 1);
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `totalImportPrice`, `amountReceived`, `amountPriceUsePoint`, `pointUse`, `pointReceive`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('Cx16RF5Ig', NULL, 419000, 380000, 419000, 0, 0, 0, 'g3W21A7SR', '2024-01-08 02:12:23', '2024-01-08 02:12:23', NULL, 1);
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `totalImportPrice`, `amountReceived`, `amountPriceUsePoint`, `pointUse`, `pointReceive`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('czzCgK5Ig', '1utsgF5SR', 536000, 500000, 536000, 0, 0, 5360, 'g3W21A7SR', '2024-01-03 06:33:23', '2024-01-03 06:33:23', NULL, 1),
('eXGlzFcSR', 'ZqeugFcSR', 205000, 160000, 194790, 10210, 10210, 1948, 'g3W21A7SR', '2024-01-16 04:47:54', '2024-01-16 04:47:54', NULL, 1),
('lslQzFcSg', '2DiuRF5IR', 195000, 165000, 188536, 6464, 6464, 1885, 'g3W21A7SR', '2024-01-16 04:47:44', '2024-01-16 04:47:44', NULL, 1),
('NAf9gKcIR', 'ZqeugFcSR', 266000, 230000, 266000, 0, 0, 2660, 'g3W21A7SR', '2024-01-02 05:33:23', '2024-01-02 05:33:23', NULL, 1),
('oSdzkKcSg', NULL, 46000, 40000, 46000, 0, 0, 0, 'g3W21A7SR', '2024-01-12 10:12:23', '2024-01-12 10:12:23', NULL, 1),
('PRSjRFcIR', '2DiuRF5IR', 360000, 270000, 360000, 0, 0, 3600, 'g3W21A7SR', '2024-01-03 08:12:23', '2024-01-03 08:12:23', NULL, 1),
('QEIeRKcSR', '1utsgF5SR', 145000, 125000, 145000, 0, 0, 1450, 'g3W21A7SR', '2024-01-07 03:12:23', '2024-01-07 03:12:23', NULL, 1),
('r_-jgKcSg', '2DiuRF5IR', 150000, 130000, 146400, 3600, 3600, 1464, 'g3W21A7SR', '2024-01-04 09:12:23', '2024-01-04 09:12:23', NULL, 1),
('X_peRK5SR', '1utsgF5SR', 85000, 70000, 78190, 6810, 6810, 782, 'g3W21A7SR', '2024-01-07 05:12:23', '2024-01-07 05:12:23', NULL, 1),
('z-zqgK5Sg', 'ZqeugFcSR', 755000, 680000, 755000, 0, 0, 7550, 'g3W21A7SR', '2024-01-06 01:12:23', '2024-01-06 01:12:23', NULL, 1),
('ZA6QzF5Ig', '1utsgF5SR', 126000, 110000, 126000, 0, 0, 1260, 'g3W21A7SR', '2024-01-16 04:47:49', '2024-01-16 04:47:49', NULL, 1);

INSERT INTO `InvoiceDetail` (`invoiceId`, `bookId`, `bookName`, `qty`, `unitPrice`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('5k3RkFcSR', 'sipm2', 'Official IELTS Practice Materials 2 with DVD', 1, 500000, '2024-01-09 04:12:23', '2024-01-09 04:12:23', NULL, 1);
INSERT INTO `InvoiceDetail` (`invoiceId`, `bookId`, `bookName`, `qty`, `unitPrice`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('bcSkzFcIg', 'smb', 'Mắt biếc', 1, 85000, '2024-01-10 10:12:23', '2024-01-10 10:12:23', NULL, 1);
INSERT INTO `InvoiceDetail` (`invoiceId`, `bookId`, `bookName`, `qty`, `unitPrice`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('Cx16RF5Ig', 'dsslsln', 'Sapiens Lược Sử Loài Người', 1, 299000, '2024-01-08 02:12:23', '2024-01-08 02:12:23', NULL, 1);
INSERT INTO `InvoiceDetail` (`invoiceId`, `bookId`, `bookName`, `qty`, `unitPrice`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('Cx16RF5Ig', 'stlbt', 'Tôi là Bêtô', 1, 120000, '2024-01-08 02:12:23', '2024-01-08 02:12:23', NULL, 1),
('czzCgK5Ig', 'sgktoan5', 'Sách giáo khoa Toán lớp 5', 1, 18000, '2024-01-03 06:33:23', '2024-01-03 06:33:23', NULL, 1),
('czzCgK5Ig', 'sgktoan7', 'Sách giáo khoa Toán lớp 7', 1, 18000, '2024-01-03 06:33:23', '2024-01-03 06:33:23', NULL, 1),
('czzCgK5Ig', 'sipm2', 'Official IELTS Practice Materials 2 with DVD', 1, 500000, '2024-01-03 06:33:23', '2024-01-03 06:33:23', NULL, 1),
('eXGlzFcSR', 'smb', 'Mắt biếc', 1, 85000, '2024-01-16 04:47:54', '2024-01-16 04:47:54', NULL, 1),
('eXGlzFcSR', 'stlbt', 'Tôi là Bêtô', 1, 120000, '2024-01-16 04:47:54', '2024-01-16 04:47:54', NULL, 1),
('lslQzFcSg', 'sdoraemont12', 'Doraemon - Tập 12', 1, 30000, '2024-01-16 04:47:44', '2024-01-16 04:47:44', NULL, 1),
('lslQzFcSg', 'sdtls', 'Đi Tìm Lẽ Sống', 1, 80000, '2024-01-16 04:47:44', '2024-01-16 04:47:44', NULL, 1),
('lslQzFcSg', 'smb', 'Mắt biếc', 1, 85000, '2024-01-16 04:47:44', '2024-01-16 04:47:44', NULL, 1),
('NAf9gKcIR', 's100ma', '100 Món Ăn Ngày Thường', 1, 46000, '2024-01-02 05:33:23', '2024-01-02 05:33:23', NULL, 1),
('NAf9gKcIR', 'sdoraemont12', 'Doraemon - Tập 12', 2, 30000, '2024-01-02 05:33:23', '2024-01-02 05:33:23', NULL, 1),
('NAf9gKcIR', 'sdtls', 'Đi Tìm Lẽ Sống', 2, 80000, '2024-01-02 05:33:23', '2024-01-02 05:33:23', NULL, 1),
('oSdzkKcSg', 's100ma', '100 Món Ăn Ngày Thường', 1, 46000, '2024-01-12 10:12:23', '2024-01-12 10:12:23', NULL, 1),
('PRSjRFcIR', 'stlbt', 'Tôi là Bêtô', 3, 120000, '2024-01-03 08:12:23', '2024-01-03 08:12:23', NULL, 1),
('QEIeRKcSR', 'sdoraemont12', 'Doraemon - Tập 12', 1, 30000, '2024-01-07 03:12:23', '2024-01-07 03:12:23', NULL, 1),
('QEIeRKcSR', 'sdoraemonv23', 'Doraemon vol23. Nobita và những pháp sư gió bí ẩn', 1, 35000, '2024-01-07 03:12:23', '2024-01-07 03:12:23', NULL, 1),
('QEIeRKcSR', 'sdtls', 'Đi Tìm Lẽ Sống', 1, 80000, '2024-01-07 03:12:23', '2024-01-07 03:12:23', NULL, 1),
('r_-jgKcSg', 'sttgbct', 'Tôi tài giỏi, bạn cũng thế!', 1, 150000, '2024-01-04 09:12:23', '2024-01-04 09:12:23', NULL, 1),
('X_peRK5SR', 'smb', 'Mắt biếc', 1, 85000, '2024-01-07 05:12:23', '2024-01-07 05:12:23', NULL, 1),
('z-zqgK5Sg', 'sipm2', 'Official IELTS Practice Materials 2 with DVD', 1, 500000, '2024-01-06 01:12:23', '2024-01-06 01:12:23', NULL, 1),
('z-zqgK5Sg', 'smb', 'Mắt biếc', 3, 85000, '2024-01-06 01:12:23', '2024-01-06 01:12:23', NULL, 1),
('ZA6QzF5Ig', 's100ma', '100 Món Ăn Ngày Thường', 1, 46000, '2024-01-16 04:47:49', '2024-01-16 04:47:49', NULL, 1),
('ZA6QzF5Ig', 'sdtls', 'Đi Tìm Lẽ Sống', 1, 80000, '2024-01-16 04:47:49', '2024-01-16 04:47:49', NULL, 1);

INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `createdAt`, `updatedAt`, `deletedAt`, `imgUrl`, `isActive`) VALUES
('AzENRKcSg', 'Quản Thị Lý', '0983219471', 'TPHCM', 'quanly@gmail.com', 'c1fbf86b28f58b7e7338d46d72aad48f', 'menicMMqJjjEdgKxqlSVOuWLSrIEZrjieYuXQhNUCgXxLzjiyz', 'manager', '2024-01-16 04:20:53', '2024-01-16 04:56:51', NULL, 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Avatar%2Fnhanvien3.png?alt=media', 1);
INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `createdAt`, `updatedAt`, `deletedAt`, `imgUrl`, `isActive`) VALUES
('g3W21A7SR', 'Nguyễn Văn A', '1234567890', 'TPHCM', 'admin@gmail.com', '5e107317df151f6e8e0015c4f2ee7936', 'mVMxRDAHpAJfyzuiXWRELghNpynUqBKueSboGBcrwHUuzEWsms', 'admin', '2023-12-02 01:52:32', '2024-01-16 03:14:32', NULL, 'https://cdn-icons-png.flaticon.com/512/149/149071.png', 1);
INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `createdAt`, `updatedAt`, `deletedAt`, `imgUrl`, `isActive`) VALUES
('gv4RgFcSR', 'Nguyễn Kim Anh Thư', '0919676723', 'TPHCM', '21521495@gm.uit.edu.vn', '082608d3164cf7081caf7ae7e98bd583', 'tKUmRRGyXMsiHzTtEHvdDIxsCQMRrUoQbveMqPXEZRDeHSEUXB', 'staff', '2024-01-16 04:16:15', '2024-01-16 04:55:51', NULL, 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Avatar%2Fnhanvien1.png?alt=media', 1);
INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `createdAt`, `updatedAt`, `deletedAt`, `imgUrl`, `isActive`) VALUES
('q837RKcIg', 'Nguyễn Thị Thu Ngân', '0585885237', 'Thủ Đức', 'thungan@gmail.com', '178aa8880a34b7eba8c5bd7fcdb62c1d', 'PzaZHlJvqHmrCGNzTgQvGAkoJnUgkDsNPGoMkGRRxsWEjNmQkH', 'staff', '2024-01-16 04:19:24', '2024-01-16 04:56:37', NULL, 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Avatar%2Fnhanvien2.png?alt=media', 1);

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
('admin', 'Admin', '2023-12-02 01:52:40', '2024-01-16 04:21:34', NULL, 1);
INSERT INTO `Role` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('manager', 'Quản lí', '2024-01-03 19:26:02', '2024-01-04 12:26:11', NULL, 1);
INSERT INTO `Role` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('staff', 'Nhân viên', '2024-01-03 19:25:45', '2024-01-03 19:26:51', NULL, 1);

INSERT INTO `RoleFeature` (`roleId`, `featureId`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('admin', 'AUTHOR_CREATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1);
INSERT INTO `RoleFeature` (`roleId`, `featureId`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('admin', 'AUTHOR_UPDATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1);
INSERT INTO `RoleFeature` (`roleId`, `featureId`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('admin', 'AUTHOR_VIEW', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1);
INSERT INTO `RoleFeature` (`roleId`, `featureId`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('admin', 'BOOK_CREATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'BOOK_TITLE_CREATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'BOOK_TITLE_UPDATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'BOOK_TITLE_VIEW', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'BOOK_UPDATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'BOOK_VIEW', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'CATEGORY_CREATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'CATEGORY_UPDATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'CATEGORY_VIEW', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'CUSTOMER_CREATE', '2024-01-13 15:43:28', '2024-01-13 15:43:28', NULL, 1),
('admin', 'CUSTOMER_UPDATE_INFO', '2024-01-13 15:43:41', '2024-01-13 15:43:41', NULL, 1),
('admin', 'CUSTOMER_VIEW', '2024-01-13 15:43:58', '2024-01-13 15:43:58', NULL, 1),
('admin', 'IMPORT_NOTE_CREATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'IMPORT_NOTE_STATUS', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'IMPORT_NOTE_VIEW', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'INVENTORY_NOTE_CREATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'INVENTORY_NOTE_VIEW', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'INVOICE_CREATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'INVOICE_VIEW', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'PUBLISHER_CREATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'PUBLISHER_UPDATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'PUBLISHER_VIEW', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'REPORT_VIEW_SALE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'REPORT_VIEW_STOCK', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'REPORT_VIEW_SUPPLIER', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'SUPPLIER_CREATE', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'SUPPLIER_PAY', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'SUPPLIER_UPDATE_INFO', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'SUPPLIER_VIEW', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'USER_UPDATE_INFO', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('admin', 'USER_VIEW', '2024-01-02 17:02:08', '2024-01-02 17:02:08', NULL, 1),
('manager', 'AUTHOR_VIEW', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'BOOK_TITLE_VIEW', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'BOOK_VIEW', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'CATEGORY_VIEW', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'CUSTOMER_CREATE', '2024-01-16 04:23:08', '2024-01-16 04:23:08', NULL, 1),
('manager', 'CUSTOMER_UPDATE_INFO', '2024-01-16 04:23:08', '2024-01-16 04:23:08', NULL, 1),
('manager', 'CUSTOMER_VIEW', '2024-01-16 04:23:08', '2024-01-16 04:23:08', NULL, 1),
('manager', 'IMPORT_NOTE_VIEW', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'INVENTORY_NOTE_CREATE', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'INVENTORY_NOTE_VIEW', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'INVOICE_CREATE', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'INVOICE_VIEW', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'PUBLISHER_VIEW', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'REPORT_VIEW_SALE', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'REPORT_VIEW_STOCK', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'REPORT_VIEW_SUPPLIER', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'SUPPLIER_VIEW', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'USER_UPDATE_INFO', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('manager', 'USER_VIEW', '2024-01-03 19:30:19', '2024-01-03 19:30:19', NULL, 1),
('staff', 'AUTHOR_VIEW', '2024-01-03 19:34:50', '2024-01-03 19:34:50', NULL, 1),
('staff', 'BOOK_TITLE_VIEW', '2024-01-03 19:34:50', '2024-01-03 19:34:50', NULL, 1),
('staff', 'BOOK_VIEW', '2024-01-03 19:34:50', '2024-01-03 19:34:50', NULL, 1),
('staff', 'CATEGORY_VIEW', '2024-01-03 19:34:50', '2024-01-03 19:34:50', NULL, 1),
('staff', 'CUSTOMER_CREATE', '2024-01-16 04:23:49', '2024-01-16 04:23:49', NULL, 1),
('staff', 'CUSTOMER_UPDATE_INFO', '2024-01-16 04:23:49', '2024-01-16 04:23:49', NULL, 1),
('staff', 'CUSTOMER_VIEW', '2024-01-16 04:23:49', '2024-01-16 04:23:49', NULL, 1),
('staff', 'INVOICE_CREATE', '2024-01-03 19:34:50', '2024-01-03 19:34:50', NULL, 1),
('staff', 'INVOICE_VIEW', '2024-01-03 19:34:50', '2024-01-03 19:34:50', NULL, 1);

INSERT INTO `ShopGeneral` (`id`, `name`, `email`, `phone`, `address`, `wifiPass`, `accumulatePointPercent`, `usePointPercent`) VALUES
('shop', 'Book Store', '', '', '', 'BookStore123', 0.01, 1);


INSERT INTO `StockChangeHistory` (`id`, `bookId`, `qty`, `qtyLeft`, `type`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('4b3WjO5Ig', 's100ma', 10, 10, 'Import', '2023-12-01 23:55:23', '2023-12-01 23:55:23', NULL, 1);
INSERT INTO `StockChangeHistory` (`id`, `bookId`, `qty`, `qtyLeft`, `type`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('5k3RkFcSR', 'sipm2', -1, 2, 'Sell', '2024-01-09 04:12:23', '2024-01-09 04:12:23', NULL, 1);
INSERT INTO `StockChangeHistory` (`id`, `bookId`, `qty`, `qtyLeft`, `type`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('9MhHCd5Ig', 'dsslsln', 10, 10, 'Import', '2023-12-02 17:10:22', '2023-12-02 17:10:22', NULL, 1);
INSERT INTO `StockChangeHistory` (`id`, `bookId`, `qty`, `qtyLeft`, `type`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('9MhHCd5Ig', 's100ma', 20, 30, 'Import', '2023-12-02 17:10:22', '2023-12-02 17:10:22', NULL, 1),
('bcSkzFcIg', 'smb', -1, 15, 'Sell', '2024-01-10 10:12:23', '2024-01-10 10:12:23', NULL, 1),
('BYwUzF5IR', 's100ma', -3, 30, 'Modify', '2024-01-16 04:47:10', '2024-01-16 04:47:10', NULL, 1),
('BYwUzF5IR', 'sipm2', -1, 6, 'Modify', '2024-01-16 04:47:10', '2024-01-16 04:47:10', NULL, 1),
('Cx16RF5Ig', 'dsslsln', -1, 9, 'Sell', '2024-01-08 02:12:23', '2024-01-08 02:12:23', NULL, 1),
('Cx16RF5Ig', 'stlbt', -1, 6, 'Sell', '2024-01-08 02:12:23', '2024-01-08 02:12:23', NULL, 1),
('czzCgK5Ig', 'sgktoan5', -1, 49, 'Sell', '2024-01-03 06:33:23', '2024-01-03 06:33:23', NULL, 1),
('czzCgK5Ig', 'sgktoan7', -1, 49, 'Sell', '2024-01-03 06:33:23', '2024-01-03 06:33:23', NULL, 1),
('czzCgK5Ig', 'sipm2', -1, 4, 'Sell', '2024-01-03 06:33:23', '2024-01-03 06:33:23', NULL, 1),
('DeT7Cd5Ig', 'smb', 10, 20, 'Import', '2023-12-07 13:22:12', '2023-12-07 13:22:12', NULL, 1),
('eFsOCOcIg', 'sipm2', 1, 5, 'Modify', '2023-12-30 12:10:12', '2023-12-30 12:10:12', NULL, 1),
('eXGlzFcSR', 'smb', -1, 13, 'Sell', '2024-01-16 04:47:54', '2024-01-16 04:47:54', NULL, 1),
('eXGlzFcSR', 'stlbt', -1, 15, 'Sell', '2024-01-16 04:47:54', '2024-01-16 04:47:54', NULL, 1),
('hAgSCd5Sg', 'sgktoan5', 50, 50, 'Import', '2023-12-04 08:34:23', '2023-12-04 08:34:23', NULL, 1),
('hAgSCd5Sg', 'sgktoan7', 50, 50, 'Import', '2023-12-04 08:34:23', '2023-12-04 08:34:23', NULL, 1),
('hAgSCd5Sg', 'sipm2', 5, 5, 'Import', '2023-12-04 08:34:23', '2023-12-04 08:34:23', NULL, 1),
('iAJGCOcIR', 'sdoraemont12', 20, 20, 'Import', '2023-12-02 19:22:14', '2023-12-02 19:22:14', NULL, 1),
('iAJGCOcIR', 'sdoraemonv23', 20, 20, 'Import', '2023-12-02 19:22:14', '2023-12-02 19:22:14', NULL, 1),
('iAJGCOcIR', 'smb', 10, 10, 'Import', '2023-12-02 19:22:14', '2023-12-02 19:22:14', NULL, 1),
('KMNnjdcSR', 'stlbt', 10, 10, 'Import', '2023-12-09 18:26:12', '2023-12-09 18:26:12', NULL, 1),
('KMNnjdcSR', 'sttgbct', 30, 30, 'Import', '2023-12-09 18:26:12', '2023-12-09 18:26:12', NULL, 1),
('lslQzFcSg', 'sdoraemont12', -1, 15, 'Sell', '2024-01-16 04:47:44', '2024-01-16 04:47:44', NULL, 1),
('lslQzFcSg', 'sdtls', -1, 12, 'Sell', '2024-01-16 04:47:44', '2024-01-16 04:47:44', NULL, 1),
('lslQzFcSg', 'smb', -1, 14, 'Sell', '2024-01-16 04:47:44', '2024-01-16 04:47:44', NULL, 1),
('NAf9gKcIR', 's100ma', -1, 29, 'Sell', '2024-01-02 05:33:23', '2024-01-02 05:33:23', NULL, 1),
('NAf9gKcIR', 'sdoraemont12', -2, 17, 'Sell', '2024-01-02 05:33:23', '2024-01-02 05:33:23', NULL, 1),
('NAf9gKcIR', 'sdtls', -2, 9, 'Sell', '2024-01-02 05:33:23', '2024-01-02 05:33:23', NULL, 1),
('oSdzkKcSg', 's100ma', -1, 28, 'Sell', '2024-01-12 10:12:23', '2024-01-12 10:12:23', NULL, 1),
('p6-LkKcIR', 's100ma', 5, 33, 'Import', '2024-01-16 04:45:30', '2024-01-16 04:45:30', NULL, 1),
('p6-LkKcIR', 'sdtls', 5, 13, 'Import', '2024-01-16 04:45:30', '2024-01-16 04:45:30', NULL, 1),
('p6-LkKcIR', 'sipm2', 5, 7, 'Import', '2024-01-16 04:45:30', '2024-01-16 04:45:30', NULL, 1),
('p6-LkKcIR', 'stlbt', 10, 16, 'Import', '2024-01-16 04:45:30', '2024-01-16 04:45:30', NULL, 1),
('p6-LkKcIR', 'sttgbct', 10, 39, 'Import', '2024-01-16 04:45:30', '2024-01-16 04:45:30', NULL, 1),
('PRSjRFcIR', 'stlbt', -3, 7, 'Sell', '2024-01-03 08:12:23', '2024-01-03 08:12:23', NULL, 1),
('PXPDCOcIg', 'sdtls', 1, 11, 'Modify', '2023-12-15 15:10:12', '2023-12-15 15:10:12', NULL, 1),
('QEIeRKcSR', 'sdoraemont12', -1, 16, 'Sell', '2024-01-07 03:12:23', '2024-01-07 03:12:23', NULL, 1),
('QEIeRKcSR', 'sdoraemonv23', -1, 18, 'Sell', '2024-01-07 03:12:23', '2024-01-07 03:12:23', NULL, 1),
('QEIeRKcSR', 'sdtls', -1, 8, 'Sell', '2024-01-07 03:12:23', '2024-01-07 03:12:23', NULL, 1),
('QFmvjd5Ig', 'sdtls', 10, 10, 'Import', '2023-12-12 22:26:12', '2023-12-12 22:26:12', NULL, 1),
('QkaOCd5SR', 'sdoraemont12', 1, 19, 'Modify', '2023-12-25 10:10:12', '2023-12-25 10:10:12', NULL, 1),
('r_-jgKcSg', 'sttgbct', -1, 29, 'Sell', '2024-01-04 09:12:23', '2024-01-04 09:12:23', NULL, 1),
('tSpOCd5Sg', 'sdoraemont12', -2, 18, 'Modify', '2023-12-20 09:05:23', '2023-12-20 09:05:23', NULL, 1),
('tSpOCd5Sg', 'sdoraemonv23', -1, 19, 'Modify', '2023-12-20 09:05:23', '2023-12-20 09:05:23', NULL, 1),
('tSpOCd5Sg', 'sipm2', -1, 4, 'Modify', '2023-12-20 09:05:23', '2023-12-20 09:05:23', NULL, 1),
('X_peRK5SR', 'smb', -1, 16, 'Sell', '2024-01-07 05:12:23', '2024-01-07 05:12:23', NULL, 1),
('z-zqgK5Sg', 'sipm2', -1, 3, 'Sell', '2024-01-06 01:12:23', '2024-01-06 01:12:23', NULL, 1),
('z-zqgK5Sg', 'smb', -3, 17, 'Sell', '2024-01-06 01:12:23', '2024-01-06 01:12:23', NULL, 1),
('ZA6QzF5Ig', 's100ma', -1, 29, 'Sell', '2024-01-16 04:47:49', '2024-01-16 04:47:49', NULL, 1),
('ZA6QzF5Ig', 'sdtls', -1, 11, 'Sell', '2024-01-16 04:47:49', '2024-01-16 04:47:49', NULL, 1);

INSERT INTO `StockReport` (`id`, `timeFrom`, `initial`, `sell`, `import`, `modify`, `final`, `timeTo`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('c-qE6O5Sg', '2023-11-30 17:00:00', 0, 0, 255, -1, 254, '2023-12-31 16:59:59', '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1);


INSERT INTO `StockReportDetail` (`reportId`, `bookId`, `initial`, `sell`, `import`, `modify`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('c-qE6O5Sg', 'dsslsln', 0, 0, 10, 0, 10, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1);
INSERT INTO `StockReportDetail` (`reportId`, `bookId`, `initial`, `sell`, `import`, `modify`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('c-qE6O5Sg', 's100ma', 0, 0, 30, 0, 30, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1);
INSERT INTO `StockReportDetail` (`reportId`, `bookId`, `initial`, `sell`, `import`, `modify`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('c-qE6O5Sg', 'sdoraemont12', 0, 0, 20, -1, 19, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1);
INSERT INTO `StockReportDetail` (`reportId`, `bookId`, `initial`, `sell`, `import`, `modify`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('c-qE6O5Sg', 'sdoraemonv23', 0, 0, 20, -1, 19, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1),
('c-qE6O5Sg', 'sdtls', 0, 0, 10, 1, 11, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1),
('c-qE6O5Sg', 'sgktoan5', 0, 0, 50, 0, 50, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1),
('c-qE6O5Sg', 'sgktoan7', 0, 0, 50, 0, 50, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1),
('c-qE6O5Sg', 'sipm2', 0, 0, 5, 0, 5, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1),
('c-qE6O5Sg', 'smb', 0, 0, 20, 0, 20, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1),
('c-qE6O5Sg', 'stlbt', 0, 0, 10, 0, 10, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1),
('c-qE6O5Sg', 'sttgbct', 0, 0, 30, 0, 30, '2024-01-16 04:11:12', '2024-01-16 04:11:12', NULL, 1);

INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('ncc1980', '1980 Book', '1980books@gmail.com', '0345689012', 0, '2023-12-19 01:09:49', '2023-12-30 11:25:23', NULL, 1);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('nccapb', 'Alpha Books', 'alphabooks@gmail.com', '0123456784', 200000, '2023-12-19 01:08:18', '2023-12-28 10:25:23', NULL, 1);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('nccfn', 'First News', 'firstnews@gmail.com', '0123456785', 5000000, '2023-12-19 01:08:18', '2024-01-16 04:46:20', NULL, 1);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('ncchnb', 'Hanoi Books', 'hanoibooks@gmail.com', '0123456782', 0, '2023-12-19 01:06:25', '2024-01-16 03:17:43', NULL, 1),
('ncchtt', 'Hoa học trò', 'hoahoctro@gmail.com', '0123456788', 0, '2023-12-19 01:10:28', '2023-12-26 10:25:23', NULL, 1),
('ncckd', 'Kim Đồng', 'kimdong@gmail.com', '0123456781', 4000000, '2023-12-19 01:06:25', '2024-01-16 04:46:36', NULL, 1),
('nccnn', 'Nhã Nam', 'nhanam@gmail.com', '0123456780', 0, '2023-12-19 01:05:13', '2024-01-16 03:19:49', NULL, 1),
('nccpdb', 'PandaBooks', 'pandabooks@gmail.com', '0123456783', 0, '2023-12-19 01:08:18', '2023-12-19 01:08:18', NULL, 1),
('nccttt', 'Tri Thức Trẻ', 'trithuctre@gmail.com', '0123456786', 4000000, '2023-12-19 01:09:49', '2024-01-16 04:46:46', NULL, 1);

INSERT INTO `SupplierDebt` (`id`, `supplierId`, `qty`, `qtyLeft`, `type`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('_Di8zKcSR', 'nccttt', -500000, 4000000, 'Pay', 'g3W21A7SR', '2024-01-16 04:46:46', '2024-01-16 04:46:46', NULL, 1);
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `qty`, `qtyLeft`, `type`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('0jUszFcSR', 'ncckd', -800000, 4000000, 'Pay', 'g3W21A7SR', '2024-01-16 04:46:36', '2024-01-16 04:46:36', NULL, 1);
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `qty`, `qtyLeft`, `type`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('4b3WjO5Ig', 'nccapb', 400000, 400000, 'Debt', 'g3W21A7SR', '2023-12-01 23:55:23', '2023-12-01 23:55:23', NULL, 1);
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `qty`, `qtyLeft`, `type`, `createdBy`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('4woaCd5Sg', 'nccapb', -200000, 200000, 'Pay', 'g3W21A7SR', '2023-12-25 09:05:23', '2023-12-25 09:05:23', NULL, 1),
('9MhHCd5Ig', 'ncc1980', 3700000, 3700000, 'Debt', 'g3W21A7SR', '2023-12-02 17:10:22', '2023-12-02 17:10:22', NULL, 1),
('DeT7Cd5Ig', 'nccttt', 700000, 700000, 'Debt', 'g3W21A7SR', '2023-12-07 13:22:12', '2023-12-07 13:22:12', NULL, 1),
('hAgSCd5Sg', 'nccttt', 3850000, 4550000, 'Debt', 'g3W21A7SR', '2023-12-04 08:34:23', '2023-12-04 08:34:23', NULL, 1),
('HNDykKcSR', 'nccfn', -100000, 5000000, 'Pay', 'g3W21A7SR', '2024-01-16 04:46:20', '2024-01-16 04:46:20', NULL, 1),
('iAJGCOcIR', 'ncckd', 1800000, 1800000, 'Debt', 'g3W21A7SR', '2023-12-02 19:22:14', '2023-12-02 19:22:14', NULL, 1),
('KMNnjdcSR', 'ncckd', 4800000, 6600000, 'Debt', 'g3W21A7SR', '2023-12-09 18:26:12', '2023-12-09 18:26:12', NULL, 1),
('OvwbjO5Ig', 'ncc1980', -3700000, 0, 'Pay', 'g3W21A7SR', '2023-12-26 10:25:23', '2023-12-26 10:25:23', NULL, 1),
('p6-LkKcIR', 'nccfn', 5100000, 5100000, 'Debt', 'g3W21A7SR', '2024-01-16 04:45:30', '2024-01-16 04:45:30', NULL, 1),
('wdf0Cd5Sg', 'ncckd', -1800000, 4800000, 'Pay', 'g3W21A7SR', '2023-12-28 10:25:23', '2023-12-28 10:25:23', NULL, 1),
('ZmD1jdcIR', 'nccttt', -50000, 4500000, 'Pay', 'g3W21A7SR', '2023-12-30 11:25:23', '2023-12-30 11:25:23', NULL, 1);

INSERT INTO `SupplierDebtReport` (`id`, `timeFrom`, `timeTo`, `initial`, `debt`, `pay`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('_poy6d5Sg', '2023-11-30 17:00:00', '2023-12-31 16:59:59', 0, 15250000, -5750000, 9500000, '2024-01-16 04:11:30', '2024-01-16 04:11:30', NULL, 1);


INSERT INTO `SupplierDebtReportDetail` (`reportId`, `supplierId`, `initial`, `debt`, `pay`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('_poy6d5Sg', 'ncc1980', 0, 3700000, -3700000, 0, '2024-01-16 04:11:30', '2024-01-16 04:11:30', NULL, 1);
INSERT INTO `SupplierDebtReportDetail` (`reportId`, `supplierId`, `initial`, `debt`, `pay`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('_poy6d5Sg', 'nccapb', 0, 400000, -200000, 200000, '2024-01-16 04:11:30', '2024-01-16 04:11:30', NULL, 1);
INSERT INTO `SupplierDebtReportDetail` (`reportId`, `supplierId`, `initial`, `debt`, `pay`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('_poy6d5Sg', 'ncckd', 0, 6600000, -1800000, 4800000, '2024-01-16 04:11:30', '2024-01-16 04:11:30', NULL, 1);
INSERT INTO `SupplierDebtReportDetail` (`reportId`, `supplierId`, `initial`, `debt`, `pay`, `final`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('_poy6d5Sg', 'nccttt', 0, 4550000, -50000, 4500000, '2024-01-16 04:11:30', '2024-01-16 04:11:30', NULL, 1);


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;