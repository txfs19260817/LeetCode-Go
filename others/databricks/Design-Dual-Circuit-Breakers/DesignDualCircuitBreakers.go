package databricks

// Server has a fixed list of boolean outcomes.
type Server struct {
	outcomes []bool
}

// NewServer creates a server with the given outcomes.
func NewServer(outcomes []bool) *Server {
	return &Server{outcomes: outcomes}
}

// Handle returns outcomes[requestId].
func (s *Server) Handle(requestId int) bool {
	return s.outcomes[requestId]
}

// CircuitBreaker wraps a server with failure/reset thresholds.
type CircuitBreaker struct {
	server              *Server
	failureThreshold    int
	resetThreshold      int
	isOpen              bool
	consecutiveFailures int
	rejectedCount       int
}

// NewCircuitBreaker creates a circuit breaker for the given server.
func NewCircuitBreaker(server *Server, failureThreshold, resetThreshold int) *CircuitBreaker {
	return &CircuitBreaker{
		server:           server,
		failureThreshold: failureThreshold,
		resetThreshold:   resetThreshold,
	}
}

// Gateway routes requests through primary and secondary circuit breakers.
type Gateway struct {
	primary   *CircuitBreaker
	secondary *CircuitBreaker
}

// NewGateway creates a gateway with the given breakers.
func NewGateway(primary, secondary *CircuitBreaker) *Gateway {
	return &Gateway{primary: primary, secondary: secondary}
}

// RouteRequests processes requests 0..totalRequests-1 and returns routing strings.
func (g *Gateway) RouteRequests(totalRequests int) []string {
	results := make([]string, totalRequests)

	for i := 0; i < totalRequests; i++ {
		primaryAttempted := false
		secondaryAttempted := false

		// Step 1-2: evaluate primary breaker
		if !g.primary.isOpen {
			// Primary closed → attempt server
			primaryAttempted = true
			success := g.primary.server.Handle(i)
			if success {
				g.primary.consecutiveFailures = 0
				results[i] = "Primary"
				continue // done, no need for secondary
			}
			// Failure path
			g.primary.consecutiveFailures++
			if g.primary.consecutiveFailures >= g.primary.failureThreshold {
				g.primary.isOpen = true
				g.primary.rejectedCount = 0
			}
			// fall through to try secondary
		} else {
			// Primary open → increment rejected, maybe close
			g.primary.rejectedCount++
			if g.primary.rejectedCount >= g.primary.resetThreshold {
				g.primary.isOpen = false
				g.primary.consecutiveFailures = 0
			}
			// fall through to try secondary
		}

		// Step 3-4: evaluate secondary breaker
		if !g.secondary.isOpen {
			// Secondary closed → attempt server
			secondaryAttempted = true
			success := g.secondary.server.Handle(i)
			if success {
				g.secondary.consecutiveFailures = 0
			} else {
				g.secondary.consecutiveFailures++
				if g.secondary.consecutiveFailures >= g.secondary.failureThreshold {
					g.secondary.isOpen = true
					g.secondary.rejectedCount = 0
				}
			}
		} else {
			// Secondary open → increment rejected, maybe close
			g.secondary.rejectedCount++
			if g.secondary.rejectedCount >= g.secondary.resetThreshold {
				g.secondary.isOpen = false
				g.secondary.consecutiveFailures = 0
			}
		}

		// Determine result string
		if primaryAttempted && secondaryAttempted {
			results[i] = "Primary -> Secondary"
		} else if primaryAttempted {
			results[i] = "Primary"
		} else if secondaryAttempted {
			results[i] = "Secondary"
		} else {
			results[i] = "Rejected"
		}
	}

	return results
}
