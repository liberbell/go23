CREATE DATABASE `company`;

USE `company`;

CREATE TABLE IF NOT EXISTS `employee` (
    `id` INT unsigned NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20),
    `city` VARCHAR(20) DEFAULT "fredericton",
    `dept` VARCHAR(20) DEFAULT "onboarding",
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CARSET=utf8mb4;