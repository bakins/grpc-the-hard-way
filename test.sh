#!/usr/bin/env bash
set -e

# this is horrible.
# requires bash 4
# quick test of servers/clients

killServer() {
    PID=$1
    pkill -P $PID
    kill $PID
    wait $PID || true
}

killChildren() {
    pkill -P $$
}

trap killChildren exit

declare -A TESTS
TESTS["go run ./standard/greeter_server/main.go"]="go run ./unary/client/main.go"
TESTS["go run ./unary/server/main.go"]="go run ./standard/greeter_client/main.go"
TESTS["go run ./standard/greeter_server/main.go"]="go run ./unary-v2/client/main.go"
TESTS["go run ./unary-v2/server/main.go"]="go run ./standard/greeter_client/main.go -gzip"
TESTS["go run ./standard/client-stream-server/main.go"]="go run ./client-stream/client/main.go"
TESTS["go run ./client-stream/server/main.go"]="go run ./standard/client-stream-client/main.go -gzip"
TESTS["go run ./standard/server-stream-server/main.go"]="go run ./server-stream/client/main.go"
TESTS["go run ./server-stream/server/main.go"]="go run ./standard/server-stream-client/main.go -gzip"
TESTS["go run ./bidi/server/main.go"]="go run ./standard/bidi-client/main.go -gzip"
TESTS["go run ./standard/bidi-server/main.go"]="go run ./bidi/client/main.go"
for SERVER in "${!TESTS[@]}"; do
    CLIENT="${TESTS[$SERVER]}"

    echo "=========="
    echo "server: $SERVER"
    echo "client: $CLIENT"

    set -e

    $SERVER &
    PID=$!
    
    sleep 1

    $CLIENT

    killServer $PID
done
