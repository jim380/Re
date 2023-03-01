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

<img width="702" alt="Screenshot 2023-03-01 at 06 26 57" src="https://user-images.githubusercontent.com/66339097/222053160-7db0a8a4-f65d-4e7a-8f76-4571cedd007c.png">

<img width="702" alt="Screenshot 2023-03-01 at 06 27 33" src="https://user-images.githubusercontent.com/66339097/222053180-9cd80547-0ed7-4059-aed3-015d4ecab5b8.png">
