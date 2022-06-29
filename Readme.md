# Sinus

This golang package can create sinus waves that are overlayed with a given configuration and sample rate.
The output is transfered into a channel the caller must provide.

## Example

Run the example

```sh
timeout 5s go run example/main.go > dat
gnuplot -e "plot \"dat\"" -p
```

![gnuplot of dat](https://github.com/siredmar/sinus/blob/main/.assets/dat.png?raw=true)


