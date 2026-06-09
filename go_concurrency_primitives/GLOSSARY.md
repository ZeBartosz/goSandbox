# Go Concurrency Glossary

This glossary will capture terms only after they are understood well enough to use accurately. For now, it is intentionally sparse.

## Terms

**WaitGroup**:
A synchronization counter from `sync` used to block until a known number of tasks have called `Done`.
_Avoid_: Thread group, task runner

**Blocked**:
Paused at an operation that cannot complete yet, such as receiving from an empty channel or waiting on a non-zero `WaitGroup`.
_Avoid_: Running in the background

**Closed channel**:
A channel state meaning no more values may be sent; receivers can continue receiving already buffered values, then receive the zero value with `ok == false`.
_Avoid_: Deleted channel, emptied channel

**Fan-in**:
A concurrency pattern where multiple goroutines send results into one shared output channel.
_Avoid_: Merge threads

**Mutex**:
A lock that protects a critical section so only one goroutine can access shared state there at a time.
_Avoid_: Thread blocker, global pause

**Atomic operation**:
A single synchronized operation on one value that cannot be interleaved with another atomic operation on that same value.
_Avoid_: Tiny mutex, magic safe variable

**select**:
A control statement that waits until one channel operation is ready, then runs one matching case.
_Avoid_: Wait for all channels, run all ready cases

**default case in select**:
A branch that runs immediately when no channel case is ready, making that select non-blocking.
_Avoid_: Timeout, else after waiting

**nil channel**:
A channel variable with no channel value assigned. Sends and receives on a nil channel are never ready and block forever unless avoided by a `select` default or disabled case logic.
_Avoid_: Closed channel, empty channel

