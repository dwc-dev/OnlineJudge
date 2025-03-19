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

	RustCode = `
		use std::io;
		fn main() {
			let mut input = String::new();
			io::stdin().read_line(&mut input).unwrap();
			let mut iter = input.split_whitespace();
			let a: i32 = iter.next().unwrap().parse().unwrap();
			let b: i32 = iter.next().unwrap().parse().unwrap();
			print!("{}", a + b);
		}`
	RustInput = []string{"1 2", "3 4", "5 6", "7 8", "9 10"}

	PythonCode  = "a, b = map(int, input().split())\nprint(a + b)"
	PythonInput = []string{"1 2", "3 4", "5 6", "7 8", "9 10"}

	GoCode = `
		package main
		import "fmt"
		func main() {
			var a, b int
			fmt.Scanf("%d%d", &a, &b)
			fmt.Print(a + b)
		}`
	GoInput = []string{"1 2", "3 4", "5 6", "7 8", "9 10"}

	JavaCode = `
		import java.util.Scanner;
		public class Main {
			public static void main(String[] args) {
				Scanner sc = new Scanner(System.in);
				int a = sc.nextInt();
				int b = sc.nextInt();
				System.out.print(a + b);
			}
		}`
	JavaInput = []string{"1 2", "3 4", "5 6", "7 8", "9 10"}

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
