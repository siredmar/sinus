# Sinus

This golang package can create sinus waves that are overlayed with a given configuration and sample rate.
The output is transfered into a channel the caller must provide. Multiple sinus waves can be overlayed in the output.

## Example

Run the example that uses the following configuration to generate an overlay of sinus waves.

```go
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
```

```sh
timeout 5s go run example/main.go > dat
gnuplot -e "plot \"dat\"" -p
```

![gnuplot of dat](https://github.com/siredmar/sinus/blob/main/.assets/dat.png?raw=true)


