# ************************************************************
# Sequel Ace SQL dump
# Version 20046
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 127.0.0.1 (MySQL 8.0.32)
# Database: fp_crowdfundchat_db
# Generation Time: 2023-06-27 18:00:43 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table chats
# ------------------------------------------------------------

DROP TABLE IF EXISTS `chats`;

CREATE TABLE `chats` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `project_id` int DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `chats` WRITE;
/*!40000 ALTER TABLE `chats` DISABLE KEYS */;

INSERT INTO `chats` (`id`, `user_id`, `project_id`, `content`, `created_at`, `user_name`)
VALUES
	(1,0,0,'8,\"','2023-04-29 06:14:38',NULL),
	(2,0,0,'l6J5','2023-04-29 06:33:35',NULL),
	(3,0,0,'l7x73iLm53WH5Iva8l26lzk6ty37lBR2U5jJAEdW1MF/zTYRlKsxomBp','2023-04-29 06:35:27',NULL),
	(4,28,29,'Ax/ScTykrBrYhdNPApju9Akasdfx39vu79e2EArj6+jE0G7N6KPwK0HA','2023-04-29 06:43:16',NULL),
	(5,28,29,'mAHVoEirRc0E87mhoS9jNxbQxTYMJG3ZaeM34Su2XHA2AveiNCdIO8QB','2023-04-29 14:35:48',NULL),
	(6,28,29,'gbZgJ4f9cR+GWKF9ihbHm9B5fuupY4q2QjLgCHnxcxn3s+ed4YbPriG4','2023-04-29 14:37:34',NULL),
	(7,29,28,'gbZgJ4f9cR+GWKF9ihfIgPNQvy4Vncc1UcbygsBXuuHWkZaRlQ==','2023-04-29 14:38:02',NULL),
	(8,29,28,'0ZxqJfiW1BM5+JGF/OEWYs7PAMUdFC1IqUkDT4ZPlevYdWvl','2023-04-29 15:33:23',NULL),
	(9,0,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 21:42:47',NULL),
	(10,0,2,'2Ilnafw=','2023-05-07 21:45:55',NULL),
	(11,0,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 21:46:54',NULL),
	(12,0,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 21:48:11',NULL),
	(13,0,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 21:48:39',NULL),
	(14,29,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 21:50:10',NULL),
	(15,29,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 22:35:33',NULL),
	(16,28,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 22:44:42',NULL),
	(17,28,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 22:44:42',NULL),
	(18,28,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 22:54:45',NULL),
	(19,29,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 22:54:53',NULL),
	(20,29,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 22:54:53',NULL),
	(21,28,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 22:55:42',NULL),
	(22,29,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 22:55:45',NULL),
	(23,29,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 22:55:45',NULL),
	(24,29,2,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 23:00:04',NULL),
	(25,0,2,'xZ9ud7Of1wFr+o2CqKgMY54ygFagVQ==','2023-05-07 23:02:50',NULL),
	(26,29,8,'0cxlYOTXwwEu5MiOs+EWbs8=','2023-05-07 23:03:07',NULL),
	(27,28,8,'2IUqJfKH11Ig94qFrqgTaovIrg==','2023-05-07 23:03:28',NULL),
	(28,28,8,'2IUqJfKH11Ig94qFrqgTaovIrg==','2023-05-07 23:03:28',NULL),
	(29,0,8,'xZ9ud7Of1wFr+o2CqKgMY54ygFagVQ==','2023-05-07 23:03:43',NULL),
	(30,29,8,'yYMrcvKT0gc7','2023-05-07 23:39:39',NULL),
	(31,29,8,'yYMrcvKT0gc7','2023-05-07 23:39:39',NULL),
	(32,29,8,'yYlyYA==','2023-05-24 18:38:25',NULL),
	(33,29,8,'yYlyYA==','2023-05-24 18:38:25',NULL),
	(34,28,8,'x41nZA==','2023-05-24 18:38:57',NULL),
	(35,28,8,'x41nZA==','2023-05-24 18:38:57',NULL),
	(36,28,8,'x41nZA==','2023-05-25 11:12:41',NULL),
	(37,29,8,'2I1jZPuW3hM=','2023-05-25 11:13:14',NULL),
	(38,29,8,'2I1jZPuW3hM=','2023-05-25 11:13:14',NULL),
	(39,29,8,'+IN8JfKF01Iy+Z0=','2023-05-25 13:33:46',NULL),
	(40,28,8,'+cxqaLOR3xwu','2023-05-25 13:34:34',NULL),
	(41,28,8,'+cxqaLOR3xwu','2023-05-25 13:34:34',NULL),
	(42,41,8,'2Ilnafw=','2023-05-25 13:56:56',NULL),
	(43,41,8,'2Ilnafw=','2023-05-25 14:25:06',NULL),
	(44,43,8,'y85qJ6nV3hcn+ofGoQ==','2023-06-06 04:53:14',NULL),
	(45,43,8,'y85qJ6nV3hcn+ofGoQ==','2023-06-06 04:53:16',NULL),
	(46,43,8,'y85qJ6nVwRMv8p2U/vU=','2023-06-06 05:08:35',NULL),
	(47,43,8,'y85qJ6nVwRMv8p2U/vU=','2023-06-06 05:08:38',NULL),
	(48,43,8,'y85qJ6nVwRMv8p2U/vU=','2023-06-06 05:08:38',NULL),
	(49,43,8,'y85qJ6nVwRMv8p2U/vU=','2023-06-06 05:08:40',NULL),
	(50,43,8,'y85qJ6nVwRMv8p2U/vU=','2023-06-06 05:08:40',NULL),
	(51,43,8,'y85qJ6nVwRMv8p2U/vU=','2023-06-06 05:08:50',NULL),
	(52,43,8,'y85qJ6nVwRMv8p2U/vU=','2023-06-06 05:09:46',NULL),
	(53,43,8,'y85qJ6nVzxcy85GBpaoF','2023-06-06 05:10:25',NULL),
	(54,43,8,'y85qJ6nV0gcv48qZ','2023-06-06 05:10:57',NULL),
	(55,43,8,'y85qJ6nV0gcv48qZ','2023-06-06 05:11:09',NULL),
	(56,43,8,'x41nZA==','2023-06-07 09:06:00','Julia Keeva'),
	(57,43,8,'y85qJ6nV3hcn+ofGoQ==','2023-06-07 13:53:33','Julia Keeva'),
	(58,43,8,'y85qJ6nV3hcn+ofGoQ==','2023-06-07 14:35:57','Julia Keeva'),
	(59,43,8,'y85qJ6nV3hcn+ofGoQ==','2023-06-07 14:36:17','Julia Keeva'),
	(60,43,8,'y85qJ6nV3hcn+ofGoQ==','2023-06-07 14:37:47','Julia Keeva'),
	(61,43,8,'y85qJ6nV3hcn+ofGoQ==','2023-06-07 14:37:48','Julia Keeva'),
	(62,43,8,'y85qJ6nV2BM4/8iDs/odZYZwaQ==','2023-06-07 14:38:08','Julia Keeva'),
	(63,43,8,'y85qJ6nV2BM4/8iDs/odZYZwaQ==','2023-06-07 14:39:08','Julia Keeva'),
	(64,43,8,'y85qJ6nV2BM4/8iDs/odZYZwaQ==','2023-06-07 14:39:10','Julia Keeva'),
	(65,43,8,'y85qJ6nV3xkq+MiQs+YfYKmmhf4=','2023-06-07 14:39:32','Julia Keeva'),
	(66,43,8,'y85qJ6nV2xM494PGoQ==','2023-06-07 14:39:39','Julia Keeva'),
	(67,43,8,'y85qJ6nVxhMn+cqZ','2023-06-07 14:45:36','Julia Keeva'),
	(68,43,8,'2Ilnafw=','2023-06-07 14:48:36','Julia Keeva'),
	(69,43,8,'y85qJ6nV3hcn+ofF/vU=','2023-06-07 14:49:14','Julia Keeva'),
	(70,43,8,'y85qJ6nV3hcn+ofF/vU=','2023-06-07 14:49:17','Julia Keeva'),
	(71,43,8,'y85qJ6nV2BM4/8iDs/odZYZwaQ==','2023-06-07 14:49:25','Julia Keeva'),
	(72,43,8,'y85qJ6nV1wsq+8jGoQ==','2023-06-07 14:49:39','Julia Keeva'),
	(73,43,8,'y85qJ6nV1wsq+8jGoQ==','2023-06-07 14:49:45','Julia Keeva'),
	(74,43,8,'y85qJ6nV1wsq+8jGoQ==','2023-06-07 14:49:51','Julia Keeva'),
	(75,43,8,'y85qJ6nV1wsq+8jGoQ==','2023-06-07 14:49:51','Julia Keeva'),
	(76,43,8,'y85qJ6nV1wsq+8jGoQ==','2023-06-07 14:51:40','Julia Keeva'),
	(77,43,8,'y85qJ6nV1wsq+8jGoQ==','2023-06-07 14:51:40','Julia Keeva'),
	(78,43,8,'y85qJ6nV2h0j88qZ','2023-06-07 14:52:36','Julia Keeva'),
	(79,43,8,'y85qJ6nV2h0qtJU=','2023-06-07 14:52:58','Julia Keeva'),
	(80,43,8,'y85qJ6nV2h0utJU=','2023-06-07 14:53:07','Julia Keeva'),
	(81,43,8,'y85qJ6nV3Qc584rGoQ==','2023-06-07 14:59:18','Julia Keeva'),
	(82,43,8,'y85qJ6nV0gUv4YzGoQ==','2023-06-07 15:06:29','Julia Keeva'),
	(83,43,8,'y85qJ6nV0gUv4dvGoQ==','2023-06-07 15:06:35','Julia Keeva'),
	(84,43,8,'y85qJ6nV2h0j88qZ','2023-06-07 15:33:43','Julia Keeva'),
	(85,43,8,'y85qJ6nV2xMg94bGoQ==','2023-06-07 15:33:51','Julia Keeva'),
	(86,43,8,'y85qJ6nV2BM4/8qZ','2023-06-07 15:34:05','Julia Keeva'),
	(87,43,8,'y85qJ6nV2BM4/8qZ','2023-06-07 15:34:32','Julia Keeva'),
	(88,43,8,'y85qJ6nV2BM4/8qZ','2023-06-07 15:36:41','Julia Keeva'),
	(89,43,8,'y85qJ6nV2BM4/8qZ','2023-06-07 15:36:43','Julia Keeva'),
	(90,43,8,'y85qJ6nV2BM4/8qZ','2023-06-07 15:36:50','Julia Keeva'),
	(91,43,8,'y85qJ6nV2R8qtoeJu6oF','2023-06-07 15:44:54','Julia Keeva'),
	(92,43,8,'y85qJ6nV2R8qtoeJuaoF','2023-06-07 15:44:59','Julia Keeva'),
	(93,43,8,'y85qJ6nV2h0j88qZ','2023-06-07 15:45:06','Julia Keeva'),
	(94,43,8,'y85qJ6nV1xMqtJU=','2023-06-07 15:48:03','Julia Keeva'),
	(95,43,8,'y85qJ6nVwhNp6w==','2023-06-07 15:48:26','Julia Keeva'),
	(96,43,8,'y85qJ6nV11A2','2023-06-07 15:51:26','Julia Keeva'),
	(97,43,8,'y85qJ6nV11A2','2023-06-07 15:51:43','Julia Keeva'),
	(98,43,8,'y85qJ6nV1xMq98qZ','2023-06-07 15:51:47','Julia Keeva'),
	(99,43,8,'y85qJ6nV1xMq98qZ','2023-06-07 15:51:59','Julia Keeva'),
	(100,43,8,'y85qJ6nVxEE5tJU=','2023-06-07 15:52:18','Julia Keeva'),
	(101,43,8,'y85qJ6nVxEE5tJU=','2023-06-07 15:52:22','Julia Keeva'),
	(102,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:02:49','Julia Keeva'),
	(103,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:02:51','Julia Keeva'),
	(104,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:03:00','Julia Keeva'),
	(105,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:03:06','Julia Keeva'),
	(106,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:03:06','Julia Keeva'),
	(107,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:03:07','Julia Keeva'),
	(108,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:03:07','Julia Keeva'),
	(109,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:03:07','Julia Keeva'),
	(110,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:03:07','Julia Keeva'),
	(111,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:03:07','Julia Keeva'),
	(112,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:03:07','Julia Keeva'),
	(113,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:03:09','Julia Keeva'),
	(114,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:04:54','Julia Keeva'),
	(115,43,8,'y85qJ6nVwRQt4Y7GoQ==','2023-06-08 04:04:55','Julia Keeva'),
	(116,43,8,'y85qJ6nVwRQt4Y6Tuu5adg==','2023-06-08 04:04:59','Julia Keeva'),
	(117,43,8,'2IlnafyRwRQ=','2023-06-08 04:05:05','Julia Keeva'),
	(118,44,8,'wp4=','2023-06-08 04:05:40','Julia Keeva'),
	(119,43,8,'2IlnafyRwRQ=','2023-06-08 04:05:49','Julia Keeva'),
	(120,43,8,'2IlnafyRwRQ=','2023-06-08 04:05:49','Julia Keeva'),
	(121,43,8,'y85qJ6nV1BAp9IrGoQ==','2023-06-08 04:07:04','Julia Keeva'),
	(122,43,8,'y85qJ6nV1BAp9IrGoQ==','2023-06-08 04:07:04','Julia Keeva'),
	(123,43,8,'kopuY/XV','2023-06-08 04:07:49','Julia Keeva'),
	(124,43,8,'kopuY/XV','2023-06-08 04:07:49','Julia Keeva'),
	(125,43,8,'ko5tb/WVlA==','2023-06-08 04:09:22','Julia Keeva'),
	(126,43,8,'ko5qbuCYlA==','2023-06-08 04:10:06','Julia Keeva'),
	(127,43,8,'koptY7E=','2023-06-08 04:10:30','Julia Keeva'),
	(128,43,8,'kotuYvWR0FA=','2023-06-08 04:10:36','Julia Keeva'),
	(129,43,8,'kotuYvWR0FA=','2023-06-08 04:10:53','Julia Keeva'),
	(130,43,8,'koRhd/udxFA=','2023-06-08 04:11:08','Julia Keeva'),
	(131,43,8,'3o14bLOQ2QAu+I8=','2023-06-08 04:24:23','Julia Keeva'),
	(132,43,8,'3o14bLOQ2QAu+I8=','2023-06-08 04:24:49','Julia Keeva'),
	(133,44,8,'3o14bLOQ2QAu+I8=','2023-06-08 04:24:58','Julia Keeva'),
	(134,43,8,'2y6ucuSF','2023-06-08 04:25:07','Julia Keeva'),
	(135,43,8,'2y6ucuSF','2023-06-08 04:25:13','Julia Keeva'),
	(136,43,8,'kptqafKSlA==','2023-06-08 04:47:48','Julia Keeva'),
	(137,43,8,'kptqafKSlA==','2023-06-08 04:47:52','Julia Keeva'),
	(138,43,8,'kptqafKSlA==','2023-06-08 04:47:52','Julia Keeva'),
	(139,43,8,'kptqafKSlA==','2023-06-08 04:47:52','Julia Keeva'),
	(140,43,8,'koBqafLV','2023-06-08 05:00:42','Julia Keeva'),
	(141,43,8,'kop8Y+SRlA==','2023-06-08 05:01:47','Julia Keeva'),
	(142,43,8,'kop8Y+SRlA==','2023-06-08 05:02:14','Julia Keeva'),
	(143,43,8,'ko1gZOOWlA==','2023-06-08 05:02:35','Julia Keeva'),
	(144,43,8,'ko1gZOOWlA==','2023-06-08 05:03:27','Julia Keeva'),
	(145,43,8,'ko1gZOOWlA==','2023-06-08 05:03:36','Julia Keeva'),
	(146,43,8,'ko1gZOOWlA==','2023-06-08 05:03:40','Julia Keeva'),
	(147,43,8,'ko1gZOOWlA==','2023-06-08 05:03:40','Julia Keeva'),
	(148,43,8,'kop8Y+SRlA==','2023-06-08 05:30:47','Julia Keeva'),
	(149,43,8,'1Jtv','2023-06-08 05:31:56','Julia Keeva'),
	(150,43,8,'1Jtv','2023-06-08 05:32:00','Julia Keeva'),
	(151,43,8,'1Jtv','2023-06-08 05:32:40','Julia Keeva'),
	(152,43,8,'0594ZuCR','2023-06-08 05:33:03','Julia Keeva'),
	(153,43,8,'0594ZuCR','2023-06-08 05:34:24','Julia Keeva'),
	(154,43,8,'x4ptY/WR','2023-06-08 05:34:30','Julia Keeva'),
	(155,43,8,'1J1vdPc=','2023-06-08 05:35:14','Julia Keeva'),
	(156,43,8,'1J1vdPc=','2023-06-08 05:35:20','Julia Keeva'),
	(157,43,8,'1J1vdPc=','2023-06-08 05:35:20','Julia Keeva'),
	(158,43,8,'wA==','2023-06-08 05:35:46','Julia Keeva'),
	(159,43,8,'05lgZOA=','2023-06-09 00:45:09','Julia Keeva'),
	(160,43,8,'05lgZOCR0A==','2023-06-09 00:45:17','Julia Keeva'),
	(161,43,8,'05lgZOCRxA==','2023-06-09 00:45:21','Julia Keeva'),
	(162,43,8,'05lgZOCRxAU84Q==','2023-06-09 00:45:26','Julia Keeva'),
	(163,43,8,'25lvZLObwx87/4aD','2023-06-09 00:46:27','Julia Keeva'),
	(164,43,8,'3o14bLOQ2QAu+I8=','2023-06-09 00:46:52','Julia Keeva'),
	(165,43,8,'3o14bLOQ2QAu+I/EsOkUag==','2023-06-09 00:47:04','Julia Keeva'),
	(166,43,8,'04VqZPI=','2023-06-09 01:32:49','Julia Keeva'),
	(167,43,8,'04VqZPKR0xQu8I2C','2023-06-09 01:34:05','Julia Keeva'),
	(168,43,8,'1optY/U=','2023-06-09 01:36:02','Julia Keeva'),
	(169,43,8,'64Npb/aUwlIE9IKBv/wl','2023-06-09 01:38:58','Julia Keeva'),
	(170,43,8,'64Npb/aUwlIE9IKBv/wl','2023-06-09 01:40:04','Julia Keeva'),
	(171,43,8,'64Npb/aUwlIE9IKBv/wl','2023-06-09 01:40:06','Julia Keeva'),
	(172,43,8,'64Npb/aUwlIE9IKBv/wl','2023-06-09 01:43:30','Julia Keeva'),
	(173,43,8,'xIl5dfKCwg==','2023-06-09 01:44:41','Julia Keeva'),
	(174,43,8,'xZ9iZA==','2023-06-09 01:45:03','Julia Keeva'),
	(175,43,8,'xZ9iZA==','2023-06-09 01:45:47','Julia Keeva'),
	(176,43,8,'3o14bA==','2023-06-09 01:45:52','Julia Keeva'),
	(177,43,8,'x41nZA==','2023-06-09 01:47:35','Julia Keeva'),
	(178,43,8,'x41nZA==','2023-06-09 01:47:39','Julia Keeva'),
	(179,43,8,'x41nZA==','2023-06-09 01:47:39','Julia Keeva'),
	(180,43,8,'x41nZA==','2023-06-09 01:49:21','Julia Keeva'),
	(181,43,8,'yYl+','2023-06-09 01:53:41','Julia Keeva'),
	(182,43,8,'yYl+','2023-06-09 01:53:41','Julia Keeva'),
	(183,43,8,'0o1gdvw=','2023-06-09 01:54:02','Julia Keeva'),
	(184,43,8,'25lmZPuW','2023-06-09 01:54:32','Julia Keeva'),
	(185,43,8,'25lmZPuW','2023-06-09 01:54:39','Julia Keeva'),
	(186,43,8,'3INjYA==','2023-06-09 01:55:20','Julia Keeva'),
	(187,43,8,'1Jtvcvc=','2023-06-09 01:55:32','Julia Keeva'),
	(188,43,8,'24Ng','2023-06-09 01:55:47','Julia Keeva'),
	(189,43,8,'24Ng','2023-06-09 01:55:48','Julia Keeva'),
	(190,43,8,'1JlgZA==','2023-06-09 01:56:51','Julia Keeva'),
	(191,43,8,'3A==','2023-06-09 02:00:53','Julia Keeva'),
	(192,43,8,'3Idg','2023-06-09 02:00:58','Julia Keeva'),
	(193,43,8,'3Idg','2023-06-09 02:01:01','Julia Keeva'),
	(194,43,8,'3Idgb/mY3B0h+YKL','2023-06-09 02:01:06','Julia Keeva'),
	(195,43,8,'krApd+GFxC5ptA==','2023-06-09 02:03:17','Julia Keeva'),
	(196,43,8,'krApYfqa1xwqtoOFtNRaKQ==','2023-06-09 02:03:39','Julia Keeva'),
	(197,43,8,'krApb/KV1y5ptA==','2023-06-09 02:03:46','Julia Keeva'),
	(198,43,8,'3YFm','2023-06-09 02:04:05','Julia Keeva'),
	(199,43,8,'3IBnaf+b2g==','2023-06-09 02:04:29','Julia Keeva'),
	(200,43,8,'1oZhYw==','2023-06-09 02:05:24','Julia Keeva'),
	(201,43,8,'1optb/mR','2023-06-09 02:05:29','Julia Keeva'),
	(202,43,8,'15hs','2023-06-09 02:06:20','Julia Keeva'),
	(203,43,8,'koZib7E=','2023-06-09 02:07:09','Julia Keeva'),
	(204,43,8,'koBnJw==','2023-06-09 02:07:25','Julia Keeva'),
	(205,43,8,'kopuY/aR0xRp','2023-06-09 02:08:28','Julia Keeva'),
	(206,43,8,'kopuY/aR0xRp','2023-06-09 02:08:33','Julia Keeva'),
	(207,43,8,'3IA=','2023-06-09 02:10:32','Julia Keeva'),
	(208,43,8,'xJh/cec=','2023-06-09 02:10:58','Julia Keeva'),
	(209,43,8,'1oltYPU=','2023-06-09 02:11:12','Julia Keeva'),
	(210,43,8,'3IVkaQ==','2023-06-09 02:11:32','Julia Keeva'),
	(211,43,8,'g544d6CF','2023-06-09 02:12:11','Julia Keeva'),
	(212,43,8,'3IBnaviY3R0=','2023-06-09 02:14:17','Julia Keeva'),
	(213,43,8,'1pljYOaR3hct','2023-06-09 02:14:43','Julia Keeva'),
	(214,43,8,'2y6ucuSF','2023-06-09 02:15:07','Julia Keeva'),
	(215,43,8,'2y6ucuSF','2023-06-09 02:15:09','Julia Keeva'),
	(216,43,8,'2y6ucuSF','2023-06-09 02:15:10','Julia Keeva'),
	(217,43,8,'2y6ucuSF','2023-06-09 02:15:14','Julia Keeva'),
	(218,43,8,'2IZjbw==','2023-06-09 02:16:34','Julia Keeva'),
	(219,43,8,'2IZjb/ue3hs=','2023-06-09 02:16:44','Julia Keeva'),
	(220,43,8,'2IZjb/ue3hst8I4=','2023-06-09 02:17:21','Julia Keeva'),
	(221,43,8,'2olhYPqd0xsu/IGBtuEd','2023-06-09 02:17:48','Julia Keeva'),
	(222,43,8,'woV4YA==','2023-06-09 02:17:54','Julia Keeva'),
	(223,43,8,'woV4YA==','2023-06-09 02:18:19','Julia Keeva'),
	(224,43,8,'woV4YA==','2023-06-09 02:18:31','Julia Keeva'),
	(225,43,8,'woV4YA==','2023-06-09 02:18:35','Julia Keeva'),
	(226,43,8,'woV4YLOCxg==','2023-06-09 02:18:51','Julia Keeva'),
	(227,43,8,'woV4YLOCxg==','2023-06-09 02:19:02','Julia Keeva'),
	(228,43,8,'0ol5bOCSlgc7','2023-06-09 02:19:11','Julia Keeva'),
	(229,43,8,'2YRibeae3g==','2023-06-09 02:20:01','Julia Keeva'),
	(230,43,8,'1oZhbOSd3xQ=','2023-06-09 02:20:25','Julia Keeva'),
	(231,43,8,'3IdnYA==','2023-06-09 02:20:38','Julia Keeva'),
	(232,43,8,'0oRhafo=','2023-06-09 02:21:08','Julia Keeva'),
	(233,43,8,'24dgbg==','2023-06-09 02:21:22','Julia Keeva'),
	(234,43,8,'24dgbg==','2023-06-09 02:21:30','Julia Keeva'),
	(235,43,8,'koJqdvrX0R0584aD/g==','2023-06-11 19:54:12','Julia Keeva'),
	(236,43,8,'koJqdvrX0R0584aD/g==','2023-06-11 19:54:31','Julia Keeva'),
	(237,43,8,'kod+cebV','2023-06-11 19:55:16','Julia Keeva'),
	(238,43,8,'koBkbfbV','2023-06-11 19:56:26','Julia Keeva'),
	(239,43,8,'2I1nag==','2023-06-11 20:05:59','Julia Keeva'),
	(240,43,8,'x414ZOM=','2023-06-11 20:07:40','Julia Keeva'),
	(241,43,8,'x414ZOM=','2023-06-11 20:07:43','Julia Keeva'),
	(242,43,8,'3Y1lZvqZ0VIm94aNvQ==','2023-06-11 20:08:35','Julia Keeva'),
	(243,43,8,'0Z9ibg==','2023-06-11 20:09:56','Julia Keeva'),
	(244,43,8,'w4V/JfeYwRw=','2023-06-11 20:12:02','Julia Keeva'),
	(245,43,8,'1INlcbOA2QA578g=','2023-06-11 20:12:20','Julia Keeva'),
	(246,43,8,'yYlnafyA','2023-06-11 20:15:34','Julia Keeva'),
	(247,43,8,'0olqcA==','2023-06-11 20:17:49','Julia Keeva'),
	(248,43,8,'0ptq','2023-06-11 20:21:16','Julia Keeva'),
	(249,43,8,'xIRiYA==','2023-06-11 20:23:33','Julia Keeva'),
	(250,43,8,'3oN/bfqZ0VIu+puB','2023-06-11 20:23:39','Julia Keeva'),
	(251,43,8,'0o1gdvw=','2023-06-11 20:27:06','Julia Keeva'),
	(252,43,8,'3oN/bfqZ0VIu+puB','2023-06-11 20:27:14','Julia Keeva'),
	(253,43,8,'24l5YP3X2h0=','2023-06-11 20:27:49','Julia Keeva'),
	(254,43,8,'24l5YP3X2h0=','2023-06-11 20:27:54','Julia Keeva'),
	(255,43,8,'24l5YP3X2h0=','2023-06-11 20:28:00','Julia Keeva'),
	(256,43,8,'24l5YP3X2h0=','2023-06-11 20:28:12','Julia Keeva'),
	(257,43,8,'3I1j','2023-06-11 20:28:26','Julia Keeva'),
	(258,43,8,'24l5YP3X2h0v4YKPq+Id','2023-06-11 20:28:37','Julia Keeva'),
	(259,43,8,'24l5YP3X2h0v4YKPq+Id','2023-06-11 20:28:38','Julia Keeva'),
	(260,43,8,'w4lnZP6WwlI794+N','2023-06-14 03:26:47','Julia Keeva'),
	(261,43,8,'3o14bA==','2023-06-14 03:29:28','Julia Keeva'),
	(262,43,8,'0ZVqaA==','2023-06-14 03:31:20','Julia Keeva'),
	(263,43,8,'155sYg==','2023-06-14 03:31:48','Julia Keeva'),
	(264,43,8,'1JtvYeQ=','2023-06-14 03:33:08','Julia Keeva'),
	(265,43,8,'1pttcvWA0A==','2023-06-14 03:34:38','Julia Keeva'),
	(266,43,8,'3o14bLOb1x4q','2023-06-14 03:35:22','Julia Keeva'),
	(267,43,8,'3INjYP+Y3hc=','2023-06-14 03:35:31','Julia Keeva'),
	(268,43,8,'3o14bLOb1x4q','2023-06-14 03:35:42','Julia Keeva'),
	(269,43,8,'25l/cA==','2023-06-14 03:36:27','Julia Keeva'),
	(270,43,8,'1YpuY/aR','2023-06-14 03:37:01','Julia Keeva'),
	(271,43,8,'2oRtb/af0Bct','2023-06-14 03:37:56','Julia Keeva'),
	(272,43,8,'x4hvcveA0g==','2023-06-15 15:18:35','Julia Keeva'),
	(273,43,8,'x4hvcveA0g==','2023-06-15 15:18:47','Julia Keeva'),
	(274,43,8,'x4hvcveA0g==','2023-06-15 15:18:58','Julia Keeva'),
	(275,42,8,'2y6ucuSF','2023-06-15 15:19:04','pepe'),
	(276,45,8,'2I1i','2023-06-15 16:17:11','lala'),
	(277,45,8,'14lnafw=','2023-06-15 16:22:07','lala'),
	(278,45,8,'1oltYPWS0A==','2023-06-15 16:22:51','lala'),
	(279,45,8,'0Zxqbg==','2023-06-15 16:24:07','lala'),
	(280,45,8,'0pllYvI=','2023-06-15 16:25:11','lala'),
	(281,45,8,'x4p8Y+SR','2023-06-15 16:25:31','lala'),
	(282,45,8,'1oltYPU=','2023-06-15 16:27:33','lala'),
	(283,45,8,'3I1nZA==','2023-06-15 16:31:05','lala'),
	(284,45,8,'2Ilnafw=','2023-06-15 16:31:11','lala'),
	(285,42,8,'w4RkcvaF','2023-06-15 16:31:25','pepe'),
	(286,45,8,'x4lid/c=','2023-06-15 16:31:36','lala'),
	(287,45,8,'x4lid/c=','2023-06-15 16:31:41','lala'),
	(288,45,8,'1Z5uY+GF','2023-06-15 16:40:27','lala'),
	(289,45,8,'yYN5','2023-06-15 16:42:07','lala'),
	(290,45,8,'3o14bA==','2023-06-15 18:06:58','lala'),
	(291,45,8,'1oltYPU=','2023-06-15 18:07:36','lala'),
	(292,45,8,'2Ydqaw==','2023-06-15 18:08:00','lala'),
	(293,45,8,'xIl5cOA=','2023-06-15 18:08:25','lala'),
	(294,45,8,'3Y1gZP0=','2023-06-15 18:09:27','lala'),
	(295,45,8,'3oNlZA==','2023-06-15 18:10:07','lala'),
	(296,45,8,'3oNlZA==','2023-06-15 18:10:12','lala'),
	(297,45,8,'w4VqdfI=','2023-06-15 18:10:36','lala'),
	(298,45,8,'w4VqdfI=','2023-06-15 18:10:47','lala'),
	(299,45,8,'3INjYA==','2023-06-16 00:44:57','lala'),
	(300,45,8,'2I1vbOE=','2023-06-16 00:45:38','lala'),
	(301,45,8,'2I1vbOE=','2023-06-16 00:45:38','lala'),
	(302,42,8,'w4RkcvaF','2023-06-16 00:45:46','pepe'),
	(303,42,8,'w4RkcvaF','2023-06-16 00:45:46','pepe'),
	(304,42,8,'w4l4ZOc=','2023-06-16 00:45:54','pepe'),
	(305,42,8,'w4l4ZOc=','2023-06-16 00:45:54','pepe'),
	(306,45,8,'0Yd+','2023-06-16 00:46:37','lala'),
	(307,45,8,'0Yd+','2023-06-16 00:46:38','lala'),
	(308,45,8,'0Yd+','2023-06-16 00:46:43','lala'),
	(309,45,8,'0Yd+','2023-06-16 00:46:43','lala'),
	(310,42,8,'w4l4ZOc=','2023-06-16 00:47:47','pepe'),
	(311,45,8,'xIl5b/KQ1w==','2023-06-16 00:47:52','lala'),
	(312,45,8,'2Ilnbvw=','2023-06-16 01:00:33','lala'),
	(313,45,8,'2Ilnbvw=','2023-06-16 01:00:34','lala'),
	(314,42,8,'w4l4ZOc=','2023-06-16 01:00:43','pepe'),
	(315,45,8,'2Ilnbvw=','2023-06-16 01:00:51','lala'),
	(316,45,8,'1YpuY/aR','2023-06-16 01:02:27','lala'),
	(317,45,8,'1YpuY/aR','2023-06-16 01:02:27','lala'),
	(318,45,8,'1YpuY/aR','2023-06-16 01:02:32','lala'),
	(319,45,8,'1YpuY/aR','2023-06-16 01:02:32','lala'),
	(320,45,8,'3oMrdvyZ0VM=','2023-06-16 01:06:30','lala'),
	(321,45,8,'3oMrdvyZ0VM=','2023-06-16 01:06:38','lala'),
	(322,45,8,'x41nZPuQ','2023-06-16 01:09:46','lala'),
	(323,42,8,'w4l4ZOc=','2023-06-16 01:10:03','pepe'),
	(324,42,8,'w4l4ZOc=','2023-06-16 01:10:03','pepe'),
	(325,42,8,'w4l4ZOc=','2023-06-16 01:10:11','pepe'),
	(326,45,8,'x41nZPuQ','2023-06-16 01:10:19','lala'),
	(327,45,8,'x41nZPuQ0gUv4Yw=','2023-06-16 01:10:30','lala'),
	(328,45,8,'0Yd+','2023-06-16 01:11:07','lala'),
	(329,42,8,'w4l4ZOc=','2023-06-16 01:11:17','pepe'),
	(330,45,8,'0Yd+','2023-06-16 01:11:21','lala'),
	(331,45,8,'1optYw==','2023-06-16 14:06:26','lala'),
	(332,45,8,'1optYw==','2023-06-16 14:06:32','lala'),
	(333,45,8,'x4hvYeST','2023-06-16 14:06:57','lala'),
	(334,45,8,'1oltYPU=','2023-06-16 14:07:44','lala'),
	(335,45,8,'1oltYPU=','2023-06-16 14:08:01','lala'),
	(336,45,8,'1oltYPWR','2023-06-16 14:08:34','lala'),
	(337,45,8,'1YpuY/aR','2023-06-16 14:08:55','lala'),
	(338,45,8,'2oRh','2023-06-16 14:09:38','lala'),
	(339,45,8,'2oRh','2023-06-16 14:09:42','lala'),
	(340,45,8,'0o14bPeA0g==','2023-06-16 14:10:32','lala'),
	(341,45,8,'2414bA==','2023-06-16 14:19:59','lala'),
	(342,45,8,'2414bA==','2023-06-16 14:20:03','lala'),
	(343,45,8,'gt45','2023-06-16 14:22:00','lala'),
	(344,45,8,'3I1nZA==','2023-06-16 14:23:15','lala'),
	(345,45,8,'3I1nZA==','2023-06-16 14:23:27','lala'),
	(346,45,8,'1YpuY/aR0xQ=','2023-06-16 14:23:39','lala'),
	(347,45,8,'24JqYg==','2023-06-16 14:24:39','lala'),
	(348,45,8,'1YVhY/qR','2023-06-16 14:27:47','lala'),
	(349,45,8,'wIlnYPuS','2023-06-16 14:37:42','lala'),
	(350,42,8,'w4l4ZOc=','2023-06-16 14:41:06','pepe'),
	(351,42,8,'w4l4ZOc=','2023-06-16 14:41:06','pepe'),
	(352,45,8,'3o14bA==','2023-06-16 14:41:17','lala'),
	(353,45,8,'3o14bA==','2023-06-16 14:41:17','lala'),
	(354,45,8,'2oVhbPk=','2023-06-16 14:42:18','lala'),
	(355,45,8,'1oltYPWS0A==','2023-06-16 14:44:06','lala'),
	(356,45,8,'1oltYPWS0A==','2023-06-16 14:44:38','lala'),
	(357,45,8,'1oVub/WR','2023-06-16 14:45:24','lala'),
	(358,45,8,'1oVub/WR','2023-06-16 14:46:42','lala'),
	(359,45,8,'3o14bA==','2023-06-16 14:47:58','lala'),
	(360,42,8,'w4l4ZOc=','2023-06-16 15:11:00','pepe'),
	(361,45,8,'1Jtvcvc=','2023-06-16 15:11:05','lala'),
	(362,45,8,'0o15cA==','2023-06-16 15:12:50','lala'),
	(363,42,8,'w4l4ZOc=','2023-06-16 15:12:59','pepe'),
	(364,42,8,'3o14bLOQ2QAu+I8=','2023-06-16 15:22:38','pepe'),
	(365,42,8,'3o14bLOQ2QAu+I8=','2023-06-16 15:22:38','pepe'),
	(366,45,8,'wp55','2023-06-16 15:28:17','lala'),
	(367,45,8,'wo18dw==','2023-06-16 15:29:52','lala'),
	(368,45,8,'3IlnYA==','2023-06-16 15:30:16','lala'),
	(369,45,8,'24l9ZA==','2023-06-16 15:32:28','lala'),
	(370,45,8,'1YBuaf+S0w==','2023-06-16 15:40:42','lala'),
	(371,45,8,'1optYw==','2023-06-16 16:32:30','lala'),
	(372,45,8,'x4Ry','2023-06-16 16:33:22','lala'),
	(373,45,8,'xI17bA==','2023-06-16 16:34:08','lala'),
	(374,42,8,'24llZOOW','2023-06-16 16:34:29','pepe'),
	(375,42,8,'24llZOOW','2023-06-16 16:34:29','pepe'),
	(376,46,8,'2I1nag==','2023-06-16 16:51:58','gege'),
	(377,46,8,'2I1nag==','2023-06-16 16:51:58','gege'),
	(378,45,8,'8ZxqJfiW1BM5+JGF','2023-06-16 16:53:04','lala'),
	(379,45,8,'8ZxqJfiW1BM5+JGF','2023-06-16 16:53:04','lala'),
	(380,46,8,'w41yZLOV1xsgtoqFteNYePeUKA==','2023-06-16 16:53:14','gege'),
	(381,46,8,'w41yZLOV1xsgtoqFteNYePeUKA==','2023-06-16 16:53:14','gege'),
	(382,45,2,'2IRjbQ==','2023-06-17 00:14:36','lala');

/*!40000 ALTER TABLE `chats` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table project_images
# ------------------------------------------------------------

DROP TABLE IF EXISTS `project_images`;

CREATE TABLE `project_images` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int DEFAULT NULL,
  `file_name` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `is_primary` tinyint DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `project_images` WRITE;
/*!40000 ALTER TABLE `project_images` DISABLE KEYS */;

INSERT INTO `project_images` (`id`, `project_id`, `file_name`, `is_primary`, `created_at`, `updated_at`)
VALUES
	(1,1,'photo1',1,'2023-03-19 00:48:36','2023-03-19 00:48:39'),
	(2,1,'photo2',0,'2023-03-19 00:48:59','2023-03-19 00:49:02'),
	(3,2,'photo3',0,'2023-03-19 00:49:14','2023-05-24 17:18:13'),
	(4,2,'photo4',0,'2023-03-19 00:49:28','2023-05-24 17:18:13'),
	(6,7,'mock/images/29-124.jpg',1,'2023-03-22 16:14:34','2023-03-22 16:14:34'),
	(8,8,'mock/images/29-124.jpg',1,'2023-03-22 16:16:32','2023-03-22 16:16:32'),
	(9,2,'mock/images/28-54b03f5a6fcedbfa6621c751d9bed891e78c7d3c.jpg',0,'2023-05-24 16:54:44','2023-05-24 17:18:13'),
	(10,2,'mock/images/28-identityphoto.jpg',0,'2023-05-24 17:15:48','2023-05-24 17:18:13'),
	(11,2,'mock/images/28-maxresdefault.jpg',1,'2023-05-24 17:18:13','2023-05-24 17:18:13'),
	(12,14,'mock/images/40-images.jpeg',1,'2023-05-24 19:55:49','2023-05-24 19:55:49'),
	(13,15,'mock/images/41-images.jpeg',1,'2023-05-25 13:25:49','2023-05-25 13:25:49');

/*!40000 ALTER TABLE `project_images` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table projects
# ------------------------------------------------------------

DROP TABLE IF EXISTS `projects`;

CREATE TABLE `projects` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `project_title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `short_description` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `long_description` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci,
  `perks` text CHARACTER SET utf8mb3,
  `contributor_count` int DEFAULT NULL,
  `goal_amount` int DEFAULT NULL,
  `current_amount` int DEFAULT NULL,
  `slug` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `projects` WRITE;
/*!40000 ALTER TABLE `projects` DISABLE KEYS */;

INSERT INTO `projects` (`id`, `user_id`, `project_title`, `short_description`, `long_description`, `perks`, `contributor_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`)
VALUES
	(1,29,'Startup','tech startup','tech startupppppppp','abc',1,50000000,50000,'slug1','2023-03-18 22:59:44','2023-03-18 22:59:49'),
	(2,28,'Startup 2','tech startup 2','tech startupppppppp 2\n','abc',50,50000000,500000,'slug2','2023-03-18 23:00:32','2023-05-24 17:25:46'),
	(7,29,'CreateCampaignTestABC','CreateCampaignTest-2','CreateCampaignTest-1','CreateCampaignTest-1, CreateCampaignTest-1a, CreateCampaignTest-1b',2,1200000,26450,'createcampaigntest-29','2023-03-21 02:32:26','2023-03-29 07:27:27'),
	(8,29,'CreateCampaignTestCGV','CreateCampaignTest-2','CreateCampaignTest-1','CreateCampaignTest-2, CreateCampaignTest-1a, CreateCampaignTest-1b',3,1500000,1512000,'createcampaigntest-29','2023-03-21 12:06:56','2023-05-25 14:10:14'),
	(9,29,'CreateCampaignTestA','CreateCampaignTest-3','CreateCampaignTest-1','CreateCampaignTest-1, CreateCampaignTest-1a, CreateCampaignTest-1b',0,1200000,0,'createcampaigntesta-29','2023-03-21 12:09:28','2023-03-21 12:09:28'),
	(10,29,'CreateCampaignTestGGG','CreateCampaignTest-5','CreateCampaignTest-1','CreateCampaignTest-1, CreateCampaignTest-1a, CreateCampaignTest-1b',0,1000000,0,'createcampaigntestggg-29','2023-03-22 14:39:51','2023-03-22 14:39:51'),
	(11,28,'nasi','font','halo bandung','front, front, nasi',0,1250000,0,'nasi-28','2023-05-24 15:56:43','2023-05-24 15:56:43'),
	(12,28,'lala','wawa','pulupululpul','like, like, like',0,12345678,0,'lala-28','2023-05-24 16:01:58','2023-05-24 16:01:58'),
	(13,28,'demi','Lala','lulalu','ayam, nasi goreng, piring',0,1230000,0,'demi-28','2023-05-24 16:18:24','2023-05-24 16:18:24'),
	(14,40,'Hello','hello','Hello','hello, hello',2,200000,20000,'hello-40','2023-05-24 19:55:23','2023-05-25 13:28:42'),
	(15,41,'CampingABC','This is a campaign good.','This is a campaign good.','Good, tasty',0,200000,0,'campaigna-41','2023-05-25 13:25:20','2023-05-25 13:38:02');

/*!40000 ALTER TABLE `projects` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table rooms
# ------------------------------------------------------------

DROP TABLE IF EXISTS `rooms`;

CREATE TABLE `rooms` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `private` tinyint DEFAULT NULL,
  `room_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



# Dump of table transactions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `transactions`;

CREATE TABLE `transactions` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `amount` int DEFAULT NULL,
  `status` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `receipt_no` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `transaction_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;

INSERT INTO `transactions` (`id`, `project_id`, `user_id`, `amount`, `status`, `receipt_no`, `transaction_url`, `created_at`, `updated_at`)
VALUES
	(1,7,29,10000,'SUCCESS','SUCCESS',NULL,'2023-03-23 01:24:45','2023-03-23 01:24:48'),
	(2,22,28,15000,'FAILED','FAILED',NULL,'2023-03-26 11:24:45','2023-03-26 11:24:45'),
	(3,7,28,20000,'FAILED','FAILED',NULL,'2023-03-26 12:24:45','2023-03-26 12:24:45'),
	(4,7,28,10000,'PENDING','',NULL,'2023-03-28 03:00:47','2023-03-28 03:00:47'),
	(5,7,28,12000,'PENDING','','','2023-03-29 03:29:00','2023-03-29 03:29:00'),
	(6,7,28,12000,'PENDING','','https://app.sandbox.midtrans.com/snap/v3/redirection/a95bdf0f-5bf2-4cde-9004-dece4aff2ce8','2023-03-29 03:31:08','2023-03-29 03:31:08'),
	(7,7,28,12000,'PENDING','','https://app.sandbox.midtrans.com/snap/v3/redirection/2849bbc4-0526-4201-a8d1-db5505fe9921','2023-03-29 03:35:43','2023-03-29 03:35:44'),
	(8,7,28,12000,'PENDING','','https://app.sandbox.midtrans.com/snap/v3/redirection/6a99f9db-611e-4687-bc62-9619c22761ea','2023-03-29 03:36:13','2023-03-29 03:36:14'),
	(9,7,28,12000,'PENDING','','https://app.sandbox.midtrans.com/snap/v3/redirection/a4ec1b40-e561-480b-aaf2-a4013d4d3a84','2023-03-29 03:53:20','2023-03-29 03:53:20'),
	(10,7,28,12000,'PENDING','','https://app.sandbox.midtrans.com/snap/v3/redirection/a384ac6c-a80d-4c52-b23c-45b0fef87e48','2023-03-29 03:55:20','2023-03-29 03:55:20'),
	(11,7,28,13200,'PENDING','','','2023-03-29 07:14:14','2023-03-29 07:14:14'),
	(12,7,28,13200,'SUCCESS','','https://app.sandbox.midtrans.com/snap/v3/redirection/c389bcc1-dd5d-4b4d-8c0d-50c10c5f93a3','2023-03-29 07:19:14','2023-03-29 07:27:27'),
	(13,7,28,13250,'SUCCESS','','https://app.sandbox.midtrans.com/snap/v3/redirection/298f5021-3a01-47ac-a7f7-05a16a28ce1a','2023-03-29 07:21:50','2023-03-29 07:22:18'),
	(14,8,28,50000,'PENDING','','','2023-05-24 13:57:32','2023-05-24 13:57:32'),
	(15,8,28,50000,'PENDING','','','2023-05-24 13:57:39','2023-05-24 13:57:39'),
	(16,8,28,50000,'PENDING','','https://app.sandbox.midtrans.com/snap/v3/redirection/30b65aa2-3605-415a-b133-c7dd695cedd0','2023-05-24 13:59:17','2023-05-24 13:59:18'),
	(17,8,28,12000,'PENDING','','https://app.sandbox.midtrans.com/snap/v3/redirection/86f8ef49-b079-4e34-b8da-8fb9106b9476','2023-05-24 13:59:46','2023-05-24 13:59:47'),
	(18,8,28,12000,'PENDING','','https://app.sandbox.midtrans.com/snap/v3/redirection/60562acb-cf2c-4190-b783-87cf80974bbb','2023-05-24 14:00:50','2023-05-24 14:00:50'),
	(19,8,28,14000,'PENDING','','https://app.sandbox.midtrans.com/snap/v3/redirection/4d762783-dbf7-4cef-bea2-b8ad7aa5834c','2023-05-24 14:01:19','2023-05-24 14:01:20'),
	(20,8,28,13250,'PENDING','','https://app.sandbox.midtrans.com/snap/v3/redirection/d6c76d16-9706-406d-8cbc-8ec07ed12ad1','2023-05-24 14:03:35','2023-05-24 14:03:36'),
	(21,8,28,12000,'SUCCESS','','https://app.sandbox.midtrans.com/snap/v3/redirection/234bb004-5560-4c91-ac2e-4f7a92da56b4','2023-05-24 14:04:07','2023-05-24 14:06:02'),
	(22,8,28,500000,'SUCCESS','','https://app.sandbox.midtrans.com/snap/v3/redirection/be3de8a0-a9e8-4fdf-8eba-a32711728870','2023-05-24 14:18:58','2023-05-24 14:19:25'),
	(23,14,28,10000,'SUCCESS','','https://app.sandbox.midtrans.com/snap/v3/redirection/62c967a4-522c-466e-b6d8-984419402c59','2023-05-25 12:50:52','2023-05-25 12:51:37'),
	(24,14,41,10000,'SUCCESS','','https://app.sandbox.midtrans.com/snap/v3/redirection/36fdbfe7-cd6a-4176-8e99-98fd71ef2728','2023-05-25 13:27:40','2023-05-25 13:28:42'),
	(25,8,28,1000000,'SUCCESS','','https://app.sandbox.midtrans.com/snap/v3/redirection/7806563a-1a25-40bb-a913-0b06260e721c','2023-05-25 14:09:52','2023-05-25 14:10:14');

/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `phone_no` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `pin` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `avatar_file_name` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `role` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `token` varchar(255) CHARACTER SET utf8mb3 DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `name`, `phone_no`, `pin`, `avatar_file_name`, `role`, `token`, `created_at`, `updated_at`, `status`)
VALUES
	(1,'TEST-NAME-1','TEST-EMAIL-1','TEST-PW-1','','TEST-ROLE-1','TEST-TOKEN-1','2023-02-16 00:17:37','2023-03-18 15:55:48',NULL),
	(2,'TEST-NAME-2','TEST-EMAIL-2','TEST-PW-2','TEST-AFN-2','TEST-ROLE-2','TEST-TOKEN-2','2023-02-16 00:17:37','2023-02-16 00:17:43',NULL),
	(3,'TEST-NAME-3','','','','',NULL,'2023-02-16 07:55:31','2023-02-16 07:55:31',NULL),
	(4,'TEST-NAME-SERVICE-4','TEST-EMAIL-SERVICE-4','$2a$04$VkNSraN4pSMlfqFRp0jGB.24XPW67wqfEqmbmVM.h4WtMMiUTd7AO','','user',NULL,'2023-02-16 09:36:30','2023-02-16 09:36:30',NULL),
	(5,'TEST-NAME-HANDLER-5','TEST-EMAIL-HANDLER-5','$2a$04$mYi7GvlA7iGXiX09yolgVOIcDkO3RUuF8n70IyJ1qh3I9FLb2dkWi','','user',NULL,'2023-02-18 01:32:58','2023-02-18 01:32:58',NULL),
	(6,'TEST-NAME-HANDLER-6','TEST-EMAIL-HANDLER-6','$2a$04$7tjuOZMhs/OOVc5CAgejpOJx30FCsV5dK1U1XyGCoQz5m753K3jXa','','user',NULL,'2023-02-18 01:36:56','2023-02-18 01:36:56',NULL),
	(7,'TEST-NAME-HANDLER-7','TEST-EMAIL-HANDLER-7','$2a$04$bwFCjfWtXmolGY7GmEDab.WB4olW9VsM12moLEg5X4pXLYA/gMBP2','','user',NULL,'2023-02-18 02:16:35','2023-02-18 02:16:35',NULL),
	(8,'TEST-NAME-HANDLER-8','TEST-EMAIL-HANDLER-8','$2a$04$O4EZlLB7gYI9xbGM4SNcgO21MbEdCE1pGQ3DSHpIuvfJ66/Z2E4Bq','','user',NULL,'2023-02-18 03:53:27','2023-02-18 03:53:27',NULL),
	(9,'TEST-NAME-HANDLER-9','TEST-EMAIL-HANDLER-9','$2a$04$Wfvny2QbYxqrrFHfkgaIkePAyS5DWAficJrI1Caq6l4dzApH6psUa','','user',NULL,'2023-02-18 05:27:07','2023-02-18 05:27:07',NULL),
	(11,'TestNEW','','','','',NULL,'2023-02-19 03:04:02','2023-02-19 03:04:02',NULL),
	(12,'TEST-NEW-2','TEST-NEW-2','7608d3b76d90745094eb9c5665e0cd999f0bbaa9145516d5604ffe8b465d329ba0966dcd9be67905e95cc5c67599c2ddc8e26a434f2eccd30aa69e87d37e52c5','','user',NULL,'2023-02-19 03:46:19','2023-02-19 03:46:19',NULL),
	(13,'TEST-NEW-3','TEST-NEW-3','$2a$10$tjro0Pb/LsO4wa5ODFlMFuCw5qw.5KG35/9LJAwn8lY2uxIdehWsK','','user',NULL,'2023-02-19 09:08:43','2023-02-19 09:08:43',NULL),
	(14,'TEST-NEW-4','TEST-NEW-4','$2a$10$dEe366Maeu0hnpKWj.m8ZulEBK6ZUi3FSgk5W0HdC72UzHU2.XtwS','','user',NULL,'2023-02-20 21:25:55','2023-02-20 21:25:55',NULL),
	(15,'TEST-NEW-4','TEST-NEW-4','$2a$10$Fu1.Kmv0ccO6jDG6jmKZTe1uoayixdh.QOJ4LoOkIh9LL7Xu2WPom','','user',NULL,'2023-02-20 21:52:36','2023-02-20 21:52:36',NULL),
	(16,'TEST-NEW-5','','$2a$10$ma9LHMw7b6YPOopJ9Dp8Q.qj1ihOKFbeLbgchxi3HSFJeSMZ2zgrS','','user',NULL,'2023-02-20 21:54:33','2023-02-20 21:54:33',NULL),
	(17,'TEST-NEW-6','TEST-NEW-6','$2a$10$bM9ptFrQm9MJl18G06dZXunW8cV9IbedHzIpQKHiKRExKkAu/cey.','','user',NULL,'2023-02-20 21:54:54','2023-02-20 21:54:54',NULL),
	(18,'TEST-NEW-7','','$2a$10$9u.1sX89pdGJXgnZ2/jP9eUvkBzIcO6tImz0BHmFsrfovitCZlmHG','','user',NULL,'2023-02-20 21:55:42','2023-02-20 21:55:42',NULL),
	(19,'TEST-NEW-8','TEST-NEW-8','$2a$10$6J4nzI1XAezY2/Tj39wpFOSJL9SEnqxgadiTUggbsNHhX.a2jc3bC','','user',NULL,'2023-02-20 21:56:03','2023-02-20 21:56:03',NULL),
	(20,'TEST-NEW-9','TEST-NEW-9','$2a$10$92NX6br4agIx9ygKyV5yUO9KpIpunRlrwiMwD7ppepHhxrU0Um7dC','','user',NULL,'2023-02-20 22:20:34','2023-02-20 22:20:34',NULL),
	(21,'TEST-NEW-10','TEST-NEW-10','$2a$10$bUG9YQk8DWM/2lkF8s54xOIev/hFHLcb1QnHNQ3VWWy8eBg5I6Iyi','','user',NULL,'2023-02-20 22:42:36','2023-02-20 22:42:36',NULL),
	(22,'TEST-NEW-11','TEST-NEW-11','$2a$10$JmhuVLtSn1uJ42.gnWGlXOfOQAuxPNzpeTl6PPvbNxLz4JZhWRDCG','','user',NULL,'2023-02-20 22:43:00','2023-02-20 22:43:00',NULL),
	(23,'TEST-NEW-23','TEST-NEW-23','$2a$10$mShSfA5wwkNPKWRZE9mPWu1/PWCS9RCoA3dtLgBU5R..PDiaemvCe','','user',NULL,'2023-03-17 00:32:08','2023-03-17 00:32:08',NULL),
	(24,'TEST-NEW-24','TEST-NEW-24','$2a$10$Yz63NoFnBMcQjxuk3CwHQORZNLMFHk3knHuQ4VVaba5faiEZY64aK','','user',NULL,'2023-03-17 00:58:10','2023-03-17 00:58:10',NULL),
	(25,'TEST-NEW-25','TEST-NEW-25','$2a$10$SCGOil58OhcVvO297t2I4OVC23c88CEzjLkK6Tod2N.mwp1.lbRGG','','user',NULL,'2023-03-17 01:05:03','2023-03-17 01:05:03',NULL),
	(26,'TEST-NEW-26','','$2a$10$BYULjLcI8PuV7VN5sw0VE.gI.unNXrFGoYwwc5wnE7cvp3PG/sFZ.','','user',NULL,'2023-03-17 01:13:34','2023-03-17 01:13:34',NULL),
	(27,'Blabla','081234567890','$2a$04$15MQYvlupSTYiwoH0CIUZuuhA3Kf9o/BVf4N8TPiB7ohFIJ3C9326','','user',NULL,'2023-03-17 02:00:25','2023-06-07 13:43:39',NULL),
	(28,'TEST-JEREMY-9','081234567899','$2a$04$.VUSXa2yLccH/.otvHS1du.PqJFBf2N6yvms5GGcv4LZOgX9905Fi','','user',NULL,'2023-03-17 02:01:23','2023-06-01 15:16:09',NULL),
	(29,'TEST-JEREMY-3','081234567892','$2a$04$WLAJrH4hqoo73tCkLho/PelsLi7IpmC9rzYEj4PQi7hSRWWFlp8PS','mock/images/54b03f5a6fcedbfa6621c751d9bed891e78c7d3c.png','user',NULL,'2023-03-17 02:24:16','2023-03-21 02:33:22',NULL),
	(33,'Jeremy','081234567890','$2a$04$eq/xHRXFu/ffNItvDIs1lu/c40OuzsmXlPSrnQtE3QuoPHhKt4Igu','mock/images/lockicon.png','user',NULL,'2023-05-24 00:58:16','2023-05-24 00:58:38',NULL),
	(34,'Jeremy','081234567890','$2a$04$A.vvhzs8l74fmAI8RZWevOGGBpMUbwsdMg4THhqz70TK9i8O9o1O6','mock/images/1677813907 copy.png','user',NULL,'2023-05-24 11:03:33','2023-05-24 11:04:02',NULL),
	(35,'jeje','081234567890','$2a$04$bvzY0Ak.i6sb/ZaOdK3B8epDC0OSp.bchqhA6vOffWVXhhamiq7Qe','mock/images/1042582.jpg','user',NULL,'2023-05-24 11:41:41','2023-05-24 11:42:14',NULL),
	(36,'hello','081234567890','$2a$04$e8lp9B03qwLCQ0sj2Tq6qeUSrJIi3UN7V7WXSSAATmeYIdp4BkyXe','','user',NULL,'2023-05-24 11:46:49','2023-05-24 11:46:49',NULL),
	(37,'TEST-JEREMY-9','081234567895','$2a$04$2GA54ztm3MjukqDKrtyChubdYZGplsSk0SYdGlGxLYK7kaRqX9ihe','','user',NULL,'2023-05-24 11:48:13','2023-05-24 11:48:13',NULL),
	(38,'Julia Keeva','081234567895','$2a$04$Y6w//g7tvKtBrFm0/7X8Q.gDSBj4IwVuVcQRuQJxIgklyA6nz/AZy','','user',NULL,'2023-05-24 11:49:39','2023-05-24 11:49:39',NULL),
	(39,'Julia Keeva','081234567899','$2a$04$1pYE3L/G7CIMnHxlxvbJKOF8aLTbgyfzVScv.3WhuZrgtzB7LiyGW','mock/images/selfie2.jpg','user',NULL,'2023-05-24 11:50:36','2023-05-24 11:50:56',NULL),
	(40,'Julia Keeva','081234567897','$2a$04$t8wm4E/2cnXpwANf0wFeKuxZcAAtRMPog0C5ooMQ4oYqiWxoikXa.','mock/images/identityphoto.jpg','user',NULL,'2023-05-24 19:52:30','2023-05-24 19:53:48',NULL),
	(41,'Hello','081234567889','$2a$04$x5CCzjpH7rM8/TSFarjqfe9D2Td4IDtiuTcCY1FrbCYz7X.OtgzM6','mock/images/Streets-Wallpaper.png','user',NULL,'2023-05-25 13:22:27','2023-05-25 13:23:11',NULL),
	(42,'pepe','081234567877','$2a$04$LBUI9ggTK5CgiHMCFERYbOaUYXy0DEjynEJXFraYO4kB25n59pSc2','','user',NULL,'2023-06-01 14:11:25','2023-06-01 16:13:05',NULL),
	(43,'Julia Keeva','081234567890','$2a$04$RKRozHUSK1l1uMSv2ar73.s0iF3arzDgHNyz2QH4eMfXOwqOtcbKm','','user',NULL,'2023-06-01 15:55:50','2023-06-01 15:55:50',NULL),
	(44,'Julia Keeva','081234567811','$2a$04$oqDJqYzfW24A2hgX4MFAN.yqvzvUZ8Op.mWR3Ug3D.ncwy0K3P9vK','','user',NULL,'2023-06-06 04:20:02','2023-06-06 04:20:02',NULL),
	(45,'lala','081234567855','$2a$04$TqmI6nlzZVkQQOQucG4lMeih6cVPL61Pp1YgqmQY5L4fnW6g4j5AK','','user',NULL,'2023-06-15 16:16:32','2023-06-15 16:16:32',NULL),
	(46,'gege','081234567856','$2a$04$qweNJZQH.6TBNNkDjXNcp.zHszLpZB0bBFEF0JUI8T5EXpDI5Hnxa','','user',NULL,'2023-06-16 16:51:12','2023-06-16 16:51:12',NULL),
	(47,'Jeremy','081234567844','$2a$04$czqpHMzfSud0c/7iADGznu89VbXQdrzUMWrxtQo3eydt0oU8Q0Rcm','','user',NULL,'2023-06-20 03:31:15','2023-06-20 03:32:39',NULL);

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
