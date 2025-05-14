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
        '["数学"]',
        'easy',
        0,
        0,
        '[{"input":"121","output":"true"},{"input":"-121","output":"false"},{"input":"10","output":"false"},{"input":"12321","output":"true"},{"input":"0","output":"true"},{"input":"12345","output":"false"},{"input":"1234321","output":"true"},{"input":"-101","output":"false"},{"input":"1001","output":"true"},{"input":"123456","output":"false"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );

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
        '两数之和',
        '## 题目描述\n编写一个程序，计算两个整数的和。\n\n## 输入格式\n两个整数 $a$ 和 $b$，以空格分隔\n\n## 输出格式\n一个整数，表示 $a + b$ 的结果\n\n## 示例\n**示例 1:**\n\n输入：\n```\n1 1\n```\n输出：\n```\n2\n```\n解释：$1 + 1 = 2$\n\n**示例 2:**\n\n输入：\n```\n5 3\n```\n输出：\n```\n8\n```\n\n**示例 3:**\n\n输入：\n```\n-2 7\n```\n输出：\n```\n5\n```\n\n## 数据范围\n$-1000 \leq a, b \leq 1000$',
        '["数学"]',
        'easy',
        0,
        0,
        '[{"input":"1 1","output":"2"},{"input":"5 3","output":"8"},{"input":"-2 7","output":"5"},{"input":"0 0","output":"0"},{"input":"100 200","output":"300"},{"input":"-5 -3","output":"-8"},{"input":"999 1","output":"1000"},{"input":"-1000 1000","output":"0"},{"input":"42 58","output":"100"},{"input":"123 456","output":"579"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );

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
        '数字的平方',
        '## 题目描述\n计算并返回一个整数的平方值。\n\n## 输入格式\n一个整数 $n$\n\n## 输出格式\n一个整数，表示 $n^2$ 的结果\n\n## 示例\n**示例 1:**\n\n输入：\n```\n3\n```\n输出：\n```\n9\n```\n解释：$3^2 = 9$\n\n**示例 2:**\n\n输入：\n```\n-4\n```\n输出：\n```\n16\n```\n\n**示例 3:**\n\n输入：\n```\n0\n```\n输出：\n```\n0\n```\n\n## 数据范围\n$-100 \leq n \leq 100$',
        '["数学","基础"]',
        'easy',
        0,
        0,
        '[{"input":"3","output":"9"},{"input":"-4","output":"16"},{"input":"0","output":"0"},{"input":"10","output":"100"},{"input":"-10","output":"100"},{"input":"15","output":"225"},{"input":"-7","output":"49"},{"input":"25","output":"625"},{"input":"-1","output":"1"},{"input":"99","output":"9801"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );

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
        '判断奇偶数',
        '## 题目描述\n给定一个整数，判断它是奇数还是偶数。如果是奇数返回"odd"，如果是偶数返回"even"。\n\n## 输入格式\n一个整数 $n$\n\n## 输出格式\n"odd"或"even"\n\n## 示例\n**示例 1:**\n\n输入：\n```\n4\n```\n输出：\n```\neven\n```\n\n**示例 2:**\n\n输入：\n```\n7\n```\n输出：\n```\nodd\n```\n\n**示例 3:**\n\n输入：\n```\n0\n```\n输出：\n```\neven\n```\n\n## 数据范围\n$-1000 \leq n \leq 1000$',
        '["数学","条件判断"]',
        'easy',
        0,
        0,
        '[{"input":"4","output":"even"},{"input":"7","output":"odd"},{"input":"0","output":"even"},{"input":"-3","output":"odd"},{"input":"-8","output":"even"},{"input":"123","output":"odd"},{"input":"456","output":"even"},{"input":"-1","output":"odd"},{"input":"1000","output":"even"},{"input":"999","output":"odd"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );

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
        '三个数的最大值',
        '## 题目描述\n给定三个整数，找出其中的最大值。\n\n## 输入格式\n三个整数 $a$, $b$, $c$，以空格分隔\n\n## 输出格式\n一个整数，表示三个数中的最大值\n\n## 示例\n**示例 1:**\n\n输入：\n```\n1 2 3\n```\n输出：\n```\n3\n```\n\n**示例 2:**\n\n输入：\n```\n-5 0 5\n```\n输出：\n```\n5\n```\n\n**示例 3:**\n\n输入：\n```\n10 10 5\n```\n输出：\n```\n10\n```\n\n## 数据范围\n$-1000 \leq a, b, c \leq 1000$',
        '["数学","比较"]',
        'easy',
        0,
        0,
        '[{"input":"1 2 3","output":"3"},{"input":"-5 0 5","output":"5"},{"input":"10 10 5","output":"10"},{"input":"0 0 0","output":"0"},{"input":"-1 -2 -3","output":"-1"},{"input":"100 200 150","output":"200"},{"input":"-100 -50 -75","output":"-50"},{"input":"999 998 997","output":"999"},{"input":"42 42 42","output":"42"},{"input":"-10 0 10","output":"10"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );

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
        '跳跃游戏',
        '## 题目描述\n给定一个非负整数数组 `nums`，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个位置。\n\n## 输入格式\n一行非负整数，表示数组 `nums`，用空格分隔\n\n## 输出格式\n"true" 或 "false"\n\n## 示例\n**示例 1:**\n\n输入：\n```\n2 3 1 1 4\n```\n输出：\n```\ntrue\n```\n解释：可以先跳 1 步从位置 0 到位置 1，然后再跳 3 步到达最后一个位置。\n\n**示例 2:**\n\n输入：\n```\n3 2 1 0 4\n```\n输出：\n```\nfalse\n```\n解释：无论如何都会到达位置 3，但该位置的最大跳跃长度是 0，所以永远无法到达最后一个位置。\n\n## 数据范围\n- $1 \leq nums.length \leq 3 \times 10^4$\n- $0 \leq nums[i] \leq 10^5$',
        '["贪心"]',
        'medium',
        0,
        0,
        '[{"input":"2 3 1 1 4","output":"true"},{"input":"3 2 1 0 4","output":"false"},{"input":"0","output":"true"},{"input":"1 0 1","output":"false"},{"input":"1 1 1 0","output":"true"},{"input":"5 0 0 0 0 0 1","output":"true"},{"input":"1 2 3 0 0 1","output":"false"},{"input":"4 2 0 0 1 1 4","output":"true"},{"input":"1 0 2 0 1","output":"false"},{"input":"2 0 2 0 1 0 0 3","output":"false"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );

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
        '组合总和',
        '## 题目描述\n给定一个无重复元素的整数数组 `candidates` 和一个目标整数 `target`，找出 `candidates` 中所有可以使数字和为 `target` 的唯一组合。\n\n`candidates` 中的数字可以无限制重复被选取。\n\n## 输入格式\n第一行：整数数组 `candidates`，用空格分隔\n第二行：目标整数 `target`\n\n## 输出格式\n所有可能的组合，每行一个组合，数字用空格分隔，按字典序排列\n\n## 示例\n**示例 1:**\n\n输入：\n```\n2 3 6 7\n7\n```\n输出：\n```\n2 2 3\n7\n```\n\n**示例 2:**\n\n输入：\n```\n2 3 5\n8\n```\n输出：\n```\n2 2 2 2\n2 3 3\n3 5\n```\n\n## 数据范围\n- $1 \leq candidates.length \leq 30$\n- $1 \leq candidates[i] \leq 200$\n- $1 \leq target \leq 500$',
        '["回溯"]',
        'medium',
        0,
        0,
        '[{"input":"2 3 6 7\n7","output":"2 2 3\n7"},{"input":"2 3 5\n8","output":"2 2 2 2\n2 3 3\n3 5"},{"input":"2\n1","output":""},{"input":"3 5 8\n11","output":"3 3 5\n3 8"},{"input":"2 4 6 8\n10","output":"2 2 2 2 2\n2 2 2 4\n2 2 6\n2 4 4\n2 8\n4 6"},{"input":"5 10 15\n20","output":"5 5 5 5\n5 5 10\n5 15\n10 10"},{"input":"1\n2","output":"1 1"},{"input":"2 3\n5","output":"2 3"},{"input":"1 2 3\n4","output":"1 1 1 1\n1 1 2\n1 3\n2 2"},{"input":"3 4 5\n9","output":"3 3 3\n4 5"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );

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
        '最长递增子序列',
        '## 题目描述\n给定一个整数数组 `nums`，找到其中最长严格递增子序列的长度。\n\n子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。\n\n## 输入格式\n一行整数，表示数组 `nums`，用空格分隔\n\n## 输出格式\n一个整数，表示最长递增子序列的长度\n\n## 示例\n**示例 1:**\n\n输入：\n```\n10 9 2 5 3 7 101 18\n```\n输出：\n```\n4\n```\n解释：最长递增子序列是 [2,3,7,101]，因此长度为 4。\n\n**示例 2:**\n\n输入：\n```\n0 1 0 3 2 3\n```\n输出：\n```\n4\n```\n\n**示例 3:**\n\n输入：\n```\n7 7 7 7 7 7 7\n```\n输出：\n```\n1\n```\n\n## 数据范围\n- $1 \leq nums.length \leq 2500$\n- $-10^4 \leq nums[i] \leq 10^4$',
        '["动态规划"]',
        'medium',
        0,
        0,
        '[{"input":"10 9 2 5 3 7 101 18","output":"4"},{"input":"0 1 0 3 2 3","output":"4"},{"input":"7 7 7 7 7 7 7","output":"1"},{"input":"1","output":"1"},{"input":"1 3 6 7 9 4 10 5 6","output":"6"},{"input":"5 8 3 7 9 2 1 6","output":"4"},{"input":"10 22 9 33 21 50 41 60 80","output":"6"},{"input":"1 2 3 4 5 6 7 8 9 10","output":"10"},{"input":"10 9 8 7 6 5 4 3 2 1","output":"1"},{"input":"3 5 6 2 5 4 7 3 6","output":"4"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );

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
        '二叉树的层序遍历',
        '## 题目描述\n给定一个二叉树的根节点 `root`，返回其节点值的层序遍历结果。（即逐层地，从左到右访问所有节点）。\n\n二叉树使用以下格式表示：\n- 空节点用 `null` 表示\n- 非空节点用 `(val left right)` 表示\n- 例如：`(3 (9 null null) (20 (15 null null) (7 null null))`\n\n## 输入格式\n一个字符串，表示二叉树的序列化形式\n\n## 输出格式\n多行输出，每行表示一层的节点值，节点值用空格分隔\n\n## 示例\n**示例 1:**\n\n输入：\n```\n(3 (9 null null) (20 (15 null null) (7 null null))\n```\n输出：\n```\n3\n9 20\n15 7\n```\n\n**示例 2:**\n\n输入：\n```\n1\n```\n输出：\n```\n1\n```\n\n**示例 3:**\n\n输入：\n```\nnull\n```\n输出：\n```\n(空输出)\n```\n\n## 数据范围\n- 节点数范围 $[0, 2000]$\n- $-1000 \leq Node.val \leq 1000$',
        '["广度优先搜索","二叉树"]',
        'medium',
        0,
        0,
        '[{"input":"(3 (9 null null) (20 (15 null null) (7 null null))","output":"3\n9 20\n15 7"},{"input":"1","output":"1"},{"input":"null","output":""},{"input":"(1 (2 (4 null null) (5 null null)) (3 null (6 null null))","output":"1\n2 3\n4 5 6"},{"input":"(A (B null null) (C null null))","output":"A\nB C"},{"input":"(1 null (2 null (3 null (4 null null))))","output":"1\n2\n3\n4"},{"input":"(1 (2 (3 (4 null null) null) null) null)","output":"1\n2\n3\n4"},{"input":"(1 (2 null null) null)","output":"1\n2"},{"input":"(1 null (2 null (3 null (4 null (5 null null)))))","output":"1\n2\n3\n4\n5"},{"input":"(1 (2 (4 null null) null) (3 (5 null null) (6 null null)))","output":"1\n2 3\n4 5 6"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );

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
        '岛屿数量',
        '## 题目描述\n给你一个由 \'1\'（陆地）和 \'0\'（水）组成的的二维网格，请你计算网格中岛屿的数量。\n\n岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。\n\n## 输入格式\n多行输入，每行表示网格的一行，用空格分隔的字符\n\n## 输出格式\n一个整数，表示岛屿数量\n\n## 示例\n**示例 1:**\n输入：\n```\n1 1 1 1 0\n1 1 0 1 0\n1 1 0 0 0\n0 0 0 0 0\n```\n输出：\n```\n1\n```\n解释：图中所有相连的1组成一个岛屿\n\n**示例 2:**\n输入：\n```\n1 1 0 0 0\n1 1 0 0 0\n0 0 1 0 0\n0 0 0 1 1\n```\n输出：\n```\n3\n```\n解释：左上角、中间和右下角各有一个岛屿\n\n## 数据范围\n- 1 ≤ 网格行数 ≤ 300\n- 1 ≤ 网格列数 ≤ 300\n- 网格元素只包含 \'0\' 或 \'1\'',
        '["深度优先搜索","矩阵"]',
        'medium',
        0,
        0,
        '[{"input":"1 1 1 1 0\\n1 1 0 1 0\\n1 1 0 0 0\\n0 0 0 0 0","output":"1"},{"input":"1 1 0 0 0\\n1 1 0 0 0\\n0 0 1 0 0\\n0 0 0 1 1","output":"3"},{"input":"1 0 1 0 1","output":"3"},{"input":"1 1 1\\n0 0 0\\n1 1 1","output":"2"},{"input":"1","output":"1"},{"input":"0","output":"0"},{"input":"1 0\\n0 1","output":"2"},{"input":"1 1 0\\n1 0 1\\n0 1 1","output":"1"},{"input":"1 1 1 0 0\\n1 0 1 0 0\\n1 1 1 0 0\\n0 0 0 0 0","output":"1"},{"input":"1 0 1 0 1\\n0 1 0 1 0\\n1 0 1 0 1\\n0 1 0 1 0","output":"10"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 1000}',
        'public'
    );

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
        '正则表达式匹配',
        '## 题目描述\n给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 \'.\' 和 \'*\' 的正则表达式匹配。\n\n- \'.\' 匹配任意单个字符\n- \'*\' 匹配零个或多个前面的那一个元素\n\n匹配应当覆盖整个字符串 s，而不是部分字符串。\n\n## 输入格式\n两行输入：\n第一行：字符串 s\n第二行：模式 p\n\n## 输出格式\n\"true\" 或 \"false\"\n\n## 示例\n**示例 1:**\n输入：\n```\naa\na\n```\n输出：\n```\nfalse\n```\n\n**示例 2:**\n输入：\n```\naa\na*\n```\n输出：\n```\ntrue\n```\n\n**示例 3:**\n输入：\n```\nab\n.*\n```\n输出：\n```\ntrue\n```\n\n## 数据范围\n- 1 <= s.length <= 20\n- 1 <= p.length <= 30\n- s 只包含小写字母\n- p 只包含小写字母、\'.\' 和 \'*\'\n- 保证每次出现字符 * 时，前面都匹配到有效的字符',
        '["动态规划","字符串"]',
        'hard',
        0,
        0,
        '[{"input":"aa\\na","output":"false"},{"input":"aa\\na*","output":"true"},{"input":"ab\\n.*","output":"true"},{"input":"aab\\nc*a*b","output":"true"},{"input":"mississippi\\nmis*is*p*.","output":"false"},{"input":"ab\\n.*c","output":"false"},{"input":"aaa\\naaaa","output":"false"},{"input":"aaa\\na*a","output":"true"},{"input":"abc\\na***abc","output":"true"},{"input":"\\n.*","output":"true"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 3000}',
        'public'
    );

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
        '滑动窗口最大值',
        '## 题目描述\n给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。\n\n返回滑动窗口中的最大值。\n\n## 输入格式\n第一行：整数数组 nums，空格分隔\n第二行：整数 k\n\n## 输出格式\n滑动窗口最大值数组，空格分隔\n\n## 示例\n**示例 1:**\n输入：\n```\n1 3 -1 -3 5 3 6 7\n3\n```\n输出：\n```\n3 3 5 5 6 7\n```\n\n**示例 2:**\n输入：\n```\n1\n1\n```\n输出：\n```\n1\n```\n\n## 数据范围\n- 1 <= nums.length <= 10^5\n- -10^4 <= nums[i] <= 10^4\n- 1 <= k <= nums.length',
        '["队列","数组","滑动窗口"]',
        'hard',
        0,
        0,
        '[{"input":"1 3 -1 -3 5 3 6 7\\n3","output":"3 3 5 5 6 7"},{"input":"1\\n1","output":"1"},{"input":"1 -1\\n1","output":"1 -1"},{"input":"9 11\\n2","output":"11"},{"input":"4 -2\\n2","output":"4"},{"input":"7 2 4\\n2","output":"7 4"},{"input":"1 3 1 2 0 5\\n3","output":"3 3 2 5"},{"input":"1 2 3 4 5 6 7 8 9 10\\n3","output":"3 4 5 6 7 8 9 10"},{"input":"10 9 8 7 6 5 4 3 2 1\\n4","output":"10 9 8 7 6 5 4"},{"input":"1 4 2 5 3 6 4 7 5 8\\n5","output":"5 5 6 6 7 7 8"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 3000}',
        'public'
    );

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
        '合并K个升序链表',
        '## 题目描述\n给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，返回合并后的链表。\n\n链表表示为：val next_val next_val ... null\n多个链表用分号分隔\n\n## 输入格式\n一行输入，多个链表用分号分隔，例如：\n1 2 4 null;1 3 4 null;2 6 null\n\n## 输出格式\n合并后的链表，格式与输入相同\n\n## 示例\n**示例 1:**\n输入：\n```\n1 4 5 null;1 3 4 null;2 6 null\n```\n输出：\n```\n1 1 2 3 4 4 5 6 null\n```\n\n**示例 2:**\n输入：\n```\nnull\n```\n输出：\n```\nnull\n```\n\n## 数据范围\n- k == lists.length\n- 0 <= k <= 10^4\n- 0 <= 单个链表长度 <= 500\n- -10^4 <= 节点值 <= 10^4\n- 保证所有链表都是升序排列的',
        '["分治","堆","链表"]',
        'hard',
        0,
        0,
        '[{"input":"1 4 5 null;1 3 4 null;2 6 null","output":"1 1 2 3 4 4 5 6 null"},{"input":"null","output":"null"},{"input":"1 2 3 null;4 5 6 null","output":"1 2 3 4 5 6 null"},{"input":"1 null;2 null;3 null","output":"1 2 3 null"},{"input":"-1 5 null;2 4 null;null","output":"-1 2 4 5 null"},{"input":"1 2 3 null;1 2 3 null;1 2 3 null","output":"1 1 1 2 2 2 3 3 3 null"},{"input":"1 10 20 null;2 11 21 null;3 12 22 null","output":"1 2 3 10 11 12 20 21 22 null"},{"input":"null;null;null","output":"null"},{"input":"1 2 null;3 4 null;5 6 null","output":"1 2 3 4 5 6 null"},{"input":"1 100 null;2 200 null;3 300 null","output":"1 2 3 100 200 300 null"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 3000}',
        'public'
    );

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
        '二叉树中的最大路径和',
        '## 题目描述\n路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中至多出现一次。该路径至少包含一个节点，且不一定经过根节点。\n\n路径和是路径中各节点值的总和。\n\n给你一个二叉树的根节点 root，返回其最大路径和。\n\n二叉树表示为：(val left right)，空节点为null\n\n## 输入格式\n一个字符串，表示二叉树的序列化形式\n\n## 输出格式\n一个整数，表示最大路径和\n\n## 示例\n**示例 1:**\n输入：\n```\n(1 (2 null null) (3 null null))\n```\n输出：\n```\n6\n```\n\n**示例 2:**\n输入：\n```\n(-10 (9 null null) (20 (15 null null) (7 null null)))\n```\n输出：\n```\n42\n```\n\n## 数据范围\n- 节点数范围 [1, 3 * 10^4]\n- -1000 <= Node.val <= 1000',
        '["深度优先搜索","二叉树"]',
        'hard',
        0,
        0,
        '[{"input":"(1 (2 null null) (3 null null))","output":"6"},{"input":"(-10 (9 null null) (20 (15 null null) (7 null null)))","output":"42"},{"input":"(1 null null)","output":"1"},{"input":"(-3 null null)","output":"-3"},{"input":"(1 (-2 null null) (3 null null))","output":"4"},{"input":"(5 (4 (11 (7 null null) (2 null null)) null) (8 (13 null null) (4 null (1 null null))))","output":"48"},{"input":"(1 (2 (3 (4 (5 null null) null) null) null)","output":"15"},{"input":"(-1 (-2 null null) (-3 null null))","output":"-1"},{"input":"(0 (-1 null null) null)","output":"0"},{"input":"(10 (9 null null) (20 (15 (30 null null) (7 null null)))","output":"72"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 3000}',
        'public'
    );

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
        '最小覆盖子串',
        '## 题目描述\n给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 \"\"。\n\n注意：\n- 对于 t 中重复字符，子串中该字符数量必须不少于 t 中该字符数量\n- s 和 t 由英文字母组成\n\n## 输入格式\n两行输入：\n第一行：字符串 s\n第二行：字符串 t\n\n## 输出格式\n满足条件的最小子串\n\n## 示例\n**示例 1:**\n输入：\n```\nADOBECODEBANC\nABC\n```\n输出：\n```\nBANC\n```\n\n**示例 2:**\n输入：\n```\na\na\n```\n输出：\n```\na\n```\n\n**示例 3:**\n输入：\n```\na\naa\n```\n输出：\n```\n\n```\n\n## 数据范围\n- 1 <= s.length, t.length <= 10^5\n- s 和 t 由英文字母组成',
        '["哈希表","字符串","滑动窗口"]',
        'hard',
        0,
        0,
        '[{"input":"ADOBECODEBANC\\nABC","output":"BANC"},{"input":"a\\na","output":"a"},{"input":"a\\naa","output":""},{"input":"ab\\na","output":"a"},{"input":"aa\\naa","output":"aa"},{"input":"abc\\nb","output":"b"},{"input":"abcdefghijklmnopqrstuvwxyz\\nxyz","output":"xyz"},{"input":"this is a test string\\ntist","output":"t stri"},{"input":"xyz\\nxyz","output":"xyz"},{"input":"aaaaaaaaaaaabbbbbcdd\\nabcdd","output":"abbbbbcdd"}]',
        '{"memory_limit_mib": 512, "stack_limit_mib": 64, "timeout_millisecond": 3000}',
        'public'
    );