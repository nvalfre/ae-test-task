package account

import (
	"ae-test-task/app/memory"
	"ae-test-task/domain"
	"ae-test-task/services"
	"ae-test-task/services/builder"
	"ae-test-task/services/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var db = memory.New()
var accountStorage = memory.NewAccountStoreService(db)
var Controller accountControllerInterface

type accountControllerInterface interface {
	CreateNewAccount(user string)
	InsertTransaction(c *gin.Context)
	FetchTransactionHistory(c *gin.Context)
	FindTransactionByID(c *gin.Context)
	GetAccountBalance(c *gin.Context)
}
type accountController struct {
	AccountService services.AccountService
}

func init() {
	Controller = &accountController{
		services.AccountService{
			AccountBuilderService: builder.AccountBuilderService{
				Store: accountStorage,
			},
			Store:                accountStorage,
			AccountValidator:     validator.AccountValidatorService{},
			TransactionValidator: validator.TransactionValidatorService{},
		},
	}
	Controller.CreateNewAccount("dev-user")
}

func (controller *accountController) GetAccountBalance(c *gin.Context) {
	account, err := controller.AccountService.GetAccount()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_controller",
			"service": "get",
			"method":  "get account",
			"err":     err,
			"message": "cannot get account",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: account,
	})
}

func (controller *accountController) CreateNewAccount(user string) {
	account, err := controller.AccountService.CreateAccount(user)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_controller",
			"service": "create",
			"method":  "CreateNewAccount",
			"err":     err,
			"message": "cannot create account",
		})
		return
	}
	fmt.Print(account)
}

func (controller *accountController) InsertTransaction(c *gin.Context) {
	var tran *domain.TransactionReq
	if err := c.ShouldBindJSON(&tran); err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_controller",
			"service": "bind",
			"method":  "InsertTransaction",
			"err":     err,
			"message": "cannot bind json transaction",
		})
		c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: "cannot bind json transaction",
		})
		return
	}
	err, transaction := controller.executeTransaction(tran)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_controller",
			"service": "start",
			"method":  "InsertTransaction",
			"err":     err,
			"message": "cannot insert transaction",
		})
		c.JSON(http.StatusInternalServerError, Response{
			Status:  http.StatusInternalServerError,
			Message: "cannot insert transaction",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: transaction,
	})
	return
}

func (controller *accountController) executeTransaction(tran *domain.TransactionReq) (error, *domain.TransactionBody) {
	trans := make(chan *domain.TransactionBody, 1)
	errs := make(chan error, 1)

	go func() {
		defer close(trans)
		defer close(errs)
		transaction, err := controller.AccountService.InsertTransaction(tran)
		if err != nil {
			errs <- err
		}
		trans <- transaction
	}()
	time.Sleep(1 * time.Millisecond)
	err := <-errs
	transaction := <-trans
	return err, transaction
}

func (controller *accountController) FetchTransactionHistory(c *gin.Context) {
	transactionHistory, err := controller.AccountService.TransactionHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprint("Can not get transaction history"),
		})
	}

	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: &transactionHistory,
	})
}

func (controller *accountController) FindTransactionByID(c *gin.Context) {
	id := c.Params.ByName("id")

	transaction, err := controller.AccountService.FindTransactionByID(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_controller",
			"service": "search by id",
			"method":  "FindTransactionByID",
			"error":   err,
		})
		c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	if transaction == nil {
		logrus.WithFields(logrus.Fields{
			"file":    "account_controller",
			"service": "search by id",
			"method":  "FindTransactionByID",
			"error":   err,
		})
		c.JSON(http.StatusNotFound, Response{
			Status:  http.StatusNotFound,
			Message: "Not found any transaction with that id",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: transaction,
	})
	return
}
