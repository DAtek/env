coverfile := ".coverage"

test:
    go test -coverprofile {{ coverfile }} .

show-coverage:
    go tool cover -html {{ coverfile }}