package typoverride

import (
	"syscall"
	"unsafe"
)

func getPage(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize()-1))))[:syscall.Getpagesize()]
}

// Do 以 interface{} 为槽接收一个 value，并获取 value 对应的 type
// 因为我的目的是自动注入 struct tag，所以这里的 type 就是 structType
// 在底层结构里， structField 的 name 和 tag 都存储在 name 这个结构里，这实际上是一个 *[?]byte
// 于是我们只需要修改 name，就能达到 注入 tag 的目的
func Do(v interface{}) struct{}{
	t := (*(*iface)(unsafe.Pointer(&v))).t
	for i := 0;i<len(t.fields);i++{
		f := &t.fields[i]
		n := newName(f.name.name(), "json:\"" + pascalToUnderline(f.name.name()) + "\"", f.name.isExported())

		page := getPage(uintptr(unsafe.Pointer(f)))
		syscall.Mprotect(page, syscall.PROT_WRITE)
		f.name.bytes = n.bytes
		syscall.Mprotect(page, syscall.PROT_READ)
	}
	return struct{}{}
}
