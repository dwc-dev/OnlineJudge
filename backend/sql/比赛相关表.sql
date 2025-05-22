-- 比赛表 (competitions) - 记录比赛的基本信息和设置
CREATE TABLE `competitions` (
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
) COMMENT = '比赛表';

-- 比赛成绩表 (competition_scores) - 记录每个用户在比赛中的得分情况
CREATE TABLE `competition_scores` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `competition_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `score_details` TEXT NOT NULL,
    `judge_ids` TEXT NOT NULL,
    `total_score` FLOAT NOT NULL DEFAULT 0,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) COMMENT = '比赛成绩表';

-- 参赛记录表 (competition_attendances) - 仅记录参赛事实，不记录分数
CREATE TABLE `competition_attendances` (
    `competition_id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `password_version` INT UNSIGNED NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`competition_id`, `user_id`)
) COMMENT = '参赛记录表';