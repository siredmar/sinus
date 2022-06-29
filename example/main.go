package main

import (
	"fmt"

	"github.com/siredmar/sinus"
)

func main() {
	config := &sinus.Config{
		Curves: []sinus.Spec{
			{
				Enabled:   true,
				Frequency: 1.0,
				Amplitude: 1.0,
			},
			{
				Enabled:   true,
				Frequency: 5.0,
				Amplitude: 0.5,
			},
			{
				Enabled:   true,
				Frequency: 4.0,
				Amplitude: 0.3,
			},
		},
		SampleRateMs: 1,
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
