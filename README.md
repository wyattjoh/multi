# multi

A simple experiment with UDP and Go to see what it would take to have several
clients communicating over websockets, brokered by a UDP connection to a central
host and have them keep in sync.

This was developed without pre-existing knowledge of common UDP techniques and
is done as a very naive implementation for experimentation purposes only.

## Installation

```bash
go get github.com/wyattjoh/multi/...
```

## Running

Run the server with:

```bash
multi-server
```

And clients with:

```bash
muli-client
```

See the `--help` flag on both binaries produced for documentation on
configuration.
