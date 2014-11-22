package kairos

import (
	"log"
	"os/exec"
	"time"
)

/* Future struct, core entry for all futures that are sent in
 * through the API */
type Future struct {
	Id        string
	Scheduled time.Time
	Cancelled bool
	Done      bool
	Error     bool
	Command   exec.Cmd
}

/* Get a Future by ID */
func GetFuture(id string) Future {
	return Future{
		Id:        id,
		Scheduled: time.Now().Add((time.Second * 1)),
		Cancelled: false,
		Done:      false,
		Error:     true,
		Command:   *exec.Command("touch", "foo"),
	}
}

/* */
func SaveFuture(f Future) {
}

/* Run a Future */
func RunFuture(f Future) {
	delta := f.Scheduled.Sub(time.Now())
	if delta < 0 {
		delta = 0
	}
	log.Println("Sleeping for ", delta)
	time.Sleep(delta)
	log.Println("Slept; running: ", f.Command.Args)

	err := f.Command.Run()
	f.Done = true
	f.Error = err == nil
	SaveFuture(f)

	if err != nil {
		log.Fatal(err)
	}
}
