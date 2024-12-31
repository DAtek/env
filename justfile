coverfile := ".coverage"

test:
    go test -coverprofile {{ coverfile }} .

show-coverage:
    go tool cover -html {{ coverfile }}

test-examples:
    #!/bin/bash
    set -eou pipefail

    for dir in $(ls examples/); do
        printf "Testing $dir..."
        go run "./examples/$dir" > /dev/null
        printf " PASS\n"
    done