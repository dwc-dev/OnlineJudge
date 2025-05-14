CREATE TABLE `questions` (
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
) COMMENT = '题目表';