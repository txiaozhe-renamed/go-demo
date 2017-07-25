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
 *     Initial: 2017/07/06     Tang Xiaoji
 */

package main

import (
	"fmt"
	"math"
)

// 类型断言

type Square struct {
	side float32
}
type Circle struct {
	radius float32
}
type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	show(sq1)

	sq2 := new(Circle)
	sq2.radius = 3

	areaIntf = sq1 //把结构实例赋值给接口变量
	// Is Square the type of areaIntf ?
	if t, ok := areaIntf.(Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
		if u, ok := areaIntf.(*Circle); ok {
			fmt.Printf("The type of areaIntf is: %T\n", u)
		} else {
			fmt.Println("areaIntf does not contain a variable of type Circle")
		}
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func show(s Shaper) {
	fmt.Println(s.Area())
}
