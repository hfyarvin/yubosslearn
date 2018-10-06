package main

import "fmt"

const (
	MAX int = 3
)

func main() {
   var a int = 10   
   var ip *int	//空指针，存值为0
   if ip == nil {
	   fmt.Println("ip 为空指针")
   }

   ip = &a

   fmt.Printf("变量的地址: %x\n", &a  )
   fmt.Printf("变量的地址: %x\n", ip  )
   fmt.Printf("*ip 变量的值: %d\n", *ip )
   
   //指向数组的指针
   arr := []int{10, 100, 1000}
   for i := 0; i < MAX; i++ {
	   fmt.Printf("a[%d] = %d\n", i, arr[i])
   }


}