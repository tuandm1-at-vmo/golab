package lab02

import "time"

const (
	Red    = "red"
	Green  = "green"
	Yellow = "yellow"
)

type Color struct {
	// Name of the color (see the declared constants).
	Current string
	// Total seconds the color exists.
	Ttl int
	// Waits for a notification for this color.
	Wait chan bool
	// Notifies to another color.
	Notify chan bool
}

func update(color Color, out chan string) {
	end := time.Now().Add(time.Duration(color.Ttl * int(time.Second)))
	for time.Now().Before(end) {
		out <- color.Current
		time.Sleep(time.Second)
	}
	color.Notify <- true
}

// Waits a color and updates it to an `out` channel every second.
func WaitColor(color Color, out chan string) {
	for {
		<-color.Wait
		update(color, out)
	}
}
