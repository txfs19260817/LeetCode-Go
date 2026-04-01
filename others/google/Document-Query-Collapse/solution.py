from __future__ import annotations

from collections import defaultdict, deque


def find_collapsed_documents(
    doc_to_queries: dict[str, list[str]],
    collapsed_queries: list[str],
) -> list[str]:
    # 反向建立 query -> docs 的索引，方便从损坏 query 找到受影响的文件。
    query_to_docs: dict[str, list[str]] = defaultdict(list)
    for doc, queries in doc_to_queries.items():
        for query in queries:
            query_to_docs[query].append(doc)

    bad_queries = set(collapsed_queries)
    bad_docs: set[str] = set()
    queue = deque(("query", query) for query in collapsed_queries)

    while queue:
        item_type, name = queue.popleft()

        if item_type == "query":
            # 损坏 query 会让所有包含它的 doc 损坏。
            for doc in query_to_docs.get(name, []):
                if doc in bad_docs:
                    continue
                bad_docs.add(doc)
                queue.append(("doc", doc))
        else:
            # 损坏 doc 会让其中所有 query 一起损坏。
            for query in doc_to_queries.get(name, []):
                if query in bad_queries:
                    continue
                bad_queries.add(query)
                queue.append(("query", query))

    return sorted(bad_docs)


if __name__ == "__main__":
    assert find_collapsed_documents(
        {
            "Doc1": ["Q1", "Q2"],
            "Doc2": ["Q2", "Q3"],
            "Doc3": ["Q4", "Q5"],
        },
        ["Q1"],
    ) == ["Doc1", "Doc2"]

    assert find_collapsed_documents(
        {
            "Doc1": ["Q1", "Q2"],
            "Doc2": ["Q2", "Q3"],
            "Doc3": ["Q3", "Q4"],
        },
        ["Q1"],
    ) == ["Doc1", "Doc2", "Doc3"]

    assert find_collapsed_documents(
        {
            "Doc1": ["Q1"],
            "Doc2": ["Q2"],
        },
        [],
    ) == []
