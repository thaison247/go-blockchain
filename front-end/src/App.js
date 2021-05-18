import { Route } from "react-router-dom";
import "./App.css";
import SignupPage from "./features/Auth/pages/SignupPage/index";
import LoginPage from "./features/Auth/pages/LoginPage/index";
import AccountInfoPage from "./features/Transaction/pages/AccountInfoPage/index";
import Explorer from "./components/Explorer/index";

function App() {
  return (
    <div className="App">
      <Route path="/" component={SignupPage} exact></Route>
      <Route path="/signup" component={SignupPage}></Route>
      <Route path="/login" component={LoginPage}></Route>
      <Route path="/dashboard" component={AccountInfoPage}></Route>
      <Route path="/explored" component={Explorer}></Route>
    </div>
  );
}

export default App;
