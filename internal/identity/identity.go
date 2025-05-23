package identity

import (
	"fmt"

	"github.com/villsource/docuval-identity/internal/models"
	authRequestRepository "github.com/villsource/docuval-identity/internal/repositories/authRequest"
	authRequestService "github.com/villsource/docuval-identity/internal/services/authRequest"
)

type Identity struct {
	AuthRequestService authRequestService.AuthRequestService
}

var repo authRequestRepository.Repository = authRequestRepository.NewRuntimeRepository()
func New() *Identity {
	return &Identity{
		AuthRequestService: *authRequestService.New(&repo),
	}
}



func (i *Identity) Hello() string {
	return "HELLO FROM IDENTITY MODULE."
}

func (i *Identity) CodeFlowRequest(req *models.AuthRequestModel) error {
	if req.ResponseType != "authorization_code" {
		return fmt.Errorf("grant type is not code flow")
	}
	if req.ChallengeCode == "" {
		return fmt.Errorf("require PKEC")
	}
	return nil
}
