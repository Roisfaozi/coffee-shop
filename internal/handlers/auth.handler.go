package handlers

import (
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/Roisfaozi/coffee-shop/pkg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"password" json:"password" form:"password"`
}
type AuthHandler struct {
	UserRepository repository.UserRepository
}

func NewAuthHandler(r repository.UserRepository) *AuthHandler {
	return &AuthHandler{r}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var data User
	if err := ctx.ShouldBind(&data); err != nil {
		pkg.NewRes(http.StatusInternalServerError, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	user, err := h.UserRepository.GetAuthUser(data.Username)
	if err != nil {
		log.Println(err.Error())
		pkg.NewRes(http.StatusUnauthorized, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	if err := pkg.VerifyPassword(user.Password, data.Password); err != nil {
		log.Println(err.Error())
		pkg.NewRes(http.StatusUnauthorized, &config.Result{
			Data: "Password is Salah",
		})
		return
	}

	jwt := pkg.NewToken(user.ID, user.Role)
	token, err := jwt.Generate()
	if err != nil {
		log.Println(err.Error())
		pkg.NewRes(http.StatusInternalServerError, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(http.StatusOK, &config.Result{Data: token}).Send(ctx)
}
