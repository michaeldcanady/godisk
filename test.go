package main

import (
	"fmt"

	"github.com/michaeldcanady/godisk/godisk"
)

func main() {
	for _, drive := range godisk.GetDrives() {
		fmt.Println(drive)
	}
}
