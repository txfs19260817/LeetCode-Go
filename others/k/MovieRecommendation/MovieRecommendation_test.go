package k

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_grouping(t *testing.T) {
	// README 中的基础事件流
	baseEvents := [][]string{
		{"CONNECT", "Alice", "Bob"},
		{"DISCONNECT", "Bob", "Alice"},
		{"CONNECT", "Alice", "Charlie"},
		{"CONNECT", "Dennis", "Bob"},
		{"CONNECT", "Pam", "Dennis"},
		{"DISCONNECT", "Pam", "Dennis"},
		{"CONNECT", "Pam", "Dennis"},
		{"CONNECT", "Edward", "Bob"},
		{"CONNECT", "Dennis", "Charlie"},
		{"CONNECT", "Alice", "Nicole"},
		{"CONNECT", "Pam", "Edward"},
		{"DISCONNECT", "Dennis", "Charlie"},
		{"CONNECT", "Dennis", "Edward"},
		{"CONNECT", "Charlie", "Bob"},
	}
	// 最终状态分析：
	// Alice: Charlie, Nicole (2)
	// Bob: Dennis, Edward, Charlie (3)
	// Charlie: Alice, Bob (2)
	// Dennis: Bob, Pam, Edward (3)
	// Pam: Dennis, Edward (2)
	// Edward: Bob, Pam, Dennis (3)
	// Nicole: Alice (1)

	// 构造大规模测试数据
	// 1000个用户围成一个圈，每人连接左右两人 -> 度数均为 2
	var circleEvents [][]string
	for i := 0; i < 1000; i++ {
		u := fmt.Sprintf("User%d", i)
		v := fmt.Sprintf("User%d", (i+1)%1000)
		circleEvents = append(circleEvents, []string{"CONNECT", u, v})
	}

	type args struct {
		events      [][]string
		connections int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Example Case N=3",
			args: args{
				events:      baseEvents,
				connections: 3,
			},
			// < 3: Alice(2), Charlie(2), Pam(2), Nicole(1)
			// >= 3: Bob(3), Dennis(3), Edward(3)
			want: [][]string{
				{"Alice", "Charlie", "Nicole", "Pam"},
				{"Bob", "Dennis", "Edward"},
			},
		},
		{
			name: "Example Case N=1 (All >= 1)",
			args: args{
				events:      baseEvents,
				connections: 1,
			},
			want: [][]string{
				{},
				{"Alice", "Bob", "Charlie", "Dennis", "Edward", "Nicole", "Pam"},
			},
		},
		{
			name: "Example Case N=10 (All < 10)",
			args: args{
				events:      baseEvents,
				connections: 10,
			},
			want: [][]string{
				{"Alice", "Bob", "Charlie", "Dennis", "Edward", "Nicole", "Pam"},
				{},
			},
		},
		{
			name: "Disconnect All",
			args: args{
				events: [][]string{
					{"CONNECT", "A", "B"},
					{"CONNECT", "B", "C"},
					{"DISCONNECT", "B", "A"}, // A-B 断开
					{"DISCONNECT", "C", "B"}, // B-C 断开
				},
				connections: 1,
			},
			// 最终所有人都度数为0。 A, B, C 都在 < 1 组
			want: [][]string{
				{"A", "B", "C"},
				{},
			},
		},
		{
			name: "Duplicate Connects (Idempotency)",
			args: args{
				events: [][]string{
					{"CONNECT", "A", "B"},
					{"CONNECT", "B", "A"}, // 重复连接，不应增加度数
					{"CONNECT", "A", "B"}, // 重复连接
				},
				connections: 2,
			},
			// A: 1, B: 1.  Both < 2.
			want: [][]string{
				{"A", "B"},
				{},
			},
		},
		{
			name: "Massive Circle N=2",
			args: args{
				events:      circleEvents,
				connections: 2,
			},
			// 所有人度数都是 2。所以 < 2 为空， >= 2 为所有人
			want: [][]string{
				{}, // filled dynamically in verify
				{}, // filled dynamically
			},
		},
		{
			name: "Massive Circle N=3",
			args: args{
				events:      circleEvents,
				connections: 3,
			},
			// 所有人度数都是 2。所以 < 3 为所有人， >= 3 为空
			want: [][]string{
				{}, // filled
				{}, // filled
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := grouping(tt.args.events, tt.args.connections)

			// 针对大规模数据的特殊断言处理
			if tt.name == "Massive Circle N=2" {
				assert.Equal(t, 0, len(got[0]), "Should have no users with < 2 connections")
				assert.Equal(t, 1000, len(got[1]), "Should have 1000 users with >= 2 connections")
				return
			}
			if tt.name == "Massive Circle N=3" {
				assert.Equal(t, 1000, len(got[0]), "Should have 1000 users with < 3 connections")
				assert.Equal(t, 0, len(got[1]), "Should have no users with >= 3 connections")
				return
			}
			assert.ElementsMatch(t, tt.want[0], got[0], "Should have users with < n connections")
			assert.ElementsMatch(t, tt.want[1], got[1], "Should have users with >= n connections")
		})
	}
}

func Test_recommendations(t *testing.T) {
	// 基础测试数据
	basicRatings := [][]string{
		{"Alice", "Frozen", "5"},
		{"Bob", "Mad Max", "5"},
		{"Dennis", "Mad Max", "4"},
		{"Bob", "Lost In Translation", "5"},
	}

	// 复杂场景数据
	// UserA: M1(5), M2(2)
	// UserB: M1(4), M3(5) -> 相似 (M1 > 3)，B 推荐 M3 给 A
	// UserC: M1(5), M3(5), M4(5) -> 相似 (M1 > 3)，C 推荐 M3, M4 给 A
	// UserD: M2(5) -> 不相似 (没有共同高分电影，虽然看过M2但UserA评分低？题目定义 Both rated it above 3)
	// UserE: M5(5) -> 不相似
	complexRatings := [][]string{
		{"UserA", "M1", "5"}, {"UserA", "M2", "2"},
		{"UserB", "M1", "4"}, {"UserB", "M3", "5"},
		{"UserC", "M1", "5"}, {"UserC", "M3", "5"}, {"UserC", "M4", "5"},
		{"UserD", "M2", "5"},
		{"UserE", "M5", "5"},
	}

	tests := []struct {
		name       string
		targetUser string
		ratings    [][]string
		want       []string
	}{
		{
			name:       "Basic Example",
			targetUser: "Dennis",
			ratings:    basicRatings,
			// Bob和Dennis都看Mad Max且>3，所以Bob是相似用户
			// Bob看过 Lost In Translation，Dennis没看 -> 推荐
			want: []string{"Lost In Translation"},
		},
		{
			name:       "No Similar Users (No Common Movie)",
			targetUser: "Alice",
			ratings:    basicRatings,
			// Alice只看Frozen，没人陪她看 -> 无推荐
			want: []string{},
		},
		{
			name:       "Shared Movie but Low Rating (Not Similar)",
			targetUser: "UserD",
			ratings:    complexRatings,
			// UserD看M2(5)，UserA看M2(2)。
			// 只有一方>3，不满足 "Both rated it above 3" -> 不相似 -> 无推荐
			want: []string{},
		},
		{
			name:       "Multiple Similar Users & Deduplication",
			targetUser: "UserA",
			ratings:    complexRatings,
			// UserA(M1>3)
			// 相似用户:
			// - UserB (M1>3): 看过 M3
			// - UserC (M1>3): 看过 M3, M4
			// 汇总推荐: M3, M4 (M3去重)
			// 注意：UserA 没看过 M3, M4，所以都有效
			want: []string{"M3", "M4"},
		},
		{
			name:       "Already Seen Movies Should Not Be Recommended",
			targetUser: "UserB",
			ratings:    complexRatings,
			// UserB(M1>3)
			// 相似用户: UserA(M1>3), UserC(M1>3)
			// UserA 看过: M2 (UserB没看 -> 推荐 M2)
			// UserC 看过: M3 (UserB已看 -> 不推), M4 (UserB没看 -> 推荐 M4)
			// 结果: M2, M4
			want: []string{"M2", "M4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, recommendations(tt.targetUser, tt.ratings))
		})
	}
}
