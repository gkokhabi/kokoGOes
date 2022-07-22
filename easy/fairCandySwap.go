/*
Alice and Bob have a different total number of candies. 
You are given two integer arrays aliceSizes and bobSizes where aliceSizes[i] is the number of candies of the ith box of candy that Alice has and bobSizes[j] is the number of candies of the jth box of candy that Bob has.

Since they are friends, they would like to exchange one candy box each so that after the exchange, they both have the same total amount of candy. 
The total amount of candy a person has is the sum of the number of candies in each box they have.

Return an integer array answer where answer[0] is the number of candies in the box that Alice must exchange, and answer[1] is the number of candies in the box that Bob must exchange. 
If there are multiple answers, you may return any one of them. 
It is guaranteed that at least one answer exists.
*/


func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumA := 0
	aliceNums := make(map[int]struct{})
	for _, n := range aliceSizes {
		sumA += n
		aliceNums[n] = struct{}{}
	}
	sumB := 0
	bobNums := make(map[int]struct{})
	for _, n := range bobSizes {
		sumB += n
		bobNums[n] = struct{}{}
	}
	for n, _ := range aliceNums {
		y := ((sumB - sumA) / 2) + n
		if _, found := bobNums[y]; found {
			return []int{n, y}
		}
	}
	return []int{}
}


/* 
Notes: 
Explanation:
x: alice candies to give.
y: bob candies to give.
sum(a) - x + y = sum(b) - y + x
- x - x = sum(b) - sum(a) - y - y
- 2x = (sum(b) - sum(a)) - 2y
2y = (sum(b) - sum(a)) + 2x
y = ((sum(b) - sum(a)) + 2x) / 2
y = (sum(b) - sum(a) / 2) + x
*/
