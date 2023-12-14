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

INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tg001', 'Nguyễn Trung Hậu', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1);
INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tg002', 'Lượng Thúc', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1);
INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tg003', 'Cao Vương', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1);
INSERT INTO `Author` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('tg004', 'Thích Nhật Từ', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg005', 'Fujimaru', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg006', 'Pha Lê', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg007', 'Diệp Lạc Vô Tâm', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg008', 'Bộ Giáo Dục và Đào Tạo', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg009', 'Nguyễn Cừ', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg010', 'Ở Đây Zui Nè', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg011', 'Huỳnh Thái Ngọc', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg012', 'Claire Belton', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg013', 'Tonton House', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg014', 'Higashino Keigo', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('tg015', 'Adam Khoo', '2023-12-02 01:51:49', '2023-12-14 14:47:40', NULL, 1),
('tg016', 'Nguyễn Nhật Ánh', '2023-12-14 14:42:49', '2023-12-14 14:47:51', NULL, 1);

INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('book001', 'Khởi Nghiệp Trên Xe Lăn', 'Một tác phẩm đầy cảm xúc với nhiều câu chuyện thú vị, cùng những chiêm nghiệm sâu sắc về hành trình khởi nghiệp trên xe lăn của tác giả.', 'tg001', 'dmkt', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1);
INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('book002', 'Tư duy kinh tế', 'Làm sao chúng ta có thể sống mà không đề cập tới vấn đề kinh tế. Trong tháp nhu cầu của Maslow, từ nhu cầu ăn uống, ngủ nghỉ đến nhu cầu an toàn như an ninh, vật chất tinh thần đến nhu cầu được quý trọng về địa vị, tôn vinh. Dù chúng ta có thừa nhận hay không, kinh tế giúp chúng ta thỏa mãn những nhu cầu đó. Mỗi một thế hệ đều phải đối mặt với vấn đề kinh tế muôn thuở, không ai có thể trốn tránh. Vậy làm thế nào để đối mặt, hiểu tường tận và vận dụng tốt những triết lý kinh tế trong cuộc sống. Câu trả lời nằm trong cuốn sách này.', 'tg001', 'dmkt', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1);
INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('book003', 'Biến Tri Thức Thành Mỏ Vàng', 'Cuốn sách này được chia thành năm phần chính, bao gồm: Chạm tay vào thị trường tỷ đô, Lộ trình xây dựng khóa học trực tuyến từ A - Z, Bí quyết thu hút học viên đến khóa học, Bí mật của webinar và bootcamp, Suy nghĩ như một nhà đào tạo triệu phú. Mỗi phần đều bao gồm các chương chi tiết, với các bước cụ thể, chiến lược và bài tập để giúp bạn áp dụng những gì đã học vào thực tế.', 'tg003', 'dmkt', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1);
INSERT INTO `BookTitle` (`id`, `name`, `desc`, `authorIds`, `categoryIds`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('book004', 'Sống Vui Sống Khỏe', 'Tất cả chúng ta ai cũng chỉ mong có được một cuộc sống hạnh phúc, an vui dài lâu. Vậy nên, với Sống vui sống khỏe, tôi mong rằng quý vị độc giả, quý vị Phật tử sẽ có thêm được nhiều thông tin bổ ích, quý báu để vững bước trên hành trình tìm kiếm hạnh phúc ấy.', 'tg004', 'dmkns', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book005', 'Thần Chết Làm Thêm 300 Yên/Giờ', 'Ẩn chứa bên trong câu chuyện là sự cảm động về những vấn vương, mong ước rất đỗi xót xa của những người đã mất.', 'tg005', 'dmln', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book006', 'Bộ Sách Ăn Lành - Nấu Sạch', 'Dinh dưỡng là một đề tài rối rắm. Không chỉ là chuyện ăn, nó còn gắn liền với khoa học, văn hóa, truyền thống, sức khỏe, cảm giác được chăm sóc, sự no ấm, sự hạnh phúc… Bên cạnh đó, nấu nướng cũng không kém phần phức tạp. Mỗi món ngon ra đời là kết quả của cả quá trình đằng sau - không đơn thuần là kỹ thuật nấu nướng mà là chuyện về muối đường tương mắm, các dụng cụ, nồi niêu, bếp nấu…', 'tg006', 'dmna', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book007', 'Gió Ngừng Thổi, Tình Còn Vương', 'Người trong nhân thế đều biết tình ái nơi cõi trần là điều đau khổ nhưng có bao nhiêu người thực sự hiểu thấu được nó? Nếu đã bước vào hồng trần, chi bằng hãy dốc hết lòng mà yêu. Bất luận là tình sâu hay duyên mỏng, chỉ mong đời này không hối hận.', 'tg007', 'dmnt', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book008', 'Giáo Trình Kinh Tế Chính Trị Mác - Lênin', 'Giáo trình do PGS.TS. Ngô Tuấn Nghĩa chủ biên, cùng tập thể tác giả là những nhà nghiên cứu, nhà giáo dục có nhiều kinh nghiệm tổ chức biên soạn.', 'tg008', 'dmsgk', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book009', 'Giáo Trình Triết Học Mác - Lênin', 'Giáo trình do Ban biên soạn gồm các tác giả là nhà nghiên cứu, nhà giáo dục thuộc Viện Triết học - Học viện Chính trị quốc gia Hồ Chí Minh, các học viện, trường đại học, Viện Triết học - Viện Hàn lâm Khoa học xã hội Việt Nam, tổ chức biên soạn trên cơ sở kế thừa những kết quả nghiên cứu trước đây, đồng thời bổ sung nhiều nội dung, kiến thức, kết quả nghiên cứu mới, gắn với công cuộc đổi mới ở Việt Nam, nhất là những thành tựu trong 35 năm đổi mới đất nước.', 'tg008', 'dmsgk', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book010', 'Giáo Trình Lịch Sử Đảng Cộng Sản Việt Nam', 'Giáo trình cung cấp cho sinh viên những tri thức vừa cơ bản, vừa chuyên sâu, mang tính hệ thống về quá trình ra đời và lãnh đạo cách mạng của Đảng Cộng sản Việt Nam; cung cấp cơ sở lịch sử, củng cố niềm tin của thế hệ trẻ vào con đường cách mạng giải phóng dân tộc và phát triển đất nước; trang bị phương pháp nhận thức biện chứng, khách quan về quá trình Đảng ra đời và lãnh đạo cách mạng;… từ đó nâng cao hiểu biết lý luận, nắm bắt thực tiễn, vận dụng vào xem xét, đánh giá vai trò và sự lãnh đạo của Đảng, góp phần thiết thực vào công tác xây dựng Đảng ngày càng vững mạnh, đáp ứng yêu cầu nhiệm vụ mới trong quá trình công nghiệp hóa, hiện đại hóa đất nước, xây dựng thành công chủ nghĩa xã hội ở Việt Nam.', 'tg008', 'dmsgk', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book011', 'Giáo Trình Tư Tưởng Hồ Chí Minh', 'Giáo trình góp phần giúp người học hiểu sâu sắc, toàn diện và đầy đủ hơn về vai trò, vị trí, ý nghĩa của tư tưởng Hồ Chí Minh, các nội dung cơ bản trong tư tưởng Hồ Chí Minh, từ đó vận dụng, liên hệ với thực tiễn học tập, rèn luyện, xây dựng nhân cách để trở thành công dân tốt, đóng góp vào công cuộc xây dựng đất nước.', 'tg008', 'dmsgk', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book012', 'Truyện Cười Việt Nam Thời @', 'Từ những cốt truyện hiện đại , tác giả đã thổi hồn mình vào trong đó , nhào nặn câu chữ , chắt lọc tình tiết , văn phòng trôi chảy , cách kể hấp dẫn , các nhân vật cười trong truyện cứ hiện ra , nhớ mãi để có dịp hội ngộ , đông người được bung ra cùng nhau cười nghiêng ngả', 'tg009', 'dmtc|dmttranh', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book013', 'Vui Vẻ Không Quạu Nha 2', 'Đúng là tuổi trẻ chưa trải sự đời, lớn rồi mới biết hóa ra cuộc đời không chỉ có màu hồng, không phải cái gì mình thích, mình muốn là a lê hấp bạn sẽ có được ngay, mà trái lại - bạn phải vật lộn với đủ thứ, dù chỉ đơn giản là để sống một đời bình thường thôi.Nhưng mà hỏi thật lòng nhé, cuộc đời thực sự “khó ở” và “buồn” tới vậy à?', 'tg010', 'dmtc|dmttranh', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book014', 'Vui Vẻ Không Quạu Nha', 'Vui vẻ không quạu nha - xin được gửi đến những bạn trẻ đang dễ giận dữ, cau có ngoài thế giới kia, những ai đang buồn phiền về rắc rối nào đó, “trái tim” nhỏ bé này còn phải dành cho niềm vui, đừng để bực bội, dỗi hờn từ những điều không đáng chiếm hết chỗ.Và hãy hét to với Thế giới rằng “Mình là một người tràn đầy năng lượng,” luôn sẵn sàng để hạnh phúc hơn.', 'tg010', 'dmtc|dmttranh', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book015', 'Thỏ Bảy Màu Và Những Người Nghĩ Nó Là Bạn', 'Vui vẻ không quạu nha - xin được gửi đến những bạn trẻ đang dễ giận dữ, cau có ngoài thế giới kia, những ai đang buồn phiền về rắc rối nào đó, “trái tim” nhỏ bé này còn phải dành cho niềm vui, đừng để bực bội, dỗi hờn từ những điều không đáng chiếm hết chỗ.Và hãy hét to với Thế giới rằng “Mình là một người tràn đầy năng lượng,” luôn sẵn sàng để hạnh phúc hơn.', 'tg011', 'dmtc|dmttranh', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book016', 'Tớ Là Mèo Pusheen', 'Tớ là mèo Pusheen - Cuốn nhật ký xoay quanh cuộc sống của Pusheen - chú mèo Icon nổi tiếng trên mạng xã hội facebook với hơn 7 triệu người hâm mộ.Hãy cùng tìm hiểu những điều khiến Pusheen thích thú và lí do hàng triệu người trót mết vẻ \"tung tăng\" của nàng mèo múp míp, mũm mĩm này nhé!', 'tg012', 'dmtc|dmttranh', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book017', 'Tonton Friends - Hội Chân Ngắn Siêu Lầy', 'Cuốn nhật kí bằng tranh ghi lại những khoảnh khắc ấm áp trong cuộc sống của những người bạn Tonton đến từ Hàn Quốc. Những câu chuyện thường ngày vui vẻ, hài hước, từng kỉ niệm trong sáng, ngọt ngào của tình bạn và tình yêu.', 'tg013', 'dmtc|dmttranh', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book018', 'Tôi là Bêtô', 'Một tác phẩm của Nguyễn Nhật Ánh', 'tg016', 'dmtt|dmtruyen', '2023-12-09 20:41:28', '2023-12-14 14:47:57', NULL, 1),
('book019', 'Tôi tài giỏi, bạn cũng thế!', 'Tôi tài giỏi, bạn cũng thế! (nhan đề gốc tiếng Anh: I Am Gifted, So Are You!) là quyển sách bán chạy nhất của doanh nhân người Singapore Adam Khoo, viết về những phương pháp học tập tiên tiến. Quyển sách đã được dịch ra hàng chục thứ tiếng, trong đó Tôi tài giỏi, bạn cũng thế! là phiên bản tiếng Việt được dịch bởi hai dịch giả nổi tiếng Trần Đăng Khoa và Uông Xuân Vy của TGM Books. Tại Việt Nam, quyển sách đã trở thành một hiện tượng giáo dục trong những năm 2009-2011 và đạt được nhiều thành tựu trong lĩnh vực xuất bản, tạo ra kỷ lục mới cho ngành xuất bản Việt Nam với hơn 200.000 bản in được bán ra và hơn 400.000 e-book được phân phối.', 'tg015', 'dmkns', '2023-12-10 16:07:24', '2023-12-14 14:47:40', NULL, 1),
('book020', 'Bạch Dạ Hành', 'Osuke, chủ một tiệm cầm đồ bị sát hại tại một ngôi nhà chưa hoàn công, một triệu yên mang theo người cũng bị cướp mất.Sau đó một tháng, nghi can Fumiyo được cho rằng có quan hệ tình ái với nạn nhân và đã sát hại ông để cướp một triệu yên, cũng chết tại nhà riêng vì ngộ độc khí ga. Vụ án mạng ông chủ tiệm cầm đồ rơi vào bế tắc và bị bỏ xó.Nhưng với hai đứa trẻ mười một tuổi, con trai nạn nhân và con gái nghi can, vụ án mạng năm ấy chưa bao giờ kết thúc. Sinh tồn và trưởng thành dưới bóng đen cái chết của bố mẹ, cho đến cuối đời, Ryoji vẫn luôn khao khát được một lần đi dưới ánh mặt trời, còn Yukiho cứ ra sức vẫy vùng rồi mãi mãi chìm vào đêm trắng.', 'tg014', 'dmttham', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1),
('book021', 'Giấc Mơ Tiên Tri', 'Đêm khuya, một gã đàn ông lẻn vào phòng của thiếu nữ mười sáu tuổi. Người mẹ phát hiện và nổ súng. Khi bị bắt, gã đàn ông khai hắn đã mơ thấy mình trở thành chồng thiếu nữ này từ mười bảy năm về trước, bằng chứng là bài văn mô tả cô gái do hắn viết từ thời tiểu học. Lẽ nào người trong mơ lại xuất hiện ngoài đời thực? Đó chỉ là sự trùng hợp quá đỗi ngẫu nhiên, hay thực sự tồn tại giấc mơ tiên tri? Một lần nữa, nhà vật lý học thiên tài Yukawa buộc phải ra tay, phá giải hàng loạt vụ án nhuốm màu huyền bí...', 'tg014', 'dmttham', '2023-12-09 20:41:28', '2023-12-11 09:54:37', NULL, 1);

INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmkns', 'Kỹ năng sống', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1);
INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmkt', 'Sách kinh tế', '2023-12-14 13:23:31', '2023-12-14 13:23:31', NULL, 1);
INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmln', 'Light Novel', '2023-12-14 13:21:05', '2023-12-14 13:21:05', NULL, 1);
INSERT INTO `Category` (`id`, `name`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('dmna', 'Sách dạy nấu ăn', '2023-12-14 13:21:22', '2023-12-14 13:21:22', NULL, 1),
('dmnt', 'Ngôn tình', '2023-12-14 13:23:16', '2023-12-14 13:23:16', NULL, 1),
('dmsgk', 'Sách giáo khoa', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('dmtc', 'Truyện cười ', '2023-12-14 13:24:25', '2023-12-14 13:24:25', NULL, 1),
('dmtct', 'Truyện cổ tích', '2023-12-02 01:52:21', '2023-12-14 13:22:37', NULL, 1),
('dmtt', 'Tiểu thuyết', '2023-12-02 01:52:21', '2023-12-02 01:52:21', NULL, 1),
('dmttham', 'Trinh thám', '2023-12-14 13:21:39', '2023-12-14 13:21:39', NULL, 1),
('dmttranh', 'Truyện tranh', '2023-12-14 13:20:21', '2023-12-14 13:20:21', NULL, 1);


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;