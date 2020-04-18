package list

// Package is called aw
import (
	"fmt"
	"github.com/yuanmomo/go-cidrman/util"
	"os/exec"
	"strings"
)

type Facebook struct{}

func (f *Facebook) Name() string {
	return  "facebook";
}


func (f *Facebook) FetchAndMerge() []string{
	var fbCIDRArray []string;
	exists, desc := util.CheckCommandExists("whois", "grep", "awk")
	fmt.Printf("%s\n", desc);
	if ! exists {
		return fbCIDRArray;
	}

	cmd :=  "whois -h whois.radb.net -- '-i origin AS32934' | grep -i  \"^route\" |grep -i  -v route6|awk -F ' ' '{print $2}'";

	whoisCmd := exec.Command("bash","-c", cmd)
	out, err := whoisCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Get Facebook ip range error : %s\n",err.Error())
		return fbCIDRArray;
	}
	fbCIDRString := string(out)
	if len(fbCIDRString) > 0{
		fbCIDRArray = append(fbCIDRArray, strings.Split(fbCIDRString,"\n")...)
	}

	return Merge(f,fbCIDRArray);
}

func init() {
	RegisterCommand(&Facebook{})
}


