package main

import (
	"fmt"
	"go-cidrman/list"
)

func main()  {
	all := list.GetAll();
	for _,cidr := range all {
		fmt.Printf("%s\n",cidr)
	}

}


