package main

import (
	"./kairos"
	"log"
)

func main() {
	f := kairos.GetFuture("0bbdae47-849c-423d-894e-7916a5fc040c")
	kairos.RunFuture(f)

	log.Println("Listening...")
}
