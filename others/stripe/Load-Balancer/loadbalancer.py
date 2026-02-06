# import heapq

# class LoadBalancer:
#     def __init__(self, servers):
#         self.servers = servers
#         self.weights = {server: 0 for server in servers}
#         self.heap = [(0, server) for server in servers]
#         heapq.heapify(self.heap)

#     def route(self, weight):
#         # Pop the server with the lowest weight
#         min_weight, server = heapq.heappop(self.heap)

#         # Update the weight
#         new_weight = min_weight + weight
#         self.weights[server] = new_weight

#         # Push the updated server weight back into the heap
#         heapq.heappush(self.heap, (new_weight, server))

#         return server

# # Example usage
# servers = ['a', 'b', 'c']
# lb = LoadBalancer(servers)

# print(lb.route(weight=2))  # Returns 'a'
# print(lb.route(weight=1))  # Returns 'b'
# print(lb.route(weight=1))  # Returns 'c'
# print(lb.route(weight=1))  # Returns 'b'

import heapq
import time

class LoadBalancer:
    def __init__(self, servers):
        self.servers = servers
        self.weights = {server: 0 for server in servers}
        self.server_heap = [(0, server) for server in servers]
        heapq.heapify(self.server_heap)
        
        # Dictionary to store active tasks with their expiry times for each server
        # self.active_tasks = {server: [] for server in servers} 

    def route(self, weight, ttl):
        # Clean up expired tasks before routing
        # self._cleanup_expired_tasks()

        # Pop the server with the lowest effective weight
        min_weight, server = heapq.heappop(self.server_heap)

        # Calculate the expiration time for the task
        # expiry_time = time.time() + ttl

        # Update the active task list and weight for this server
        # self.active_tasks[server].append((weight, expiry_time))
        new_weight = min_weight + weight
        # self.weights[server] = new_weight

        # Push the updated server back into the heap
        heapq.heappush(self.server_heap, (new_weight, server))

        return server

    def _cleanup_expired_tasks(self):
        # Get current time
        current_time = time.time()

        for server in self.servers:
            # Get active tasks for the server
            new_task_list = []
            total_weight = 0

            for weight, expiry_time in self.active_tasks[server]:
                if expiry_time > current_time:
                    # Task is still active
                    new_task_list.append((weight, expiry_time))
                    total_weight += weight

            # Update the server's weight and active tasks list
            self.active_tasks[server] = new_task_list
            self.weights[server] = total_weight

        # Rebuild the heap to reflect updated weights
        self.server_heap = [(self.weights[server], server) for server in self.servers]
        heapq.heapify(self.server_heap)

# Example usage
lb = LoadBalancer(['server1', 'server2', 'server3'])

# Route requests
print(lb.route(weight=5, ttl=10))  # Expect: server1 (initial lowest)
print(lb.route(weight=3, ttl=5))   # Expect: server2 (lowest after server1)
print(lb.route(weight=2, ttl=20))  # Expect: server3 (lowest after server2)

# After a few seconds (before TTLs expire), routing should account for the weights
time.sleep(3)
print(lb.route(weight=1, ttl=10))  # Expect: server3 (since it has the lowest weight at 2)

# After 5 seconds (server2's first request should expire)
time.sleep(5)
print(lb.route(weight=1, ttl=10))  # Expect: server2 (server2's weight resets after its task expires)

# After 10 seconds (all weights should expire except for server3's task)
time.sleep(5)
print(lb.route(weight=1, ttl=10))  # Expect: server1 (since server1's task expires first)
