#!/bin/bash
# run.sh - Script to run Node.js

# Start time (nanoseconds for precision)
start=$(date +%s%N)

echo ""
echo "▶ Running script_promise.js"
node script_promise.js

# End time and duration
end=$(date +%s%N)
elapsed=$(( (end - start) / 1000000 ))

echo "⏱ Execution time: $elapsed ms"
echo ""
