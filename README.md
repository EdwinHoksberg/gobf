# Gobf - A JIT powered [Brainfuck](https://wikipedia.org/wiki/Brainfuck) interpreter written in Go

## Support

- [x] Apple Silicon (aarch64)
- [ ] Linux x86_64

## Benchmarks

Ran on a 2021 Apple M1 Pro (16GB):

```
Benchmark 1: ./gobf examples/bench.b
  Time (mean ± σ):      1.086 s ±  0.007 s    [User: 1.077 s, System: 0.008 s]
  Range (min … max):    1.078 s …  1.099 s    10 runs

Benchmark 2: ./gobf examples/mandelbrot.b
  Time (mean ± σ):      3.117 s ±  0.015 s    [User: 3.091 s, System: 0.022 s]
  Range (min … max):    3.098 s …  3.145 s    10 runs
```


## How to use

Execute brainfuck instructions from a file:
```shell
$ ./gobf examples/hello-world.b
```
Or piping instructions into it:
```shell
$ echo "+-[..." | ./gobf -
```
