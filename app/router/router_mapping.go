package router

import (
	"ae-test-task/controllers/account"
	"ae-test-task/controllers/ping"
	"github.com/gin-gonic/gin"
)

const pingEndpoint = "/ping"
const insertTransaction = "/account/transaction"
const fetchAccountTransactions = "/account/transactions"
const findTransactionById = "/transactions/:id"

func InitRoutes(r *gin.Engine) {
	r.GET(pingEndpoint, ping.Ping())
	r.POST(fetchAccountTransactions, account.Controller.InsertTransaction)
	r.GET(insertTransaction, account.Controller.FetchTransactionHistory)
	r.GET(findTransactionById, account.Controller.FindTransactionByID)
}
