# Concurrency In Golang

## Introductions

### Race Conditions

- A race condition occurs when 2 or more operations must execute in the correct order but the program has not been written so that this order is guaranteed to be maintained
- Data race is when 1 concurrent operations attempts to read a variable while at some undetermined time another concurrent operations is attempting to write to the same variable

```golang
    /*
        There will be 3 possible outcomes:
        - Nothing is printed => data++ happen before if the condition
        - Value 0 is printed => the if condition and print happen before  data++
        - Value 1 is printed => the if condition happen but data++ before print
    */
    var data int
    go func() {
        data++
    }
    if data == 0 {
        fmt.Printf("The value is %v.\n", data)
    }
```

=> We need to iterate through the possible scenarios.

### Atomicity

- `context`. Something may be atomic in one context but not another. The atomicity of an operation can change depending on the currently defined code.
- Atomicity is important because if something is atomic => it is same in concurrent contexts

### Memory Access Synchronization

- `critical section` is when we want to exclusive access to a shard resource. To solve this problem problem we can use synchronize access to the memory

```golang
    /*
        There will be 2 scenario
        - The goroutine function will run 1st and it will have exclusive access to the memory. After done will will release the exclusive access
        - The if condition will run 1st and it will have exclusive access to the memory. After done will will release the exclusive access
        => This will solve the problem data race but we still cant solve the race condition
    */

    var memoryAccess sync.Mutex var value int
    go func() {
            memoryAccess.Lock()
            value++
            memoryAccess.Unlock()
    }()

    memoryAccess.Lock()
    if value == 0 {
        fmt.Printf("the value is %v.\n", value)
    } else {
        fmt.Printf("the value is %v.\n", value)
    }
    memoryAccess.Unlock()
```

### Deadlocks, Livelocks, adn Starvation

#### Deadlocks

```golang
    /*
        The below example will cause deadlock:
        - printSum(&a, &b) will lock a and attempt to lock b. In the meantime printSum(&b, &a) will lock b and attempt to lock a => 2 goroutine will wait for other infinitely
    */

    type value struct {
        mu sync.Mutex
        value int
    }

    var wg synce.WaitGroup
    printSum := func(v1, v2 *value) {
        defer wg.Done()
        v1.mu.Lock()
        defer v1.v1.Unlock()

        time.Sleep(2*time.Second)
        v2.mu.Lock()
        defer v2.mu.Unlock()

        fmt.Printf("sum=%v\n", v1.value + v2.value)
    }

    var a, b value
    wg.Add(2)
    go printSum(&a, &b)
    go printSum(&b, &a)
    wg.Wait
```

- A `deadlock` program is one in which all concurrent process are waiting on one another => The program will never recover without outside intervention.

- Techniques to help detect, prevent and correct deadlock:
  - A concurrent process holds exclusive rights to a resource at any onetime
  - A concurrent process must simultaneously hold a resource and be waiting for a additional resource
  - A resource held by a concurrent process can only be released by that process
  - A concurrent process (P1) must be waiting on a chain of other concurrent process (p2) - which are in turn waiting on it (P1)

#### Livelocks

- `Livelocks` are programs that are actively performing concurrent operations but these operations do nothing to move the state of the program forward.
  Ex: You walk in a hallway and you meet a woman. You move to 1 side to let her pass but she do the same. So she move to the other side and you do the same. And this go on forever.

- `Livelocks` happen when 2 or more concurrent process attempt to prevent a deadlock without coordination.

### Starvation

```golang
    var wg sync.WaitGroup
    var sharedLock sync.Mutex
    const runtime = 1*time.Second

    greedyWorker := func() {
        defer wg.Done()
        var count int
        for begin := time.Now(); time.Since(begin) <= runtime; {
            sharedLock.Lock()
            time.Sleep(3*time.Nanosecond)
            sharedLock.Unlock()
            count++
        }
        fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
    }

    politeWorker := func() {
        defer wg.Done()
        var count int
        for begin := time.Now(); time.Since(begin) <= runtime; {
            sharedLock.Lock()
            time.Sleep(1*time.Nanosecond)
            sharedLock.Unlock()

            sharedLock.Lock()
            time.Sleep(1*time.Nanosecond)
            sharedLock.Unlock()

            sharedLock.Lock()
            time.Sleep(1*time.Nanosecond)
            sharedLock.Unlock()

            count++
        }
        fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
    }

    /*
        In below examples. Both worker do the same simulated works (sleep for 3 nano seconds)
        But the greedy worker finishes 471287 works while the polite worker finishes 289777 (half the work of greedy worker)
        Explain: greedyWorker hold the lock for all 3 seconds while the politeWorker only hold the lock when it needed (for 1 second) => greedyWorker is preventing politeWorker from working
    */

    wg.Add(2)
    go greedyWorker()
    go politeWorker()
    wg.Wait()
```

- `Starvation` is any situation where a concurrent process cannot get all the resources it needs to perform work or 1 or more concurrent process are preventing other concurrent processes from accomplishing work
- `Livelocks` is a special case of `Starvation` because all concurrent processes are starved equally and no work is accomplished

## Communicating Sequential Processes

### The Difference Between Concurrency and Parallelism
