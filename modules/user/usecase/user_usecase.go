package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/ditdittdittt/backend-sitpi/config"
	"github.com/ditdittdittt/backend-sitpi/domain"

	"github.com/dgrijalva/jwt-go"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func (uc *userUsecase) Register(ctx context.Context, request *domain.RegisterUserRequest) (jwtToken string, err error) {
	panic("implement me")
}

func (uc *userUsecase) Login(ctx context.Context, request *domain.LoginUserRequest) (jwtToken string, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.userRepository.GetByUsername(ctx, request.Username)
	if err != nil {
		return "", nil
	}
	if existedUser == (domain.User{}) {
		return "", domain.ErrNotFound
	}

	if request.Password != existedUser.Password {
		return "", errors.New("wrong password")
	}

	jwtToken, err = generateJwt(existedUser)
	if err != nil {
		return "", nil
	}

	return jwtToken, err
}

func (uc *userUsecase) ChangePassword(ctx context.Context, request *domain.ChangePasswordRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	// Get current login information

	// Check old password

	// Hash new password

	// Update with new password

	return nil
}

func NewUseUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: u,
		contextTimeout: timeout,
	}
}

func generateJwt(user domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = user.ID
	claims["role_id"] = user.RoleID
	claims["status_id"] = user.StatusID
	claims["tpi_id"] = user.TpiID
	claims["nik"] = user.Nik
	claims["name"] = user.Name
	claims["address"] = user.Address
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(config.JwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
