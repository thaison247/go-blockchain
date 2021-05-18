module.exports = class Block {
  constructor(
    index,
    timestamp,
    date,
    transactions,
    nonce,
    hash,
    previousBlockHash
  ) {
    this.index = index;
    this.timestamp = timestamp;
    this.date = date;
    this.transactions = transactions;
    this.nonce = nonce;
    this.hash = hash;
    this.previousBlockHash = previousBlockHash;
  }
};
