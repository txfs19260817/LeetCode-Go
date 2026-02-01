from dataclasses import dataclass
from typing import Dict, List, Set, Tuple


@dataclass
class Loan:
    customer_id: str
    requested_company: str
    loan_company: str


def build_parent_map(relations: Dict[str, List[str]]) -> Dict[str, str]:
    """
    relations: parent -> [child1, child2, ...]
    returns: parent_of[child] = parent
    """
    parent_of: Dict[str, str] = {}
    for parent, children in relations.items():
        for child in children:
            if child in parent_of and parent_of[child] != parent:
                raise ValueError(
                    f"Child {child} has multiple parents: {parent_of[child]} and {parent}"
                )
            parent_of[child] = parent
    return parent_of


def find_root(company: str, parent_of: Dict[str, str]) -> str:
    """
    Returns topmost parent (root) of company.
    Path-compress along the way.
    """
    seen: Set[str] = set()
    path: List[str] = []
    cur = company

    while cur in parent_of:
        if cur in seen:
            raise ValueError(f"Cycle detected in parent chain at {cur}")
        seen.add(cur)
        path.append(cur)
        cur = parent_of[cur]

    root = cur
    for node in path:
        parent_of[node] = root

    return root


def create_loan(
    customer_id: str, requested_company: str, parent_of: Dict[str, str]
) -> Loan:
    loan_company = find_root(requested_company, parent_of)
    return Loan(
        customer_id=customer_id,
        requested_company=requested_company,
        loan_company=loan_company,
    )


if __name__ == "__main__":
    relations = {
        "AA": ["BB", "CC"],
        "DD": ["AA"],
        "EE": [],
    }
    parent_of = build_parent_map(relations)

    loans = [
        create_loan("cust1", "CC", parent_of),
        create_loan("cust2", "AA", parent_of),
        create_loan("cust3", "EE", parent_of),
        create_loan("cust4", "ZZ", parent_of),
    ]

    for loan in loans:
        print(
            "Loan(customer={0}, requested={1}, loan_company={2})".format(
                loan.customer_id, loan.requested_company, loan.loan_company
            )
        )

    assert loans[0].loan_company == "DD"
    assert loans[1].loan_company == "DD"
    assert loans[2].loan_company == "EE"
    assert loans[3].loan_company == "ZZ"
