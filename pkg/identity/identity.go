package identity

import (
	"fmt"

	"github.com/villsource/docuval-identity/pkg/models"
	authRequestRepository "github.com/villsource/docuval-identity/pkg/repositories/authRequest"
	authRequestService "github.com/villsource/docuval-identity/pkg/services/authRequest"
)

type Identity struct {
    authRequestService authRequestService.AuthRequestService
}

func New() *Identity {
    var repo authRequestRepository.Repository = authRequestRepository.NewRuntimeRepository()
    return &Identity{
        authRequestService: *authRequestService.New(&repo),
    }
}

func (i *Identity) Hello() string{
    return "HELLO FROM IDENTITY MODULE."
}


func (i *Identity) CodeFlowRequest(req *models.AuthRequestModel) error{
    if req.GrantType != "authorization_code" {
        return fmt.Errorf("grant type is not code flow")
    }
    if req.ChallengeCode == "" {
        return fmt.Errorf("require PKEC")
    }
    return nil
}