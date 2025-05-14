package utils

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppError struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}

func (e AppError) Error() string {
	if msg, ok := e.Message.(string); ok {
		return fmt.Sprintf("App Error %d: %s", e.Code, msg)
	}

	return fmt.Sprintf("App Error %d", e.Code)
}

func NewAppError(code int, message any) AppError {
	return AppError{
		Code:    code,
		Message: message,
	}
}

func InvalidRequest(message string) AppError {
	return NewAppError(http.StatusBadRequest, message)
}

func NotFoundError(message string) AppError {
	return NewAppError(http.StatusNotFound, message)
}

func ValidationError(message string) AppError {
	return NewAppError(http.StatusUnprocessableEntity, message)
}

func ErrorCatcherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err != nil {
			if appErr, ok := err.(AppError); ok {
				return c.JSON(appErr.Code, appErr.Message)
			}

			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %s", err.Error()))
		}

		return nil
	}
}
