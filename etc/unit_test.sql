/*
 Navicat Premium Data Transfer

 Source Server         : mysqlWSL
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : unit_test

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 14/11/2022 18:09:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for host
-- ----------------------------
DROP TABLE IF EXISTS `host`;
CREATE TABLE `host`  (
  `resource_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `cpu` tinyint NOT NULL,
  `memory` int NOT NULL,
  `gpu_amount` tinyint NULL DEFAULT NULL,
  `gpu_spec` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `os_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `os_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `serial_number` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`resource_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of host
-- ----------------------------

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource`  (
  `id` char(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '资源的实例Id',
  `vendor` tinyint(1) NOT NULL,
  `region` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `create_at` bigint NOT NULL,
  `expire_at` bigint NOT NULL,
  `type` varchar(120) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `update_at` bigint NOT NULL,
  `sync_at` bigint NOT NULL,
  `account` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `public_ip` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `private_ip` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `name`(`name` ASC) USING BTREE,
  INDEX `status`(`status` ASC) USING BTREE,
  INDEX `private_ip`(`public_ip` ASC) USING BTREE,
  INDEX `public_ip`(`public_ip` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of resource
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
