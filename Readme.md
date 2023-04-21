# How a server handles concurrent TCP requests?

1. Starts listening on a port
2. Wait for a client to connect `listener.Accept()` - this is a blocking call
3. Read, Write and Close the connection
4. To serve concurrent requests - create a new thread to serve that request


## Improvements:

1. Limiting the number of threads
2. Add the thread pool to save thread creation time
3. Add timeouts to the connection objects [think `context`]
4. TCP backlog queue config

------------------