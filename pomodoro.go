package main

import (
	"fmt"
	"time"
)

func main() {
	prod := 25 * time.Minute
	short := 3 * time.Minute
	long := 15 * time.Minute

	for {
		for i := 0; i < 4; i++ {
			wait(prod, "productive phase")
			alert()
			if i != 3 {
				wait(short, "short break")
				alert()
			}
		}
		wait(long, "long break")
		alert()
	}

}

func alert() {
	for i := 0; i < 10; i++ {
		fmt.Print("\a")
		time.Sleep(time.Second)
	}
}

func wait(d time.Duration, phase string) {
	t := time.Now()
	f := t.Add(d)
	c := make(chan bool)
	go func() {
	loop:
		for {
			select {
			case <-c:
				break loop
			case <-time.Tick(250 * time.Millisecond):
				s := time.Since(f)
				min := int(s.Minutes()) * -1
				sec := (int(s.Seconds()) * -1) - (min * 60)
				fmt.Print("\r\t\t\t\t\t")
				fmt.Print("\rPhase: ", phase, "\t", fmt.Sprintf("%02dm%02ds", min, sec), "\t\t\t")
			}
		}
	}()
	time.Sleep(d)
	close(c)
}
