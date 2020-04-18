package list

import (
	"go-cidrman/util"
)
type Amazon struct{}

func (f *Amazon) Name() string {
	return "amazon"
}

type Result struct {
	CreateDate string       `json:"createDate"`
	Prefixes   []CidrDetail `json:"prefixes"`
	SyncToken  string       `json:"syncToken"`
}

type CidrDetail struct {
	IPPrefix           string `json:"ip_prefix"`
	NetworkBorderGroup string `json:"network_border_group"`
	Region             string `json:"region"`
	Service            string `json:"service"`
}

func (a *Amazon) FetchAndMerge() []string {
	var amazonJson *Result
	var amazonCIDRArray []string
	util.HttpGetJSON("https://ip-ranges.amazonaws.com/ip-ranges.json", &amazonJson)
	if amazonJson != nil && len(amazonJson.Prefixes) > 0 { // cannot find in local cache
		for _, cidr := range amazonJson.Prefixes {
			amazonCIDRArray = append(amazonCIDRArray, cidr.IPPrefix)
		}
	}
	return Merge(a,amazonCIDRArray);
}

func init() {
	RegisterCommand(&Amazon{})
}
