/*
Given an n x n binary matrix image, flip the image horizontally, then invert it, and return the resulting image.
To flip an image horizontally means that each row of the image is reversed.
For example, flipping [1,1,0] horizontally results in [0,1,1].
To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0.
For example, inverting [0,1,1] results in [1,0,0].
*/

func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
			for i, j := 0, len(row) - 1; i < j; i, j = i + 1, j - 1 {
					row[i], row[j] = row[j], row[i]
			}
			
			for k, _ := range row {
					row[k] ^= 1
			}
	}
	
	return image
}
