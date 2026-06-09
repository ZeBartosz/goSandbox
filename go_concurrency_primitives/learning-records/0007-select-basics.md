# Select basics

The learner started `select` and understood that `default` makes a `select` non-blocking when no channel case is ready. They correctly identified that a buffered channel with a value makes a receive case ready, and that a nil channel is never ready.

Corrections to reinforce next time:

- `select` runs only one ready case, not all ready cases.
- `time.After(d)` returns a channel that becomes ready after the duration, so a timeout case can prevent blocking forever.
- Receiving from a closed channel is ready immediately and yields the zero value with `ok == false`; `default` does not run in that case.

Next session should start with a quick snippet review, then create/fix a select debugging lab.
