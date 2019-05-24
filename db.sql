CREATE DATABASE `artists` DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;

CREATE TABLE `artists` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
    `poster` VARCHAR(600) CHARACTER SET utf8 COLLATE utf8_general_ci,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `albums` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `artist_id` INT NOT NULL,
    `name` VARCHAR(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    UNIQUE KEY `idx_art_id_name` (`name`,`artist_id`) USING BTREE,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `store_type` (
    `name` VARCHAR(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    PRIMARY KEY (`name`)
) ENGINE=InnoDB;

CREATE TABLE `stores` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `artist_id` INT NOT NULL,
    `store_name` VARCHAR(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `store_id` VARCHAR(300) NOT NULL,
    UNIQUE KEY `idx_artist_store_info` (`artist_id`,`store_name`,`store_id`) USING BTREE,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;
