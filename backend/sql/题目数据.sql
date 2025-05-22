INSERT INTO
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
VALUES (
        '回文数',
        '## 题目描述\n给你一个整数 `x`，如果 `x` 是一个回文整数，返回 `true`；否则，返回 `false`。\n\n回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。\n\n## 示例\n\n**示例 1:**\n\n输入：\n```\n121\n```\n输出：\n```\ntrue\n```\n\n**示例 2:**\n\n输入：\n```\n-121\n```\n输出：\n```\nfalse\n```\n解释：从左向右读, 为 $-121$ 。 从右向左读, 为 $121-$ 。因此它不是一个回文数。\n\n**示例 3:**\n\n输入：\n```\n10\n```\n输出：\n```\nfalse\n```\n解释：从右向左读, 为 $01$ 。因此它不是一个回文数。\n\n## 提示\n\n- $-2^{31} <= x <= 2^{31} - 1$\n\n## 进阶\n\n你能不将整数转为字符串来解决这个问题吗？',
        '["数学"、"字符串"]',
        'easy',
        0,
        0,
        '[{"input":"121","output":"true"},{"input":"-121","output":"false"},{"input":"10","output":"false"},{"input":"12321","output":"true"},{"input":"0","output":"true"},{"input":"12345","output":"false"},{"input":"1234321","output":"true"},{"input":"-101","output":"false"},{"input":"1001","output":"true"},{"input":"123456","output":"false"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );