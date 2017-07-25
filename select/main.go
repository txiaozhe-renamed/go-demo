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
 *     Initial: 2017/07/10     Tang Xiaoji
 */

package main

import (
	"fmt"
)

func main() {

	// demo 1
	//ch1 := make(chan int, 1)
	//ch2 := make(chan int, 1)
	//
	//go func() {
	//	ch2 <- 1
	//}()
	//
	//select {
	//case <- ch1:
	//	fmt.Println("ch1")
	//case <- ch2:
	//	fmt.Println("ch2")
	//}

	// demo2
	//timeout := make(chan bool, 1)
	//ch := make(chan int)
	//go func() {
	//	ch <- 1 // 如果在这里给一个值，则无法产生延迟效果
	//	time.Sleep(1e9 * 15)
	//	timeout <- true
	//}()
	//
	//select {
	//case <- ch:
	//	fmt.Println("ch")
	//case <- timeout:
	//	fmt.Println("timeout")
	//}

	// demo3
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("channel is full")
	}
}
