# 通过修改底层类型信息来注入 struct tag `json`

具体逻辑详见 `./override.go`

程序效果请见 `./override_test.go` 测试程序的输出

## 目的

你不用花费时间在重复书写 struct tag 上，比如

```go
// BAD
type T struct {
    InputType           string   `json:"inputType"`
    ValueType           string   `json:"valueType"`
    ValueUnit           string   `json:"valueUnit"`
    SelectList          []string `json:"selectList"`
    SelectListInterface string   `json:"selectListInterface"`
}
```

现在只需要

```go
// GOOD

type T struct {
	InputType           string
	ValueType           string
	ValueUnit           string
	SelectList          []string
	SelectListInterface string
}
var _ = typoverride.Do(T{})
```

## 存在的问题

原先类型信息是存放在 0x112 开头的地址下的（貌似是代码段

覆写类型信息后，新的 name 值的地址在 0xc 开头的地址下

类型信息分开存放后，可能会因为 gc 产生问题 或者 影响内存读取次数

我个人不太清楚这样做的后果。不建议生产环境使用