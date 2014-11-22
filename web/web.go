package web

import (
	"../kairos"
	"fmt"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "(root) Hitting: %s!", r.URL.Path[1:])
}

func Foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "(foo ) Hitting: %s!", r.URL.Path[1:])
}

func FooBar(f kairos.Future) {
}
