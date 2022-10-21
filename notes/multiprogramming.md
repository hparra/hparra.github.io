---
---

Multi-Programming
=================

Concurrency and Parallelism.

## tl;dr

concurrency - dealing with many things at once
parallelism - doing many things at once

Threads:
- 1:1 -- (kernel-level) each user thread corresponds to a kernel thread
- N:1 (user-level) all user threads correspond to a single kernel thread
- M:N (hybrid)

## IMPLEMENATIONS

### C++

TODO

### Java

Every java object has its own lock

```java
new Thread(() -> {
  //
}).start();
```

### Go

Lighter weight than threads.
Goroutines get multiplexed onto OS threads as needed.

Channels are like unix pipes but they have names and types.

`select` allows you to listen to channels. Like a switch statement.

### Javascript

TODO 

## RESOURCES

[Coordinated Concurrency: Reactive (Observables) vs. CSP](https://vimeo.com/144325523). Kyle Simpson. JSLA.

[Bell Labs and CSP Threads](https://swtch.com/~rsc/thread/). Russ Cox.

[Higher-level threading interface](https://docs.python.org/2/library/threading.html). The Python Standard Library.

[Chapter 17. Threads and Locks](https://docs.oracle.com/javase/specs/jls/se7/html/jls-17.html). The Java Language Specification. Oracle.

[What is the difference between concurrency and parallelism?](https://www.quora.com/What-is-the-difference-between-concurrency-and-parallelism). Quora.

[Concurrency model and Event Loop](https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop). Mozilla Developer Network.

[Concurrency Is Not Parallelism](https://vimeo.com/49718712). Rob Pike.

[Go channels are bad and you should feel bad](http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/)
