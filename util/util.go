package util

import (
	"fmt"
	"os/exec"
)

func CheckCommandExists(commandList ... string) (bool,string){

	for _,cmd := range commandList {
		path, err := exec.LookPath(cmd)
		if err != nil {
			return false,fmt.Sprintf("didn't find '%s' executable\n",cmd)
		} else {
			return true,fmt.Sprintf("%s executable is in '%s'\n", cmd, path)
		}
	}
	return true,""
}

