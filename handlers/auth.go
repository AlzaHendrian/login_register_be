package handlers

import (
	"fmt"
	"log"
	"net/http"
	authdto "test_fullstack/dto/auth"
	dto "test_fullstack/dto/result"
	"test_fullstack/models"
	"test_fullstack/pkg/bcrypt"
	jwtToken "test_fullstack/pkg/jwt"
	"test_fullstack/repositories"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c echo.Context) error {
	request := new(authdto.AuthRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateOfBirth: request.DateOfBird,
		Gender:      request.Gender,
		Email:       request.Email,
		Password:    password,
		Role:        request.Role,
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	log.Printf("Login request data: %+v\n", request)

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// Check email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	log.Printf("User retrieved from database: %+v\n", user)

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		log.Printf("Password does not match for user %s\n", request.Email)
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"})
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginResponse{
		Name:     user.FirstName,
		LastName: user.LastName,
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
		IsAdmin:  user.IsAdmin,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: loginResponse})
}

func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}

func SeedDummyCredentials(db *gorm.DB) error {
	// Define the admin email and password
	adminEmail := "admin@mail.com"
	password := "admin123"

	// Check if the admin already exists in the database
	var admin models.User
	if err := db.Where("email = ?", adminEmail).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// If admin does not exist, create the admin account
			hashedPassword, err := bcrypt.HashingPassword(password)
			if err != nil {
				return err
			}

			admin = models.User{
				Email:    adminEmail,
				Password: hashedPassword,
				IsAdmin:  true,
			}

			if err := db.Create(&admin).Error; err != nil {
				return err
			}

			fmt.Println("Admin account created successfully")
		} else {
			return err
		}
	} else {
		fmt.Println("Admin account already exists")
	}

	return nil
}
