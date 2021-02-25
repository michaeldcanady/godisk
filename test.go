package main

func main() {
  for _, drive := godisk.GetDrives() {
    fmt.Println(drive)
  }
}
