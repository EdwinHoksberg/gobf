# Gobf - A JIT powered [Brainfuck](https://wikipedia.org/wiki/Brainfuck) interpreter written in Go

## Support

- [x] Apple Silicon (aarch64)
- [ ] Linux x86_64

## Benchmarks

Ran on a 2021 Apple M1 Pro (16GB):

```
Benchmark 1: ./gobf examples/bench.b
  Time (mean ± σ):     639.9 ms ±   4.7 ms    [User: 633.2 ms, System: 5.8 ms]
  Range (min … max):   635.1 ms … 649.1 ms    10 runs

Benchmark 2: ./gobf examples/mandelbrot.b
  Time (mean ± σ):     740.4 ms ±   3.8 ms    [User: 730.4 ms, System: 8.7 ms]
  Range (min … max):   735.8 ms … 746.8 ms    10 runs
```

## Optimizations

### Jump linking

At parsing time, the program figures out which jumps are linked to each-other and stores that offset with the instruction,
so jump instructions can be executed without calculating the required jump location at runtime.

### Instruction minification

An instruction optimizer can be run to find any consecutive `>`, `<`, `+` or `-` instructions and will merge them into a single instruction.
This can save many operations and (usually a lot of) time for large programs.

Can be disabled with the `-disable-instruction-optimizer` flag.

## How to use

Execute brainfuck instructions from a file:
```shell
$ ./gobf examples/hello-world.b
```
Or piping instructions into it:
```shell
$ echo "+-[..." | ./gobf -
```

Usage:
```
-disable-instruction-optimizer
    Disable optimizer of JIT code

-dump-jit
    Dump generated JIT code to stderr

-memory-size uint
    Size (in bytes) of the memory available to the program (default 30000)
```