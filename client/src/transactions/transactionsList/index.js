import React, { useEffect, useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import ExpansionPanel from "@material-ui/core/ExpansionPanel";
import TransactionDescription from "../transactionDescription";
import TransactionSummary from "../transactionSummary";
import axios from "axios";

const useStyles = makeStyles((theme) => ({
  heading: {
    fontSize: theme.typography.pxToRem(15),
    fontWeight: theme.typography.fontWeightRegular,
  },
}));

export default function TransactionsList() {
  const classes = useStyles();
  const [transactions, setTransactions] = useState([]);

  async function fetchData() {
    const result = await axios("/api/transaction");
    setTransactions(result.data);
  }

  useEffect(() => {
    fetchData();
  }, []);

  const list = transactions.map((transaction) => {
    return (
      <li key={transaction.id} style={ { "listStyleType": "none" } }>
        <ExpansionPanel>
          <TransactionSummary {...transaction}></TransactionSummary>
          <TransactionDescription {...transaction}></TransactionDescription>
        </ExpansionPanel>
      </li>
    );
  });

  return <ul className={classes.root}>{list}</ul>;
}
