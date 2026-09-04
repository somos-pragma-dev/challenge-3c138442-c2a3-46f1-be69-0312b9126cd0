package handlers

import (
	"github.com/gin-gonic/gin"
	"internal/app/models"
	"internal/app/services"
	"net/http"
)

type AccountHandler struct {
	accountService *services.AccountService
}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{
		accountService: services.NewAccountService(),
	}
}

func (h *AccountHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/accounts", h.CreateAccount)
	r.GET("/accounts/:id", h.GetAccount)
	r.PUT("/accounts/:id", h.UpdateAccount)
	r.DELETE("/accounts/:id", h.DeleteAccount)
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.accountService.CreateAccount(&account); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, account)
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	account, err := h.accountService.GetAccount(id)
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.Account
	if err := c.ShouldBindJSON(&account); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.accountService.UpdateAccount(id, &account); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if err := h.accountService.DeleteAccount(id); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}