import React, { Component } from "react";
import { connect } from "./api";
import "./App.css";
import logo from "./logo.svg";

class App extends Component {
  constructor(props) {
    super(props);
    // call our connect function and define
    // an anonymous callback function that
    // simply console.log's the received
    // message
    connect(message => {
      console.log(message);
    });
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;