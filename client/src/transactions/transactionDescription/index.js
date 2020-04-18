import React from "react";
import ExpansionPanelDetails from "@material-ui/core/ExpansionPanelDetails";
import Typography from "@material-ui/core/Typography";

export default function TransactionDescription({
  id,
  type,
  amount,
  effective_date,
}) {
  return (
    <ExpansionPanelDetails>
      <div>
          <Typography>{`Id: ${id}`}</Typography>
          <Typography>{`Type: ${type}`}</Typography>
          <Typography>{`Amount: ${amount}`}</Typography>
          <Typography>{`Effective Date: ${effective_date}`}</Typography>
      </div>
    </ExpansionPanelDetails>
  );
}
