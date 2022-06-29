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
				Amplitude: 0.2,
			},
			{
				Enabled:   true,
				Frequency: 2.2,
				Amplitude: 0.6,
			},
			{
				Enabled:   true,
				Frequency: 3.3,
				Amplitude: 0.5,
			},
			{
				Enabled:   true,
				Frequency: 4.1,
				Amplitude: 0.4,
			},
			{
				Enabled:   true,
				Frequency: 5.7,
				Amplitude: 0.3,
			},
			{
				Enabled:   true,
				Frequency: 6.3,
				Amplitude: 0.2,
			},
			{
				Enabled:   true,
				Frequency: 7.1,
				Amplitude: 0.1,
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
