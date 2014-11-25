package main

import (
	"./future"
	"log"
)

func main() {
	f := future.GetFuture("0bbdae47-849c-423d-894e-7916a5fc040c")
	f.Run()
	log.Println("Listening...")
}
