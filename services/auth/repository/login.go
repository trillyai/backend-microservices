package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/services/auth/shared"
	"golang.org/x/crypto/bcrypt"
)

// Login method performs user login process.
func (r repository) Login(ctx context.Context, request shared.LoginRequest) (shared.LoginResponse, error) {
	// Check if the user exists in the database.
	request.Username = strings.ToLower(request.Username)
	user, err := postgres.Read[tables.User, tables.User](ctx, map[string]interface{}{"Username": request.Username})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.LoginResponse{}, err
	}
	if user.Username == "" {
		return shared.LoginResponse{}, errors.New("user does not exist") // Return an error if the user does not exist.
	}

	// Compare the hashed password with the provided password.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return shared.LoginResponse{}, errors.New("incorrect password") // Return an error if the password is incorrect.
	}

	// Read the user's session from the database.
	session, err := postgres.Read[tables.Session, tables.Session](ctx, map[string]interface{}{"UserId": user.Id, "EndDate": nil})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.LoginResponse{}, err
	}

	// Update the session's end date to the current time if it exists.
	if session.UserId == user.Id {
		now := time.Now()
		session.EndDate = &now
		_, err := postgres.Update[tables.Session, tables.Session](ctx, map[string]interface{}{"Id": session.Id}, session)
		if err != nil {
			r.logger.Error(err.Error())
			return shared.LoginResponse{}, err
		}
	}

	// Create a new session for the user.
	respSession, err := postgres.Create[tables.Session, tables.Session](ctx, tables.Session{UserId: user.Id})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.LoginResponse{}, err
	}

	// Create a JWT token for the user.
	token, err := auth.CreateJwtToken(user, respSession.Id.String())
	if err != nil {
		r.logger.Error(err.Error())
		return shared.LoginResponse{}, err
	}

	// Return the login response with the username and token.
	return shared.LoginResponse{
		Username: request.Username,
		Token:    token,
	}, nil
}
