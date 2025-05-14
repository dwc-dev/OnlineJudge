package codes

var (
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
)
