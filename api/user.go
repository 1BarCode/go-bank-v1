package api

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	"github.com/1BarCode/go-bank-v1/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=4,max=25"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}

type createUserResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, intServerErrorResponse())
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Email: req.Email,
		HashedPassword: hashedPassword,
		FirstName: req.FirstName,
		LastName: req.LastName,
	}

	user, err := s.services.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(fmt.Errorf("username or email already exists")))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userResDTO := createUserResponse{
		Username: user.Username,
		Email: user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		FirstName: user.FirstName,
		LastName: user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, userResDTO)
}