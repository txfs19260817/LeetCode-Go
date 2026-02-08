package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteRequests_MainExample(t *testing.T) {
	primary := NewServer([]bool{true, false, false, true, true, false, true})
	secondary := NewServer([]bool{false, true, false, false, true, true, true})
	primaryBreaker := NewCircuitBreaker(primary, 2, 2)
	secondaryBreaker := NewCircuitBreaker(secondary, 2, 2)
	gw := NewGateway(primaryBreaker, secondaryBreaker)

	result := gw.RouteRequests(7)
	expected := []string{
		"Primary",              // Req 0: primary success
		"Primary -> Secondary", // Req 1: primary fail, secondary success
		"Primary -> Secondary", // Req 2: primary fail→opens, secondary fail
		"Secondary",            // Req 3: primary open rej=1, secondary fail→opens
		"Rejected",             // Req 4: primary open rej=2→closes, secondary open rej=1
		"Primary",              // Req 5: primary fail, secondary open rej=2→closes
		"Primary",              // Req 6: primary success
	}
	assert.Equal(t, expected, result)
}

func TestRouteRequests_AllPrimarySuccess(t *testing.T) {
	primary := NewServer([]bool{true, true, true, true, true})
	secondary := NewServer([]bool{false, false, false, false, false})
	primaryBreaker := NewCircuitBreaker(primary, 2, 2)
	secondaryBreaker := NewCircuitBreaker(secondary, 2, 2)
	gw := NewGateway(primaryBreaker, secondaryBreaker)

	result := gw.RouteRequests(5)
	expected := []string{"Primary", "Primary", "Primary", "Primary", "Primary"}
	assert.Equal(t, expected, result)
}

func TestRouteRequests_BothAlwaysFail(t *testing.T) {
	primary := NewServer([]bool{false, false, false, false, false, false, false, false})
	secondary := NewServer([]bool{false, false, false, false, false, false, false, false})
	primaryBreaker := NewCircuitBreaker(primary, 2, 2)
	secondaryBreaker := NewCircuitBreaker(secondary, 2, 2)
	gw := NewGateway(primaryBreaker, secondaryBreaker)

	result := gw.RouteRequests(8)
	expected := []string{
		"Primary -> Secondary", // Req 0: pri fail(1), sec fail(1)
		"Primary -> Secondary", // Req 1: pri fail(2)→opens, sec fail(2)→opens
		"Rejected",             // Req 2: pri rej=1, sec rej=1
		"Rejected",             // Req 3: pri rej=2→closes, sec rej=2→closes
		"Primary -> Secondary", // Req 4: pri fail(1), sec fail(1)
		"Primary -> Secondary", // Req 5: pri fail(2)→opens, sec fail(2)→opens
		"Rejected",             // Req 6: pri rej=1, sec rej=1
		"Rejected",             // Req 7: pri rej=2→closes, sec rej=2→closes
	}
	assert.Equal(t, expected, result)
}
