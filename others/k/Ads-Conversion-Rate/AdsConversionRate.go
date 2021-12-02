package k

import (
	"fmt"
	"strings"
)

func adsConversionRate(completedPurchaseUserIds, adClicks, allUserIps []string) []string {
	// slice 2 map
	uid2ip, ip2txt, txt2clk, txt2bgt := map[string]string{}, map[string]string{}, map[string]int{}, map[string]int{}
	for _, e := range allUserIps {
		entry := strings.Split(e, ",")
		uid2ip[entry[0]] = entry[1]
	}
	for _, e := range adClicks {
		entry := strings.Split(e, ",")
		ip2txt[entry[0]] = entry[2]
		txt2clk[entry[2]]++
	}
	for _, userId := range completedPurchaseUserIds {
		txt2bgt[ip2txt[uid2ip[userId]]]++
	}
	// output
	ans := []string{"Bought Clicked Ad Text"}
	for txt, clk := range txt2clk {
		bgt := txt2bgt[txt]
		ans = append(ans, fmt.Sprintf("%d of %d  %s", bgt, clk, txt))
	}
	return ans
}
