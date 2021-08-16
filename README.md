# etime

Encapsulate golang time operations, time operations are easier

# Installation

````
GO111MODULE=on
go get github.com/dkys/etime
````

# Using

```go
package main

import (
	"fmt"
	"github.com/dkys/etime"
)

func main() {
	start, end := etime.Today()
	timeStr := etime.Format(start, "Y-m-d H/i/s")
	fmt.Printf("start %v\n end %v\n timeStr %s\n", start, end, timeStr)
}

```