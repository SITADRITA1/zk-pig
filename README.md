# ZK-PIG

**ZK-PIG** is a ZK-EVM Prover Input generator responsible for generating the data inputs necessary for proving Execution Layer (EL) blocks. Those prover inputs can later be consumed by proving infrastructures to generate EL block proofs.

From an architecture perspective, ZK-PIG connects to an Ethereum compatible EL full or archive node via JSON-RPC to fetch the necessary data.

> **Note about Prover Inputs:** ZK proving engines operate in isolated & stateless environments without a direct access to a full blockchain node. The **Prover Input** refer to the minimal EVM data required by such a ZK-EVM proving engine to effectively prove an EL block. For more information on prover inputs, you can refer to this [article](https://ethresear.ch/t/zk-evm-prover-input-standardization/21626).

## Installation

### Homebrew

`zkpig` is distributed with Homebrew.

If installing for the first time you'll need to add the `kkrt-labs/kkrt` tap

```sh
brew tap kkrt-labs/kkrt
```

Then to install `zkpig`, run 

```sh
brew install zkpig
```

You can test the installation by running

```sh
zkpig version
```

## Usage

### Prerequisites

- You have an Ethereum Execution Layer compatible node accessible via JSON-RPC (e.g., Geth, Erigon, Infura, etc.).

    > **⚠️ Warning ⚠️:** If generating prover inputs for an old block, you must use an Ethereum archive node, that effectively exposes `eth_getProof` JSON-RPC for the block in question. Otherwise, ZK-PIG will fail at generating the prover inputs due to missing data.

    > **Note:** ZK-PIG is compatible with both HTTP and WebSocket JSON-RPC endpoints.

### Generate Prover Inputs

First, you can set the `CHAIN_RPC_URL` environment variable to the URL of the Ethereum node to collect the data from.

```sh
export CHAIN_RPC_URL=<rpc-url>
```

To generate prover inputs for a given block, you can use the following command:

```sh
zkpig generate --block-number <block-number>
```

> **Note:** The command takes around 1 minute to complete, mainly due to the time it takes to fetch the necessary data from the Ethereum node (around 2,000 requests/block).

On successful completion, the prover inputs are stored in the `/data` directory.

To generate prover inputs for the `latest` block, you can use the following command:

```sh
zkpig generate --block-number latest
```

For more information on the commands, you can use the following command:

```sh
zkpig generate --help
```

### Commands Overview 

To get all available commands, and flags, you can use the following command:

```sh
zkpig help
```

### Logging

To configure logging you can set
- `--log-level` to configure verbosity (`debug`, `info`, `warn`, `error`)
- `--log-format` to switch between `json` and `text`. For example:

```sh
zkpig generate \
  --block-number 1234 \
  --chain-rpc-url http://127.0.0.1:8545 \
  --data-dir ./data \
  --store-content-type json \
  --log-level debug \
  --log-format text
```

## Architecture

For a more detailed architecturedocumentation, you can refer to the [Documentation](https://kkrt-labs/zkpig/docs/prover-inputs-generation).

## Contributing

Interested in contributing? Check out our [Contributing Guidelines](CONTRIBUTING.md) to get started! 

## Other commands

### `zkpig preflight`

> Description: Only fetches and locally stores the the necessary data (e.g. pre-state, block, transactions, state proofs, etc.) but does not run block validation. This is useful if you want to collect the data for a block and run block validation separately. It is also useful for debugging purposes.

#### Usage

```sh
zkpig preflight \
  --block-number 1234 \
  --chain-rpc-url http://127.0.0.1:8545 \
  --data-dir ./data
  --store-content-type json
```

### `zkpig prepare`

> Description: Converts the data collected during preflight data into the minimal, final prover input.
> Can be ran offline without a chain-rpc-url. In which case it needs to be provided with a chain-id.

#### Usage

```sh
kkrtctl prepare \
  --block-number 1234 \
  --chain-id 1 \
  --chain-rpc-url http://127.0.0.1:8545 \
  --data-dir ./data \
  --store-content-type json
```

### `zkpig execute`

> Description: Re-executes the block over the previously generated prover inputs.
> Can be ran offline without a chain-rpc-url. In which case it needs to be provided with a chain-id.

#### Usage

```sh
kkrtctl prover-inputs execute \
  --chain-id 1 \
  --block-number 1234 \
  --chain-rpc-url http://127.0.0.1:8545 \
  --data-dir ./data \
  --store-content-type json
```

### Commands List

1. `zkpig version` - Print the version
1. `zkpig help` - Print the help message
1. `zkpig generate --block-number <block-number> --chain-rpc-url <rpc-url> --data-dir <data-dir>` - Generate prover input for a specific block.
1. `zkpig preflight --block-number <block-number> --chain-rpc-url <rpc-url> --data-dir <data-dir>` - Collect necessary data to generate prover inputs from a remote JSON-RPC Ethereum Execution Layer node
1. `zkpig prepare --block-number <block-number> --chain-rpc-url <rpc-url> --data-dir <data-dir>` - Prepare prover inputs by basing on data previously collected during preflight.
1. `zkpig execute --block-number <block-number> --chain-rpc-url <rpc-url> --data-dir <data-dir>` - Execute block by basing on prover inputs previously generated during prepare.
