package main

import (
	"fmt"
	"github.com/yuanmomo/go-cidrman/list"
)

func main()  {
	all := list.GetAll();
	for _,cidr := range all {
		fmt.Printf("%s\n",cidr)
	}

}


