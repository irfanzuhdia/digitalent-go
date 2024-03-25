package utils

import (
	"fmt"
	"mygram/domain"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string
	Error string
}

type ValidationErrors []ValidationError

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUser(user domain.User) error {
	if err := validate.Var(user.Email, "required,email"); err != nil {
		return fmt.Errorf("email must be filled and valid")
	}

	if err := validate.Var(user.Username, "required"); err != nil {
		return fmt.Errorf("username must be filled")
	}

	if err := validate.Var(user.Password, "required,min=6"); err != nil {
		return fmt.Errorf("password must be filled and at least 6 characters long")
	}

	if err := validate.Var(user.Age, "required,min=8"); err != nil {
		return fmt.Errorf("age must be filled and at least 8 years old")
	}

	pattern := `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	match, _ := regexp.MatchString(pattern, user.ProfileImageUrl)
	if !match {
		err := ValidationErrors{
			{Field: "PhotoUrl", Error: "invalid url"},
		}
		fmt.Println(err)
	}
	return nil
}
func ValidateUpdateUser(user domain.User) error {
	if err := validate.Var(user.Email, "email"); err != nil {
		return fmt.Errorf("email must be filled and valid")
	}
	if err := validate.Var(user.Username, "required"); err != nil {
		return fmt.Errorf("username must be filled")
	}
	if err := validate.Var(user.Age, "min=8"); err != nil {
		return fmt.Errorf("age must be greater than 8")
	}
	pattern := `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	match, _ := regexp.MatchString(pattern, user.ProfileImageUrl)
	if !match {
		err := ValidationErrors{
			{Field: "PhotoUrl", Error: "invalid url"},
		}
		fmt.Println(err)
	}
	return nil
}

func ValidatePhoto(photo domain.Photo) error {
	if err := validate.Var(photo.Title, "required"); err != nil {
		return fmt.Errorf("title must be filled")
	}

	if err := validate.Var(photo.PhotoUrl, "required"); err != nil {
		return fmt.Errorf("photo url must be filled")
	}

	pattern := `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	match, _ := regexp.MatchString(pattern, photo.PhotoUrl)
	if !match {
		err := ValidationErrors{
			{Field: "PhotoUrl", Error: "invalid url"},
		}
		fmt.Println(err)
	}

	return nil
}

func ValidateComment(comment domain.Comment) error {
	if err := validate.Var(comment.Message, "required"); err != nil {
		return err
	}

	return nil
}

func ValidateSocialMedia(socialMedia domain.SocialMedia) error {
	if err := validate.Var(socialMedia.Name, "required"); err != nil {
		return err
	}

	if err := validate.Var(socialMedia.SocialMediaUrl, "required"); err != nil {
		return err
	}

	pattern := `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	match, _ := regexp.MatchString(pattern, socialMedia.SocialMediaUrl)
	if !match {
		err := ValidationErrors{
			{Field: "SocialMediaUrl", Error: "invalid url"},
		}
		fmt.Println(err)
	}

	return nil
}
