package leetcode

func isSelfCrossing(distance []int) bool {
	for i := 3; i < len(distance); i++ {
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}
		if i >= 4 && distance[i]+distance[i-4] >= distance[i-2] && distance[i-1] == distance[i-3] {
			return true
		}
		if i >= 5 && distance[i-4] <= distance[i-2] && distance[i]+distance[i-4] >= distance[i-2] && distance[i-1] <= distance[i-3] && distance[i-1]+distance[i-5] >= distance[i-3] {
			return true
		}
	}
	return false
}

/* https://leetcode.com/problems/self-crossing/discuss/79131/Java-Oms-with-explanation/84006
			 i-2
case 1 : i-1┌─┐
            └─┼─>i
             i-3

				 i-2
case 2 : i-1 ┌────┐
             └─══>┘i-3
             i  i-4      (i overlapped i-4)

case 3 :    i-4
           ┌──┐
           │i<┼─┐
        i-3│ i-5│i-1
           └────┘
            i-2
*/
