## test
```shell
cd your-path/gear

# single file
go test -v ./gearmap

# all files and output test report
go test -v ./... -coverprofile=coverage.txt -covermode=atomic

# test all and without output
go test -cover ./...
``` 