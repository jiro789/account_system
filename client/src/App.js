import React from "react";
import TransactionsList from "./transactions/transactionsList";
import "./App.css";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <TransactionsList></TransactionsList>
      </header>
    </div>
  );
}

export default App;
