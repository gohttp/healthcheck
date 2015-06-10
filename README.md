# healthcheck

Health check endpoint for your HTTP server.

By default endpoint is located at `/internal/health"`.

# Example

```
// ...
healthcheck.AddHandler()

http.ListenAndServe(addr, nil)
```

# License

MIT
