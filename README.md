# gear     
> gear is a tiny front-end friendly go tool library, like lodash.

## install
```shell
go get github.com/CiroLee/gear
```
## uses     
all function modules startwidth gearXXX.     
such as gearslice that expands the ability of slice
```go
import 
  (
    "fmt",
    "github.com/CiroLee/gearslice"
  )

func main() {
  s := []int{1,2,3,3,4,5,6,6}
  r := gearslice.Uniq(s)
  // []int{1,2,3,4,5,6}
}
```


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
## apis    

### gearslice     
> slice expansion functions    

- [IndexOf](#indexof)




### IndexOf    
return the index of the element in the slice, if the element is not in the slice, return -1    
```go
s := []string{"a", "b", "c"}
i := gearslice.IndexOf(s, "c")
// 2

```
[back](#gearslice)



