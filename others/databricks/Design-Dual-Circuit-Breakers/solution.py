class Server:
    def __init__(self, outcomes: list[bool]):
        self.outcomes = outcomes

    def handle(self, request_id: int) -> bool:
        return self.outcomes[request_id]


class CircuitBreaker:
    def __init__(self, server: Server, failure_threshold: int, reset_threshold: int):
        self.server = server
        self.failure_threshold = failure_threshold
        self.reset_threshold = reset_threshold
        self.is_open = False
        self.consecutive_failures = 0
        self.rejected_count = 0


class Gateway:
    def __init__(self, primary: CircuitBreaker, secondary: CircuitBreaker):
        self.primary = primary
        self.secondary = secondary

    def route_requests(self, total_requests: int) -> list[str]:
        results: list[str] = []

        for i in range(total_requests):
            primary_attempted = False
            secondary_attempted = False

            # Evaluate primary breaker
            if not self.primary.is_open:
                primary_attempted = True
                success = self.primary.server.handle(i)
                if success:
                    self.primary.consecutive_failures = 0
                    results.append("Primary")
                    continue
                # Failure
                self.primary.consecutive_failures += 1
                if self.primary.consecutive_failures >= self.primary.failure_threshold:
                    self.primary.is_open = True
                    self.primary.rejected_count = 0
            else:
                self.primary.rejected_count += 1
                if self.primary.rejected_count >= self.primary.reset_threshold:
                    self.primary.is_open = False
                    self.primary.consecutive_failures = 0

            # Evaluate secondary breaker
            if not self.secondary.is_open:
                secondary_attempted = True
                success = self.secondary.server.handle(i)
                if success:
                    self.secondary.consecutive_failures = 0
                else:
                    self.secondary.consecutive_failures += 1
                    if self.secondary.consecutive_failures >= self.secondary.failure_threshold:
                        self.secondary.is_open = True
                        self.secondary.rejected_count = 0
            else:
                self.secondary.rejected_count += 1
                if self.secondary.rejected_count >= self.secondary.reset_threshold:
                    self.secondary.is_open = False
                    self.secondary.consecutive_failures = 0

            # Determine result
            if primary_attempted and secondary_attempted:
                results.append("Primary -> Secondary")
            elif primary_attempted:
                results.append("Primary")
            elif secondary_attempted:
                results.append("Secondary")
            else:
                results.append("Rejected")

        return results


if __name__ == "__main__":
    # Test 1: Main example
    primary = Server([True, False, False, True, True, False, True])
    secondary = Server([False, True, False, False, True, True, True])
    pb = CircuitBreaker(primary, 2, 2)
    sb = CircuitBreaker(secondary, 2, 2)
    gw = Gateway(pb, sb)
    result = gw.route_requests(7)
    assert result == [
        "Primary",
        "Primary -> Secondary",
        "Primary -> Secondary",
        "Secondary",
        "Rejected",
        "Primary",
        "Primary",
    ], f"Test 1 failed: {result}"

    # Test 2: All primary successes
    primary2 = Server([True, True, True, True, True])
    secondary2 = Server([False, False, False, False, False])
    pb2 = CircuitBreaker(primary2, 2, 2)
    sb2 = CircuitBreaker(secondary2, 2, 2)
    gw2 = Gateway(pb2, sb2)
    result2 = gw2.route_requests(5)
    assert result2 == ["Primary"] * 5, f"Test 2 failed: {result2}"

    # Test 3: Both always fail
    primary3 = Server([False] * 8)
    secondary3 = Server([False] * 8)
    pb3 = CircuitBreaker(primary3, 2, 2)
    sb3 = CircuitBreaker(secondary3, 2, 2)
    gw3 = Gateway(pb3, sb3)
    result3 = gw3.route_requests(8)
    assert result3 == [
        "Primary -> Secondary",
        "Primary -> Secondary",
        "Rejected",
        "Rejected",
        "Primary -> Secondary",
        "Primary -> Secondary",
        "Rejected",
        "Rejected",
    ], f"Test 3 failed: {result3}"

    print("All tests passed!")
