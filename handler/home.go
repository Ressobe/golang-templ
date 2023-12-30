package handler

import (
	user "github.com/Ressobe/golang+templ/view"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {}

func (h UserHandler) HandleUserShow(c echo.Context) error {
  return render(c, user.Show()) 
}
