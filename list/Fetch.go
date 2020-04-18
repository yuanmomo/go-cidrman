package list

import (
	"fmt"
	"github.com/yuanmomo/go-cidrman/cidr"
	"strings"
)

var (
	typeRegistry = make(map[string] FetchType)
)


type FetchType interface {
	Name() string
	FetchAndMerge() []string
}

func RegisterCommand(fetchType FetchType) string {
	entry := strings.ToLower(fetchType.Name())
	if entry == "" {
		return "empty command name"
	}
	typeRegistry[entry] = fetchType;
	return ""
}

func GetType(name string) FetchType {
	cmd, found := typeRegistry[name]
	if !found {
		return nil
	}
	return cmd
}

func Merge(fetchType FetchType, cidrArray []string) []string {
	//res, err := MergeCIDRs(strings.Split(ipList, "\n"))
	res, err := cidr.MergeCIDRs(cidrArray)
	if err != nil {
		fmt.Printf("[%s] merge error : %s\n", fetchType.Name(), err.Error())
		return []string{};
	}
	return res;
}


func GetAll() []string{
	var cidrList []string;

	for _, fetchType := range typeRegistry {
		merged := fetchType.FetchAndMerge();
		cidrList = append(cidrList, merged...);
	}
	return  cidrList;
}