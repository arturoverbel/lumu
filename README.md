# 🧩 Fragment Puzzle

This project solves a puzzle by retrieving and assembling fragments of text from an HTTP service. It includes:

- A JavaScript solution using Promises. `./run.sh`
- A benchmarking script to measure performance.
- An experimental (and deprecated) Go version in `/golang_script`.

---

## 🚀 How to Run

### 🔧 Run the script
This command exec the script once

```bash
$ npm i
$ chmod +x run.sh
$ ./run.sh

▶ Running script_promise.js
hello world quick brown fox jumps over lazy dog you have to call all request at same time if you want to see the puzzle fragments fast enough
⏱ Execution time: 532 ms

$
```

### ⚒️ Benchmark
Default run (100 repetitions):
```bash
$ npm i
$ chmod +x benchmark.sh
$ ./benchmark.sh
```
Custom repetitions:
```bash
$ ./benchmark.sh 10

▶ Running script_promise.js 10 times...
⏳ Iteration 1...
✔ Time: 475 ms
⏳ Iteration 2...
✔ Time: 532 ms
⏳ Iteration 3...
✔ Time: 506 ms
⏳ Iteration 4...
✔ Time: 585 ms
⏳ Iteration 5...
✔ Time: 570 ms
⏳ Iteration 6...
✔ Time: 521 ms
⏳ Iteration 7...
✔ Time: 568 ms
⏳ Iteration 8...
✔ Time: 527 ms
⏳ Iteration 9...
✔ Time: 511 ms
⏳ Iteration 10...
✔ Time: 534 ms
📊 Average execution time: 532 ms (over 10 executions)

$
```
# 💡 Why Promises in JS?
The JavaScript solution leverages Promises because:

- fetch is asynchronous and non-blocking by design.

- Promise.all() fires off all requests in parallel and waits for all of them.

- Node.js is highly optimized for concurrent I/O operations like HTTP.

This makes it **fast**, **clean**, and **reliable** for use cases like retrieving many small HTTP fragments quickly.
Additionally, the script is designed to:

- Exit immediately after collecting all 28 puzzle fragments.

- Avoid unnecessary waiting or polling.


# ⚠️ Why Go Didn't Work Well Here
An attempt to solve the same puzzle using Go with goroutines was made and stored in /golang_script.
Despite Go being great for concurrency, in this specific challenge it faced issues such as:

- **Unstable performance** when firing hundreds of HTTP requests in parallel.

- Extra complexity in managing goroutine lifecycles and channels.

- Higher sensitivity to rate limits and TCP/connection errors under load.

As a result, the JavaScript version was preferred due to:

- Simpler implementation.

- More predictable execution.

- Better performance for HTTP-heavy workloads.