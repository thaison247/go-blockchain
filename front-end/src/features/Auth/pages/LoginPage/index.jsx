import React from "react";
import PropTypes from "prop-types";
import LoginForm from "./../../components/LoginForm/index";
import authApi from "./../../../../api/authApi";
import { useHistory } from "react-router";

LoginPage.propTypes = {};

function LoginPage(props) {
  const history = useHistory();
  const handleSubmitLoginForm = async (data) => {
    const wordListArray = Object.values(data);
    const wordListString = wordListArray.join(" ");
    const result = await authApi.login({ mnemonic: wordListString });
    console.log("resssssssssssss: ", result);
    localStorage.setItem("hdKey", JSON.stringify(result.data.hdKey));
    localStorage.setItem("publicKey", result.data.publicKey);
    history.push({
      pathname: "/dashboard",
    });
  };
  return (
    <div>
      <LoginForm handleSubmitLoginForm={handleSubmitLoginForm}></LoginForm>
    </div>
  );
}

export default LoginPage;
