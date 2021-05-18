const { v4: uuidv4 } = require("uuid");
const sha256 = require("sha256");
const Block = require("./Block");
const Transaction = require("./Transaction");
const currentNodeUrl = process.env.URL || process.argv[3];

module.exports = class Blockchain {
  constructor(socketId = "") {
    this.socketId = socketId;
    this.chain = [];
    this.pendingTransactions = [];
    this.transactions = [];
    this.currentNodeUrl = currentNodeUrl;
    this.networkNodes = []; // TODO: weird?
    this.createNewBlock(100, "0", "0", []);
  }

  createNewBlock(nonce, previousBlockHash, hash, transactions) {
    const newBlock = new Block(
      this.chain.length + 1,
      Date.now(),
      new Date().toString(),
      transactions,
      nonce,
      hash,
      previousBlockHash
    );
    this.pendingTransactions = [];
    this.chain.push(newBlock);
    return newBlock;
  }

  createNewBlockWithoutAdd(nonce, previousBlockHash, hash, transactions) {
    const newBlock = new Block(
      this.chain.length + 1,
      Date.now(),
      new Date().toString(),
      transactions,
      nonce,
      hash,
      previousBlockHash
    );
    return newBlock;
  }

  getLastBlock() {
    return this.chain[this.chain.length - 1];
  }

  createNewTransaction(amount, sender, recipient) {
    const newTransaction = {
      transactionId: uuidv4().split("-").join(""),
      amount: amount,
      date: new Date(),
      sender: sender,
      recipient: recipient,
    };

    return newTransaction;
  }

  addTransactionToPendingTransactions(transactionObject) {
    this.pendingTransactions.push(transactionObject);
    return this.getLastBlock()["index"] + 1;
  }

  proofOfWork(previousBlockHash, currentBlockData) {
    let nonce = 0;
    let hash = this.hashBlock(previousBlockHash, currentBlockData, nonce);
    while (hash.substring(0, 4) !== "0000") {
      nonce++;
      hash = this.hashBlock(previousBlockHash, currentBlockData, nonce);
    }
    return nonce;
  }

  hashBlock(previousBlockHash, currentBlockData, nonce) {
    const dataAsString =
      previousBlockHash + nonce.toString() + JSON.stringify(currentBlockData);
    const hash = sha256(dataAsString);
    return hash;
  }

  chainIsValid(blockchain) {
    // TODO: change
    let validChain = true;

    for (var i = 1; i < blockchain.length; i++) {
      const currentBlock = blockchain[i];
      const prevBlock = blockchain[i - 1];
      const blockHash = this.hashBlock(
        prevBlock["hash"],
        {
          transactions: currentBlock["transactions"],
          index: currentBlock["index"],
        },
        currentBlock["nonce"]
      );
      if (blockHash.substring(0, 4) !== "0000") validChain = false;
      if (currentBlock["previousBlockHash"] !== prevBlock["hash"])
        validChain = false;
    }

    //check genesis block validation
    const genesisBlock = blockchain[0];
    const correctNonce = genesisBlock["nonce"] === 100;
    const correctPreviousBlockHash = genesisBlock["previousBlockHash"] === "0";
    const correctHash = genesisBlock["hash"] === "0";
    const correctTransactions = genesisBlock["transactions"].length === 0;

    if (
      !correctNonce ||
      !correctPreviousBlockHash ||
      !correctHash ||
      !correctTransactions
    )
      validChain = false;

    return validChain;
  }

  getBlock(blockHash) {
    let correctBlock = null;
    this.chain.forEach((block) => {
      if (block.hash === blockHash) correctBlock = block;
    });
    return correctBlock;
  }

  getTransaction(transactionId) {
    let correctTransaction = null;
    let correctBlock = null;

    this.chain.forEach((block) => {
      block.transactions.forEach((transaction) => {
        if (transaction.transactionId === transactionId) {
          correctTransaction = transaction;
          correctBlock = block;
        }
      });
    });

    return {
      transaction: correctTransaction,
      block: correctBlock,
    };
  }

  getPendingTransactions() {
    return this.pendingTransactions;
  }

  getAddressData(address) {
    const addressTransactions = [];
    this.chain.forEach((block) => {
      block.transactions.forEach((transaction) => {
        if (
          transaction.sender === address ||
          transaction.recipient === address
        ) {
          addressTransactions.push(transaction); //push all tranasction by sender or recipient into array.
        }
      });
    });

    if (addressTransactions == null) {
      return false;
    }

    var amountArr = [];

    let balance = 0;
    addressTransactions.forEach((transaction) => {
      if (transaction.recipient === address) {
        balance += transaction.amount;
        amountArr.push(balance);
      } else if (transaction.sender === address) {
        balance -= transaction.amount;
        amountArr.push(balance);
      }
    });

    return {
      addressTransactions: addressTransactions,
      addressBalance: balance,
      amountArr: amountArr,
    };
  }
};
