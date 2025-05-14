package codes

var (
	GolangCode = `
		package main
		import "fmt"
		func main() {
			var a, b int
			fmt.Scanf("%d%d", &a, &b)
			fmt.Print(a + b)
		}`
	GolangInput = []string{"1 2", "3 4", "5 6", "7 8", "9 10"}
)
