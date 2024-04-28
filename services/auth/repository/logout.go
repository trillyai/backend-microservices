package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/services/auth/shared"
)

func (r repository) Logout(ctx context.Context, request shared.LogoutRequest) (shared.LogoutResponse, error) {

	claims := ctx.Value("user").(*auth.Claims)
	if claims.Name == "" {
		return shared.LogoutResponse{}, errors.New("context not found")
	}

	// get session
	session, err := postgres.Read[tables.Session, tables.Session](ctx, map[string]interface{}{"Id": claims.SessionId, "EndDate": nil})
	if err != nil {
		return shared.LogoutResponse{}, err
	}

	if session.Id == uuid.Nil {
		return shared.LogoutResponse{}, errors.New("session not found")
	}

	// end session
	now := time.Now()
	session.EndDate = &now

	resp, err := postgres.Update[shared.LogoutResponse, tables.Session](ctx, map[string]interface{}{"Id": session.Id}, session)
	if session.Id != claims.SessionId {
		return shared.LogoutResponse{}, err
	}
	resp.Username = claims.UserName
	return resp, nil
}
