# WaitGroup counter model

The learner understood `WaitGroup` as a counter: `Add` increases the number of tasks being waited on, `Done` decreases it, and `Wait` blocks until the counter reaches zero. They also identified that an unmatched `Add(2)` with only one `Done` causes the program to wait forever, which is essential for debugging WaitGroup deadlocks.
