package identity

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/villsource/docuval-identity/configs/endpoints"
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

var identityService *Identity = New()

func NewFiberMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		if c.Path() == endpoints.OIDC.Auth {
			authRequest := &models.AuthRequestModel{
				RequestTime:     c.Context().ConnTime().Unix(),
				ChallengeCode:   c.Query("code_challenge"),
				ChallengeMethod: c.Query("code_challenge_method"),
				ClientSecret:    c.Query("client_secret"),
				RedirectURI:     c.Query("redirect_uri"),
				Scope:           c.Query("scope"),
				State:           c.Query("state"),
				ClientID:        c.Query("client_id"),
				ResponseType:    c.Query("response_type"),
			}

			if authRequest.ClientID == "" {
				return c.Status(fiber.StatusBadRequest).SendString("require client_id")
			}

			if authRequest.ResponseType == "" {
				return c.Status(fiber.StatusBadRequest).SendString("require response_type")
			}

			if authRequest.RedirectURI == "" {
				return c.Status(fiber.StatusBadRequest).SendString("require redirect_uri")
			}

			// Not Implemented yet another flow
			if authRequest.ResponseType != "code" {
				return c.Status(fiber.StatusBadRequest).SendString("not support response type ")
			}

			// Code Flow Logic
			authRequest.AuthCode = fmt.Sprintf("mock-%x", uuid.New().String())

			if authRequest.ChallengeCode == "" {
				return c.Status(fiber.StatusBadRequest).SendString("require PKEC")
			}

			if identityService.authRequestService.AddAuthRequest(authRequest) != nil {
				return c.Status(fiber.StatusBadRequest).SendString("failed to add auth request")
			}

			callbackURI := authRequest.RedirectURI + "?code=" + authRequest.AuthCode + "&state=" + authRequest.State

			return c.Redirect(callbackURI, fiber.StatusFound)
		}

		if c.Path() == endpoints.OIDC.Token {
            // CODE FLOW
			tokenRequest := &models.AuthRequestModel{
				RequestTime:  c.Context().ConnTime().Unix(),
				RedirectURI:  c.FormValue("redirect_uri"),
				State:        c.FormValue("state"),
				ClientID:     c.FormValue("client_id"),
				ResponseType: c.FormValue("grant_type"),
				AuthCode:     c.FormValue("code"),
			}

            // if tokenRequest.ResponseType != "authorization_code"{
            //     return c.Status(fiber.StatusBadRequest).SendString("auth flow not support")
            // }

            fmt.Println("tokenRequest", tokenRequest)

            codeVerify := c.FormValue("code_verifier")

            auth, err := identityService.authRequestService.GetCodeFlowRequest(tokenRequest.AuthCode)
            if err != nil {
                return c.Status(fiber.StatusBadRequest).SendString("Code not found" + err.Error())
            }

            if auth.ChallengeCode != codeVerify {
                return c.Status(fiber.StatusBadRequest).SendString("Code verifier not match")
            }

            // if auth.ClientID != tokenRequest.ClientID {
            //     return c.Status(fiber.StatusBadRequest).SendString("Client ID not match")
            // }

            return c.JSON(fiber.Map{
                "access_token":  "MOCK_ACCESS_TOKEN" + tokenRequest.ResponseType,
                "token_type":    "bearer",
                "expires_in":    3600,
                "refresh_token": "MOCK_REFRESH_TOKEN",
                "scope":         tokenRequest.Scope,
            })

		}

		if c.Path() == "/identity-health-check" {
			checkCode := c.Query("check_code")
			return c.SendString(c.Method() + " " + c.Path() + " " + checkCode)
		}
		c.Locals("userID", "anonymous")
		return c.Next()
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
