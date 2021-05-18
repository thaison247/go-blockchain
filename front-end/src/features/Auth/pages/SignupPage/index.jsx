import React from "react";
import "./styles.scss";
import SignupForm from "./../../components/SignupForm/index";

SignupPage.propTypes = {};

function SignupPage(props) {
  return (
    <div className="container-signuppage">
      <SignupForm></SignupForm>;
    </div>
  );
}

export default SignupPage;
