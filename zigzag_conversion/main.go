// https://leetcode.com/problems/zigzag-conversion/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("1: ", convert("PAYPALISHIRING", 1))
	fmt.Println("2: ", convert("PAYPALISHIRING", 2))
	fmt.Println("3: ", convert("PAYPALISHIRING", 3))
	fmt.Println("4: ", convert("PAYPALISHIRING", 4))
	fmt.Println("5: ", convert("PAYPALISHIRING", 5))
	//PAYPALISHIRING
	//PYAIHRNAPLSIIG
	//PAHNAPLSIIGYIR
	//PINALSIGYAHRPI
	//PHASIYIRPLIGAN
}

func getCurRowWidths(curOffset, maxWidth int) (int, int) {
	curWidth := maxWidth - 2*curOffset
	return curWidth, maxWidth - curWidth
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	answer := make([]byte, 0, len(s))
	maxWidth := numRows*2 - 2
	for curRow := 0; curRow < numRows; curRow++ {
		for jdx := 0; jdx+curRow < len(s); jdx += maxWidth {
			answer = append(answer, s[jdx+curRow])
			if curRow != 0 && curRow != numRows-1 && ((jdx + maxWidth - curRow) < len(s)) {
				answer = append(answer, s[jdx+maxWidth-curRow])
			}
		}

		// my own solution
		//curOffset := int(math.Abs(math.Abs(float64(numRows-1)/2-float64(curRow)) - float64(numRows-1)/2))
		//if curOffset == 0 {
		//	for idx := curRow; idx < len(s); idx += maxWidth {
		//		answer = append(answer, s[idx])
		//	}
		//} else {
		//	width1, width2 := getCurRowWidths(curOffset, maxWidth)
		//	if curRow >= numRows/2 {
		//		width1, width2 = width2, width1
		//	}
		//	for idx, jdx := curRow, 1; idx < len(s); jdx++ {
		//		answer = append(answer, s[idx])
		//		if jdx%2 == 1 {
		//			idx += width1
		//		} else {
		//			idx += width2
		//		}
		//	}
		//}
	}
	return string(answer)
}
