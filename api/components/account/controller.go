package account

import (
	"account_system/api/api_errors"
	"account_system/api/common"
	"account_system/api/components/transaction"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	systemAccount *Account = createMainAccount();
)

type modificationInput struct {
	Type transaction.Type `form:"type" json:"Type" binding:"required,oneof=credit debit"`
	Amount common.AmountType `form:"amount" json:"Amount" binding:"required,gt=0"`
}

type getTransactionInput struct {
	Id string `uri:"transactionId" binding:"required,uuid4"`
}

func HandleModification(c *gin.Context) {
	input := modificationInput{}
	inputErr := c.ShouldBindJSON(&input)
	if inputErr != nil {
		c.JSON(http.StatusBadRequest, api_errors.InvalidInputError())
		return
	}
	var result *transaction.Transaction
	var err *api_errors.ApiError
	if input.Type == transaction.CREDIT {
		result, err = systemAccount.Credit(input.Amount)
	} else {
		result, err = systemAccount.Debit(input.Amount)
	}
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func HandleGetAccount(c *gin.Context) {
	balance := systemAccount.GetBalance()
	c.JSON(http.StatusOK, map[string]interface{}{
		"balance": balance,
	})
}

func HandleGetTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, systemAccount.GetTransactions())
}

func HandleGetTransaction(c *gin.Context) {
	input := getTransactionInput{}
	inputErr := c.ShouldBindUri(&input)
	if inputErr != nil {
		c.JSON(http.StatusBadRequest, api_errors.InvalidIdError())
		return
	}
	transaction, err := systemAccount.GetTransaction(common.NewIdFromString(input.Id))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, transaction)
}

