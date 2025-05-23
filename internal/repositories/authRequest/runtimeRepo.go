package authRequestRepository


import (
	"fmt"

	"github.com/villsource/docuval-identity/internal/models"
)

type RuntimeRepository struct {
    storage map[string]*models.AuthRequestModel
}

func NewRuntimeRepository() *RuntimeRepository {
    return &RuntimeRepository{
        storage: make(map[string]*models.AuthRequestModel),
    }
}

func (l *RuntimeRepository) SaveAuthRequest(authRequest *models.AuthRequestModel) error {
    if authRequest == nil {
        return fmt.Errorf("authRequest is nil")
    }
    if authRequest.AuthCode == "" {
        return fmt.Errorf("authCode is empty")
    }
    l.storage[authRequest.AuthCode] = authRequest
    return nil
}

func (l *RuntimeRepository) GetAuthRequest(authCode string) (*models.AuthRequestModel, error) {
    if l.storage == nil {
        return nil, fmt.Errorf("storage is nil")
    }
    if authCode == "" {
        return nil, fmt.Errorf("authCode is empty")
    }
    authRequest, ok := l.storage[authCode]
    if !ok {
        return nil, fmt.Errorf("authRequest not found for authCode: %s", authCode)
    }
    return authRequest, nil
}