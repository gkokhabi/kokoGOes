/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".
*/

func longestCommonPrefix(strs []string) string {
	var res string
	for i,ch := range strs[0] {
			for _,str := range strs {
					if i > len(str)-1 return res

					if rune(str[i]) != ch return res

			}
			res += string(ch)
	}
	return res
}
