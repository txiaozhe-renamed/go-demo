package main

import "fmt"

//两个结构 stockPosition 和 car , 具有相同的方法 getValue()
//定义一个接口包含getValue() 方法
//定义一个以接口为参数的方法，意味着它接受任何实现了这个接口的参数

//结构1
type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

//结构2
type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

//接口
type valuable interface {
	getValue() float32
}

/* anything that satisfies the "valuable" interface is accepted */
func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	//stockPosition结构实现 getValue() 方法返回两个数的积
	//valuable 指明类型
	var o valuable = stockPosition{"GOOG", 577.20, 4 }
	//showValue() 函数接受任何值，接口中有getValue() 函数已被stockPosition结构实现
	showValue(o)
	//o = car{"BMW", "M3", 66500 }
	//showValue(o)
}
