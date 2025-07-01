# EVM-Compatible Cosmos-Based Blockchain

## Prerequisites

- **Go** version **1.22+** installed

## Installation

1. **Install Go dependencies:**

   ```bash
   go mod tidy

2. **Build and install the blockchain binary:**
   ```bash
   make install
   ```
## Initialize the Blockchain
  Run the following script to initialize the local blockchain configuration, accounts, and genesis state:
   
   ```bash
   ./scripts/init.sh
   ```

## Start the Local Node
  To launch the blockchain node locally:
   
   ```bash
   ./scripts/run-local.sh
   ```

## Interacting with the Local Network
1. **Add a new custom network in MetaMask with the following settings:**
    - RPC URL: http://localhost:8545
    - Chain ID: 929
    - Currency Symbol: COSE
2. **Import an account:**
    - Use the mnemonic phrases printed during the blockchain initialization (e.g., for alice or bob) to import the corresponding wallets into MetaMask.
3. **You should see the pre-funded balances for the imported accounts.**
4. **You can now send transactions between accounts using MetaMask.**