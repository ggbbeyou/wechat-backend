/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : chat

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 18/09/2022 21:42:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `cid` int(0) NOT NULL COMMENT '评论编号id',
  `p_cid` int(0) NOT NULL COMMENT ' 父级评论id   0代表评论说说',
  `uid` bigint(0) NOT NULL COMMENT 'uid',
  `tid` bigint(0) NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '评论时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, 666, 0, 123, NULL, '哈哈哈', NULL);
INSERT INTO `comments` VALUES (2, 667, 0, 124, NULL, 'xixixi', NULL);
INSERT INTO `comments` VALUES (3, 668, 666, 124, NULL, '测试回复', NULL);

-- ----------------------------
-- Table structure for friendships
-- ----------------------------
DROP TABLE IF EXISTS `friendships`;
CREATE TABLE `friendships`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id 雪花算法',
  `uid1` bigint(0) NOT NULL COMMENT '用户1 uid',
  `uid2` bigint(0) NOT NULL COMMENT '用户2 uid',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '加好友时间',
  `is_dalete` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '逻辑删除 1 删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uid1_uid2_idx`(`uid1`, `uid2`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of friendships
-- ----------------------------
INSERT INTO `friendships` VALUES (1, 123, 124, '2022-09-09 14:45:19', 0);
INSERT INTO `friendships` VALUES (2, 124, 123, '2022-09-09 14:45:44', 0);
INSERT INTO `friendships` VALUES (3, 123, 125, '2022-09-11 11:00:34', 0);
INSERT INTO `friendships` VALUES (4, 125, 123, '2022-09-11 11:00:45', 0);
INSERT INTO `friendships` VALUES (5, 123, 126, '2022-09-11 11:00:58', 0);
INSERT INTO `friendships` VALUES (6, 126, 123, '2022-09-11 11:01:06', 0);
INSERT INTO `friendships` VALUES (7, 126, 124, '2022-09-16 10:46:07', 0);
INSERT INTO `friendships` VALUES (8, 124, 126, '2022-09-16 10:46:16', 0);

-- ----------------------------
-- Table structure for groupfriends
-- ----------------------------
DROP TABLE IF EXISTS `groupfriends`;
CREATE TABLE `groupfriends`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id ',
  `uid` bigint(0) NOT NULL COMMENT '用户id',
  `gid` bigint(0) NOT NULL COMMENT '群聊id',
  `indentity` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 群主 1 成员',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '加入群聊时间',
  `exit_time` datetime(0) NULL DEFAULT NULL COMMENT '退出群聊时间',
  `is_delete` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '逻辑删除 用户退出群聊时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uid_gid_idx`(`uid`, `gid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of groupfriends
-- ----------------------------
INSERT INTO `groupfriends` VALUES (1, 123, 333, 1, '2022-09-09 14:52:52', NULL, 0);
INSERT INTO `groupfriends` VALUES (2, 123, 335, 1, '2022-09-09 14:55:31', NULL, 0);
INSERT INTO `groupfriends` VALUES (3, 124, 334, 1, '2022-09-09 14:55:41', NULL, 0);
INSERT INTO `groupfriends` VALUES (4, 123, 334, 0, '2022-09-11 12:39:05', NULL, 0);
INSERT INTO `groupfriends` VALUES (5, 124, 333, 0, NULL, NULL, 0);

-- ----------------------------
-- Table structure for groups
-- ----------------------------
DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id ',
  `gid` bigint(0) NOT NULL COMMENT '群聊id 雪花算法',
  `uid` bigint(0) NOT NULL COMMENT '用户id 创建者',
  `gname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '群聊名称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '群聊头像',
  `count` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '人数 通过触发器维护',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `gid_uid_idx`(`gid`, `uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of groups
-- ----------------------------
INSERT INTO `groups` VALUES (1, 333, 123, '东风快递群1', 'https://img2.baidu.com/it/u=3624909041,2918634060&fm=253&fmt=auto&app=138&f=JPEG?w=520&h=500', 2, '2022-09-11 13:07:01', NULL);
INSERT INTO `groups` VALUES (2, 334, 124, '西北风快递群1', 'https://img2.baidu.com/it/u=3624909041,2918634060&fm=253&fmt=auto&app=138&f=JPEG?w=520&h=500', 2, '2022-09-11 13:07:05', NULL);
INSERT INTO `groups` VALUES (3, 335, 123, '哈哈群', 'https://img2.baidu.com/it/u=3624909041,2918634060&fm=253&fmt=auto&app=138&f=JPEG?w=520&h=500', 1, '2022-09-11 13:07:08', NULL);

-- ----------------------------
-- Table structure for recodes
-- ----------------------------
DROP TABLE IF EXISTS `recodes`;
CREATE TABLE `recodes`  (
  `id` bigint(0) NOT NULL COMMENT '主键',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息内容',
  `from` bigint(0) NOT NULL COMMENT '发送用户id',
  `to` bigint(0) NOT NULL COMMENT '接受用户id',
  `type` int(0) NOT NULL DEFAULT 1 COMMENT '消息类型 1用户消息',
  `send_time` datetime(0) NOT NULL COMMENT '发送时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of recodes
-- ----------------------------
INSERT INTO `recodes` VALUES (6969797308397387776, '3123131', 123, 126, 0, '2022-09-17 13:29:23');
INSERT INTO `recodes` VALUES (6969797403272544256, '达瓦达瓦', 126, 123, 0, '2022-09-17 13:29:46');
INSERT INTO `recodes` VALUES (6969797612643811328, '大王的', 123, 124, 0, '2022-09-17 13:30:36');
INSERT INTO `recodes` VALUES (6969797612773834752, 'awdawd', 123, 124, 0, '2022-09-17 13:30:36');
INSERT INTO `recodes` VALUES (6969797616846503936, 'dawdawd', 123, 124, 0, '2022-09-17 13:30:37');
INSERT INTO `recodes` VALUES (6969797621313437696, '', 123, 124, 0, '2022-09-17 13:30:38');
INSERT INTO `recodes` VALUES (6969797629458776064, 'awdawdaw', 123, 124, 0, '2022-09-17 13:30:40');
INSERT INTO `recodes` VALUES (6969797633632108544, 'awdawdawd', 123, 124, 0, '2022-09-17 13:30:41');
INSERT INTO `recodes` VALUES (6969797914155548672, 'awdawdawd', 123, 124, 0, '2022-09-17 13:31:48');
INSERT INTO `recodes` VALUES (6969797914231046144, 'awdawdaw', 123, 124, 0, '2022-09-17 13:31:48');
INSERT INTO `recodes` VALUES (6969797914243629056, '', 123, 124, 0, '2022-09-17 13:31:48');
INSERT INTO `recodes` VALUES (6969797914264600576, 'dawdawd', 123, 124, 0, '2022-09-17 13:31:48');
INSERT INTO `recodes` VALUES (6969797914277183488, 'awdawd', 123, 124, 0, '2022-09-17 13:31:48');
INSERT INTO `recodes` VALUES (6969797914289766400, '大王的', 123, 124, 0, '2022-09-17 13:31:48');
INSERT INTO `recodes` VALUES (6969798218817208320, 'llll', 124, 123, 0, '2022-09-17 13:33:01');
INSERT INTO `recodes` VALUES (6970185269605564416, '哈哈哈', 123, 126, 0, '2022-09-18 15:12:33');
INSERT INTO `recodes` VALUES (6970185889452392448, '看看看看', 126, 123, 0, '2022-09-18 15:15:01');
INSERT INTO `recodes` VALUES (6970186456031559680, '擦哇', 126, 123, 0, '2022-09-18 15:17:16');
INSERT INTO `recodes` VALUES (6970186461421240320, 'awdaw1', 123, 126, 0, '2022-09-18 15:17:18');

-- ----------------------------
-- Table structure for talks
-- ----------------------------
DROP TABLE IF EXISTS `talks`;
CREATE TABLE `talks`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id ',
  `tid` bigint(0) NOT NULL COMMENT '说说id uuid',
  `uid` bigint(0) NOT NULL COMMENT 'uid用户id',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '说说内容',
  `imgs` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '图片地址',
  `like` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数量',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '发表时间',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `is_delate` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '逻辑删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uid_tid_idx`(`uid`, `tid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of talks
-- ----------------------------
INSERT INTO `talks` VALUES (1, 444, 123, '哈哈哈', NULL, 0, NULL, NULL, 0);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '主键id ',
  `uid` bigint(0) NOT NULL COMMENT '用户id 雪花算法',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码 32位加密',
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `age` int(0) NULL DEFAULT NULL COMMENT '年龄',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像地址',
  `introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `state` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0正常 1=禁用',
  `is_delete` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '逻辑删除',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新信息时间',
  `last_time` datetime(0) NULL DEFAULT NULL COMMENT '上次登陆时间',
  `last_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '上次登录地址',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniq_uid_idx`(`uid`) USING BTREE,
  UNIQUE INDEX `uniq_username_idx`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 123, '1', '1', '李四', 15, 'https://img2.baidu.com/it/u=3624909041,2918634060&fm=253&fmt=auto&app=138&f=JPEG?w=520&h=500', '我最帅', '342@163.com', 0, 0, '2022-09-10 21:01:16', '2022-09-10 21:01:19', '2022-09-10 21:01:23', '湖北黄冈');
INSERT INTO `users` VALUES (2, 124, 'zhangsan', '12345678', '张三', 18, 'https://img2.baidu.com/it/u=3624909041,2918634060&fm=253&fmt=auto&app=138&f=JPEG?w=520&h=500', '我最帅', '34@163.com2', 0, 0, '2022-09-11 10:35:25', '2022-09-11 10:35:27', '2022-09-11 10:35:30', '湖北武汉');
INSERT INTO `users` VALUES (3, 125, 'zhaoliu', '12345678', '赵六', 14, 'https://img2.baidu.com/it/u=3624909041,2918634060&fm=253&fmt=auto&app=138&f=JPEG?w=520&h=500', '我最帅', '34@163.com2', 0, 0, '2022-09-11 10:59:09', '2022-09-11 10:59:12', '2022-09-11 10:59:17', '湖北武汉');
INSERT INTO `users` VALUES (4, 126, 'liqi', '12345678', '李七', 13, 'https://img2.baidu.com/it/u=3624909041,2918634060&fm=253&fmt=auto&app=138&f=JPEG?w=520&h=500', '我最帅', '34@163.com2', 0, 0, '2022-09-11 10:59:49', '2022-09-11 10:59:52', '2022-09-11 10:59:55', '上海');

SET FOREIGN_KEY_CHECKS = 1;
