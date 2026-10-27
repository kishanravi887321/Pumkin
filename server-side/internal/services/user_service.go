package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/kishanravi887321/pumkin/server-side/dto"
	"github.com/kishanravi887321/pumkin/server-side/internal/models"
	"github.com/kishanravi887321/pumkin/server-side/internal/repository"
	"github.com/kishanravi887321/pumkin/server-side/utils"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(
	repo *repository.UserRepository,
) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(
	ctx context.Context,
	req dto.RegisterUserRequest,
) (*models.User, error) {
	username := strings.TrimSpace(req.Username)
	email := strings.ToLower(strings.TrimSpace(req.Email))

	if len(username) < 3 || len(username) > 30 {
		return nil, errors.New("username must be 3-30 characters")
	}

	if email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	existing, err := s.repo.FindByEmail(ctx, email)

	if err == nil && existing != nil {
		return nil, errors.New("email is already registered")
	}

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("check existing user: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	userId, err := utils.GenerateUserID()
	if err != nil {
		return nil, fmt.Errorf("generate user ID: %w", err)
	}

	user := &models.User{
		UserName:  username,
		UserId:    userId,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().UTC().Unix(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
