package _079_Apply_Discount_to_Prices

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if word[0] != 36 {
			continue
		}
		num, err := strconv.Atoi(word[1:])
		if err != nil {
			continue
		}
		discountPrice := fmt.Sprintf("%.2f", float64(num)*(100-float64(discount))/100)
		words[i] = "$" + string(discountPrice)
	}
	return strings.Join(words, " ")
}
