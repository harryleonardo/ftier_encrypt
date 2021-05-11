package vo

type (
	// EncryptRequest ...
	EncryptRequest struct {
		Message string `json:"message"`
	}

	// EncryptResponse ...
	EncryptResponse struct {
		EncryptedMessage interface{} `json:"encrypted_message"`
	}
)
