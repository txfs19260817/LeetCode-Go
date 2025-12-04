package leetcode

import (
	"reflect"
	"testing"
)

func TestThroneInheritance(t *testing.T) {
	t.Run("Example 1: Basic inheritance with death", func(t *testing.T) {
		obj := Constructor("king")
		obj.Birth("king", "andy")
		obj.Birth("king", "bob")
		obj.Birth("king", "catherine")
		obj.Birth("andy", "matthew")
		obj.Birth("bob", "alex")
		obj.Birth("bob", "asha")

		got1 := obj.GetInheritanceOrder()
		want1 := []string{"king", "andy", "matthew", "bob", "alex", "asha", "catherine"}
		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("GetInheritanceOrder() before death = %v, want %v", got1, want1)
		}

		obj.Death("bob")

		got2 := obj.GetInheritanceOrder()
		want2 := []string{"king", "andy", "matthew", "alex", "asha", "catherine"}
		if !reflect.DeepEqual(got2, want2) {
			t.Errorf("GetInheritanceOrder() after death = %v, want %v", got2, want2)
		}
	})

	t.Run("Single king with no children", func(t *testing.T) {
		obj := Constructor("king")
		got := obj.GetInheritanceOrder()
		want := []string{"king"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetInheritanceOrder() = %v, want %v", got, want)
		}
	})

	t.Run("King dies", func(t *testing.T) {
		obj := Constructor("king")
		obj.Birth("king", "prince")
		obj.Death("king")
		got := obj.GetInheritanceOrder()
		want := []string{"prince"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetInheritanceOrder() = %v, want %v", got, want)
		}
	})

	t.Run("Multiple generations", func(t *testing.T) {
		obj := Constructor("king")
		obj.Birth("king", "child1")
		obj.Birth("king", "child2")
		obj.Birth("child1", "grandchild1")
		obj.Birth("child1", "grandchild2")
		obj.Birth("child2", "grandchild3")

		got := obj.GetInheritanceOrder()
		want := []string{"king", "child1", "grandchild1", "grandchild2", "child2", "grandchild3"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetInheritanceOrder() = %v, want %v", got, want)
		}
	})
}
