> [!tip]
> This repository has been archived and is no longer maintained.  
> [Ignite ClI](https://github.com/ignite/cli) is now supports building minimal chains.  
> Run `ignite s chain mini --minimal` to create a minimal chain.

# Mini - A minimal Cosmos SDK chain

This repository contains an example of a tiny, but working Cosmos SDK chain.
It uses the least modules possible and is intended to be used as a starting point for building your own chain, without all the boilerplate that other tools generate. It is a simpler version of Cosmos SDK's [simapp](https://github.com/cosmos/cosmos-sdk/tree/main/simapp).

This branch of `Minid` uses [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) v0.50.x.

> [!note]
> Minid for [Cosmos SDK v0.50.x](https://github.com/cosmonity/chain-minimal/blob/v0.50.x)  
> Minid for [Cosmos SDK v0.52.x](https://github.com/cosmonity/chain-minimal/blob/v0.52.x)  
> Minid for [Cosmos SDK v2](https://github.com/cosmonity/chain-minimal/blob/v2). Learn what v2 was [in this comparison](https://gist.github.com/julienrbrt/6f737f275f784c5540dcbd64dfa982a5) and [talk](https://www.youtube.com/watch?v=Am6eDVRVZSw).  

## How to use

In addition to learn how to build a chain thanks to `minid`, you can as well directly run `minid`.

### Prerequisites

* Install Go as described [here](https://go.dev/doc/install).
* Add `GOPATH` to your `PATH`:
  * `export PATH="$PATH:/usr/local/go/bin:$(/usr/local/go/bin/go env GOPATH)/bin"`

You are all set!

### Installation

Install and run `minid`:

```sh
git clone git@github.com:cosmosregistry/chain-minimal.git
cd chain-minimal
make install # install the minid binary
make init # initialize the chain
minid start # start the chain
```

### Troubleshoot

After running `make install`, verify `minid` has been installed by doing `which minid`.
If `minid` is not found, verify that your `$PATH` is configured correctly.

## Useful links

* [Cosmos-SDK Documentation](https://docs.cosmos.network/)
