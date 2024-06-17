package models

import (
	"time"
)

// webhook type
type Webhook struct {
	ID         string `json:"id" validate:"required,uuid"`
	Encryption bool   `json:"encryption" validate:"omitempty"`
	// The public key is stored in plaintext
	PublicKey string `json:"publicKey" validate:"omitempty"`
	// The private key is stored using aes256 encryption. The encrypted password is the user's password and the salt is the webhook's Id.
	PrivateKey string    `json:"privateKey" validate:"omitempty"`
	CreatedAt  time.Time `json:"createdAt" validate:"required"`
	UpdatedAt  time.Time `json:"updatedAt" validate:"required"`
}
