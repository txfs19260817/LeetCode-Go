package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type examOp struct {
	kind string
	arg  int
}

func TestExamRoomSeatLeave(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		ops       []examOp
		wantSeats []int
	}{
		{
			name: "example flow",
			n:    10,
			ops: []examOp{
				{kind: "seat"},
				{kind: "seat"},
				{kind: "seat"},
				{kind: "seat"},
				{kind: "leave", arg: 4},
				{kind: "seat"},
			},
			wantSeats: []int{0, 9, 4, 2, 5},
		},
		{
			name: "leave left edge",
			n:    3,
			ops: []examOp{
				{kind: "seat"},
				{kind: "seat"},
				{kind: "leave", arg: 0},
				{kind: "seat"},
			},
			wantSeats: []int{0, 2, 0},
		},
		{
			name: "leave right edge",
			n:    4,
			ops: []examOp{
				{kind: "seat"},
				{kind: "seat"},
				{kind: "seat"},
				{kind: "leave", arg: 3},
				{kind: "seat"},
			},
			wantSeats: []int{0, 3, 1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			room := Constructor(tt.n)
			gotSeats := make([]int, 0, len(tt.wantSeats))
			for i, op := range tt.ops {
				switch op.kind {
				case "seat":
					gotSeats = append(gotSeats, room.Seat())
				case "leave":
					room.Leave(op.arg)
				default:
					t.Fatalf("op[%d]: unknown kind %q", i, op.kind)
				}
			}
			assert.Equal(t, tt.wantSeats, gotSeats)
		})
	}
}
