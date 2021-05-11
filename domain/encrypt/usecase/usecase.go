package usecase

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"

	"github.com/ftier-encrypt/domain/encrypt"
	SharedContext "github.com/ftier-encrypt/shared/context"
	SharedVO "github.com/ftier-encrypt/shared/vo"
	"github.com/labstack/echo"
)

type usecase struct {
}

func NewEncryptUsecase() encrypt.Usecase {
	return &usecase{}
}

func (u usecase) AESEncrypterMessage(ctx echo.Context) (interface{}, error) {
	ac := ctx.(*SharedContext.ApplicationContext)
	config := ac.Config

	dto := &SharedVO.EncryptRequest{}
	// bind the payload
	if err := ac.Bind(dto); err != nil {
		return nil, ctx.JSON(http.StatusBadRequest, err)
	}

	fmt.Println("Incoming Request with Message : ", dto.Message)
	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher([]byte(config.GetKey()))
	// if there are any errors, handle them
	if err != nil {
		return nil, ctx.JSON(http.StatusInternalServerError, err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, ctx.JSON(http.StatusInternalServerError, err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, ctx.JSON(http.StatusInternalServerError, err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	return &SharedVO.EncryptResponse{
		EncryptedMessage: gcm.Seal(nonce, nonce, []byte(dto.Message), nil),
	}, nil
}
