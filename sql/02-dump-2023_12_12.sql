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
CREATE TABLE `Author` (
  `id` varchar(12) NOT NULL,
  `name` text NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
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

DROP TABLE IF EXISTS `BookChangeHistory`;
CREATE TABLE `BookChangeHistory` (
  `id` varchar(12) NOT NULL,
  `bookId` varchar(12) NOT NULL,
  `amount` float NOT NULL,
  `amountLeft` float NOT NULL,
  `type` enum('Sell','Import','Modify') NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`,`bookId`)
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Feature`;
CREATE TABLE `Feature` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` text,
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
  `totalPrice` float DEFAULT '0',
  `status` enum('InProgress','Done','Cancel') DEFAULT 'InProgress',
  `closedBy` varchar(12) DEFAULT NULL,
  `closedAt` datetime DEFAULT NULL,
  `createdBy` varchar(12) NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ImportNoteDetail`;
CREATE TABLE `ImportNoteDetail` (
  `importNoteId` varchar(12) NOT NULL,
  `bookId` varchar(12) NOT NULL,
  `price` float NOT NULL,
  `qtyImport` float DEFAULT '0',
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`importNoteId`,`bookId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `InventoryCheckNote`;
CREATE TABLE `InventoryCheckNote` (
  `id` varchar(12) NOT NULL,
  `qtyDifferent` float NOT NULL,
  `qtyAfterAdjust` float NOT NULL,
  `createBy` varchar(12) NOT NULL,
  `createAt` datetime DEFAULT (now()),
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `InventoryCheckNoteDetail`;
CREATE TABLE `InventoryCheckNoteDetail` (
  `inventoryCheckNoteId` varchar(12) NOT NULL,
  `bookId` varchar(12) NOT NULL,
  `initial` float NOT NULL,
  `difference` float NOT NULL,
  `final` float NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`inventoryCheckNoteId`,`bookId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Invoice`;
CREATE TABLE `Invoice` (
  `id` varchar(13) NOT NULL,
  `totalPrice` float NOT NULL,
  `qtyReceived` float NOT NULL,
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
  `qty` float NOT NULL,
  `unitPrice` float NOT NULL,
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
  `isActive` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`) USING BTREE,
  KEY `roleId` (`roleId`),
  CONSTRAINT `MUser_ibfk_1` FOREIGN KEY (`roleId`) REFERENCES `Role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Publisher`;
CREATE TABLE `Publisher` (
  `id` varchar(12) NOT NULL,
  `name` varchar(50) NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
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
  `id` varchar(12) NOT NULL,
  `name` varchar(12) NOT NULL,
  `email` float NOT NULL,
  `phone` float NOT NULL,
  `address` text,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `StockReport`;
CREATE TABLE `StockReport` (
  `id` varchar(12) NOT NULL,
  `year` int NOT NULL,
  `month` int NOT NULL,
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
  `initial` float NOT NULL,
  `sell` float NOT NULL,
  `import` float NOT NULL,
  `modify` float NOT NULL,
  `final` float NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`reportId`,`bookId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Supplier`;
CREATE TABLE `Supplier` (
  `id` varchar(12) NOT NULL,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `phone` varchar(11) NOT NULL,
  `debt` float DEFAULT '0',
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
  `qty` float NOT NULL,
  `qtyLeft` float NOT NULL,
  `type` enum('Debt','Pay') NOT NULL,
  `createAt` datetime DEFAULT (now()),
  `createBy` varchar(9) NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`,`supplierId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `SupplierDebtDetail`;
CREATE TABLE `SupplierDebtDetail` (
  `reportId` varchar(12) NOT NULL,
  `supplierId` varchar(12) NOT NULL,
  `initial` float NOT NULL,
  `arise` float NOT NULL,
  `final` float NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`reportId`,`supplierId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `SupplierDebtReport`;
CREATE TABLE `SupplierDebtReport` (
  `id` varchar(12) NOT NULL,
  `year` int NOT NULL,
  `month` int NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tgak', 'Adam Khoo', '2023-12-02 01:51:49', '2023-12-02 01:51:49', NULL, 1);
INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tgnna', 'Nguyễn Nhật Ánh', '2023-12-02 01:51:49', '2023-12-02 01:51:49', NULL, 1);
INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tgvef', 'Viktor E Frankl', '2023-12-02 01:51:49', '2023-12-02 01:51:49', NULL, 1);





INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('stlbt', 'Tôi là Bêtô', 'Một tác phẩm của Nguyễn Nhật Ánh', 'tgnna', 'dmtt|dmtruyen', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1);
INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('sttgbct', 'Tôi tài giỏi, bạn cũng thế!', 'Tôi tài giỏi, bạn cũng thế! (nhan đề gốc tiếng Anh: I Am Gifted, So Are You!) là quyển sách bán chạy nhất của doanh nhân người Singapore Adam Khoo, viết về những phương pháp học tập tiên tiến. Quyển sách đã được dịch ra hàng chục thứ tiếng, trong đó Tôi tài giỏi, bạn cũng thế! là phiên bản tiếng Việt được dịch bởi hai dịch giả nổi tiếng Trần Đăng Khoa và Uông Xuân Vy của TGM Books. Tại Việt Nam, quyển sách đã trở thành một hiện tượng giáo dục trong những năm 2009-2011 và đạt được nhiều thành tựu trong lĩnh vực xuất bản, tạo ra kỷ lục mới cho ngành xuất bản Việt Nam với hơn 200.000 bản in được bán ra và hơn 400.000 e-book được phân phối.', 'tgak', 'dmkns', '2023-12-10 16:07:24', '2023-12-11 09:50:45', NULL, 1);
INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('UDKbP0vIR', 'Tôi tài giỏi, bạn cũng thế!', 'Tôi tài giỏi, bạn cũng thế! (nhan đề gốc tiếng Anh: I Am Gifted, So Are You!) là quyển sách bán chạy nhất của doanh nhân người Singapore Adam Khoo, viết về những phương pháp học tập tiên tiến. Quyển sách đã được dịch ra hàng chục thứ tiếng, trong đó Tôi tài giỏi, bạn cũng thế! là phiên bản tiếng Việt được dịch bởi hai dịch giả nổi tiếng Trần Đăng Khoa và Uông Xuân Vy của TGM Books. Tại Việt Nam, quyển sách đã trở thành một hiện tượng giáo dục trong những năm 2009-2011 và đạt được nhiều thành tựu trong lĩnh vực xuất bản, tạo ra kỷ lục mới cho ngành xuất bản Việt Nam với hơn 200.000 bản in được bán ra và hơn 400.000 e-book được phân phối.', 'tgak', 'dmkns', '2023-12-12 02:39:46', '2023-12-12 02:39:46', NULL, 1);

INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmkns', 'Kỹ năng sống', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1);
INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmsgk', 'Sách giáo khoa', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1);
INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmtruyen', 'Truyện', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1);
INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmtt', 'Tiểu thuyết', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1);

INSERT INTO `Feature` (`id`, `description`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_CREATE', 'Tạo tác giả', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1);
INSERT INTO `Feature` (`id`, `description`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_DELETE', 'Xóa tác giả', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1);
INSERT INTO `Feature` (`id`, `description`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_UPDATE', 'Chỉnh sửa thông tin tác giả', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1);
INSERT INTO `Feature` (`id`, `description`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('AUTHOR_VIEW', 'Xem tác giả', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('BOOK_CREATE', 'Tạo sách', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('BOOK_DELETE', 'Xóa sách', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('BOOK_TITLE_CREATE', 'Tạo đầu sách', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('BOOK_TITLE_DELETE', 'Xóa đầu sách', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('BOOK_TITLE_UPDATE', 'Chỉnh sửa thông tin đầu sách', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('BOOK_TITLE_VIEW', 'Xem đầu sách', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('BOOK_UPDATE', 'Chỉnh sửa thông tin sách', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('BOOK_VIEW', 'Xem sách', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('CATEGORY_CREATE', 'Tạo danh mục', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('CATEGORY_DELETE', 'Xóa danh mục', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('CATEGORY_UPDATE', 'Chỉnh sửa thông tin danh mục', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('CATEGORY_VIEW', 'Xem danh mục', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('IMPORT_NOTE_CREATE', 'Tạo phiếu nhập', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('IMPORT_NOTE_STATUS', 'Chỉnh sửa trạng thái phiếu nhập', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('IMPORT_NOTE_VIEW', 'Xem phiếu nhập', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('INVENTORY_NOTE_CREATE', 'Tạo phiếu kiểm kho', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('INVENTORY_NOTE_VIEW', 'Xem phiếu kiểm kho', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('PUBLISHER_CREATE', 'Tạo nhà sản xuất', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('PUBLISHER_VIEW', 'Xem nhà sản xuất', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('SUPPLIER_CREATE', 'Tạo nhà cung cấp', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('SUPPLIER_PAY', 'Trả nợ nhà cung cấp', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('SUPPLIER_UPDATE_INFO', 'Chỉnh sửa thông tin nhà cung cấp', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('SUPPLIER_VIEW', 'Xem nhà cung cấp', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('USER_UPDATE_INFO', 'Chỉnh sửa thông tin người dùng', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('USER_UPDATE_STATE', 'Chỉnh sửa trạng thái', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1),
('USER_VIEW', 'Xem người dùng', '2023-12-13 08:54:39', '2023-12-13 08:54:39', NULL, 1);













INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('bgIqwQSIg', 'user', '', '', 'user@gmail.com', '0dd71ba5a82e98ccdc6f5edb6fb870a5', 'ByVwWucjSGZkozLFeQcopssBrHPbCHoqRuUCFUbpfIhhqGUujj', 'user', '2023-12-02 01:52:32', '2023-12-04 01:24:10', NULL, 1);
INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('g3W21A7SR', 'admin', '1234567890', '', 'admin@gmail.com', '5e107317df151f6e8e0015c4f2ee7936', 'mVMxRDAHpAJfyzuiXWRELghNpynUqBKueSboGBcrwHUuzEWsms', 'admin', '2023-12-02 01:52:32', '2023-12-04 01:24:10', NULL, 1);


INSERT INTO `Publisher` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('nxbdk', 'Kim Đồng', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1);
INSERT INTO `Publisher` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('nxbgd', 'Giáo dục', '2023-12-02 01:52:21', '2023-12-10 16:07:18', NULL, 1);


INSERT INTO `Role` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('admin', 'admin', '2023-12-02 01:52:40', '2023-12-02 01:52:40', NULL, 1);
INSERT INTO `Role` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('user', 'user', '2023-12-02 01:52:40', '2023-12-02 01:52:40', NULL, 1);


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
('admin', 'PUBLISHER_CREATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'PUBLISHER_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'SUPPLIER_CREATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'SUPPLIER_PAY', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'SUPPLIER_UPDATE_INFO', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'SUPPLIER_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'USER_UPDATE_INFO', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'USER_UPDATE_STATE', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('admin', 'USER_VIEW', '2023-12-12 08:46:33', '2023-12-12 08:46:33', NULL, 1),
('user', 'AUTHOR_CREATE', '2023-12-12 08:48:06', '2023-12-12 08:48:06', NULL, 1);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-12 10:44:18
