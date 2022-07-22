/*
Given an integer number n, return the difference between the product of its digits and the sum of its digits.
*/

func subtractProductAndSum(n int) int {
	product := 1;
	sum := 0;
	for n > 0 {
			item := n % 10;
			product *=item;
			sum +=item;
			n /=10;
	}
	return product - sum;
}
