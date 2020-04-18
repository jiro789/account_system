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
        <p>
          <Typography>{`Id: ${id}`}</Typography>
        </p>
        <p>
          <Typography>{`Type: ${type}`}</Typography>
        </p>
        <p>
          <Typography>{`Amount: ${amount}`}</Typography>
        </p>
        <p>
          <Typography>{`Effective Date: ${effective_date}`}</Typography>
        </p>
      </div>
    </ExpansionPanelDetails>
  );
}
