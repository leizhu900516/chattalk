package main

import (
	"fmt"
	"reflect"
	"time"
)
func main() {
	GetNowTime()

}

func GetNowTime() {
	fmt.Println("This is my first go")
	timeStr := time.Now().String()[:19]


	fmt.Println(timeStr)
	fmt.Println(reflect.TypeOf(timeStr))
}