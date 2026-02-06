from collections import deque
from time import time
from typing import Dict, Deque, Tuple


class RateLimiter:
    def __init__(self, max_requests: int = 5, window_seconds: int = 2):
        """
        Initialize rate limiter with configurable parameters

        Args:
            max_requests: Maximum number of requests allowed in the time window
            window_seconds: Time window in seconds
        """
        self.max_requests = max_requests
        self.window_seconds = window_seconds
        # Dictionary to store request timestamps for each user
        self.user_requests: Dict[str, Deque[float]] = {}

    def is_allowed(self, user_id: str) -> bool:
        """
        Check if a request from a user is allowed

        Args:
            user_id: Unique identifier for the user

        Returns:
            bool: True if request is allowed, False if rate limit exceeded
        """
        current_time = time()

        # Initialize queue for new users
        if user_id not in self.user_requests:
            self.user_requests[user_id] = deque()

        requests = self.user_requests[user_id]

        # Remove timestamps outside the window
        while requests and current_time - requests[0] > self.window_seconds:
            requests.popleft()

        # Check if request limit is reached
        if len(requests) >= self.max_requests:
            return False

        # Add current request timestamp
        requests.append(current_time)
        return True

    def get_remaining_requests(self, user_id: str) -> Tuple[int, float]:
        """
        Get remaining requests and time until window reset

        Args:
            user_id: Unique identifier for the user

        Returns:
            Tuple containing:
                - Number of remaining requests
                - Seconds until oldest request expires
        """
        if user_id not in self.user_requests:
            return self.max_requests, 0.0

        current_time = time()
        requests = self.user_requests[user_id]

        # Remove expired timestamps
        while requests and current_time - requests[0] > self.window_seconds:
            requests.popleft()

        if not requests:
            return self.max_requests, 0.0

        remaining = self.max_requests - len(requests)
        time_until_reset = max(0, self.window_seconds - (current_time - requests[0]))

        return remaining, time_until_reset


# Example usage
if __name__ == "__main__":
    # Create rate limiter: 5 requests per 2 seconds
    limiter = RateLimiter(max_requests=5, window_seconds=2)

    # Simulate requests
    user_id = "user123"

    for i in range(7):
        allowed = limiter.is_allowed(user_id)
        remaining, reset_time = limiter.get_remaining_requests(user_id)
        print(
            f"Request {i+1}: Allowed: {allowed}, Remaining: {remaining}, Reset in: {reset_time:.2f}s"
        )
