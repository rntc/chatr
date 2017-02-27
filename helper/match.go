package helper

import (
	"fmt"
	"io"
	"log"
)

var partner = make(chan io.ReadWriteCloser)

func chat(u1, u2 io.ReadWriteCloser) {
	fmt.Fprintln(u1, "Found one! say hi.")
	fmt.Fprintln(u2, "Found one! say hi.")

	errch := make(chan error, 1)

	go cp(u1, u2, errch)
	go cp(u2, u1, errch)

	if err := <-errch; err != nil {
		log.Println(err)
	}

	u1.Close()
	u2.Close()
}

func cp(w io.Writer, r io.Reader, errch chan<- error) {
	_, err := io.Copy(w, r)
	errch <- err
}

func Match(c io.ReadWriteCloser) {
	fmt.Fprint(c, "Waiting for a connection...")

	select {
	case partner <- c:
	case p := <-partner:
		chat(p, c)
	}
}
