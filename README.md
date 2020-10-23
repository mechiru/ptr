# ptr

[![ci](https://github.com/mechiru/ptr/workflows/ci/badge.svg)](https://github.com/mechiru/ptr/actions?query=workflow:ci)

A pointer utilities for golang.

before:
```go
func TakeReferenceFunc(s *string) {
	value = "fuga"
	if s != nil {
		value = *s
	}
	fmt.Println(value)
}

func main() {
	tmp := "hoge"
	TakeReferenceFunc(&tmp)
}
```

after:
```go
func TakeReferenceFunc(s *string) {
	fmt.Println(ptr.DerefString(s, "fuga"))
}

func main() {
	TakeReferenceFunc(ptr.RefString("hoge"))
}
```
