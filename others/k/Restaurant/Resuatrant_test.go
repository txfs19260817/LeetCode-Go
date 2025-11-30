package k

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummarizeCustomer(t *testing.T) {
	type args struct {
		transactions [][]int
		customerID   int
	}
	tests := []struct {
		name            string
		args            args
		wantEventCount  int
		wantTotalAmount int
	}{
		{
			name: "Example 1",
			args: args{
				transactions: [][]int{
					{1, 100, 300},
					{2, 200, 300},
					{1, 101, 300},
					{1, 100, 300},
				},
				customerID: 1,
			},
			wantEventCount:  2,
			wantTotalAmount: 900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEventCount, gotTotalAmount := SummarizeCustomer(tt.args.transactions, tt.args.customerID)
			assert.Equal(t, tt.wantEventCount, gotEventCount)
			assert.Equal(t, tt.wantTotalAmount, gotTotalAmount)
		})
	}
}

func TestRecommendRestaurants(t *testing.T) {
	type args struct {
		liked        []int
		friendsLiked [][]int
	}
	tests := []struct {
		name            string
		args            args
		wantRecommended []int
		wantMaxCount    int
	}{
		{
			name: "Example 1",
			args: args{
				liked: []int{1, 2, 3},
				friendsLiked: [][]int{
					{2, 3, 4}, // friend 0: {2,3} common, new {4} -> count 1
					{1, 4, 5}, // friend 1: {1} common -> ignore
					{3, 6},    // friend 2: {3} common -> ignore
				},
			},
			wantRecommended: []int{4},
			wantMaxCount:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecommended, gotMaxCount := RecommendRestaurants(tt.args.liked, tt.args.friendsLiked)
			if !reflect.DeepEqual(gotRecommended, tt.wantRecommended) {
				// Handle empty slice vs nil slice if necessary, but DeepEqual usually strict
				if len(gotRecommended) == 0 && len(tt.wantRecommended) == 0 {
					// pass
				} else {
					t.Errorf("RecommendRestaurants() gotRecommended = %v, want %v", gotRecommended, tt.wantRecommended)
				}
			}
			assert.Equal(t, tt.wantMaxCount, gotMaxCount)
		})
	}
}
