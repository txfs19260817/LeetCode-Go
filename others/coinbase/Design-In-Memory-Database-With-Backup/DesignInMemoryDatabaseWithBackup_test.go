package designinmemorydatabasewithbackup

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type step struct {
	op   string
	args []string
	want interface{}
}

func runSteps(t *testing.T, steps []step) {
	t.Helper()
	db := Constructor()
	mustInt := func(input string) int {
		value, err := strconv.Atoi(input)
		if err != nil {
			t.Fatalf("invalid int arg %q: %v", input, err)
		}
		return value
	}
	for i, s := range steps {
		var got interface{}
		switch s.op {
		case "setData":
			db.SetData(s.args[0], s.args[1], s.args[2])
			got = nil
		case "getData":
			got = db.GetData(s.args[0], s.args[1])
		case "deleteData":
			got = db.DeleteData(s.args[0], s.args[1])
		case "scanData":
			got = db.ScanData(s.args[0])
		case "scanDataByPrefix":
			got = db.ScanDataByPrefix(s.args[0], s.args[1])
		case "setDataAt":
			db.SetDataAt(s.args[0], s.args[1], s.args[2], mustInt(s.args[3]))
			got = nil
		case "setDataAtWithTtl":
			db.SetDataAtWithTtl(s.args[0], s.args[1], s.args[2], mustInt(s.args[3]), mustInt(s.args[4]))
			got = nil
		case "deleteDataAt":
			got = db.DeleteDataAt(s.args[0], s.args[1], mustInt(s.args[2]))
		case "getDataAt":
			got = db.GetDataAt(s.args[0], s.args[1], mustInt(s.args[2]))
		case "scanDataAt":
			got = db.ScanDataAt(s.args[0], mustInt(s.args[1]))
		case "scanDataByPrefixAt":
			got = db.ScanDataByPrefixAt(s.args[0], s.args[1], mustInt(s.args[2]))
		case "backup":
			got = db.Backup(mustInt(s.args[0]))
		case "restore":
			db.Restore(mustInt(s.args[0]), mustInt(s.args[1]))
			got = nil
		default:
			t.Fatalf("unknown op at step %d: %s", i, s.op)
		}
		assert.Equal(t, s.want, got, "step %d (%s)", i, s.op)
	}
}

func TestInMemoryDB(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		runSteps(t, []step{
			{op: "setData", args: []string{"A", "B", "E"}, want: nil},
			{op: "setData", args: []string{"A", "C", "F"}, want: nil},
			{op: "getData", args: []string{"A", "B"}, want: "E"},
			{op: "getData", args: []string{"A", "D"}, want: ""},
			{op: "deleteData", args: []string{"A", "B"}, want: true},
			{op: "deleteData", args: []string{"A", "D"}, want: false},
		})
	})

	t.Run("Example 2", func(t *testing.T) {
		runSteps(t, []step{
			{op: "setData", args: []string{"user1", "name", "John"}, want: nil},
			{op: "setData", args: []string{"user1", "age", "25"}, want: nil},
			{op: "getData", args: []string{"user1", "name"}, want: "John"},
			{op: "getData", args: []string{"user1", "age"}, want: "25"},
			{op: "deleteData", args: []string{"user1", "age"}, want: true},
			{op: "getData", args: []string{"user1", "age"}, want: ""},
			{op: "getData", args: []string{"user1", "name"}, want: "John"},
		})
	})

	t.Run("Example 3", func(t *testing.T) {
		runSteps(t, []step{
			{op: "setData", args: []string{"record1", "field1", "value1_1"}, want: nil},
			{op: "setData", args: []string{"record1", "field2", "value2_1"}, want: nil},
			{op: "setData", args: []string{"record1", "field3", "value3_1"}, want: nil},
			{op: "setData", args: []string{"record50", "field1", "value1_50"}, want: nil},
			{op: "setData", args: []string{"record50", "field2", "value2_50"}, want: nil},
			{op: "setData", args: []string{"record50", "field3", "value3_50"}, want: nil},
			{op: "setData", args: []string{"record100", "field1", "value1_100"}, want: nil},
			{op: "setData", args: []string{"record100", "field2", "value2_100"}, want: nil},
			{op: "setData", args: []string{"record100", "field3", "value3_100"}, want: nil},
			{op: "getData", args: []string{"record50", "field2"}, want: "value2_50"},
			{op: "getData", args: []string{"record1", "field1"}, want: "value1_1"},
			{op: "getData", args: []string{"record100", "field3"}, want: "value3_100"},
			{op: "setData", args: []string{"record50", "field2", "updated_value"}, want: nil},
			{op: "getData", args: []string{"record50", "field2"}, want: "updated_value"},
			{op: "deleteData", args: []string{"record1", "field1"}, want: true},
			{op: "getData", args: []string{"record1", "field1"}, want: ""},
			{op: "getData", args: []string{"record1", "field2"}, want: "value2_1"},
		})
	})

	t.Run("Example 4", func(t *testing.T) {
		runSteps(t, []step{
			{op: "setData", args: []string{"largeRecord", "field9", "initialValue9"}, want: nil},
			{op: "setData", args: []string{"largeRecord", "field10", "initialValue10"}, want: nil},
			{op: "setData", args: []string{"largeRecord", "field100", "initialValue100"}, want: nil},
			{op: "setData", args: []string{"largeRecord", "field250", "initialValue250"}, want: nil},
			{op: "setData", args: []string{"largeRecord", "field750", "initialValue750"}, want: nil},
			{op: "setData", args: []string{"largeRecord", "field999", "initialValue999"}, want: nil},
			{op: "setData", args: []string{"largeRecord", "field10", "overwrittenValue10"}, want: nil},
			{op: "setData", args: []string{"largeRecord", "field100", "overwrittenValue100"}, want: nil},
			{op: "getData", args: []string{"largeRecord", "field10"}, want: "overwrittenValue10"},
			{op: "getData", args: []string{"largeRecord", "field9"}, want: "initialValue9"},
			{op: "getData", args: []string{"largeRecord", "field100"}, want: "overwrittenValue100"},
			{op: "getData", args: []string{"largeRecord", "field999"}, want: "initialValue999"},
			{op: "deleteData", args: []string{"largeRecord", "field250"}, want: true},
			{op: "deleteData", args: []string{"largeRecord", "field750"}, want: true},
			{op: "getData", args: []string{"largeRecord", "field250"}, want: ""},
			{op: "getData", args: []string{"largeRecord", "field750"}, want: ""},
		})
	})

	t.Run("Example 5", func(t *testing.T) {
		runSteps(t, []step{
			{op: "getData", args: []string{"nonExistent", "field"}, want: ""},
			{op: "deleteData", args: []string{"nonExistent", "field"}, want: false},
			{op: "setData", args: []string{"", "", ""}, want: nil},
			{op: "getData", args: []string{"", ""}, want: ""},
			{op: "deleteData", args: []string{"", ""}, want: true},
			{op: "setData", args: []string{"testKey", "nullField", "null"}, want: nil},
			{op: "getData", args: []string{"testKey", "nullField"}, want: "null"},
			{op: "setData", args: []string{"existingKey", "field1", "value1"}, want: nil},
			{op: "deleteData", args: []string{"existingKey", "nonExistentField"}, want: false},
			{op: "getData", args: []string{"existingKey", "field1"}, want: "value1"},
			{op: "deleteData", args: []string{"existingKey", "field1"}, want: true},
			{op: "getData", args: []string{"existingKey", "field1"}, want: ""},
			{op: "setData", args: []string{"key@#$%", "field!@#", "value~`"}, want: nil},
			{op: "getData", args: []string{"key@#$%", "field!@#"}, want: "value~`"},
		})
	})

	t.Run("Follow-up 1", func(t *testing.T) {
		runSteps(t, []step{
			{op: "setData", args: []string{"A", "BC", "E"}, want: nil},
			{op: "setData", args: []string{"A", "BD", "F"}, want: nil},
			{op: "setData", args: []string{"A", "C", "G"}, want: nil},
			{op: "scanDataByPrefix", args: []string{"A", "B"}, want: []string{"BC(E)", "BD(F)"}},
			{op: "scanData", args: []string{"A"}, want: []string{"BC(E)", "BD(F)", "C(G)"}},
			{op: "scanDataByPrefix", args: []string{"B", "B"}, want: []string{}},
		})
	})

	t.Run("Follow-up 2", func(t *testing.T) {
		runSteps(t, []step{
			{op: "setDataAtWithTtl", args: []string{"A", "BC", "E", "1", "9"}, want: nil},
			{op: "setDataAtWithTtl", args: []string{"A", "BC", "E", "5", "10"}, want: nil},
			{op: "setDataAt", args: []string{"A", "BD", "F", "5"}, want: nil},
			{op: "scanDataByPrefixAt", args: []string{"A", "", "14"}, want: []string{"BC(E)", "BD(F)"}},
			{op: "scanDataByPrefixAt", args: []string{"A", "", "15"}, want: []string{"BD(F)"}},
		})
	})

	t.Run("Follow-up 3", func(t *testing.T) {
		runSteps(t, []step{
			{op: "setDataAtWithTtl", args: []string{"A", "B", "C", "1", "10"}, want: nil},
			{op: "backup", args: []string{"3"}, want: 1},
			{op: "setDataAt", args: []string{"A", "D", "E", "4"}, want: nil},
			{op: "backup", args: []string{"5"}, want: 1},
			{op: "deleteDataAt", args: []string{"A", "B", "8"}, want: true},
			{op: "backup", args: []string{"9"}, want: 1},
			{op: "restore", args: []string{"10", "7"}, want: nil},
			{op: "backup", args: []string{"11"}, want: 1},
			{op: "scanDataAt", args: []string{"A", "15"}, want: []string{"B(C)", "D(E)"}},
			{op: "scanDataAt", args: []string{"A", "16"}, want: []string{"D(E)"}},
		})
	})
}
