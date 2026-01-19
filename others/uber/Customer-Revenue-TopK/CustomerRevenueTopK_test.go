package uber

import (
	"reflect"
	"testing"
)

func TestCustomerRevenueTopK(t *testing.T) {
	t.Run("Basic add and referral", func(t *testing.T) {
		rt := NewRevenueTracker()

		id1 := rt.Add(10)
		if id1 != 1 {
			t.Fatalf("Add(10) id = %d, want 1", id1)
		}
		id2 := rt.Add(20)
		if id2 != 2 {
			t.Fatalf("Add(20) id = %d, want 2", id2)
		}
		id3 := rt.AddByReferral(30, 1)
		if id3 != 3 {
			t.Fatalf("AddByReferral(30, 1) id = %d, want 3", id3)
		}

		if got := rt.ShowRevenue(1); got != 40 {
			t.Fatalf("ShowRevenue(1) = %d, want 40", got)
		}
		if got := rt.ShowRevenue(3); got != 10 {
			t.Fatalf("ShowRevenue(3) = %d, want 10", got)
		}

		got := rt.TopSmallestKCustomer(9, 2)
		want := []int{3, 2}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("TopSmallestKCustomer(9, 2) = %v, want %v", got, want)
		}

		got = rt.TopSmallestKCustomer(10, 2)
		want = []int{2, 1}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("TopSmallestKCustomer(10, 2) = %v, want %v", got, want)
		}
	})

	t.Run("Tie breaks by ID", func(t *testing.T) {
		rt := NewRevenueTracker()
		id1 := rt.Add(10)
		id2 := rt.Add(10)
		id3 := rt.Add(10)

		if id1 != 1 || id2 != 2 || id3 != 3 {
			t.Fatalf("expected IDs 1,2,3 got %d,%d,%d", id1, id2, id3)
		}

		got := rt.TopSmallestKCustomer(9, 2)
		want := []int{1, 2}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("TopSmallestKCustomer(9, 2) = %v, want %v", got, want)
		}
	})

	t.Run("Threshold excludes equals", func(t *testing.T) {
		rt := NewRevenueTracker()
		rt.Add(10)
		rt.Add(20)
		rt.Add(30)

		got := rt.TopSmallestKCustomer(20, 5)
		want := []int{3}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("TopSmallestKCustomer(20, 5) = %v, want %v", got, want)
		}
	})

	t.Run("Referral uses pre-increase revenue for new customer", func(t *testing.T) {
		rt := NewRevenueTracker()
		rt.Add(10)
		rt.AddByReferral(5, 1)  // customer 2 should be 10
		rt.AddByReferral(20, 1) // customer 3 should be 15

		if got := rt.ShowRevenue(1); got != 35 {
			t.Fatalf("ShowRevenue(1) = %d, want 35", got)
		}
		if got := rt.ShowRevenue(2); got != 10 {
			t.Fatalf("ShowRevenue(2) = %d, want 10", got)
		}
		if got := rt.ShowRevenue(3); got != 15 {
			t.Fatalf("ShowRevenue(3) = %d, want 15", got)
		}
	})
}
