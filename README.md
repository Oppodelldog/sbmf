# Simple Binary Message Format - Code Generator

![image](image.png)

This is a simple binary message format code generator.
It generates code for encoding and decoding messages in a binary format.

The code generator is written in Go and can be used to generate code in Go, C#.

### Preparations
Actually there are no releases, so you have to build the code generator yourself.

Therefore you need to have Go installed.
Once you have Go installed, you can build and install the code generator with the following command:

```bash
go install github.com/Oppodelldog/sbmf
```

## Usage

The code generator is a command line tool which takes a file as input.
The file contains the definition of the message format and meta information where to generate the code.

```bash
sbmf sample.yaml
```

See the [example/sample.yaml](example/sample.yaml) for an example.