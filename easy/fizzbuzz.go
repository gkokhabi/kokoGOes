/*
Given an integer n, return a string array answer (1-indexed) where:
answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.
*/

func fizzBuzz(n int) []string {

	var ans []string
	for i:=1;i <= n;i++ {
			mod3 := (i % 3 ==0)
			mod5 := (i % 5 ==0)
			if mod3&&mod5 {
					ans=append(ans,"FizzBuzz")
					continue
			}
			if mod3 {
					ans=append(ans,"Fizz")
					continue
			}
			if mod5 {
					ans=append(ans,"Buzz")
					continue
			}
			ans = append(ans,strconv.Itoa(i))
	}
	return ans
}

