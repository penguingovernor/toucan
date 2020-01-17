# toucan
A simple stream cipher for educational purposes only.

## Installation

### Universal

Install appropriate version from the [release page](https://github.com/penguingovernor/toucan/releases).

### From source

1. Clone the repo

`git clone https://github.com/penguingovernor/toucan.git`

2. Run go build

`go build`


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

64 -- for 512-bit security.

32 -- for 256-bit security.

16 -- for 128-bit security.

Windows:

Download a POSIX shell like [git bash](https://git-scm.com/) and repeat above steps.

#### How can I use standard input?

Linux/MacOS:
```bash
echo -n "key" | toucan encrypt file /dev/stdin IV file.out
echo -n "key" | toucan decrypt file.out /dev/stdin IV file.decrypted
```
