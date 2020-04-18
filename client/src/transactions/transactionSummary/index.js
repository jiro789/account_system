import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import ExpansionPanelSummary from "@material-ui/core/ExpansionPanelSummary";
import Typography from "@material-ui/core/Typography";
import ExpandMoreIcon from "@material-ui/icons/ExpandMore";

const useStyles = makeStyles((theme) => ({
  content: {
    fontSize: theme.typography.pxToRem(15),
    fontWeight: theme.typography.fontWeightRegular,
  },
  container: {
    backgroundColor: (props) =>
      props.type === "credit" ? "aquamarine" : "tomato",
  },
}));

export default function TransactionSummary({ type, amount }) {
  const classes = useStyles({ type });

  return (
    <ExpansionPanelSummary
      className={classes.container}
      expandIcon={<ExpandMoreIcon />}
      aria-controls="panel1a-content"
      id="panel1a-header"
    >
      <Typography
        className={classes.content}
      >{`Type: ${type} Amount: ${amount}`}</Typography>
    </ExpansionPanelSummary>
  );
}
