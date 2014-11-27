package future

import (
	"code.google.com/p/go-uuid/uuid"
	"os/exec"
	"time"
)

/* Future struct, core entry for all futures that are sent in
 * through the API */
type Future struct {
	Id        string
	Scheduled time.Time
	Started   bool
	Cancelled bool
	Done      bool
	Error     bool
	Command   string
}

/* Create a new Future */
func New(scheduled time.Time, command string) Future {
	f := Future{
		Id:        uuid.New(),
		Scheduled: scheduled,
		Cancelled: false,
		Done:      false,
		Error:     true,
		Command:   command,
	}
	return f
}

/* Get a Future by ID */
func Get(id string) Future {
	/* Fill this out with doing a sqlite call... */
	return Future{
		Id: id,
	}
}

/* Save a Future */
func (f Future) Save() {
	/* Fill this out with doing a sqlite call... */
}

/* Schedule a Future */
func (f Future) Schedule() {
	delta := f.Scheduled.Sub(time.Now())
	if delta < 0 {
		delta = 0
	}
	time.Sleep(delta)
	f.Run()
}

/* Run a Future */
func (f Future) Run() {
	if f.Done || f.Started {
		return
	}
	f.Started = true
	f.Save()

	command := exec.Command("/bin/sh", "-c", f.Command)
	err := command.Run()

	f.Done = true
	f.Error = (err == nil)
	f.Save()
}
