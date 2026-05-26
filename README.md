# ARImpRec Go Bindings

基于 [ARImpRec](https://github.com/hasherezade/ARImpRec) 的 Go 语言绑定，**纯 Go 实现，无 CGO 依赖**。

## 特性

- **🚀 无 CGO**: 纯 Go 实现
- **DLL 嵌入**: `ARImpRec.dll` 嵌入到 Go 二进制
- **PE 导入表重建**: 用于 PE 文件修改

## 使用方法

```go
package main

import "github.com/ddkwork/ARImpRec"

func main() {
    r := &ipmrec.ARImpRec{}
    // 使用导入表重建功能...
}
```

## 测试

```bash
go test -v
```

## 许可证

MIT License
