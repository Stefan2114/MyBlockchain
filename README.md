# Simple Blockchain in Go
This project implements a basic blockchain in Go. Was made in order to get used to the language.

* **Block Structure:** Each block contains an index, timestamp, data, the hash of the previous block, the hash of the current block, and a nonce used for mining.
* **Mining Mechanism:** Blocks are mined by iterating the nonce until the block's hash meets the difficulty requirements (leading zeros).
* **Blockchain** A simple blockchain that stores blocks and validates new blocks before adding them to the chain.

## Features
* **Genesis Block:** The first block in the chain, created with no previous block hash.
* **Difficulty Adjustments:** Each block’s hash must begin with a specified number of leading zeros to ensure proof-of-work.
* **Block Validation:** Blocks are validated by ensuring the index, previous block hash, and the calculated hash match.

# Exemple

```
BlockChain:
Difficuilty: 2
Blocks:
Index: 0
Timestamp: 1700000000
Data: Genesis block
PrevBlockHash: 
Hash: 0000e2d6a7c0f1b520f9294c5e1b88a5a8f63bb7ad86c24abf68b08f1b240f7a
Nonce: 1234

Index: 1
Timestamp: 1700000002
Data: First block
PrevBlockHash: 0000e2d6a7c0f1b520f9294c5e1b88a5a8f63bb7ad86c24abf68b08f1b240f7a
Hash: 0000f52b13ac594d4b87bdb17b03d6b2c7c64c5a8ab5c01bb04ac032f5fa1202
Nonce: 5678

Index: 2
Timestamp: 1700000004
Data: Second block
PrevBlockHash: 0000f52b13ac594d4b87bdb17b03d6b2c7c64c5a8ab5c01bb04ac032f5fa1202
Hash: 0000d8b28f2da8d5b97116036ac0281d8f057fbde9ab088c1c418da1189578a3
Nonce: 91011
```
