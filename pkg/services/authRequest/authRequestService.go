package authRequestService

import (
	"fmt"

	"github.com/villsource/docuval-identity/pkg/models"
	"github.com/villsource/docuval-identity/pkg/repositories/authRequest"
)

type AuthRequestService struct {
    repo authRequestRepository.Repository
}

func New(authRequestRepository *authRequestRepository.Repository) *AuthRequestService{
    return &AuthRequestService{
        repo: *authRequestRepository,
    }
}

func (s *AuthRequestService) AddAuthRequest(authRequest *models.AuthRequestModel) error{
    ok := s.repo.SaveAuthRequest(authRequest)
    if ok != nil {
        return ok
    }
    return nil
}
func (s *AuthRequestService)GetCodeFlowRequest(authCode string) (*models.AuthRequestModel, error){
    val, ok :=s.repo.GetAuthRequest(authCode)
    return val, ok
}
func (s *AuthRequestService) GetTokenFlowRequest(authCode string) (*models.AuthRequestModel, error){
    return nil, fmt.Errorf("not implement")
}
