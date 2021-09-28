package _929_Unique_Email_Addresses

import "strings"

func numUniqueEmails(emails []string) int {
	m := map[string]bool{}
	for _, email := range emails {
		names := strings.Split(email, "@")
		legalName := strings.Split(names[0], "+")[0]
		localName := strings.Join(strings.Split(legalName, "."), "")
		m[localName+"@"+names[1]] = true
	}
	return len(m)
}
