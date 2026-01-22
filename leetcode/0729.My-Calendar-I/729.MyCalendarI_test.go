package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCalendarBook(t *testing.T) {
	type op struct {
		start int
		end   int
		want  bool
	}
	tests := []struct {
		name       string
		ops        []op
		wantStarts []int
		wantEnds   []int
	}{
		{
			name: "leetcode-example",
			ops: []op{
				{start: 10, end: 20, want: true},
				{start: 15, end: 25, want: false},
				{start: 20, end: 30, want: true},
			},
			wantStarts: []int{10, 20},
			wantEnds:   []int{20, 30},
		},
		{
			name: "out-of-order-non-overlap-and-touching",
			ops: []op{
				{start: 30, end: 40, want: true},
				{start: 5, end: 10, want: true},
				{start: 15, end: 20, want: true},
				{start: 10, end: 15, want: true},
				{start: 9, end: 12, want: false},
			},
			wantStarts: []int{5, 10, 15, 30},
			wantEnds:   []int{10, 15, 20, 40},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cal := Constructor()
			for i, op := range tt.ops {
				got := cal.Book(op.start, op.end)
				assert.Equalf(t, op.want, got, "op %d Book(%d, %d)", i, op.start, op.end)
			}
			assert.Equal(t, tt.wantStarts, cal.start, "calendar start mismatch")
			assert.Equal(t, tt.wantEnds, cal.end, "calendar end mismatch")
		})
	}
}

func TestMyCalendarDelete(t *testing.T) {
	cal := Constructor()
	assert.True(t, cal.Book(10, 20), "Book(10, 20) should succeed")
	assert.True(t, cal.Book(30, 40), "Book(30, 40) should succeed")
	assert.True(t, cal.Delete(10, 20), "Delete(10, 20) should succeed")
	assert.False(t, cal.Delete(10, 20), "Delete(10, 20) should fail for missing booking")
	assert.False(t, cal.Delete(30, 50), "Delete(30, 50) should fail for non-matching end time")
	assert.True(t, cal.Book(15, 25), "Book(15, 25) should succeed after delete")
	assert.True(t, cal.Delete(15, 25), "Delete(15, 25) should succeed")
	assert.Equal(t, []int{30}, cal.start, "calendar start mismatch after deletes")
	assert.Equal(t, []int{40}, cal.end, "calendar end mismatch after deletes")
}
