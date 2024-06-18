package models

import (
	"time"
)

// channel type
type Channel struct {
	ID         string `json:"id" validate:"required"`
	Encryption bool   `json:"encryption" validate:"omitempty"`
	// The public key is stored in plaintext
	PublicKey string `json:"publicKey" validate:"omitempty"`
	// The private key is stored using aes256 encryption. The encrypted password is the user's password and the salt is the channel's Id.
	PrivateKey string    `json:"privateKey" validate:"omitempty"`
	CreatedAt  time.Time `json:"createdAt" validate:"required"`
	UpdatedAt  time.Time `json:"updatedAt" validate:"required"`
}

type ChannelMessage struct {
	ChannelId string `json:"channel" validate:"required"`
	ID        string `json:"id" validate:"required"`
	Body      []byte `json:"body,omitempty" validate:"omitempty"`
}
