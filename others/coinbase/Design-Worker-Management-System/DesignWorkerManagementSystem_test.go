package coinbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type step struct {
	op   string
	args []interface{}
	want interface{}
}

func runSteps(t *testing.T, steps []step) {
	t.Helper()
	manager := Constructor()
	for i, s := range steps {
		var got interface{}
		switch s.op {
		case "addWorker":
			workerID := s.args[0].(string)
			position := s.args[1].(string)
			compensation := s.args[2].(int)
			got = manager.AddWorker(workerID, position, compensation)
		case "registerWorker":
			workerID := s.args[0].(string)
			timestamp := s.args[1].(int)
			got = manager.RegisterWorker(workerID, timestamp)
		case "get":
			workerID := s.args[0].(string)
			got = manager.Get(workerID)
		case "topNWorkers":
			count := s.args[0].(int)
			position := s.args[1].(string)
			got = manager.TopNWorkers(count, position)
		case "promote":
			workerID := s.args[0].(string)
			newPosition := s.args[1].(string)
			newCompensation := s.args[2].(string)
			startTimestamp := s.args[3].(int)
			got = manager.Promote(workerID, newPosition, newCompensation, startTimestamp)
		case "calcSalary":
			workerID := s.args[0].(string)
			startTimestamp := s.args[1].(int)
			endTimestamp := s.args[2].(int)
			got = manager.CalcSalary(workerID, startTimestamp, endTimestamp)
		case "setDoublePaid":
			startTimestamp := s.args[0].(int)
			endTimestamp := s.args[1].(int)
			manager.SetDoublePaid(startTimestamp, endTimestamp)
			got = nil
		default:
			t.Fatalf("unknown op at step %d: %s", i, s.op)
		}
		assert.Equal(t, s.want, got, "step %d (%s)", i, s.op)
	}
}

func TestOfficeManager(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		runSteps(t, []step{
			{op: "addWorker", args: []interface{}{"Ashley", "Middle Developer", 150}, want: true},
			{op: "addWorker", args: []interface{}{"Ashley", "Junior Developer", 100}, want: false},
			{op: "registerWorker", args: []interface{}{"Ashley", 10}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Ashley", 25}, want: "registered"},
			{op: "get", args: []interface{}{"Ashley"}, want: 15},
			{op: "registerWorker", args: []interface{}{"Ashley", 40}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Ashley", 67}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Ashley", 100}, want: "registered"},
			{op: "get", args: []interface{}{"Ashley"}, want: 42},
			{op: "get", args: []interface{}{"Walter"}, want: -1},
			{op: "registerWorker", args: []interface{}{"Walter", 120}, want: "invalid_request"},
		})
	})

	t.Run("Example 2", func(t *testing.T) {
		runSteps(t, []step{
			{op: "addWorker", args: []interface{}{"John", "Senior Developer", 200}, want: true},
			{op: "get", args: []interface{}{"Ashely"}, want: -1},
			{op: "registerWorker", args: []interface{}{"John", 15}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"John", 30}, want: "registered"},
			{op: "get", args: []interface{}{"John"}, want: 15},
		})
	})

	t.Run("Example 3", func(t *testing.T) {
		runSteps(t, []step{
			{op: "addWorker", args: []interface{}{"Oliver", "Middle Developer", 150}, want: true},
			{op: "registerWorker", args: []interface{}{"Oliver", 25}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Oliver", 55}, want: "registered"},
			{op: "get", args: []interface{}{"Oliver"}, want: 30},
			{op: "addWorker", args: []interface{}{"Oliver", "Middle Developer", 150}, want: false},
		})
	})

	t.Run("Example 4", func(t *testing.T) {
		runSteps(t, []step{
			{op: "addWorker", args: []interface{}{"Sophia", "Senior Developer", 200}, want: true},
			{op: "registerWorker", args: []interface{}{"Sophia", 30}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Sophia", 60}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Sophia", 90}, want: "registered"},
			{op: "get", args: []interface{}{"Sophia"}, want: 30},
		})
	})

	t.Run("Example 5", func(t *testing.T) {
		runSteps(t, []step{
			{op: "addWorker", args: []interface{}{"Emma", "Junior Developer", 100}, want: true},
			{op: "registerWorker", args: []interface{}{"Emma", 20}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Emma", 50}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Emma", 60}, want: "registered"},
			{op: "get", args: []interface{}{"Emma"}, want: 30},
			{op: "registerWorker", args: []interface{}{"Emma", 80}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Emma", 100}, want: "registered"},
			{op: "get", args: []interface{}{"Emma"}, want: 50},
		})
	})

	t.Run("Follow-up 1", func(t *testing.T) {
		runSteps(t, []step{
			{op: "addWorker", args: []interface{}{"John", "Junior Developer", 120}, want: true},
			{op: "addWorker", args: []interface{}{"Jason", "Junior Developer", 120}, want: true},
			{op: "addWorker", args: []interface{}{"Ashley", "Junior Developer", 120}, want: true},
			{op: "registerWorker", args: []interface{}{"John", 100}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"John", 150}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Jason", 200}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Jason", 250}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Jason", 275}, want: "registered"},
			{op: "topNWorkers", args: []interface{}{5, "Junior Developer"}, want: "Jason(50), John(50), Ashley(0)"},
			{op: "topNWorkers", args: []interface{}{1, "Junior Developer"}, want: "Jason(50)"},
			{op: "registerWorker", args: []interface{}{"Ashley", 400}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Ashley", 500}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"Jason", 575}, want: "registered"},
			{op: "topNWorkers", args: []interface{}{5, "Junior Developer"}, want: "Jason(350), Ashley(100), John(50)"},
			{op: "topNWorkers", args: []interface{}{5, "Middle Developer"}, want: ""},
		})
	})

	t.Run("Follow-up 2", func(t *testing.T) {
		runSteps(t, []step{
			{op: "addWorker", args: []interface{}{"John", "Middle Developer", 200}, want: true},
			{op: "registerWorker", args: []interface{}{"John", 100}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"John", 125}, want: "registered"},
			{op: "promote", args: []interface{}{"John", "Senior Developer", "500", 200}, want: "success"},
			{op: "registerWorker", args: []interface{}{"John", 150}, want: "registered"},
			{op: "promote", args: []interface{}{"John", "Senior Developer", "350", 250}, want: "invalid_request"},
			{op: "registerWorker", args: []interface{}{"John", 300}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"John", 325}, want: "registered"},
			{op: "calcSalary", args: []interface{}{"John", 0, 500}, want: 35000},
			{op: "topNWorkers", args: []interface{}{3, "Senior Developer"}, want: "John(0)"},
			{op: "registerWorker", args: []interface{}{"John", 400}, want: "registered"},
			{op: "get", args: []interface{}{"John"}, want: 250},
			{op: "topNWorkers", args: []interface{}{10, "Senior Developer"}, want: "John(75)"},
			{op: "topNWorkers", args: []interface{}{10, "Middle Developer"}, want: ""},
			{op: "calcSalary", args: []interface{}{"John", 110, 350}, want: 45500},
			{op: "calcSalary", args: []interface{}{"John", 900, 1400}, want: 0},
		})
	})

	t.Run("Follow-up 3", func(t *testing.T) {
		runSteps(t, []step{
			{op: "addWorker", args: []interface{}{"John", "Middle Developer", 100}, want: true},
			{op: "registerWorker", args: []interface{}{"John", 100}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"John", 200}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"John", 500}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"John", 600}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"John", 900}, want: "registered"},
			{op: "registerWorker", args: []interface{}{"John", 1000}, want: "registered"},
			{op: "setDoublePaid", args: []interface{}{50, 170}, want: nil},
			{op: "setDoublePaid", args: []interface{}{530, 650}, want: nil},
			{op: "setDoublePaid", args: []interface{}{580, 900}, want: nil},
			{op: "calcSalary", args: []interface{}{"John", 0, 250}, want: 17000},
			{op: "calcSalary", args: []interface{}{"John", 0, 1500}, want: 44000},
		})
	})
}
