# Installation


## Install Go

[Go 1.19+](https://golang.org/doc/install) is required.

## Install the `red`

If you want to install the `red` binary to run your node or to communicate with other nodes,
please clone the `Re` project and build it.

```bash
# Make sure to checkout to decentralized-identifier branch
git clone https://github.com/jim380/Re.git
cd Re
make install  # All binaries are installed in $GOPATH/bin
```

Verify that the `red` binary is installed successfully.
```bash
$ red version