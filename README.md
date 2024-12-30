# Simple Blockchain in Go
This project implements a basic blockchain in Go, featuring:

* **Block Structure:** Each block contains an index, timestamp, data, the hash of the previous block, the hash of the current block, and a nonce used for mining.
* **Mining Mechanism:** Blocks are mined by iterating the nonce until the block's hash meets the difficulty requirements (leading zeros).
* **Blockchain** A simple blockchain that stores blocks and validates new blocks before adding them to the chain.

## Features
* **Genesis Block:** The first block in the chain, created with no previous block hash.
* **Difficulty Adjustments:** Each block’s hash must begin with a specified number of leading zeros to ensure proof-of-work.
* **Block Validation:** Blocks are validated by ensuring the index, previous block hash, and the calculated hash match.
