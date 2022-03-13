package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/kinmaBackend/db/sqlc"
	"github.com/kinmaBackend/util"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username  string 	`json:"username" binding:"required,min=2"`
	Password 	string 	`json:"password" binding:"required,min=6"`
	Email			string 	`json:"email" binding:"required,email"`
	Phone			string  `json:"phone" binding:"required,min=10"`
}

type userResponse struct{
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	Phone							string  	`json:"phone" binding:"required,min=10"`
}

func newUserResponse(user db.User) userResponse{
	return userResponse{
		Username					: user.Username,
		Email							: user.Email,
		PasswordChangedAt	: user.PasswordChangedAt,
		CreatedAt					: user.CreatedAt,
		Phone							: user.Phone,
	}
}
// Server expose method for API
func (server *Server) createUser(ctx *gin.Context){
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedpassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}
	arg := db.CreateUserParams{
		Username			: req.Username,
		HashedPassword: hashedpassword,
		Email					: req.Email,
		Phone					: req.Phone,
	}
	//Implement the DB CRUD
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}
	//We don't want to return password to user
	rsp := newUserResponse(user)

	ctx.JSON(http.StatusOK, rsp)
}

// type loginUserRequest struct {
// 	Username  string 	`json:"username" binding:"required,alphanum"`
// 	Password 	string 	`json:"password" binding:"required,min=6"`
// }

// type loginUserResponse struct {
// 	AccessToken string 				`json:"access_token"`
// 	User				userResponse 	`json:"user"`
// }

type loginUserByMailRequest struct {
	Email			string 	`json:"email" binding:"required,email"`
	Password 	string 	`json:"password" binding:"required,min=6"`
}

type loginUserByMailResponse struct {
	AccessToken string 				`json:"access_token"`
	User				userResponse 	`json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context){
	var req loginUserByMailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByMail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	rsp := loginUserByMailResponse{
		AccessToken	: accessToken,
		User				: newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}