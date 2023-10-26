package models

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Conn struct {
	db *gorm.DB
}

// CreatJob implements Service.
func (*Conn) CreatJob(ctx context.Context, ni NewJob, CompanyID uint) (Job, error) {
	panic("unimplemented")
}

func NewService(db *gorm.DB) (*Conn, error) {
	if db == nil {
		return nil, errors.New("please provide a valid connection")
	}
	s := &Conn{db: db}
	return s, nil
}
func (s *Conn) CreateUser(ctx context.Context, nu NewUser) (User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("generating password hash: %w", err)
	}
	u := User{
		Name:         nu.Name,
		PhoneNumber:  nu.PhoneNumber,
		Email:        nu.Email,
		PasswordHash: string(hashedPass),
	}
	err = s.db.Create(&u).Error
	if err != nil {
		return User{}, err
	}
	return u, nil
}
func (s *Conn) Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims,
	error) {
	var u User
	tx := s.db.Where("email = ?", email).First(&u)
	if tx.Error != nil {
		return jwt.RegisteredClaims{}, tx.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	c := jwt.RegisteredClaims{
		// Issuer:    "service project",
		Subject: strconv.FormatUint(uint64(u.ID), 10),
		// Audience:  jwt.ClaimStrings{"students"},
		// ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		// IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	return c, nil
}
func (s *Conn) AutoMigrate() error {
	//if s.db.Migrator().HasTable(&User{}) {
	//	return nil
	//}

	err := s.db.Migrator().DropTable(&User{}, &Job{})
	if err != nil {
		return err
	}

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err = s.db.Migrator().AutoMigrate(&User{}, &Job{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	return nil
}
