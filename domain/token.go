package domain

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

// constant for token scope.
const (
    ScopeActivation = "activation"
)

// Token struct to hold data for individual token
type Token struct {
    Plaintext string
    Hash []byte
    CustomerID int64
    Expiry time.Time
    Scope string
}

// TokenService represents a service for managing tokens
type TokenService interface {
    // Create
    // GetByCustomerID
}

func GenerateToken(customerID int64, ttl time.Duration, scope string) (*Token, error) {
    // create a Token instance containing the userID, expiry, and scope info.
    // We add the ttl (time-to-live) duration parameter to the current time to get the expiry time
    token := &Token{
        CustomerID: customerID,
        Expiry: time.Now().Add(ttl),
        Scope: scope,
    }

    // initialize zero value byte slice with a length of 16 bytes
    randomBytes := make([]byte, 16)

    // fill the byte slice with random bytes from your operating system's CSPRNG. 
    // This returns an error if the CSPRNG fails to function properly.
    _, err := rand.Read(randomBytes)
    if err != nil {
        return nil, err
    }

    // Encode our byte slice to a base-32-encoded string and assign it to the token 
    // Plaintext field
    token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

    // Generate SHA-256 hash of the plaintext token string. This will be the value
    // that we store in the `hash` field of our db table.
    hash := sha256.Sum256([]byte(token.Plaintext))
    token.Hash = hash[:]

    return token, nil
}



