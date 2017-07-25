/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/07/21     Tang Xiaoji
 */

package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}

	//for index, value := range arr {
	//	fmt.Printf("arr[%d]=%d \n", index, value)
	//}
	//
	//for index := 0; index < len(arr); index++ {
	//	fmt.Printf("arr[%d]=%d \n", index, arr[index])
	//}

	//var s1 = arr[1:3]
	//var s2 = arr[2:3]
	//fmt.Println(s1, s2) // [1 3] [3]
	//
	//arr[2] = 9
	//fmt.Println(s1, s2) // [2 9] [9]

	var s1 = arr[1:3]
	var s2 = arr[1:4]
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := append(s1, 9, 10, 11)
	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Println(s2)
	fmt.Println(arr)

}
