import React from "react";
import PropTypes from "prop-types";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Paper from "@material-ui/core/Paper";
import Avatar from "@material-ui/core/Avatar";
import "./styles.scss";
const useStyles = makeStyles((theme) => ({
  table: {
    minWidth: 650,
  },
  detail_root: {
    width: "1350px",
    margin: "auto",
    backgroundColor: "#F8F9FA",
  },
  avatar_root: {
    width: "20px",
    height: "20px",
  },
  custom_cell: {
    color: "#3498db",
  },
}));

function createData(txnhash, block, age, from, to, value, txnfee) {
  return { txnhash, block, age, from, to, value, txnfee };
}

const rows = [
  createData(
    "0xd3939e2e60cdf81613",
    "12404373",
    "8 hrs 47 mins ago",
    "0xc8f595e2084db484f8a",
    "0xc8f595e2084db484f8a",
    "0.000401 Ether",
    "0.006846"
  ),
  createData(
    "0xd3939e2e60cdf81613",
    "12404373",
    "8 hrs 47 mins ago",
    "0xc8f595e2084db484f8a",
    "0xc8f595e2084db484f8a",
    "0.000401 Ether",
    "0.006846"
  ),
  createData(
    "0xd3939e2e60cdf81613",
    "12404373",
    "8 hrs 47 mins ago",
    "0xc8f595e2084db484f8a",
    "0xc8f595e2084db484f8a",
    "0.000401 Ether",
    "0.006846"
  ),
  createData(
    "0xd3939e2e60cdf81613",
    "12404373",
    "8 hrs 47 mins ago",
    "0xc8f595e2084db484f8a",
    "0xc8f595e2084db484f8a",
    "0.000401 Ether",
    "0.006846"
  ),
  createData(
    "0xd3939e2e60cdf81613",
    "12404373",
    "8 hrs 47 mins ago",
    "0xc8f595e2084db484f8a",
    "0xc8f595e2084db484f8a",
    "0.000401 Ether",
    "0.006846"
  ),
];
BlockExplorerDetail.propTypes = {};

function BlockExplorerDetail(props) {
  const classes = useStyles();
  return (
    <div className={classes.detail_root}>
      <div className="address-container">
        <div>
          <Avatar
            className={classes.avatar_root}
            alt="Remy Sharp"
            src="https://gitcoin.co/dynamic/avatar/github-changelog-generator"
          />
        </div>
        <span className="address-title">Address</span>
        <span className="address-code">
          0xc8F595E2084DB484f8A80109101D58625223b7C9
        </span>
      </div>
      <div className="balance-container">
        <span className="balance-title">Balance: </span>
        <span className="balance-price">0 Ether</span>
      </div>
      <TableContainer component={Paper}>
        <Table className={classes.table} size="small">
          <TableHead>
            <TableRow>
              <TableCell>TxnHash</TableCell>
              <TableCell>Block</TableCell>
              <TableCell className={classes.custom_cell}>Age</TableCell>
              <TableCell>From</TableCell>
              <TableCell>To</TableCell>
              <TableCell>Value</TableCell>
              <TableCell className={classes.custom_cell}>Txn Fee</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow key={row.name}>
                <TableCell className={classes.custom_cell}>
                  {row.txnhash}
                </TableCell>
                <TableCell className={classes.custom_cell}>
                  {row.block}
                </TableCell>
                <TableCell>{row.age}</TableCell>
                <TableCell>{row.from}</TableCell>
                <TableCell className={classes.custom_cell}>{row.to}</TableCell>
                <TableCell>{row.value}</TableCell>
                <TableCell>{row.txnfee}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
}

export default BlockExplorerDetail;
