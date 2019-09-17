CREATE DATABASE  IF NOT EXISTS `go-mysql-crud` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `go-mysql-crud`;

/* Create table for `events` */

DROP TABLE IF EXISTS `events`;
CREATE TABLE `events` (

  `event_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `description` longtext COLLATE utf8_unicode_ci NOT NULL,
  `location` longtext COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
);

/* Create table for `event_subscription` */
DROP TABLE IF EXISTS `event_subscription`;
CREATE TABLE `event_subscription` (

  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `event_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`user_id`)
  FOREIGN KEY (event_id) REFERENCES events(event_id)
);