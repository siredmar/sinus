package sinus

import (
	"fmt"
	"math"
	"time"

	"github.com/siredmar/sinus/internal/timer"
)

// Spec defines one single sinus curve
type Spec struct {
	Enabled   bool
	Frequency float64
	Amplitude float64
}

// Config is passed to the factory methode NewSinus
type Config struct {
	Curves       []Spec
	SampleRateMs float64
}

// Sinus represents the sinus instance
type Sinus struct {
	curves      []Spec
	currentMs   float64
	sampleTimer *timer.Timer
	buf         chan float64
}

// NewSinus creates a new Sinus instance with a configuration and a channel where the results are being sent to
func NewSinus(config *Config, buf chan float64) (*Sinus, error) {
	s := &Sinus{
		currentMs:   0,
		sampleTimer: nil,
		buf:         buf,
		curves:      config.Curves,
	}
	if buf == nil {
		return nil, fmt.Errorf("passed buffer is nil")
	}
	s.sampleTimer = &timer.Timer{
		Function: func() {
			s.sample()
		},
		Duration: time.Duration(config.SampleRateMs * float64(time.Millisecond)),
		Times:    0,
	}
	return s, nil
}

func (s *Sinus) sample() {
	s.currentMs += float64(s.sampleTimer.Duration.Milliseconds())
	s.buf <- s.calculate()
}

func (s *Sinus) calculate() float64 {
	var output float64 = 0.0
	for _, c := range s.curves {
		if c.Enabled {
			output += c.Amplitude * math.Sin(2*math.Pi*c.Frequency*math.Remainder(s.currentMs/1000, 1/(c.Frequency)))
		}
	}
	return output
}

// Start starts the calculation and delivery of values
func (s *Sinus) Start() {
	s.sampleTimer.Start()
}

// Start stops the calculation and delivery of values
func (s *Sinus) Stop() {
	s.sampleTimer.Stop()
}
