package main

import (
	"fmt"
	"strconv"
	"time"
)

type Person interface {
	getName() string
}
type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

//func main() {
//	var p Person = &Student{
//		name: "Tom",
//		age:  18,
//	}
//	stu := p.(*Student)
//	fmt.Println(stu.age)
//}

//func main() {
//	m := make(map[string]interface{})
//	m["name"] = "Tom"
//	m["age"] = 18
//	m["scores"] = [3]int{98, 99, 85}
//	fmt.Println(m)
//}

var ch = make(chan string, 10)

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url //将url发送到信道
}
func main() {
	for i := 0; i < 3; i++ {
		go download("a.com/" + strconv.Itoa(i))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch //等待信道返回信息
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}
