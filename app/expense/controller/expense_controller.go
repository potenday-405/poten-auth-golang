package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"poten-invitation-golang/app/expense/model"
	"poten-invitation-golang/domain"
	"strconv"
)

type expenseController struct {
	service domain.ExpenseService
}

func NewExpenseController(service domain.ExpenseService) domain.ExpenseController {
	return &expenseController{
		service: service,
	}
}

func (c *expenseController) CreateExpense(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.CreateExpense
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.CreateExpense(ctx, &expense)
	if err != nil {
		log.Printf("error: CreateExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) UpdateExpense(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.UpdateExpense
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.UpdateExpense(ctx, &expense)
	if err != nil {
		log.Printf("error: UpdateExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) DeleteExpense(ctx *gin.Context) {
	var expense model.DeleteExpense
	expense.UserID = ctx.Request.Header.Get("user_id")
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	expense.EventID = ctx.Query("event_id")
	if expense.EventID == "" {
		log.Printf("error: parameter event_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errors.New("parameter event_id not exist")})
		return
	}
	if err := c.service.DeleteExpense(ctx, &expense); err != nil {
		log.Printf("error: DeleteExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}

func (c *expenseController) GetExpense(ctx *gin.Context) {
	var expense model.GetExpense
	expense.UserID = ctx.Request.Header.Get("user_id")
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	expense.EventID = ctx.Query("event_id")
	if expense.EventID == "" {
		log.Printf("error: event_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errors.New("event_id not exist")})
		return
	}
	res, err := c.service.GetExpense(ctx, &expense)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("error: GetExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpenseList(ctx *gin.Context) {
	var expense model.GetExpenseList
	expense.UserID = ctx.Request.Header.Get("user_id")
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	expense.IsInvited = ctx.Query("is_invited")
	if expense.IsInvited == "" {
		log.Printf("error: parameter is_invited not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errors.New("parameter is_invited not exist")})
		return
	}
	offsetOrderType, _ := strconv.Atoi(ctx.Query("offset_order_type"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	expense.Offset = ctx.Query("offset")
	expense.OffsetOrderType = int8(offsetOrderType)
	expense.Order = ctx.Query("order")
	expense.Limit = limit
	expense.Page = page
	res, err := c.service.GetExpenseList(ctx, &expense)
	if err != nil {
		log.Printf("error: GetExpenseList API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpenseTotal(ctx *gin.Context) {
	var expense model.GetExpenseTotal
	expense.UserID = ctx.Request.Header.Get("user_id")
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	offsetOrderType, _ := strconv.Atoi(ctx.Query("offset_order_type"))
	expense.Offset = ctx.Query("offset")
	expense.OffsetOrderType = int8(offsetOrderType)
	res, err := c.service.GetExpenseTotal(ctx, &expense)
	if err != nil {
		log.Printf("error: GetExpenseTotal API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpenseSearch(ctx *gin.Context) {
	var expense model.GetExpenseSearch
	expense.UserID = ctx.Request.Header.Get("user_id")
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	expense.IsInvited = ctx.Query("is_invited")
	expense.Name = ctx.Query("name")
	expense.Order = ctx.Query("order")
	list, err := c.service.GetExpenseSearch(ctx, &expense)
	if err != nil {
		log.Printf("error: GetExpenseSearch API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, list)
}
