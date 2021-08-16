# etime

Encapsulate golang time operations, time operations are easier

# Installation

````
GO111MODULE=on
go get github.com/dkys/etime
````

# Using

````go
import (
  "github.com/dkys/etime"
)

start, end := etime.Today()
timeStr := etime.Format(start,"Y-m-d H/i/s")
````go