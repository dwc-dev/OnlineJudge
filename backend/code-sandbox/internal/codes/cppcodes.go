package codes

var (
	CppCode = `
		#include <iostream>
		using namespace std;
		int main() {
			int a, b;
			cin >> a >> b;
			cout << a + b;
			return 0;
		}`
	CppInput = []string{"1 2", "3 4", "5 6", "7 8", "9 10"}
)
