package main

import (
	"fmt"
	"github.com/shijting/cache/src/cache"
	"log"
	"time"
)
func main()  {
	c := cache.NewCache()
	err :=c.SetMaxMemory("1kb")
	if err!=nil {
		log.Fatal(err)
	}
	//for i :=0;i<1000;i++ {
	//	data := map[int]int{1:1, 2:2,3:3}
	//	key := i
	//	err = c.Set(string(key), data, time.Second * 2)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	err = c.Set("a", 1, time.Second * 2)
	if err!=nil {
		log.Fatal(err)
	}
	c.Set("a", 11, time.Second * 10)
	c.Set("b", 2, time.Second * 20)
	c.Set("c", 3, time.Second * 5)
	//c.Del("c")
	//c.Flush()
	for  {
		v, _ := c.Get("a")
		fmt.Printf("key[a]=%v\n", v)
		v, _ = c.Get("b")
		fmt.Printf("key[b]=%v\n", v)
		v, _ = c.Get("c")
		fmt.Printf("key[c]=%v\n", v)
		fmt.Println("length: ",c.Keys())
		fmt.Printf("size:%#v\n", cache.SizeOf(c))
		fmt.Println("----------------------------------")
		time.Sleep(time.Second * 1)
	}
}
