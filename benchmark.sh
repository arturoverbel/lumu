#!/bin/bash
# benchmark.sh - Runs script_promise.js multiple times and calculates the average execution time

# Number of repetitions (default 100)
runs="${1:-100}"

# Accumulator of total time
total=0

echo ""
echo "▶ Running script_promise.js $runs times..."

for ((i = 1; i <= runs; i++)); do
  echo "⏳ Iteration $i..."

  start=$(date +%s%N)
  node script_promise.js > /dev/null
  end=$(date +%s%N)

  elapsed=$(( (end - start) / 1000000 )) # in milliseconds
  echo "✔ Time: ${elapsed} ms"

  total=$((total + elapsed))
done

average=$((total / runs))

echo "📊 Average execution time: $average ms (over $runs executions)"
echo ""
