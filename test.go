package main

import(
  "github.com/michaeldcanady/godisk/godisk"
)

func main() {
  for _, drive := godisk.GetDrives() {
    fmt.Println(drive)
  }
}
