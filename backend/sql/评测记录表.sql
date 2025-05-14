CREATE TABLE `judges` (
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
    ) COMMENT '语言',
    `code` TEXT NOT NULL COMMENT '代码',
    `exec_result` TEXT NOT NULL COMMENT '执行结果',
    `accepted` BOOLEAN NOT NULL COMMENT '是否通过',
    `judge_type` ENUM('normal', 'competition') NOT NULL COMMENT '评测类型',
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id)
) COMMENT = '评测记录表';