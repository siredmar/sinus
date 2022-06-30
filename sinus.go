package sinus

import (
	"fmt"
	"math"
	"time"

	"github.com/siredmar/sinus/internal/timer"
)

var NaN float64 = math.NaN()

// Spec defines one single sinus curve
type Spec struct {
	// Enabled enables the sine wave
	Enabled bool
	// Frequency specifies the frequency for the sine wave
	Frequency float64
	// Amplitude is the factor the sine wave gets multiplied with
	Amplitude float64
	// Offset specifies the offset for the sine wave
	Offset float64
	// PhaseShift (in degrees) is the phase shift for the sine wave
	PhaseShift float64
}

// Config is passed to the factory methode NewSinus
type Config struct {
	// Curves describe the single sinus waves
	Curves []Spec
	// OverallOffset is the overall offset after calculation of the overlay of the signals
	OverallOffset float64
	// MinCap is the minimum limit of calculated values. Lower values are capped to this value.
	// Set to NaN if not needed
	MinCap float64
	// MaxCap is the maximum limit of calculated values. Higher values are capped to this value.
	// Set to NaN if not needed
	MaxCap float64
	// SampleRate is the time between two calculated points
	SampleRate time.Duration
}

// Sinus represents the sinus instance
type Sinus struct {
	currentUs   float64
	sampleTimer *timer.Timer
	buf         chan float64
	config      *Config
}

// NewSinus creates a new Sinus instance with a configuration and a channel where the results are being sent to
func NewSinus(config *Config, buf chan float64) (*Sinus, error) {
	s := &Sinus{
		currentUs:   0,
		sampleTimer: nil,
		buf:         buf,
		config:      config,
	}
	if buf == nil {
		return nil, fmt.Errorf("passed buffer is nil")
	}
	s.sampleTimer = &timer.Timer{
		Function: func() {
			s.sample()
		},
		Duration: config.SampleRate,
		Times:    0,
	}
	return s, nil
}

func (s *Sinus) sample() {
	s.currentUs += float64(s.sampleTimer.Duration.Microseconds())
	s.buf <- s.calculate()
}

func (s *Sinus) calculate() float64 {
	var output float64 = 0.0
	for _, c := range s.config.Curves {
		if c.Enabled {
			radians := (c.PhaseShift / 360.0) * 2 * math.Pi
			output += c.Amplitude*math.Sin(2*math.Pi*c.Frequency*math.Remainder(s.currentUs/1000000, 1/(c.Frequency))+radians) + c.Offset
		}
	}
	output += s.config.OverallOffset
	if output < s.config.MinCap {
		output = s.config.MinCap
	} else if output > s.config.MaxCap {
		output = s.config.MaxCap
	}
	return output
}

// Start starts the calculation and delivery of values. You can now receive the calculated values
// reading the channel provided
func (s *Sinus) Start() {
	s.sampleTimer.Start()
}

// Start stops the calculation and delivery of values
func (s *Sinus) Stop() {
	s.sampleTimer.Stop()
}
