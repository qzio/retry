# RETRY


```
RETRIES=5 BACKOFF=10 retry ./my-flaky-script.sh
```

`retry` retries a command a configurable amount of times, with configurable backoff in seconds.

The defaults are 3 retries and 3s backoff
