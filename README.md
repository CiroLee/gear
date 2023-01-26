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
    "github.com/CiroLee/gear/gearslice"
  )

func main() {
  s := []int{1,2,3,3,4,5,6,6}
  r := gearslice.Uniq(s)
  // []int{1,2,3,4,5,6}
}
```
## docs    

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
- [Count](#count)
- [CountBy](#countby)

### gearstring    
> string expansion functions    

- [Substring](#Substring)
- [CharAt](#charat)
- [Contact](#contact)
- [ToUpperCase](#touppercase)
- [ToLowerCase](#tolowercase)
- [DesensitizeData](#desensitizedata)
- [DesensitizePhone](#desensitizephone)

### gearmap
> map expansion functions

- [Pick](#pick)
- [PickBy](#pickby)
- [Omit](#omit)
- [OmitBy](#omitby)
- [Values](#values)
- [Keys](#keys)
- [Assign](#assign)

### gearmath    
> math expansion functions     

- [Sum](#sum)
- [SumBy](#sumby)
- [Min](#min)
- [Max](#max)
- [Mean](#mean)
- [isPrime](isprime)
- [IsSubset](#issubset)
- [Union](#union)

### geardate     
> date expansion function      
- [Format](#format)
- [IsLeap](#isleap)

### IndexOf    
return the index of the element in the slice, if the element is not in the slice, return -1    
signature:     
```go
func IndexOf[T comparable](s []T, el T) int
```
example:
```go
s := []string{"a", "b", "c"}
i := gearslice.IndexOf(s, "c")
// 2

```

[⬆️ back](#gearslice)
### FindIndex
return the index of the first element in the slice that passed the test implemented by the provided function. return -1 if no corresponding element is found.        
signature:     
```go
func FindIndex[T any](s []T, fn func(el T, index int) bool) int 
```
example:    
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
signature:     
```go
func Find[T any](s []T, fn func(el T, index int) bool) (T, bool)
```
example:     
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
signature:       
```go
func Includes[T comparable](s []T, el T) bool
```
example:    
```go
s := []string{"a", "b", "c"}
r := gearslice.Includes(s, "d")
//  false
```

[⬆️ back](#gearslice)

### Every     
weather all elements in the slice passed the test implemented by the provided function    
signature:       
```go
func Every[T any](s []T, fn func(el T, index int) bool) bool 
```
example:    
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
signature:     
```go
func Some[T any](s []T, fn func(el T, index int) bool) bool
```
example:     
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
signature:       
```go
func Uniq[T constraints.Ordered | constraints.Complex](s []T) []T 
```
example:    
```go
s := []int{1, 2, 3, 4, 4, 5}
r := gearslice.Uniq(s)
// []int{1,2,3,4,5}
```
[⬆️ back](#gearslice)

### Map     
create a new slice populated with the results of calling the provide function on every element in the calling slice      
signature:     
```go
func Map[T, K any](s []T, cb func(el T, index int) K) []K 
```
example:      
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
signature:     
```go
func ForEach[T any](s []T, cb func(el T, index int))
```
example:    
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
signature:    
```go
func Filter[T any](s []T, filter func(el T, index int) bool) []T 
```
example:    
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
signature:    
```go
func Contact[T any](args ...[]T) []T
```
example:     
```go
s1 := []int{1, 2, 3, 4}
s2 := []int{5, 6}
s3 := []int{6, 7, 8}

r := gearslice.Contact(s1, s2, s3)
// []]int{1, 2, 3, 4, 5, 6, 6, 7, 8}

```
[⬆️ back](#gearslice)

### Pop     
remove the last element from a slice and return that element, it will change the length of the slice       
signature:    
```go
func Pop[T any](s *[]T) T
```
example:       
```go
s := []int{1, 2, 3, 4}
r := gearslice.Pop(&s)
// r: 4
// s: []int{1, 2, 3}
```
[⬆️ back](#gearslice)

### Shift      
remove the first element from a slice and return that removed element. This method changes the length of the slice       
signature:      
```go
func Shift[T any](s *[]T) T
```
example:     
```go
s := []int{1, 2, 3, 4}
r := gearslice.Shift(&s)
// r: 1
// s: []int{2, 3, 4}
```
[⬆️ back](#gearslice)

### Remove      
remove a value in the slice at a given index. it will modify the origin slice      
signature:     
```go
func Remove[T any](s *[]T, index int) error
```
example:     
```go
s := []int{1, 2, 3}
r := gearslice.Remove(&s, 1)
// []int{1, 3}
```
[⬆️ back](#gearslice)

### Insert     
insert a value in the slice at a given index. it will modify the origin slice       
signature:      
```go
func Insert[T any](s *[]T, index int, value T) error
```
example:     
```go
s := []int{1, 2, 3}
gearslice.Insert(&s, 1, 20)
// s: []int{1, 20, 2, 3}
```
[⬆️ back](#gearslice)

### Drop    
drop n elements from the beginning(if n greater than zero) or the end(if n less than zero) of the slice       
signature:      
```go
func Drop[T any](s []T, n int) ([]T, error)
```
example:        
```go
s := []int{1, 2, 3, 4, 5, 6, 7}
r, _ := gearslice.Drop(s, 2)
// []int{3, 4, 5, 6, 7}
```
[⬆️ back](#gearslice)

### Sample    
get a random element from the slice       
signature:       
```go
func Sample[T any](s []T) T
```
example:     
```go
s := []int{1, 2, 3, 4, 5, 6, 7, 8}
r := gearslice.Includes(s, gearslice.Sample(s))
// true
```
[⬆️ back](#gearslice)

### Reverse     
reverse a slice, return a new slice       
signature:     
```go
func Reverse[T any](s []T) []T 
```
example:     
```go
s := []int{1, 2, 3, 4, 5}
r := gearslice.Reverse(s)
// []int{5, 4, 3, 2, 1}
```
[⬆️ back](#gearslice)

### Count     
return the number of elements in the slice that equal to value      
signature:     
```go
func Count[T comparable](s []T, value T) int
```
example:     
```go
s := []int{1, 2, 3, 4, 4}
r := gearslice.Count(s, 4)
// 2
```
[⬆️ back](#gearslice)

### CountBy      
return the number of the elements in the slice that pass the test implemented by the provided the function         
signature:    
```go
func CountBy[T comparable](s []T, fn func(el T, index int) bool) int
```
example:    
```go
s := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
r := gearslice.CountBy(s, func(el int, _ int) bool {
    return el > 5
})
// 3
```
[⬆️ back](#gearslice)

### Substring     
return the part of the string from the start and excluding the end, or to the end of the string if no end index is supplied. Not include the index element        
signature:    
```go
func Substring(str string, start, end int) string
```
example:    
```go
str := "hello world"
r := gearstring.Substring(s, 1, 5)
// "ello"
```
[⬆️ back](#gearstring)
### CharAt     
return a specified character from a string     
signature:    
```go
func CharAt(str string, index int) string
```
example:    
```go
str := "hello world"
r := gearstring.CharAt(str, 2)
// "l"
```
[⬆️ back](#gearstring)

### Contact     
Concatenate multiple strings and return a new string        
signature:    
```go
func Contact(args ...string) string
```
example:     
```go
r := gearslice.Contact("hello ", "world")
// "hello world"

```
[⬆️ back](#gearstring)
### ToUpperCase     
change the first letter of the string to upper     
signature:    
```go
func ToUpperCase(s string) string
```
example:    
```go
str := "hello world"
r := gearstring.ToUpperCase(str)
// "Hello world"
```
[⬆️ back](#gearstring)
### ToLowerCase     
change the first letter of the string to lower       
sigmature:     
```go
func ToLowerCase(s string) string
```
example:     
```go
str := "HELLO WORLD"
r := gearstring.ToLowerCase(str)
// "hELLO WORLD"
```
[⬆️ back](#gearstring)

### DesensitizeData    
make data insensitive via hidden part of the data       
signature:    
```go
func DesensitizeData
```
example:    
```go
str := "123这段文字加密00000"
r, _ := gearstring.DesensitizeData(str, 3, 9, "@")
// "123@@@@@@00000"
```
[⬆️ back](#gearstring)

### DesensitizePhone      
hidden middle 4 numbers of the mobile phone, default placeholder is '*'       
signature:    
```go
func DesensitizePhone(val string, placeholder string) (string, error)
```
example:    
```go
phone := "13299889988"
r := gearstring.DesensitizePhone(phone, "#")
// "132####9988"
```
[⬆️ back](#gearstring)

### Pick
return parts of the map according to keys     
signature:     
```go
func Pick[K comparable, V any](m map[K]V, keys []K) map[K]V 
```
example:    
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
r := gearmap.Pick(m, []string{"a", "b"})
// map[string]int{"a":1, "b": 2}
```
[⬆️ back](#gearmap)

### PickBy     
return parts of a map passed the test implemented by the provided function       
signature:     
```go
func PickBy[K comparable, V any](m map[K]V, fn func(k K, v V) bool) map[K]V 
```
example:    
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
r := gearmap.PickBy(m, func(_ string, v int) bool {
  return v > 2
})
// map[string]int{"c": 3}
```
[⬆️ back](#gearmap)

### Omit     
remove parts of a map according to keys      
signature:    
```go
func Omit[K comparable, V any](m map[K]V, keys []K) map[K]V
```
example:      
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
r := gearmap.Omit(m, []string{"a", "b"})
// map[string]int{"c": 3}
```
[⬆️ back](#gearmap)

### OmitBy     
remove parts of a map passed the test implemented by the provided function      
signature:     
```go
func OmitBy[K comparable, V any](m map[K]V, fn func(k K, v V) bool) map[K]V
```
example:   
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
r := gearmap.OmitBy(m, func(_ string, v int) bool {
  return v < 2
})
// map[string]int{"b": 2, "c": 3}
```
[⬆️ back](#gearmap)

### Values     
return values of the map      
signature:     
```go
func Values[K comparable, V any](m map[K]V) []V
```
example:      
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
r := gearmap.Values(m)
sort.Strings(r)
// []int{1, 2, 3}
```
[⬆️ back](#gearmap)

### Keys     
return keys of the map         
signature:     
```go
func Keys[K comparable, V any](m map[K]V) []K
```
example:     
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
r := gearmap.Keys(m)
sort.Strings(r)
// []string{"a", "b", "c"}
```
[⬆️ back](#gearmap)

### Assign     
merge multiple maps into a new map      
signature:      
```go
func Assign[K comparable, V any](maps ...map[K]V) map[K]V
```
example:    
```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"d": 4}
r := gearmap.Assign(m1, m2)
// map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
```
[⬆️ back](#gearmap)

### Sum     
return a sum of the slice       
signature:     
```go
func Sum[T constraints.Integer | constraints.Float | constraints.Complex](s []T) T
```
example:      
```go
s := []int{1, 2, 3, 4, 5}
r := gearmath.Sum(s)
// 15
```
[⬆️ back](#gearmath)

### SumBy     
summarize the values in the slice using the given return value from the function      
signature:    
```go
func SumBy[T any, V constraints.Integer | constraints.Float](s []T, fn func(el T, index int) V) V
```
example:    
```go
s := []string{"hello", "world"}
r := gearmath.SumBy(s, func(el string, _ int) int {
  return len(el)
})
// 10
```
[⬆️ back](#gearmath)

### Min     
return the minimum value of the slice, return zero value if the slice is empty     
signature:    
```go
func Min[T constraints.Integer | constraints.Float](s []T) T
```
example:    
```go
s := []int{1, 2, 3, 4, -4, 5, 6}
r := gearmath.Min(s)
// -4
```
[⬆️ back](#gearmath)

### Max      
return the minimum value of the slice, return zero value if the slice is empty      
signature:     
```go
func Max[T constraints.Integer | constraints.Float](s []T) T 
```
example:     
```go
s := []int{1, 2, 3, 4, -4, 5, 6}
r := gearmath.Max(s)
// 6
```
[⬆️ back](#gearmath)

### Mean     
return the mean value of the slice           
signature:    
```go
func Mean(s []float64) float64
```
example:    
```go
s := []float64{2, 4, 6, 8}
r := gearmath.Mean(s)
// 5.0
```
[⬆️ back](#gearmath)

### IsPrime     
weather a number is a prime   
signature:    
```go
func IsPrime(num int) bool
```
example:    
```go
gearmath.IsPrime(4)
// false
```
[⬆️ back](#gearmath)

### Format    
format a unix timestamp according the layout      
signature:    
```go
func Format(t int64, layout string) string
```
example:    
```go
var timestamp int64 = 1673259412 // 2023-01-09 18:16:52
r := geardate.Format(timestamp, geardate.DefaultLayout)
// "2023-01-09 18:16:52"
```
[⬆️ back](#gearmath)

### IsLeap     
weather the year is leap      
signature:     
```go
func IsLeap(year int) bool
```
example:   
```go
geardate.IsLeap(2023)
// false
```
[⬆️ back](#gearmath)

### IsSubset      
return true if the slice contains all the elements in the subset

```go
s1 := []int{1, 2, 3, 4}
s2 := []int{1, 3}
r := gearslice.IsSubset(s1, s2)
// true
```
[⬆️ back](#gearmath)

### Union     
return the union values of slices      

```go
s1, s2, s3 := []int{1, 2, 3, 4}, []int{2, 5, 7}, []int{-1, 0, 0}
r := Union(s1, s2, s3)
sort.Ints(r)
// []int{-1, 0, 1, 2, 3, 4, 5, 7}
```
[⬆️ back](#gearmath)