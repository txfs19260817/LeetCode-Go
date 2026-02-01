package affirm

import "fmt"

/*
Create loans based on a company hierarchy. Given a mapping of parent -> children, assign each
loan to the topmost parent company of the requested company. If a company has no parent, it is
its own topmost parent.
*/

type Loan struct {
	CustomerID       string
	RequestedCompany string
	LoanCompany      string
}

type LoanService struct {
	parentOf map[string]string
}

func NewLoanService(relations map[string][]string) (*LoanService, error) {
	parentOf, err := buildParentMap(relations)
	if err != nil {
		return nil, err
	}
	return &LoanService{parentOf: parentOf}, nil
}

func (ls *LoanService) CreateLoan(customerID, requestedCompany string) (Loan, error) {
	root, err := findRoot(requestedCompany, ls.parentOf)
	if err != nil {
		return Loan{}, err
	}
	return Loan{
		CustomerID:       customerID,
		RequestedCompany: requestedCompany,
		LoanCompany:      root,
	}, nil
}

func buildParentMap(relations map[string][]string) (map[string]string, error) {
	parentOf := make(map[string]string)
	for parent, children := range relations {
		for _, child := range children {
			if existing, ok := parentOf[child]; ok && existing != parent {
				return nil, fmt.Errorf("child %s has multiple parents: %s and %s", child, existing, parent)
			}
			parentOf[child] = parent
		}
	}
	return parentOf, nil
}

func findRoot(company string, parentOf map[string]string) (string, error) {
	seen := make(map[string]struct{})
	path := make([]string, 0, 8)
	cur := company

	for {
		parent, ok := parentOf[cur]
		if !ok {
			root := cur
			for _, node := range path {
				parentOf[node] = root
			}
			return root, nil
		}

		if _, ok := seen[cur]; ok {
			return "", fmt.Errorf("cycle detected in parent chain at %s", cur)
		}
		seen[cur] = struct{}{}
		path = append(path, cur)
		cur = parent
	}
}
