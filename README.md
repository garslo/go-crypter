# crypter

A simple cli program to encrypt phrases.

# Installation

```bash
$ go get github.com/garslo/crypter
$ go install github.com/garslo/crypter
```

# Usage

```bash
$ crypter help
Usage:
  crypter [command]

Available Commands:
  md5 <phrases>             :: encrypt using md5
  sha1 <phrases>            :: encrypt using sha1
  sha256 <phrases>          :: encrypt using sha256
  sha512 <phrases>          :: encrypt using sha512
  help [command]            :: Help about any command


Use "crypter help [command]" for more information about that command.

$ crypter md5 zoom zing
15913c103a5238e5a80ac2f498ee090d
2f5117cb8211933814c3da646e0e4dde
```

# Credits

`crypter` uses the excellent [cobra](https://github.com/spf13/cobra)
library for command line goodness.
