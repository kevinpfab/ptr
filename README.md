# ptr - work with pointers to basic Go types inline

[![Travis CI](https://travis-ci.org/kevinpfab/ptr.svg?branch=master)](https://travis-ci.org/kevinpfab/ptr) [![GoDoc](https://godoc.org/github.com/kevinpfab/ptr?status.svg)](https://godoc.org/github.com/kevinpfab/ptr)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevinpfab/ptr)](https://goreportcard.com/report/github.com/kevinpfab/ptr)

`go get github.com/kevinpfab/ptr`

In Go, to get a pointer to a basic literal you need to first set the literal to a variable.
```go
var namePtr *string
name := "Octocat"
namePtr = &name
```

The `ptr` package let's you do this inline, which is particularly useful when working with pointer heavy structs.

```go
type Employee struct {
    FirstName *string
    LastName *string
    HourlyRate *float32
}

github := &Employee{
    FirstName: ptr.String("Octo"),
    LastName: ptr.String("Cat"),
    HourlyRate: ptr.Float32(15.00),
}
```

All basic types and `time.Time` are supported. Other commonly _pointered_ types in the standard library would be welcome additions.

An extra helper `S(v interface{}) string` is defined to safely get the string representation of the underlying value of a pointer.

```go
var namePtr *string
fmt.Println(ptr.S(namePtr)) // Writing fmt.Println(*namePtr) would crash the program.
// Output: <nil>

name := "Octocat"
namePtr = &name
fmt.Println(ptr.S(namePtr))
// Output: Octocat
```

Package `ptr` is MIT licensed.
