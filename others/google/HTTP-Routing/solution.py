class HttpRouter:
    class TrieNode:
        def __init__(self):
            self.children = {}  # str -> TrieNode
            self.param_name = None
            self.param_child = None
            self.pattern = None

    class MatchResult:
        def __init__(self, pattern, params):
            self.pattern = pattern
            self.params = params

        def __str__(self):
            return f"MatchResult({self.pattern}, {self.params})"

    def __init__(self):
        self.root = self.TrieNode()

    def register_handler(self, path: str) -> None:
        path_segments = [p for p in path.split("/") if p]
        cur = self.root
        for segment in path_segments:
            if segment.startswith("{"):
                if cur.param_name is None:
                    cur.param_name = segment
                    cur.param_child = self.TrieNode()
                elif cur.param_name != segment:
                    raise ValueError(
                        f"Parameter segment mismatch: {cur.param_name} != {segment}"
                    )
                cur = cur.param_child
            else:
                cur = cur.children.setdefault(
                    segment, self.TrieNode()
                )  # if segment not in children, add it
        cur.pattern = path  # set the pattern for the current node

    def get_handler(self, path: str) -> MatchResult | None:
        path_segments = [p for p in path.split("/") if p]
        cur = self.root
        params: list[str] = []
        for segment in path_segments:
            static_child = cur.children.get(segment)
            if static_child is not None:
                cur = static_child
            elif cur.param_child is not None:
                cur = cur.param_child
                params.append(segment)
            else:
                return None
        return self.MatchResult(cur.pattern, params) if cur.pattern else None


if __name__ == "__main__":
    s = HttpRouter()
    s.register_handler("/home")
    s.register_handler("/users/new")
    s.register_handler("/users/{id}")
    s.register_handler("/users/{id}/pics/{picId}")
    assert str(s.get_handler("/home")) == "MatchResult(/home, [])"
    assert str(s.get_handler("/users/new")) == "MatchResult(/users/new, [])"
    assert str(s.get_handler("/users/101")) == "MatchResult(/users/{id}, ['101'])"
    assert (
        str(s.get_handler("/users/101/pics/7"))
        == "MatchResult(/users/{id}/pics/{picId}, ['101', '7'])"
    )
    assert s.get_handler("/users/") is None
    assert s.get_handler("/users") is None
