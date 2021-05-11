package encrypt

import "github.com/labstack/echo"

type Usecase interface {
	AESEncrypterMessage(ctx echo.Context) (interface{}, error)
}
