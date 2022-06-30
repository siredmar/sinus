# Sinus

This golang package can create sinus waves that are overlayed with a given configuration and sample rate.
The output is transfered into a channel the caller must provide. Multiple sinus waves can be overlayed in the output.

# Installation

Run `go get github.com/siredmar/sinus@latest`.

## Example

See the example that uses the following configuration to generate an overlay of sinus waves.

```go
config := &sinus.Config{
		Curves: []sinus.Spec{
			{
				Enabled:    true,
				Frequency:  0.0,
				Amplitude:  1,
				Offset:     0.0,
				PhaseShift: 0.0,
			},
			{
				Enabled:    true,
				Frequency:  1.0,
				Amplitude:  0.2,
				Offset:     0.0,
				PhaseShift: 0.0,
			},
			{
				Enabled:    true,
				Frequency:  5.2,
				Amplitude:  0.6,
				Offset:     0.0,
				PhaseShift: 0.0,
			},
			{
				Enabled:    true,
				Frequency:  12.3,
				Amplitude:  0.5,
				Offset:     0.0,
				PhaseShift: 0.0,
			},
			{
				Enabled:    true,
				Frequency:  21.1,
				Amplitude:  0.4,
				Offset:     0.0,
				PhaseShift: 0.0,
			},
			{
				Enabled:    true,
				Frequency:  55.7,
				Amplitude:  0.3,
				Offset:     0.0,
				PhaseShift: 0.0,
			},
			{
				Enabled:    true,
				Frequency:  66.3,
				Amplitude:  0.2,
				Offset:     0.0,
				PhaseShift: 0.0,
			},
			{
				Enabled:    true,
				Frequency:  77.1,
				Amplitude:  0.1,
				Offset:     0.0,
				PhaseShift: 0.0,
			},
		},
		MinCap:        sinus.NaN,
		MaxCap:        sinus.NaN,
		OverallOffset: 0.0,
		SampleRate:    time.Microsecond * 1,
	}
```

Run the example and plot the data

```sh
$ timeout 4s go run example/noise/main.go > noise
$ gnuplot -e "plot \"noise\"" -p
```

![gnuplot of noise](https://github.com/siredmar/sinus/blob/main/.assets/noise.png?raw=true)


