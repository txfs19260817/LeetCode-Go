package k

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shopping(t *testing.T) {
	products := [][]string{
		{"Cheese", "Dairy"},
		{"Carrots", "Produce"},
		{"Potatoes", "Produce"},
		{"Canned Tuna", "Pantry"},
		{"Romaine Lettuce", "Produce"},
		{"Chocolate Milk", "Dairy"},
		{"Flour", "Pantry"},
		{"Iceberg Lettuce", "Produce"},
		{"Coffee", "Pantry"},
		{"Pasta", "Pantry"},
		{"Milk", "Dairy"},
		{"Blueberries", "Produce"},
		{"Pasta Sauce", "Pantry"},
	}

	type args struct {
		products [][]string
		list     []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "list1",
			args: args{
				products: products,
				list:     []string{"Blueberries", "Milk", "Coffee", "Flour", "Cheese", "Carrots"},
			},
			want: 2,
		},
		{
			name: "list2",
			args: args{
				products: products,
				list:     []string{"Blueberries", "Carrots", "Coffee", "Milk", "Flour", "Cheese"},
			},
			want: 2,
		},
		{
			name: "list3",
			args: args{
				products: products,
				list:     []string{"Blueberries", "Carrots", "Romaine Lettuce", "Iceberg Lettuce"},
			},
			want: 0,
		},
		{
			name: "list4",
			args: args{
				products: products,
				list:     []string{"Milk", "Flour", "Chocolate Milk", "Pasta Sauce"},
			},
			want: 2,
		},
		{
			name: "list5",
			args: args{
				products: products,
				list:     []string{"Cheese", "Potatoes", "Blueberries", "Canned Tuna"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, shopping(tt.args.products, tt.args.list))
		})
	}
}

func Test_carpool(t *testing.T) {
	// 构造庞大输入的辅助数据
	// 场景：一条长长的主路，和一条从中间汇入的支路
	// MainLine: M0 -> M1 -> M2 ... -> M100 -> Campground
	// SideLine: S0 -> S1 ... -> S50 -> M50 (汇入点)
	var hugeRoads [][]string
	var hugePeople [][]string

	// 构建主路 0-100
	for i := 0; i < 100; i++ {
		// M0->M1, M1->M2...
		// 每段路耗时 10 分钟
		start := "M" + strconv.Itoa(i)
		end := "M" + strconv.Itoa(i+1)
		if i == 99 {
			end = "Campground"
		}
		hugeRoads = append(hugeRoads, []string{start, end, "10"})

		// 每个主路节点放一个人 "P_Mi"
		hugePeople = append(hugePeople, []string{"P_" + start, start})
	}

	// 构建支路 0-50，汇入 M50
	for i := 0; i <= 50; i++ {
		start := "S" + strconv.Itoa(i)
		var end string
		if i == 50 {
			end = "M50" // 汇入主路
		} else {
			end = "S" + strconv.Itoa(i+1)
		}
		// 支路走得快一点，每段 5 分钟
		hugeRoads = append(hugeRoads, []string{start, end, "5"})

		// 每个支路节点放一个人 "P_Si"
		hugePeople = append(hugePeople, []string{"P_" + start, start})
	}

	// Car1 从 M0 出发
	// Car2 从 S0 出发

	// 预期逻辑分析：
	// Car1 到达 M50 的时间：50段 * 10分钟 = 500分钟
	// Car2 到达 M50 的时间：51段(S0..S50->M50) * 5分钟 = 255分钟
	// 所以在 M50 及其之后的节点 (M50...Campground)，Car2 都会比 Car1 先到（或者时间更短）。

	// 乘客分配：
	// 1. M0...M49 的人：只有 Car1 能到 -> Car1
	// 2. S0...S50 的人：只有 Car2 能到 -> Car2
	// 3. M50...M99 的人：
	//    Car1 到 M50 是 500分
	//    Car2 到 M50 是 255分
	//    对于 M(50+k)，Car1 时间 500 + 10k，Car2 时间 255 + 10k。
	//    恒定 Car2 更快 -> Car2

	var wantCar1Huge []string
	var wantCar2Huge []string

	for i := 0; i < 50; i++ {
		wantCar1Huge = append(wantCar1Huge, "P_M"+strconv.Itoa(i))
	}
	for i := 0; i <= 50; i++ {
		wantCar2Huge = append(wantCar2Huge, "P_S"+strconv.Itoa(i))
	}
	for i := 50; i < 100; i++ {
		wantCar2Huge = append(wantCar2Huge, "P_M"+strconv.Itoa(i))
	}

	tests := []struct {
		name   string
		roads  [][]string
		starts []string
		people [][]string
		want   [][]string
	}{
		{
			name: "Test Case 1 (Standard)",
			roads: [][]string{
				{"Bridgewater", "Caledonia", "30"},
				{"Caledonia", "New Grafton", "15"},
				{"New Grafton", "Campground", "5"},
				{"Liverpool", "Milton", "10"},
				{"Milton", "New Grafton", "30"},
			},
			starts: []string{"Bridgewater", "Liverpool"},
			people: [][]string{
				{"Jessie", "Bridgewater"}, {"Travis", "Caledonia"},
				{"Jeremy", "New Grafton"}, {"Katie", "Liverpool"},
			},
			want: [][]string{{"Jessie", "Travis"}, {"Katie", "Jeremy"}},
		},
		{
			name: "Test Case 2 (Merge & Different Distances)",
			roads: [][]string{
				{"Riverport", "Chester", "40"},
				{"Chester", "Campground", "60"},
				{"Halifax", "Chester", "40"},
			},
			starts: []string{"Riverport", "Halifax"},
			people: [][]string{
				{"Colin", "Riverport"}, {"Sam", "Chester"}, {"Alyssa", "Halifax"},
			},
			// Riverport->Chester(40)
			// Halifax->Chester(40)
			// Both arrive at Chester at time 40. Sam can go in either.
			// We expect Sam to be in either list, but for deterministic test check we might need flexibility.
			// Based on code logic: if t1 < t2 -> car1; if t2 < t1 -> car2; else -> car1 (usually).
			// My implementation defaults equal time to car1.
			want: [][]string{{"Colin", "Sam"}, {"Alyssa"}},
		},
		{
			name: "Test Case 3_1 (Same Line, Car1 leads)",
			roads: [][]string{
				{"Riverport", "Bridgewater", "1"},
				{"Bridgewater", "Liverpool", "1"},
				{"Liverpool", "Campground", "1"},
			},
			starts: []string{"Riverport", "Bridgewater"},
			people: [][]string{
				{"Colin", "Riverport"}, {"Jessie", "Bridgewater"}, {"Sam", "Liverpool"},
			},
			want: [][]string{{"Colin"}, {"Jessie", "Sam"}},
		},
		{
			name: "Test Case 3_2 (Same Line, Car2 leads)",
			roads: [][]string{
				{"Riverport", "Bridgewater", "1"},
				{"Bridgewater", "Liverpool", "1"},
				{"Liverpool", "Campground", "1"},
			},
			starts: []string{"Bridgewater", "Riverport"}, // Swapped starts
			people: [][]string{
				{"Colin", "Riverport"}, {"Jessie", "Bridgewater"}, {"Sam", "Liverpool"},
			},
			// Car1 (Bri) starts ahead.
			// Car2 (Riv) is behind.
			// Colin(Riv): Only Car2 can reach.
			// Jessie(Bri): Car1(0), Car2(1). Car1 wins.
			// Sam(Liv): Car1(1), Car2(2). Car1 wins.
			want: [][]string{{"Jessie", "Sam"}, {"Colin"}},
		},
		{
			name: "Test Case 3_3 (Gaps)",
			roads: [][]string{
				{"Riverport", "Bridgewater", "1"},
				{"Bridgewater", "Liverpool", "1"},
				{"Liverpool", "Campground", "1"},
			},
			starts: []string{"Riverport", "Liverpool"},
			people: [][]string{
				{"Colin", "Riverport"}, {"Jessie", "Bridgewater"}, {"Sam", "Liverpool"},
			},
			// Car1(Riv): Riv(0)->Bri(1)->Liv(2)
			// Car2(Liv): Liv(0)
			// Colin(Riv): Car1(0) -> Car1
			// Jessie(Bri): Car1(1) -> Car1
			// Sam(Liv): Car1(2) vs Car2(0) -> Car2
			want: [][]string{{"Colin", "Jessie"}, {"Sam"}},
		},
		{
			name:   "Test Case Massive (Scalability)",
			roads:  hugeRoads,
			starts: []string{"M0", "S0"},
			people: hugePeople,
			want:   [][]string{wantCar1Huge, wantCar2Huge},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := carpool(tt.roads, tt.starts, tt.people)

			// Helper to check if a slice contains an element
			contains := func(slice []string, item string) bool {
				for _, s := range slice {
					if s == item {
						return true
					}
				}
				return false
			}

			// For "Test Case 2", Sam can be in either car.
			// We handle this specific ambiguity or rely on implementation detail (Car1 preference).
			if tt.name == "Test Case 2 (Merge & Different Distances)" {
				// Strict check for non-ambiguous ones
				assert.True(t, contains(got[0], "Colin"))
				assert.True(t, contains(got[1], "Alyssa"))
				// Sam must be in one of them
				inCar1 := contains(got[0], "Sam")
				inCar2 := contains(got[1], "Sam")
				assert.True(t, inCar1 || inCar2, "Sam must be picked up")
				assert.False(t, inCar1 && inCar2, "Sam cannot be in both cars")
			} else {
				assert.ElementsMatch(t, tt.want[0], got[0], "Car 1 passengers mismatch")
				assert.ElementsMatch(t, tt.want[1], got[1], "Car 2 passengers mismatch")
			}
		})
	}
}
