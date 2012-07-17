package main

import(
	"log"
	"github.com/opesun/copyrecur"
)

func main() {
    err = copyrecur.CopyDir("/home/jaybill/data", "/home/jaybill/backup")
    if err != nil {
        log.Fatal(err)
    } else {
        log.Print("Files copied.")
    }
}