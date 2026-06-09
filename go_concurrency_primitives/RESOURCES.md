# Go Concurrency Resources

## Knowledge

- [Go Tour: Concurrency](https://go.dev/tour/concurrency/1)
  Official beginner-friendly introduction to goroutines, channels, buffering, select, and mutexes. Use for: first-pass mental models.
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
  Official idiomatic guidance, including the principle that goroutines are cheap and channels are communication mechanisms. Use for: style and idioms.
- [Package `sync` documentation](https://pkg.go.dev/sync)
  Official API docs for `WaitGroup`, `Mutex`, `RWMutex`, `Once`, and related primitives. Use for: exact method contracts and caveats.
- [Package `sync/atomic` documentation](https://pkg.go.dev/sync/atomic)
  Official API docs for atomic operations and typed atomics. Use for: counters, flags, and low-level synchronization.
- [The Go Memory Model](https://go.dev/ref/mem)
  Official specification for when one goroutine is guaranteed to observe another goroutine's writes. Use for: understanding races, happens-before, channels, locks, and atomics.
- [Go Blog: Share Memory By Communicating](https://go.dev/blog/codelab-share)
  Official Go blog article explaining the channel-oriented design philosophy. Use for: deciding channels vs shared memory.
- [Go Blog: Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
  Official Go blog article on pipelines, cancellation, and avoiding goroutine leaks. Use for: real-world channel patterns.
- [Data Race Detector](https://go.dev/doc/articles/race_detector)
  Official guide to `go test -race` and race debugging. Use for: practical detection of unsafe shared memory.

## Wisdom (Communities)

- [Go Forum](https://forum.golangbridge.org/)
  Beginner-friendly Go community. Use for: asking design questions and getting code review on small examples.
- [Gophers Slack](https://invite.slack.golangbridge.org/)
  Large Go practitioner community. Use for: real-world advice once you have a specific concurrency question.
- [r/golang](https://www.reddit.com/r/golang/)
  Broad Go subreddit. Use for: seeing common production questions, but verify answers against official docs.
