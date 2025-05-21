package authRequestRepository

import "github.com/villsource/docuval-identity/pkg/models"

type Repository interface {
	SaveAuthRequest(authRequest *models.AuthRequestModel) error
	GetAuthRequest(authCode string) (*models.AuthRequestModel, error)
}