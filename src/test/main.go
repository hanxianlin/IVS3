// test project main.go
package main

import (
	"database/sql"
	"fmt"
	"test/pkg_test"
)

func main() {
	pkg_test.Test("dd")
	fmt.Println("Hello World!")
}
