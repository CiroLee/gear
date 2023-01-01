# gear     
> gear is a tiny front-end friendly go tool library, like lodash.


## test
```shell
# single file
go test -v ./gearmap
# all files and output test report
go test -v ./... -coverprofile=coverage.txt -covermode=atomic 

```