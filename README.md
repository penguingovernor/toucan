![](https://user-images.githubusercontent.com/13544676/72605461-a96c7880-38d1-11ea-859c-7538f32a1623.png)

Toucan is a simple stream cipher for educational purposes only.
It is both a CLI and a library.

[![GoDoc](https://godoc.org/github.com/penguingovernor/toucan/crypto/toucan?status.svg)](https://godoc.org/github.com/penguingovernor/toucan/crypto)
[![Go Report Card](https://goreportcard.com/badge/github.com/PenguinGovernor/toucan)](https://goreportcard.com/report/github.com/PenguinGovernor/toucan)

## API Installation

Grab the `toucan/crypto` package using `go get`

```shell
go get github.com/penguingovernor/toucan/crypto
```

See the above `godoc` badge for the API specification.

To avoid naming conflicts when using this package (which again, you really shouldn't be using this package), import the package under a different name, like so:

```go
import toucan "github.com/penguingovernor/toucan/crypto"
```

## CLI Installation

### Universal

Install appropriate version from the [release page](https://github.com/penguingovernor/toucan/releases).

### From source

1. Clone the repo

`git clone https://github.com/penguingovernor/toucan.git`

2. Run go build

`cd cmd && go build`

### Usage

```
FOR EDUCATIONAL PURPOSES ONLY.

Usage:
        toucan [command] msgFile keyFile IVFile [outputFile]

Available Commands:
        encrypt         Encrypt files
        decrypt         Decrypt files

Flags:
        -h, --help              print this help message.

Notes:
        If [outputFile is omitted] then stdout is used.
```

### F.A.Q.

#### Is this secure?

No.
Toucan is for educational purposes only.

#### How does this work?

Read this [Wiki](https://en.wikipedia.org/wiki/Stream_cipher) page, it can explain it much better than I can!

#### What's an IV?

An IV or initialization vector allows the reuse of keys.
It doesn't have to be securely stored and can be sent along with the cipher text.
As long as the IV isn't reused all should be dandy.
For practical purpose, consider an IV as a secondary key that doesn't have to be secret and must only be used once.

#### How do I come up with an IV/Key?

Use a cryptographically secure pseudo random number generator!

Linux/MacOS:

```bash
head -c nBytes < /dev/urandom > file.IV
# You can also do this for your key!
```

Where I suggest nBytes be one of the following:

64 - for 512-bit security.

32 - for 256-bit security.

16 - for 128-bit security.

Windows:

Download a POSIX shell like [git bash](https://git-scm.com/) and repeat above steps.

#### How can I use standard input?

Linux/MacOS:

```bash
echo -n "key" | toucan encrypt file /dev/stdin IV file.out
echo -n "key" | toucan decrypt file.out /dev/stdin IV file.decrypted
```
