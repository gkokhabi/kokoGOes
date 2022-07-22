/*
Given an integer numRows, return the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it
*/


func generate(numRows int) [][]int {
	arr:=make([][]int,numRows)  //make a 2d array
	for i:=0;i<numRows;i++{
			arr[i]=make([]int,i+1)      //size of each row will be equal to the row number
	}
	for i:=0;i<numRows;i++{
			for j:=0;j<len(arr[i]);j++{
					if (j==0 || j==len(arr[i])-1){
							arr[i][j]=1     //if the element is first or last then it should be 1
					}else{
							arr[i][j]=arr[i-1][j-1]+arr[i-1][j]   //sum of values directly above and to the left of the current element
					}
			}
	}
	return arr
}
