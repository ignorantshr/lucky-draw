DROP DATABASE IF EXISTS `lucky_draw`;
CREATE DATABASE IF NOT EXISTS `lucky_draw`;

USE `lucky_draw`;

SET FOREIGN_KEY_CHECKS=0;

SELECT 'create table prize';
DROP TABLE IF EXISTS `prize`;
CREATE TABLE `prize` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '名称',
  `url` varchar(255) DEFAULT NULL COMMENT '图片链接',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name_INDEX` (`name`) USING BTREE
) ENGINE=InnoDB COMMENT='奖品表';
INSERT INTO `prize` VALUES(1, "空奖", 'https://img12.360buyimg.com/n1/jfs/t3772/105/602238500/79292/dd8c8f6f/580ea45eN030695f3.jpg');

SELECT 'create table prize_pool';
DROP TABLE IF EXISTS `prize_pool`;
CREATE TABLE `prize_pool` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '名称',
  `type` int(4) NOT NULL DEFAULT 1 COMMENT '抽奖类型',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name_INDEX` (`name`) USING BTREE
) ENGINE=InnoDB COMMENT='奖池表';

SELECT 'create table prize_pool_prize';
DROP TABLE IF EXISTS `prize_pool_prize`;
CREATE TABLE `prize_pool_prize` (
  `prize_pool_id` bigint(20) NOT NULL COMMENT '奖池id',
  `prize_id` bigint(20) NOT NULL COMMENT '奖品id',
  `prize_probability` int(5) DEFAULT 0 COMMENT '中奖概率',
  `prize_number` bigint(20) DEFAULT 0 COMMENT '奖品数量',
  PRIMARY KEY (`prize_pool_id`, `prize_id`) USING BTREE
) ENGINE=InnoDB COMMENT='奖池-奖品关联表';

SELECT 'create table user';
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '名称',
  `external_id` varchar(500) DEFAULT NULL COMMENT '其他系统的账户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB COMMENT='账户表';

SELECT 'create table user_prize';
DROP TABLE IF EXISTS `user_prize`;
CREATE TABLE `user_prize` (
  `user_id` bigint(20) DEFAULT 0 COMMENT '用户id',
  `prize_pool_id` bigint(20) NOT NULL COMMENT '奖池id',
  `prize_id` bigint(20) NOT NULL COMMENT '奖品id',
  PRIMARY KEY (`user_id`, `prize_pool_id`, `prize_id`) USING BTREE
) ENGINE=InnoDB COMMENT='账户-奖品关联表';
