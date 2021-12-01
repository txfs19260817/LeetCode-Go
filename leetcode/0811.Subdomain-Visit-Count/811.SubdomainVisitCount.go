package _811_Subdomain_Visit_Count

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domain2cnt := map[string]int{}
	for _, cpdomain := range cpdomains {
		cpSplit := strings.Split(cpdomain, " ")
		cntStr, domainsStr := cpSplit[0], cpSplit[1]
		count, _ := strconv.Atoi(cntStr)
		domains := strings.Split(domainsStr, ".")
		var subDomain string
		for i := len(domains) - 1; i >= 0; i-- {
			if len(subDomain) == 0 {
				subDomain = domains[i]
			} else {
				subDomain = domains[i] + "." + subDomain
			}
			domain2cnt[subDomain] += count
		}
	}
	ans := make([]string, 0, len(domain2cnt))
	for d, c := range domain2cnt {
		ans = append(ans, strconv.Itoa(c)+" "+d)
	}
	return ans
}
