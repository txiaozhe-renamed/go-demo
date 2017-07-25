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
 *     Initial: 2017/07/08     Tang Xiaoji
 */

package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
)

var flagvar int

type name string

func (n *name) Set(val string) error {
	if val == "" {
		return errors.New("name is nil")
	}
	fmt.Println(val)
	return nil
}

func (i *name) String() string {
	return fmt.Sprintf("%v", *i)
}

func main() {
	// type 1
	//var ip = flag.Int("name", 1234, "help message for flagname")
	//var a = flag.String("addr", "ddd", "HTTP address of addsvc")
	//
	//flag.Parse()
	//
	//fmt.Println(strconv.Itoa(*ip))
	//
	//fmt.Println(*a)

	// type2
	//flag.IntVar(&flagvar, "flagname", 1234, "help message for flag")
	//flag.Parse()
	//fmt.Println(flagvar)

	// type3
	var n name
	flag.Var(&n, "name", "show flag name")
	flag.Parse()
}
