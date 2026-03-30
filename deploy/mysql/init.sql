SET NAMES utf8mb4;

CREATE DATABASE IF NOT EXISTS `online_judge` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `online_judge`;

GRANT ALL PRIVILEGES ON `online_judge`.* TO 'oj' @'%';

-- 比赛表 (competitions) - 记录比赛的基本信息和设置
CREATE TABLE IF NOT EXISTS `competitions` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL,
    `description` TEXT NOT NULL,
    `start_time` TIMESTAMP NOT NULL,
    `end_time` TIMESTAMP NOT NULL,
    `questions` TEXT NOT NULL,
    `type` ENUM('acm', 'oi') NOT NULL,
    `password` VARCHAR(255) NULL,
    `password_version` INT UNSIGNED NOT NULL DEFAULT 0,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 比赛成绩表 (competition_scores) - 记录每个用户在比赛中的得分情况
CREATE TABLE IF NOT EXISTS `competition_scores` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `competition_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `score_details` TEXT NOT NULL,
    `judge_ids` TEXT NOT NULL,
    `total_score` FLOAT NOT NULL DEFAULT 0,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 参赛记录表 (competition_attendances) - 仅记录参赛事实，不记录分数
CREATE TABLE IF NOT EXISTS `competition_attendances` (
    `competition_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `password_version` INT UNSIGNED NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`competition_id`, `user_id`)
);

-- 对话会话表 (sessions) - 记录每次对话会话的基本信息
CREATE TABLE IF NOT EXISTS `sessions` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `session_id` VARCHAR(64) NOT NULL UNIQUE,
    `title` VARCHAR(255) NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `question_id` BIGINT UNSIGNED NOT NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 对话消息表 (messages) - 记录每条对话消息的详细内容
CREATE TABLE IF NOT EXISTS `messages` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `session_id` VARCHAR(64) NOT NULL,
    `round` BIGINT UNSIGNED NOT NULL,
    `message_role` ENUM('system', 'assistant', 'user') NOT NULL,
    `content` TEXT NOT NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 评测表 (judges) - 记录每次代码评测的详细信息
CREATE TABLE IF NOT EXISTS `judges` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评测ID',
    `question_id` BIGINT UNSIGNED NOT NULL COMMENT '题目ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `language` ENUM(
        'c',
        'cpp',
        'java',
        'python',
        'golang',
        'rust'
    ) NOT NULL COMMENT '语言',
    `code` TEXT NOT NULL COMMENT '代码',
    `exec_result` TEXT NOT NULL COMMENT '执行结果',
    `accepted` BOOLEAN NOT NULL COMMENT '是否通过',
    `judge_type` ENUM('normal', 'competition') NOT NULL COMMENT '评测类型',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id)
);

-- 题目表 (questions) - 记录每道题目的详细信息
CREATE TABLE IF NOT EXISTS `questions` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `title` VARCHAR(512) NOT NULL COMMENT '标题',
    `content` TEXT NOT NULL COMMENT '内容（markdown格式）',
    `tags` TEXT NOT NULL COMMENT '标签',
    `difficulty` ENUM('easy', 'medium', 'hard') NOT NULL COMMENT '难度',
    `submit_num` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '题目提交数',
    `accepted_num` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '题目通过数',
    `judge_case` TEXT NOT NULL COMMENT '判题用例（json格式）',
    `judge_config` TEXT NOT NULL COMMENT '判题配置（json格式）',
    `visible_scope` ENUM('public', 'competition_only') NOT NULL DEFAULT 'public' COMMENT '题目可见范围',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id)
);

-- 用户表 (users) - 记录用户的基本信息和权限角色
CREATE TABLE IF NOT EXISTS `users` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `name` VARCHAR(50) NOT NULL UNIQUE COMMENT '用户昵称',
    `email` VARCHAR(255) NOT NULL UNIQUE COMMENT '用户邮箱',
    `password` VARCHAR(255) NOT NULL COMMENT '密码',
    `avatar_url` VARCHAR(1024) NOT NULL COMMENT '用户头像URL',
    `profile` TEXT NOT NULL COMMENT '个人简介',
    `role` ENUM('user', 'admin') NOT NULL DEFAULT 'user' COMMENT '用户角色：user/admin',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) AUTO_INCREMENT = 100000;

-- 预插入一条题目数据
INSERT IGNORE INTO
    `questions` (
        `title`,
        `content`,
        `tags`,
        `difficulty`,
        `submit_num`,
        `accepted_num`,
        `judge_case`,
        `judge_config`,
        `visible_scope`
    )
SELECT '回文数', '## 题目描述\n给你一个整数 `x`，如果 `x` 是一个回文整数，返回 `true`；否则，返回 `false`。\n\n回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。\n\n## 示例\n\n**示例 1:**\n\n输入：\n```\n121\n```\n输出：\n```\ntrue\n```\n\n**示例 2:**\n\n输入：\n```\n-121\n```\n输出：\n```\nfalse\n```\n解释：从左向右读, 为 $-121$ 。 从右向左读, 为 $121-$ 。因此它不是一个回文数。\n\n**示例 3:**\n\n输入：\n```\n10\n```\n输出：\n```\nfalse\n```\n解释：从右向左读, 为 $01$ 。因此它不是一个回文数。\n\n## 提示\n\n- $-2^{31} <= x <= 2^{31} - 1$\n\n## 进阶\n\n你能不将整数转为字符串来解决这个问题吗？', '["数学","字符串"]', 'easy', 0, 0, '[{"input":"121","output":"true"},{"input":"-121","output":"false"},{"input":"10","output":"false"},{"input":"12321","output":"true"},{"input":"0","output":"true"},{"input":"12345","output":"false"},{"input":"1234321","output":"true"},{"input":"-101","output":"false"},{"input":"1001","output":"true"},{"input":"123456","output":"false"}]', '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}', 'public'
WHERE
    NOT EXISTS (
        SELECT 1
        FROM `questions`
        LIMIT 1
    );

-- 预插入管理员用户数据，密码为 12345678，使用 bcrypt 加密后的哈希值
INSERT IGNORE INTO
    `users` (
        `name`,
        `email`,
        `password`,
        `avatar_url`,
        `profile`,
        `role`
    )
SELECT 'admin', 'admin@example.com', '$2a$10$QlED1QdJ8o1JT8UdOEghC.3ChEjKeKdjTKrRWT3KYPEGZhFYX9Kw.', 'http://localhost/online-judge/avatar/default_avatar.jpg', '这个人很懒，什么都没有留下', 'admin'
WHERE
    NOT EXISTS (
        SELECT 1
        FROM `users`
        WHERE `role` = 'admin'
        LIMIT 1
    );

-- 预插入测试用户数据，密码为 12345678，使用 bcrypt 加密后的哈希值
INSERT IGNORE INTO
    `users` (
        `name`,
        `email`,
        `password`,
        `avatar_url`,
        `profile`,
        `role`
    )
SELECT 'test', 'test@example.com', '$2a$10$cycSpK91fEE3SyO0UGScPeoaXYQkcn9c9Po9T82d796olPUrjg8Nu', 'http://localhost/online-judge/avatar/default_avatar.jpg', '这个人很懒，什么都没有留下', 'user'
WHERE
    NOT EXISTS (
        SELECT 1
        FROM `users`
        WHERE `role` = 'user'
        LIMIT 1
    );