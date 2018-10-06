package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func main() {
	/*函数作为值*/
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	/*闭包 */
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	//函数方法
	var c1 Circle
	c1.radius = 10.0
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}                                                        

/*函数作为闭包*/
func getSequence() func() int {
	i := 0

	return func() int {
 		i += 1
		return i
	}
}

/*函数方法*/
//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}