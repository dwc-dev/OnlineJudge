package codes

var (
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
)

var (
	JavaCodeStackOverflow = `
		import java.util.Scanner;
		public class Main {
			public static void main(String[] args) {
				recurse();
			}
			public static void recurse() {
				recurse();
			}
		}`
	JavaInputStackOverflow = []string{"1 2"}
)
