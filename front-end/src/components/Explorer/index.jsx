import React, { useEffect } from "react";
import "./styles.scss";
import { Grid, makeStyles, Paper } from "@material-ui/core";
import socketIOClient from "socket.io-client";
import blockchainApi from "../../api/blockchainApi";
import { useState } from "react";
import DayJS from "react-dayjs";
import Button from "@material-ui/core/Button";
import axios from "axios";
import { toast, ToastContainer } from "react-toastify";

const useStyles = makeStyles((theme) => ({
  explorer_root: {
    width: "100%",
    height: "100%",
  },
  paper: {
    margin: "20px 5px",
    height: "80px",
    display: "flex",
    justifyContent: "center",
  },
  lastest_block: {
    display: "flex",
    margin: "10px 10px",
    height: "50px",
    justifyContent: "space-between",
    alignItems: "center",
  },
  lastest_transaction: {
    display: "flex",
    margin: "10px 10px",
    height: "50px",
    justifyContent: "space-between",
    alignItems: "center",
  },
  icon: {
    height: 48,
  },
}));
Explorer.propTypes = {};

function Explorer(props) {
  const classes = useStyles();
  const [chain, setChain] = useState([]);
  const [pendingTransactions, setPendingTransactions] = useState([]);
  const [transactions, setTransactions] = useState([]);
  useEffect(() => {
    const socket = socketIOClient("http://localhost:4000");
    socket.on("PT", (data) => {
      setPendingTransactions(data);
      console.log(data);
    });
    socket.on("B", (data) => {
      setChain(data);
    });
    socket.on("T", (data) => {
      setTransactions(data);
    });
    return () => socket.disconnect();
  }, []);

  useEffect(() => {
    const fetchBlockChain = async () => {
      const result = await blockchainApi.getBlockChain();
      const { chain, pendingTransactions, transactions } = result.data;
      setChain(chain);
      setPendingTransactions(pendingTransactions);
      setTransactions(transactions);
      console.log(pendingTransactions);
    };

    fetchBlockChain();
  }, []);

  const onClickMine = async () => {
    const hdKey = JSON.parse(localStorage.getItem("hdKey"));
    const res = await axios.post("http://localhost:4000" + "/mine", {
      clientHdKey: hdKey,
    });

    if (res.data.note) {
      toast.success(
        `New block has index of ${res.data.block.index} is mined and broadcaset successfuly!`
      );
    }
  };

  return (
    <div className={classes.explorer_root}>
      <div className="transactions">
        <div className={classes.root}>
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <Paper className={classes.paper}>
                <span className="transaction_title">TRANSACION: </span>
                <span className="transaction_price">{transactions.length}</span>
              </Paper>
            </Grid>

            <Grid item xs={12} sm={6}>
              <span className="latest_block-first-title">Lastest Blocks</span>
              {chain.map((block, index) => {
                return (
                  <Paper key={index} className={classes.lastest_block}>
                    <div className="lastest_block-first">
                      <span className="lastest_block-first-icon">BK</span>
                      <div className="lastest_block-first-content">
                        <span className="lastest_block-first-content-info">
                          {block.index}
                        </span>
                        <DayJS format="DD-MM-YYYY">{block.date}</DayJS>
                      </div>
                    </div>
                    <div className="lastest_block-second">
                      <div className="lastest_block-second-content">
                        <div className="lastest_block-second-content-top">
                          <span className="lastest_block-second-content-info">
                            Miner
                          </span>
                          <a href="">Spark pool</a>
                        </div>

                        <div className="lastest_block-second-content-bottom">
                          <a href="">{`${block.transactions.length}txns`}</a>
                        </div>
                      </div>
                    </div>
                    <div className="lastest_block-third">
                      <span className="lastest_block-third-price">
                        {`${
                          block.transactions.reduce((total, transaction) => {
                            return total + transaction.amount;
                          }, 0) - 10
                        } Eth`}
                      </span>
                    </div>
                  </Paper>
                );
              })}
            </Grid>

            <Grid item xs={12} sm={6}>
              <span className="latest_transaction-title">
                Lastest Transaction
              </span>
              {transactions.map((transaction, index) => {
                return (
                  <Paper key={index} className={classes.lastest_transaction}>
                    <div className="latest-transaction-first">
                      <span className="latest-transaction-first-icon">TX</span>
                      <div className="latest-transaction-first-content">
                        <span className="latest-transaction-first-content-info">
                          {transaction.transactionId}
                        </span>
                        <DayJS format="DD-MM-YYYY">{transaction.date}</DayJS>
                      </div>
                    </div>
                    <div className="latest-transaction-second">
                      <div className="latest-transaction-second-content">
                        <div className="latest-transaction-second-content-top">
                          <span className="latest-transaction-second-content-info">
                            From
                          </span>
                          <a href="">{transaction.sender}</a>
                        </div>

                        <div className="latest-transaction-second-content-bottom">
                          <span className="latest-transaction-second-content-info">
                            To
                          </span>
                          <a href="">{transaction.recipient}</a>
                        </div>
                      </div>
                    </div>
                    <div className="latest-transaction-third">
                      <span className="latest-transaction-third-price">
                        {`${transactions.reduce((total, transaction) => {
                          return total + transaction.amount;
                        }, 0)} Eth`}
                      </span>
                    </div>
                  </Paper>
                );
              })}
            </Grid>
            <Grid item xs={12}>
              <span className="latest_transaction-title">
                Pending Transaction
              </span>
              {pendingTransactions.map((transaction, index) => {
                return (
                  <Paper key={index} className={classes.lastest_transaction}>
                    <div className="latest-transaction-first">
                      <span className="latest-transaction-first-icon">TX</span>
                      <div className="latest-transaction-first-content">
                        <span className="latest-transaction-first-content-info">
                          {transaction.transactionId}
                        </span>
                        <DayJS format="DD-MM-YYYY">{transaction.date}</DayJS>
                      </div>
                    </div>
                    <div className="latest-transaction-second">
                      <div className="latest-transaction-second-content">
                        <div className="latest-transaction-second-content-top">
                          <span className="latest-transaction-second-content-info">
                            From
                          </span>
                          <a href="">{transaction.sender}</a>
                        </div>

                        <div className="latest-transaction-second-content-bottom">
                          <span className="latest-transaction-second-content-info">
                            To
                          </span>
                          <a
                            style={{ display: "block", width: "600px" }}
                            href=""
                          >
                            {transaction.recipient}
                          </a>
                        </div>
                      </div>
                    </div>
                    <div className="latest-transaction-third">
                      <span className="latest-transaction-third-price">
                        {`${pendingTransactions.reduce((total, transaction) => {
                          return total + transaction.amount;
                        }, 0)} Eth`}
                      </span>
                    </div>
                    <Button
                      variant="contained"
                      color="primary"
                      href="#contained-buttons"
                      onClick={onClickMine}
                    >
                      Mine
                    </Button>
                  </Paper>
                );
              })}
            </Grid>
          </Grid>
        </div>
      </div>
    </div>
  );
}

export default Explorer;
