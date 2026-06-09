# Blocked goroutine mental model

The learner corrected the "background task" model into a blocked/concurrent model: a goroutine can be running concurrently, but pause at a channel send or receive until another goroutine performs the matching operation. This unlocks clearer reasoning about unbuffered channels, select, and deadlocks.
