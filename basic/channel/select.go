package channel

import "time"

func process(ch chan<- string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}
