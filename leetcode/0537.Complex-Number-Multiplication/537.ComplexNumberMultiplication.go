package _537_Complex_Number_Multiplication

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	r1, i1 := getRealAndImg(num1)
	r2, i2 := getRealAndImg(num2)
	ri, ii := r1*r2-i1*i2, r1*i2+r2*i1
	return strconv.Itoa(ri) + "+" + strconv.Itoa(ii) + "i"
}

func getRealAndImg(num string) (int, int) {
	ss := strings.Split(num, "+")
	rs, is := ss[0], ss[1]
	ri, _ := strconv.Atoi(rs)
	ii, _ := strconv.Atoi(is[:len(is)-1])
	return ri, ii
}
