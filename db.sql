CREATE DATABASE `muxisite_auth`;

USE `muxisite_auth`;

CREATE TABLE `roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL,
  `default` tinyint(1) DEFAULT NULL,
  `permissions` int(11) DEFAULT NULL,

  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `ix_roles_default` (`default`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO `roles` (`name`, `default`, `permissions`)
VALUES ('Moderator', 0, 14);

INSERT INTO `roles` (`name`, `default`, `permissions`)
VALUES ('Administrator', 0, 255);

INSERT INTO `roles` (`name`, `default`, `permissions`)
VALUES ('User', 1, 6);

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(164) DEFAULT NULL,
  `info` text,
  `username` varchar(164) DEFAULT NULL,
  `avatar_url` text,
  `personal_blog` text,
  `github` text,
  `flickr` text,
  `weibo` text,
  `zhihu` text,
  `password_hash` varchar(164) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `birthday` varchar(164) DEFAULT NULL,
  `group` varchar(164) DEFAULT NULL,
  `hometown` varchar(164) DEFAULT NULL,
  `left` tinyint(1) DEFAULT NULL,
  `timejoin` varchar(164) DEFAULT NULL,
  `timeleft` varchar(164) DEFAULT NULL,
  `reset_t` varchar(164) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `role_id` (`role_id`),
  CONSTRAINT `users_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_identities` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `provider` varchar(32) NOT NULL,
  `provider_subject` varchar(164) NOT NULL,
  `email` varchar(164) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `provider_subject` (`provider`, `provider_subject`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_identities_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `member_profiles` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '关联 users.id',
  `real_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `student_id` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '学号',
  `group` VARCHAR(32) NOT NULL COMMENT '所属组别: Frontend, Backend, Design, Product, Android 等',
  `join_year` INT NOT NULL COMMENT '加入年份/届数 (如 2023)',
  `personal_blog` VARCHAR(255) DEFAULT '',
  `github` VARCHAR(255) DEFAULT '',
  `zhihu` VARCHAR(255) DEFAULT '',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_student_id` (`student_id`),
  KEY `idx_group` (`group`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;