import React, { useState } from "react";
import PropTypes from "prop-types";
import "./styles.scss";
import Avatar from "@material-ui/core/Avatar";
import { Button, makeStyles, TextField } from "@material-ui/core";
import AccountBalanceWalletOutlinedIcon from "@material-ui/icons/AccountBalanceWalletOutlined";
import AutorenewIcon from "@material-ui/icons/Autorenew";
import ExploreIcon from "@material-ui/icons/Explore";
import LocalAtmIcon from "@material-ui/icons/LocalAtm";
import socketIOClient from "socket.io-client";
import { useEffect } from "react";
import transactionApi from "../../../../api/transactionApi";
import { useForm } from "react-hook-form";
import { ToastContainer, toast } from "react-toastify";
import { Link } from "react-router-dom";

const useStyles = makeStyles((theme) => ({
  root_account_info: {
    // width: "80%",
    backgroundColor: "#fff",
    padding: "20px",
    height: "94vh",
    display: "flex",
    flexDirection: "column",
  },
  address: {
    height: theme.spacing(10),
    width: theme.spacing(10),
    marginLeft: "20px",
    border: "solid 5px #fff",
  },
  balance_icon: {
    color: "#fff",
    height: theme.spacing(10),
    width: theme.spacing(10),
    fontWeight: "400",
    marginLeft: "20px",
  },
  refresh_icon: {
    marginTop: "10px",

    width: "30px",
    height: "30px",
  },
  explorer_icon: {
    color: "#fff",
    height: theme.spacing(10),
    width: theme.spacing(10),
    fontWeight: "400",
    marginLeft: "20px",
    marginRight: "20px",
  },
  input: {
    backgroundColor: "#f9f9f9",
    margin: "5px 0",
  },
  button: {
    margin: "20px 0",
    width: "200px",
    background: "linear-gradient(45deg, #2196F3 30%, #21CBF3 90%)",
  },
}));

AccountInfoPage.propTypes = {};

function AccountInfoPage(props) {
  const classes = useStyles();
  const [addressBalance, setAddressBalance] = useState(0);
  const { reset, register, handleSubmit } = useForm({});

  useEffect(() => {
    const socket = socketIOClient("http://localhost:4000");
    socket.on("PT", (data) => {});
    return () => socket.disconnect();
  }, []);

  useEffect(() => {
    const fetchAddressData = async () => {
      const result = await transactionApi.getAddressData();
      const addressData = result.data;
      setAddressBalance(addressData.addressData.addressBalance);
    };
    fetchAddressData();
  }, [addressBalance]);

  const onSubmitSendTransaction = async (data) => {
    console.log(data);
    const result = await transactionApi.createPendingTransaction({
      ...data,
      sender: localStorage.getItem("publicKey"),
    });
    if (result.data.note) {
      toast.success("Success");
      reset({});
    }
  };

  return (
    <div className={classes.root_account_info}>
      <div className="dashboard">
        <div className="address">
          {/* <Avatar
            className={classes.address}
            alt="Remy Sharp"
            src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/59/User-avatar.svg/1024px-User-avatar.svg.png"
          /> */}
          <AccountBalanceWalletOutlinedIcon
            className={classes.balance_icon}
          ></AccountBalanceWalletOutlinedIcon>
          <div className="address-info">
            <h6 className="title">Address</h6>
            <div className="content">{localStorage.getItem("publicKey")}</div>
          </div>
        </div>
        <div className="balance">
          <LocalAtmIcon className={classes.balance_icon}></LocalAtmIcon>
          <div className="balance-info">
            <h6 className="title">Balance</h6>
            <span className="blance-content">{`${addressBalance} COINS`}</span>
          </div>
        </div>
        <div className="explorer">
          <Link
            className="explorer__body"
            style={{ textDecoration: "none" }}
            to="/explored"
          >
            <ExploreIcon className={classes.explorer_icon}></ExploreIcon>
            <h6 className="title">Explorer</h6>
          </Link>
        </div>
      </div>
      <form
        className="form-transaction"
        onSubmit={handleSubmit(onSubmitSendTransaction)}
      >
        <h3 className="transaction-title">Send Transaction</h3>
        <div className="amount">
          <span>Amount</span>
          <TextField
            className={classes.input}
            id="standard-search"
            label=""
            {...register("amount")}
          />
        </div>
        <div className="to-address">
          <span> To Adress</span>
          <TextField
            className={classes.input}
            id="standard-search"
            label=""
            {...register("recipient")}
          />
        </div>
        <Button
          className={classes.button}
          variant="contained"
          color="primary"
          disableElevation
          type="submit"
        >
          Send
        </Button>
      </form>
      <ToastContainer />
    </div>
  );
}

export default AccountInfoPage;
