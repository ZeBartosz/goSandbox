# Buffered channel full blocks before later code

The learner understood that a send to a full buffered channel blocks immediately, so later receive code in the same goroutine cannot run to free space. Future exercises can rely on this execution-order model when diagnosing deadlocks.
