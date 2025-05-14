-- 比赛表 (competitions) - 记录比赛的基本信息和设置
CREATE TABLE `competitions` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL,
    `description` TEXT NOT NULL,
    `start_time` TIMESTAMP NULL COMMENT '比赛开始时间',
    `end_time` TIMESTAMP NULL COMMENT '比赛结束时间',
    `questions` TEXT NOT NULL COMMENT '题目列表，JSON格式',
    `type` ENUM('acm', 'oi') NOT NULL COMMENT '比赛类型：ACM赛制或OI赛制',
    `password` VARCHAR(255) DEFAULT NULL,
    `password_version` INT UNSIGNED NOT NULL DEFAULT 0,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) COMMENT = '比赛表';

-- 比赛成绩表 (competition_scores) - 记录每个用户在比赛中的得分情况
CREATE TABLE `competition_scores` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `competition_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `score_details` TEXT NOT NULL COMMENT '题目得分详情',
    `judge_ids` TEXT NOT NULL COMMENT '评测详情，记录所有评测ID，JSON格式：[judge_id1, judge_id2,...]',
    `total_score` FLOAT NOT NULL DEFAULT 0,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) COMMENT = '比赛成绩表';

-- 参赛记录表 (competition_attendances) - 仅记录参赛事实，不记录分数
CREATE TABLE `competition_attendances` (
    `competition_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `password_version` INT UNSIGNED DEFAULT NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`competition_id`, `user_id`)
) COMMENT = '参赛记录表';