package main

import (
	"fmt"
	"github.com/shijting/cache/src/cache"
	"time"
)

func main()  {
	c := cache.NewCache()
	c.Set("a", 1, time.Second * 2)
	c.Set("a", 11, time.Second * 10)
	fmt.Println(c.Get("a"))
	c.Set("b", 2, time.Second * 2)
	fmt.Println(c.Get("b"))
	c.Set("c", 3, time.Second * 2)
	fmt.Println("length: ",c.Keys())
	data := map[int]int{1:1, 2:2,3:3}
	c.Set("d", data, time.Second * 2)
	fmt.Printf("size:%#v", cache.SizeOf(c))
}
