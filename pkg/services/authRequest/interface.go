package authRequestService

import "github.com/villsource/docuval-identity/pkg/models"


type Service interface {
    AddAuthRequest(authRequest *models.AuthRequestModel) error
    GetCodeFlowRequest(authCode string) (*models.AuthRequestModel, error)
    GetTokenFlowRequest(authCode string) (*models.AuthRequestModel, error)
}
