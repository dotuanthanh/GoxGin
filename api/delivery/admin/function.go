package admin

import (
	"api-server/pck"
	usecase "api-server/usecase/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	var request SignUpForm
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func LoginHandler(c *gin.Context) {
	var request LoginForm
	if err := c.ShouldBindJSON(&request); err != nil {
		res := pck.HttpErrorResponse{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	response := usecase.CheckLogin(request.Email, request.Password)
	res := pck.HttpResponse{
		Success: true,
		Data:    response,
	}
	c.JSON(http.StatusOK, res)
}
