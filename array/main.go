package main

import "fmt"

func main() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var i,j int

   /* 为数组 n 初始化元素 */         
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }

   fmt.Println("The average of n = ", getAverage(n))
   /* 数组 - 5 行 2 列,先行后列*/
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
   for i := 0; i < 5; i++{
	   for j := 0; j < 2; j++ {
		   fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
	   }
   }
}

func getAverage(arr [10]int) (ret float64){
	l := len(arr)
	sum := 0

	for i := 0; i < l; i++ {
		sum += arr[i]
	}

	if l != 0 {
		ret = float64(sum)/float64(l)
	} else {
		ret = 0
	}

	return ret
}