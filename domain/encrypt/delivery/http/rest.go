package http

import (
	"net/http"

	"github.com/ftier-encrypt/domain/encrypt"
	"github.com/ftier-encrypt/shared/vo"
	"github.com/labstack/echo"
)

type encryptHandler struct {
	usecase encrypt.Usecase
}

func EncryptHandler(e *echo.Echo, usecase encrypt.Usecase) {
	handler := encryptHandler{
		usecase: usecase,
	}

	e.POST("aes/encrypt", handler.EncryptMessage)
}

func (handler encryptHandler) EncryptMessage(e echo.Context) error {
	res, err := handler.usecase.AESEncrypterMessage(e)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, vo.EncryptResponse{
		EncryptedMessage: res,
	})
}
