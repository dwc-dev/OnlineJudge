package codes

var (
	PythonCode = `
def main():
	a, b = map(int, input().split())
	print(a + b)
if __name__ == "__main__":
	main()
`
	PythonInput = []string{"1 2", "3 4", "5 6", "7 8", "9 10"}
)

var (
	PythonCodeStackOverflow = `
def crash():
	crash()
crash()
`
	PythonInputStackOverflow = []string{"1 2"}
)
