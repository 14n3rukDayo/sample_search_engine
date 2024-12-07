// main.go
package main

import (
	"main/wire"
)

func main() {
	e, err := wire.InitializeEcho()
	if err != nil {
		return
	}
	e.Logger.Fatal(e.Start(":8000"))
}
