const { v4: uuidv4 } = require("uuid");

module.exports = class Transaction {
  constructor(amount, sender, recipient) {
    this.transactionId = uuidv4().split("-").join("");
    this.amount = amount;
    this.sender = sender;
    this.recipient = recipient;
    this.date = new Date().toString();
  }
};
