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
- [FindIndex](#findindex)
- [Find](#find)
- [Includes](#includes)
- [Every](#every)
- [Some](#some)
- [Uniq](#uniq)
- [Map](#map)
- [ForEach](#foreach)
- [Filter](#filter)
- [Contact](#contact)
- [Pop](#pop)
- [Shift](shift)
- [Remove](#remove)
- [Insert](#insert)
- [Drop](#drop)
- [Sample](#sample)
- [Reverse](#reverse)



### IndexOf    
return the index of the element in the slice, if the element is not in the slice, return -1    

```go
s := []string{"a", "b", "c"}
i := gearslice.IndexOf(s, "c")
// 2

```

[⬆️ back](#gearslice)
### FindIndex
return the index of the first element in the slice that passed the test implemented by the provided function. return -1 if no corresponding element is found.

```go
s := []int{0,1,2,3,4,5,6}
r := FindIndex(s, func(el int, _ int) bool {
  return el > 0
})
// 1

```

[⬆️ back](#gearslice)

### Find     
return the value of the first element of the slice that passed the test function provided, return zero value if no corresponding element is found

```go
type Person struct {
  Name   string
  Age    int
  Gender int
  Grade  int
}

s := []Person{
  {Name: "Tom", Age: 12, Gender: 1, Grade: 2},
  {Name: "Jim", Age: 11, Gender: 1, Grade: 1},
  {Name: "Dave", Age: 13, Gender: 0, Grade: 3},
}

r, ok := Find(s, func(el Person, _ int) bool {
  return el.Age > 11 && el.Gender ==0
})
// Person{Name: "Dave", Age: 13, Gender: 0, Grade: 3}, true


```

[⬆️ back](#gearslice)

### Includes      
weather the slice contains a certain element      

```go
s := []string{"a", "b", "c"}
r := gearslice.Includes(s, "d")
//  false
```

[⬆️ back](#gearslice)

### Every     
weather all elements in the slice passed the test implemented by the provided function    

```go
s := []int{1, 2, 3, -1}
r := gearslice.Every(s, func(el int, _ int) bool {
  return el > 0
})
// false
```

[⬆️ back](#gearslice)

### Some    
 weather at least one element in the slice passed the test implemented by the provide function

 ```go
 s := []int{1, 2, 3, -1}

r := gearslice.Some(s, func(el int, _ int) bool {
  return el > 0
})
// true
 ```

[⬆️ back](#gearslice)

### Uniq     
remove duplicate elements in the slice

```go
s := []int{1, 2, 3, 4, 4, 5}
r := gearslice.Uniq(s)
// []int{1,2,3,4,5}
```

[⬆️ back](#gearslice)

### Map     
create a new slice populated with the results of calling the provide function on every element in the calling slice    

```go
s := []int{1, 2, 3, 4, 5}
r := gearslice.Map(s, func(el int, _ int) int {
  return el * 2
})
// []int{2, 4, 6, 8, 10}
```
[⬆️ back](#gearslice)


### ForEach    
execute a provided function once for each element in the slice    

```go
s := []int{1, 2, 3, 4, 5}
var r = make([]string, 0, len(s))
gearslice.ForEach(s, func(el int, _ int) {
  r = append(r, fmt.Sprintf("%d", el))
})
// r: []string{"1", "2", "3", "4", "5"}
```
[⬆️ back](#gearslice)

### Filter     
filtered down to just the elements from the given slice that passed the test implemented by the provided function     

```go
s := []int{1, 2, 3, 4, -1}
r := gearslice.Filter(s, func(el int, _ int) bool {
  return el > 0
})
// []int{1, 2, 3, 4}
```
[⬆️ back](#gearslice)

### Contact      
Concatenate multiple slices and return a new slice      

```go
s1 := []int{1, 2, 3, 4}
s2 := []int{5, 6}
s3 := []int{6, 7, 8}

r := gearslice.Contact(s1, s2, s3)
// []]int{1, 2, 3, 4, 5, 6, 6, 7, 8}

```
[⬆️ back](#gearslice)

### Pop     
remove the last element from a slice and return that element, remove the removed element. it will change the length of the slice     

```go
s := []int{1, 2, 3, 4}
r := gearslice.Pop(&s)
// r: 4
// s: []int{1, 2, 3}
```
[⬆️ back](#gearslice)

### Shift      
remove the first element from a slice and return that removed element. This method changes the length of the slice

```go
s := []int{1, 2, 3, 4}
r := gearslice.Shift(&s)
// r: 1
// s: []int{2, 3, 4}
```
[⬆️ back](#gearslice)

### Remove      
remove a value in the slice at a given index. it will modify the origin slice

```go
s := []int{1, 2, 3}
r := gearslice.Remove(&s, 1)
// []int{1, 3}
```
[⬆️ back](#gearslice)

### Insert     
insert a value in the slice at a given index. it will modify the origin slice

```go
s := []int{1, 2, 3}
gearslice.Insert(&s, 1, 20)
// s: []int{1, 20, 2, 3}
```
[⬆️ back](#gearslice)

### Drop    
drop n elements from the beginning(if n greater than zero) or the end(if n less than zero) of the slice    

```go
s := []int{1, 2, 3, 4, 5, 6, 7}
r, _ := gearslice.Drop(s, 2)
// []int{3, 4, 5, 6, 7}
```
[⬆️ back](#gearslice)

### Sample    
get a random element from the slice

```go
s := []int{1, 2, 3, 4, 5, 6, 7, 8}
r := gearslice.Includes(s, gearslice.Sample(s))
// true
```
[⬆️ back](#gearslice)

### Reverse     
reverse a slice, return a new slice

```go
s := []int{1, 2, 3, 4, 5}
r := gearslice.Reverse(s)
// []int{5, 4, 3, 2, 1}
```
[⬆️ back](#gearslice)

