import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import Typography from "@material-ui/core/Typography";
import "./styles.scss";
import { makeStyles } from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import authApi from "./../../../../api/authApi";
import { Link } from "react-router-dom";

const useStyles = makeStyles((theme) => ({
  title: {
    fontWeight: "500",
    fontSize: "30px",
    lineHeight: "42px",
    marginBottom: "15px",
    margin: "auto",
  },
  root: {
    display: "flex",
    justifyContent: "center",

    "& > *": {
      margin: theme.spacing(1),
      width: "50%",
    },
  },
  button: {
    background: "linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)",
    width: "200px",
    height: "50px",
    color: "#fff",
    textDecoration: "none",
  },
}));

SignupForm.propTypes = {};

function SignupForm(props) {
  const classes = useStyles();

  const [wordList, setWordList] = useState([]);
  useEffect(() => {
    const fetchSignUpWords = async () => {
      const wordArray = await authApi.signup();
      setWordList(wordArray.data.mnemonic.split(" "));
      console.log(wordArray.data.mnemonic.split(" "));
    };
    fetchSignUpWords();
  }, []);

  return (
    <div className="root-login">
      <Typography className={classes.title} variant="h1" component="h2">
        Create A New Wallet
      </Typography>
      <div className="sub-root">
        <span className="sub_title">Already have a wallet?</span>
        <a className="sub_title-link" href="">
          Access My Wallet
        </a>
      </div>
      <h4>
        Write 12 words down on your pager. Resist temptation to email it to
        yourself or screenshot it.
      </h4>
      <form className={classes.root} noValidate autoComplete="off">
        <Grid container spacing={3}>
          {wordList.map((item, index) => {
            return (
              <Grid item xs={4} key={index}>
                <TextField
                  id="standard-basic"
                  label={`${index + 1}. ${item}`}
                />
              </Grid>
            );
          })}
        </Grid>
      </form>
      <Link style={{ textDecoration: "none" }} to="/login">
        <Button className={classes.button} variant="contained" color="primary">
          Start
        </Button>
      </Link>
    </div>
  );
}

export default SignupForm;
