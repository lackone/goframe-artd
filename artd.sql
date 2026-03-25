/*
 Navicat Premium Dump SQL

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80029 (8.0.29)
 Source Host           : 127.0.0.1:3306
 Source Schema         : artd

 Target Server Type    : MySQL
 Target Server Version : 80029 (8.0.29)
 File Encoding         : 65001

 Date: 23/03/2026 15:41:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_admin_auths
-- ----------------------------
DROP TABLE IF EXISTS `tb_admin_auths`;
CREATE TABLE `tb_admin_auths`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '类型(1:菜单,2:按钮)',
  `is_enable` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用(-1:否,1:是)',
  `keepalive` tinyint(1) NOT NULL DEFAULT 1 COMMENT '页面缓存(-1:否,1:是)',
  `is_hide` tinyint(1) NOT NULL DEFAULT -1 COMMENT '隐藏菜单(-1:否,1:是)',
  `is_hide_tab` tinyint(1) NOT NULL DEFAULT -1 COMMENT '标签隐藏(-1:否,1:是)',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '外部链接',
  `is_iframe` tinyint(1) NOT NULL DEFAULT -1 COMMENT '是否内嵌(-1:否,1:是)',
  `show_badge` tinyint(1) NOT NULL DEFAULT -1 COMMENT '显示徽章(-1:否,1:是)',
  `show_text_badge` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '文本徽章',
  `fixed_tab` tinyint(1) NOT NULL DEFAULT -1 COMMENT '固定标签(-1:否,1:是)',
  `active_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '激活路径',
  `is_full_page` tinyint(1) NOT NULL DEFAULT -1 COMMENT '全屏页面(-1:否,1:是)',
  `pid` int NOT NULL DEFAULT 0 COMMENT '父级',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标',
  `sort` smallint NOT NULL DEFAULT 0 COMMENT '排序(越小越前)',
  `created` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted` int NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '后台权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_admin_auths
-- ----------------------------
INSERT INTO `tb_admin_auths` VALUES (1, 'System', '系统管理', '/system', '/index/index', 1, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 0, 'ri:user-3-line', 100, 1774016335, 1774020547, NULL);
INSERT INTO `tb_admin_auths` VALUES (2, 'User', '用户管理', 'user', '/system/user', 1, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 1, 'ri:user-line', 0, 1774016335, 1774016335, NULL);
INSERT INTO `tb_admin_auths` VALUES (3, 'Role', '角色管理', 'role', '/system/role', 1, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 1, 'ri:user-settings-line', 0, 1774016335, 1774016335, NULL);
INSERT INTO `tb_admin_auths` VALUES (4, 'Menus', '菜单管理', 'menu', '/system/menu', 1, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 1, 'ri:menu-line', 0, 1774016335, 1774016335, NULL);
INSERT INTO `tb_admin_auths` VALUES (5, 'add', '用户新增', '', '', 2, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 2, '', 1, 1774016335, 1774016335, NULL);
INSERT INTO `tb_admin_auths` VALUES (6, 'edit', '用户编辑', '', '', 2, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 2, '', 1, 1774016335, 1774016335, NULL);
INSERT INTO `tb_admin_auths` VALUES (7, 'delete', '用户删除', '', '', 2, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 2, '', 1, 1774016335, 1774016335, NULL);
INSERT INTO `tb_admin_auths` VALUES (8, 'add', '角色新增', '', '', 2, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 3, '', 1, 1774016335, 1774016335, NULL);
INSERT INTO `tb_admin_auths` VALUES (9, 'edit', '角色编辑', '', '', 2, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 3, '', 1, 1774016335, 1774016335, NULL);
INSERT INTO `tb_admin_auths` VALUES (10, 'delete', '角色删除', '', '', 2, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 3, '', 1, 1774016335, 1774016335, NULL);
INSERT INTO `tb_admin_auths` VALUES (11, 'add', '菜单新增', '', '', 2, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 4, '', 1, 1774016764, 1774016764, NULL);
INSERT INTO `tb_admin_auths` VALUES (12, 'edit', '菜单编辑', '', '', 2, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 4, '', 1, 1774016774, 1774016774, NULL);
INSERT INTO `tb_admin_auths` VALUES (13, 'delete', '菜单删除', '', '', 2, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 4, '', 1, 1774016780, 1774016780, NULL);
INSERT INTO `tb_admin_auths` VALUES (14, 'Dashboard', '仪表盘', '/dashboard', '/index/index', 1, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 0, 'ri:pie-chart-line', 1, 1774020328, 1774020328, NULL);
INSERT INTO `tb_admin_auths` VALUES (15, 'Console', '工作台', 'console', '/dashboard/console', 1, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 14, 'ri:home-smile-2-line', 1, 1774020367, 1774020367, NULL);
INSERT INTO `tb_admin_auths` VALUES (16, 'setRole', '设置角色', '', '', 2, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 2, '', 1, 1774022659, 1774022659, NULL);
INSERT INTO `tb_admin_auths` VALUES (17, 'UserCenter', '个人中心', 'user-center', '/system/user-center', 1, 1, -1, 1, 1, '', -1, -1, '', -1, '', -1, 1, 'ri:user-line', 1, 1774025001, 1774025001, NULL);
INSERT INTO `tb_admin_auths` VALUES (18, 'SystemConfig', '系统设置', 'config', '', 1, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 1, 'ri:function-line', 1, 1774083647, 1774083647, NULL);
INSERT INTO `tb_admin_auths` VALUES (19, 'Website', '网站设置', 'website', '/system/config/website', 1, 1, -1, -1, -1, '', -1, -1, '', -1, '', -1, 18, 'ri:pages-line', 1, 1774083837, 1774085899, NULL);

-- ----------------------------
-- Table structure for tb_admin_role_assocs
-- ----------------------------
DROP TABLE IF EXISTS `tb_admin_role_assocs`;
CREATE TABLE `tb_admin_role_assocs`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `admin_id` int NOT NULL DEFAULT 0 COMMENT '用户ID',
  `role_id` int NOT NULL DEFAULT 0 COMMENT '角色ID',
  `created` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted` int NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniq_role_admin`(`role_id` ASC, `admin_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '后台用户角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_admin_role_assocs
-- ----------------------------

-- ----------------------------
-- Table structure for tb_admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `tb_admin_roles`;
CREATE TABLE `tb_admin_roles`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `auth_ids` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '权限ID(逗号分割)',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态(-1:禁用,1:启用)',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted` int NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '后台用户角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_admin_roles
-- ----------------------------
INSERT INTO `tb_admin_roles` VALUES (1, '管理员', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19', 1, '管理员', 1774250421, 1774250637, NULL);

-- ----------------------------
-- Table structure for tb_admins
-- ----------------------------
DROP TABLE IF EXISTS `tb_admins`;
CREATE TABLE `tb_admins`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `account` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '账号',
  `nick_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `real_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `sex` tinyint(1) NOT NULL DEFAULT 0 COMMENT '性别(0:未知,1:男,2:女)',
  `salt` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '加密盐',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码(md5(md5(salt) . password))',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `tel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '座机',
  `email` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `weixin` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '微信',
  `last_login_ip` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` int NOT NULL DEFAULT 0 COMMENT '最后登录时间',
  `is_super` tinyint(1) NOT NULL DEFAULT -1 COMMENT '超级管理员(-1:否,1:是)',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态(-1:禁用,1:启用)',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '地址',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted` int NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniq_account`(`account` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '后台用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_admins
-- ----------------------------
INSERT INTO `tb_admins` VALUES (1, 'admin', 'admin', 'admin', '', 0, 'Opxcw6', '89f9685671b6d878998ca09de7878397', '1', '1', '1', '1', '127.0.0.1', 1773123780, 1, 1, '1', '1', 1689660985, 1774027357, NULL);

-- ----------------------------
-- Table structure for tb_configs
-- ----------------------------
DROP TABLE IF EXISTS `tb_configs`;
CREATE TABLE `tb_configs`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '中文名',
  `key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '键',
  `sub_key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '子键',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '内容',
  `pid` int NOT NULL DEFAULT 0 COMMENT '父级',
  `sort` smallint NOT NULL DEFAULT 0 COMMENT '排序(越小越前)',
  `admin_id` int NOT NULL DEFAULT 0 COMMENT '操作人',
  `created` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted` int NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniq_key_sub`(`key` ASC, `sub_key` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_configs
-- ----------------------------
INSERT INTO `tb_configs` VALUES (1, '', 'website', 'website_name', NULL, 0, 0, 0, 1774089896, 1774250285, NULL);
INSERT INTO `tb_configs` VALUES (2, '', 'website', 'website_domain', NULL, 0, 0, 0, 1774089896, 1774250285, NULL);
INSERT INTO `tb_configs` VALUES (3, '', 'website', 'website_title', NULL, 0, 0, 0, 1774089896, 1774250285, NULL);
INSERT INTO `tb_configs` VALUES (4, '', 'website', 'website_keywords', NULL, 0, 0, 0, 1774089896, 1774250285, NULL);
INSERT INTO `tb_configs` VALUES (5, '', 'website', 'website_description', NULL, 0, 0, 0, 1774089896, 1774250285, NULL);
INSERT INTO `tb_configs` VALUES (6, '', 'website', 'website_logo', NULL, 0, 0, 0, 1774089896, 1774250285, NULL);
INSERT INTO `tb_configs` VALUES (7, '', 'website', 'website_favicon', NULL, 0, 0, 0, 1774089896, 1774250285, NULL);
INSERT INTO `tb_configs` VALUES (8, '', 'filing', 'filing_number', NULL, 0, 0, 0, 1774096120, 1774250347, NULL);
INSERT INTO `tb_configs` VALUES (9, '', 'filing', 'filing_license', NULL, 0, 0, 0, 1774096120, 1774250347, NULL);
INSERT INTO `tb_configs` VALUES (10, '', 'filing', 'filing_subject', NULL, 0, 0, 0, 1774096120, 1774250347, NULL);
INSERT INTO `tb_configs` VALUES (11, '', 'filing', 'filing_link', NULL, 0, 0, 0, 1774096120, 1774250347, NULL);
INSERT INTO `tb_configs` VALUES (12, '', 'filing', 'filing_remark', NULL, 0, 0, 0, 1774096120, 1774250347, NULL);
INSERT INTO `tb_configs` VALUES (13, '', 'policy', 'user_agreement', '<p><br></p>', 0, 0, 0, 1774096193, 1774250137, NULL);
INSERT INTO `tb_configs` VALUES (14, '', 'policy', 'privacy_policy', '<p><br></p>', 0, 0, 0, 1774096193, 1774250137, NULL);
INSERT INTO `tb_configs` VALUES (15, '', 'policy', 'service_terms', '<p><br></p>', 0, 0, 0, 1774096193, 1774250137, NULL);
INSERT INTO `tb_configs` VALUES (16, '', 'policy', 'disclaimer', '<p><br></p>', 0, 0, 0, 1774096193, 1774250137, NULL);
INSERT INTO `tb_configs` VALUES (17, '', 'policy', 'copyright_notice', '<p><br></p>', 0, 0, 0, 1774096193, 1774250137, NULL);
INSERT INTO `tb_configs` VALUES (18, '', 'policy', 'cookie_policy', '<p><br></p>', 0, 0, 0, 1774096193, 1774250137, NULL);

SET FOREIGN_KEY_CHECKS = 1;
