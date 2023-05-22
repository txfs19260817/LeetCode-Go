package leetcode

func compress(chars []byte) int {
	var l, r, w int // left, right pointers, and a writer pointer
	for r < len(chars) {
		for chars[l] == chars[r] && r <= len(chars) && r+1 < len(chars) && chars[l] == chars[r+1] { // let r points to the rightmost byte that equals to chars[l]
			r++
		}
		chars[w] = chars[l]             // write this byte
		if cnt := r - l + 1; cnt != 1 { // with O(1) space, convert the number of chars[l] to string format and write them into chars
			var digits int // get the number of digits
			for k := cnt; k > 0; k /= 10 {
				digits++
			}
			w += digits // write them from right to left
			for ; cnt > 0; cnt, w = cnt/10, w-1 {
				chars[w] = byte('0' + cnt%10)
			}
			w += digits
		}
		w++ // move w to the next position that can be written
		r++ // move both pointers to the next different byte
		l = r
	}
	return w
}
