package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// validateBody ...
func validateBody(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindJSON(data); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return err
	}

	return nil
}
