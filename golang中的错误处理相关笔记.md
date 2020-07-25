## golang 中的错误处理相关

#### 好用的相关包

- errors: github.com/pkg/errors  对错误进行包装
- errcheck:  github.com/kisielk/errcheck 检查没有被处理的错误

Error()：将错误转为字符串

---

```go
err := errors.New("error test")
check_str_1 := "error test"
check_str_2 := "bad"

fmt.Printf("%T\n", err.Error())
// Output: string
fmt.Println(err.Error())
// Output: error test
fmt.Println(err.Error() == check_str_1)
// Output: true
fmt.Println(err.Error() == check_str_2)
// Output: false
```



