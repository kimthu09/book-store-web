/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


INSERT INTO `Role` (`id`, `name`) VALUES
('admin', 'user');
INSERT INTO `Role` (`id`, `name`) VALUES
('fl_9lf4Ig', 'haha');
INSERT INTO `Role` (`id`, `name`) VALUES
('user', 'user');

INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAN_CREATE');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAN_VIEW');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_CREATE');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_UP_INFO'),
('admin', 'CAT_VIEW'),
('admin', 'CUS_CREATE'),
('admin', 'CUS_PAY'),
('admin', 'CUS_UP_INFO'),
('admin', 'CUS_VIEW'),
('admin', 'EXP_CREATE'),
('admin', 'EXP_VIEW'),
('admin', 'FOD_CREATE'),
('admin', 'FOD_UP_INFO'),
('admin', 'FOD_UP_STATE'),
('admin', 'FOD_VIEW'),
('admin', 'IMP_CREATE'),
('admin', 'IMP_UP_STATE'),
('admin', 'IMP_VIEW'),
('admin', 'ING_CREATE'),
('admin', 'ING_VIEW'),
('admin', 'INV_CREATE'),
('admin', 'INV_VIEW'),
('admin', 'SUP_CREATE'),
('admin', 'SUP_PAY'),
('admin', 'SUP_UP_INFO'),
('admin', 'SUP_VIEW'),
('admin', 'TOP_CREATE'),
('admin', 'TOP_UP_INFO'),
('admin', 'TOP_UP_STATE'),
('admin', 'TOP_VIEW'),
('admin', 'USE_UP_INFO'),
('admin', 'USE_UP_STATE'),
('admin', 'USE_VIEW'),
('fl_9lf4Ig', 'IMP_CREATE'),
('user', 'CAN_CREATE'),
('user', 'CAT_CREATE'),
('user', 'CAT_UP_INFO'),
('user', 'CUS_CREATE'),
('user', 'CUS_PAY'),
('user', 'CUS_UP_INFO'),
('user', 'EXP_CREATE'),
('user', 'FOD_CREATE'),
('user', 'FOD_UP_INFO'),
('user', 'FOD_UP_STATE'),
('user', 'IMP_CREATE'),
('user', 'IMP_UP_STATE'),
('user', 'ING_CREATE'),
('user', 'INV_CREATE'),
('user', 'SUP_CREATE'),
('user', 'SUP_PAY'),
('user', 'SUP_UP_INFO'),
('user', 'TOP_CREATE'),
('user', 'TOP_UP_INFO'),
('user', 'TOP_UP_STATE'),
('user', 'USE_UP_INFO'),
('user', 'USE_UP_STATE');


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;