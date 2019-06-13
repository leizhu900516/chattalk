package utils

import (
	"fmt"
	"testing"
)

func TestIpgps( t *testing.T){
	if result,err := IpLocation("1.192.63.98");err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}

}
