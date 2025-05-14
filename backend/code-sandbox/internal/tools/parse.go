package tools

import (
	"bufio"
	"fmt"
	"strings"
)

// Parse 提取错误输出流中的最大内存使用和运行时间
func Parse(output string) (memoryUsage int64, elapsedTime int64) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Maximum resident set size (kbytes):") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				var kb int64
				fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &kb)
				memoryUsage = kb
			}
		} else if strings.Contains(line, "Time elapsed(ms):") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				var ms int64
				fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &ms)
				elapsedTime = ms
			}
		}
	}
	return memoryUsage, elapsedTime
}
