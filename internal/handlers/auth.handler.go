package handlers

import (
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/repository"
	"github.com/Roisfaozi/coffee-shop/pkg"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthHandlerInterface interface {
	Login(ctx *gin.Context)
}

type User struct {
	Username string `db:"username" json:"username" form:"username" valid:"type(string),required"`
	Password string `db:"password" form:"password" json:"password,omitempty" valid:"stringlength(6|100)~Password minimal 6,required"`
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

	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		log.Println(err)
		pkg.NewRes(http.StatusBadRequest, &config.Result{Data: err.Error()}).Send(ctx)
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
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data: "Password is Salah",
		}).Send(ctx)
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
