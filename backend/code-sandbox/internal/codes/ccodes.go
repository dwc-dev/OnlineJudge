package codes

var (
	CCode = `
		#include <stdio.h>
		int main() {
			int a, b;
			scanf("%d%d", &a, &b);
			printf("%d", a + b);
			return 0;
		}`
	CInput = []string{"1 2", "3 4", "5 6", "7 8", "9 10"}
)
