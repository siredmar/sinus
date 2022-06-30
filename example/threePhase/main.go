package main

import (
	"fmt"
	"time"

	"github.com/siredmar/sinus"
)

func main() {
	config := &sinus.Config{
		Curves: []sinus.Spec{
			{
				Enabled:    true,
				Frequency:  50.0,
				Amplitude:  230,
				Offset:     0.0,
				PhaseShift: 0.0,
			},
			{
				Enabled:    true,
				Frequency:  50.0,
				Amplitude:  230,
				Offset:     0.0,
				PhaseShift: 120.0,
			},
			{
				Enabled:    true,
				Frequency:  50.0,
				Amplitude:  230,
				Offset:     0.0,
				PhaseShift: 240.0,
			},
		},
		MinCap:        sinus.NaN,
		MaxCap:        sinus.NaN,
		OverallOffset: 0.0,
		SampleRate:    time.Microsecond * 10,
	}
	buf := make(chan float64, 10)
	go func() {
		s, err := sinus.NewSinus(config, buf)
		if err != nil {
			panic(err)
		}
		s.Start()
	}()
	for {
		fmt.Println(<-buf)
	}
}
