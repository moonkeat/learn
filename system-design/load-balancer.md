## Load Balance Algorithm

### Round Robin

- Distribute incoming requests to servers in cyclic order.

Use case:
- All servers have similar capabilities and performance
- Stateless applications

Pros:
- Equal distribution of requests
- Easy to implement
- Works well with server with same capabilities

Cons:
- No load awareness (all servers treated equally)
- No session affinity (problematic for stateful applications)
- Potentially be exploited by attackers by predicting which server will handle their requests

### Weighted Round Robin

- Assign weights to each server based on capacity or performance

Use cases:
- Heterogeneous server environment (servers with different capabilities)\
- Database clusters where nodes have different processing power

Pros:
- Load distribution according to server capacity
- Prevent overloading less powerful servers

Cons:
- Complexity in weight assignment (requires accurate metrics)
- Increased overhead
- Not ideal for highly variable loads

### IP Hash

Use cases:
- Stateful applications
- Geographically distributed clients

### Least Response Time

Use cases:
- Real time applications
- Web services that requires quick responses

### Least Bandwidth

Use cases:
- Real time applications
- High Bandwidth Applications
- CDN
