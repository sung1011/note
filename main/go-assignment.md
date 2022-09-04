# go assignment

## 变量遮蔽

    switch代码块独立赋值了a, err

```go
package main

import (
	"errors"
	"fmt"
)

var a int = 2020

func main() {
	err := errors.New("eeeeee")

	switch a, err := get(); a {
	case 2045:
		fmt.Println("[inner]", a, err)
	case 2020:
		fmt.Println("[inner]", a, err)
	}
	fmt.Println("[outer]", a, err)
}

func get() (int, error) {
	return 2045, nil
}

// 输出:
// [inner] 2045 <nil>
// [outer] 2020 eeeeee
```