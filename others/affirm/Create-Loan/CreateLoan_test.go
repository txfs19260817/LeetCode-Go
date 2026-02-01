package affirm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoanServiceCreateLoan(t *testing.T) {
	relations := map[string][]string{
		"AA": {"BB", "CC"},
		"DD": {"AA"},
		"EE": {},
	}

	service, err := NewLoanService(relations)
	assert.NoError(t, err)

	loan, err := service.CreateLoan("cust1", "CC")
	assert.NoError(t, err)
	assert.Equal(t, "DD", loan.LoanCompany)

	loan, err = service.CreateLoan("cust2", "AA")
	assert.NoError(t, err)
	assert.Equal(t, "DD", loan.LoanCompany)

	loan, err = service.CreateLoan("cust3", "EE")
	assert.NoError(t, err)
	assert.Equal(t, "EE", loan.LoanCompany)

	loan, err = service.CreateLoan("cust4", "ZZ")
	assert.NoError(t, err)
	assert.Equal(t, "ZZ", loan.LoanCompany)
}

func TestBuildParentMapMultipleParents(t *testing.T) {
	relations := map[string][]string{
		"AA": {"BB"},
		"CC": {"BB"},
	}
	_, err := buildParentMap(relations)
	assert.Error(t, err)
}

func TestFindRootCycle(t *testing.T) {
	relations := map[string][]string{
		"AA": {"BB"},
		"BB": {"AA"},
	}
	parentOf, err := buildParentMap(relations)
	assert.NoError(t, err)

	_, err = findRoot("AA", parentOf)
	assert.Error(t, err)
}
