# Gobf - A [Brainfuck](https://wikipedia.org/wiki/Brainfuck) interpreter written in Go

## How to use

Execute brainfuck instructions from a file:
```shell
$ ./gobf examples/hello-world.b
```
Or piping instructions into it:
```shell
$ echo "+-[..." | ./gobf -
```
