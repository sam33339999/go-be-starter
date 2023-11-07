# Golang String And Byte Deal

- [How To Assign String To Bytes Array](https://stackoverflow.com/questions/8032170/how-to-assign-string-to-bytes-array)

### string -> []byte
```go
[]byte("Hello world")
```

### [20]byte -> []byte
```go
arr[:]
```

### string -> [20]byte
```go
copy(arr[:], "hello world")
```

### 與上之相同，但首先明確地先將字符串轉化成 []byte
```go
copy(arr[:], []byte("hello world"))
```


#### example code:
```go
package main

import (
	"fmt"
)

func main() {
	var arr [20]byte
	copy(arr[:], "abcdefghijklmnopqrstuvwxyz")
	fmt.Printf("array: %v (%T)\n", arr, arr)
}
```

#### output
```
array: [97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116] ([20]uint8)
```
> 97: a, 98: b, ... 116: t 因為 arr 只有開 20 個空間存儲， byte 可以將字符轉化成 ascii 存儲。 故 a-z 的 20 位到 `t`
