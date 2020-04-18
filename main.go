package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)
type Result struct {
	CreateDate string            `json:"createDate"`
	Prefixes   []CidrDetail `json:"prefixes"`
	SyncToken  string            `json:"syncToken"`
}

type CidrDetail struct {
	IPPrefix           string `json:"ip_prefix"`
	NetworkBorderGroup string `json:"network_border_group"`
	Region             string `json:"region"`
	Service            string `json:"service"`
}

func GetAmazonCIDR() (cidrList []string){
	var amazonJson *Result;
	var amazonCIDRArray []string;

	HttpGetJSON("https://ip-ranges.amazonaws.com/ip-ranges.json",&amazonJson);
	if amazonJson != nil && len(amazonJson.Prefixes) > 0 { // cannot find in local cache
		for _, cidr := range amazonJson.Prefixes {
			amazonCIDRArray = append(amazonCIDRArray,cidr.IPPrefix);
		}
	}
	return amazonCIDRArray;
}
func checkCommandExists(commandList ... string) (bool,string){

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

//
func GetFacebookCIDR() (cidrList []string){
	exists, desc := checkCommandExists("whois", "grep", "awk")
	fmt.Printf("%s\n", desc);
	if ! exists {
		os.Exit(1);
	}

	cmd :=  "whois -h whois.radb.net -- '-i origin AS32934' | grep -i  \"^route\" |grep -i  -v route6|awk -F ' ' '{print $2}'";
	var fbCIDRArray []string;
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

	return fbCIDRArray;
}

func main()  {

	var cidrArray []string;

	cidrArray = append(cidrArray,GetFacebookCIDR()...);

	//res, err := MergeCIDRs(strings.Split(ipList, "\n"))
	res, err := MergeCIDRs(cidrArray)
	if err != nil {
		fmt.Printf("Merge error : %s\n",err.Error())
		return
	}
	if len(res) > 0 {
		for _,cidr := range res {
			fmt.Printf("%s\n",cidr)
		}
	}
}


