package v1

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/DoWithLogic/golang-clean-architecture/internal/users/dtos"
	"github.com/DoWithLogic/golang-clean-architecture/internal/users/entities"
	usecases "github.com/DoWithLogic/golang-clean-architecture/internal/users/usecase"
	"github.com/DoWithLogic/golang-clean-architecture/pkg/middleware"
	"github.com/DoWithLogic/golang-clean-architecture/pkg/otel/zerolog"
	"github.com/DoWithLogic/golang-clean-architecture/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type (
	Handlers interface {
		Login(c echo.Context) error
		CreateUser(c echo.Context) error
		UserDetail(c echo.Context) error
		UpdateUser(c echo.Context) error
		UpdateUserStatus(c echo.Context) error
	}

	handlers struct {
		uc  usecases.Usecase
		log *zerolog.Logger
	}
)

const (
	BooleanTextTrue  = "true"
	BooleanTextFalse = "false"
)

var (
	ErrInvalidIsActive = errors.New("invalid is_active")
)

func NewHandlers(uc usecases.Usecase, log *zerolog.Logger) Handlers {
	return &handlers{uc, log}
}

func (h *handlers) Login(c echo.Context) error {
	var (
		request dtos.UserLoginRequest
	)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	authData, httpCode, err := h.uc.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(httpCode, response.NewResponseError(httpCode, response.MsgFailed, err.Error()))
	}

	return c.JSON(httpCode, response.NewResponse(httpCode, response.MsgSuccess, authData))
}

func (h *handlers) CreateUser(c echo.Context) error {
	var (
		ctx, cancel = context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
		payload     dtos.CreateUserRequest
	)
	defer cancel()

	if err := c.Bind(&payload); err != nil {
		h.log.Z().Err(err).Msg("[handlers]CreateUser.Bind")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	if err := payload.Validate(); err != nil {
		h.log.Z().Err(err).Msg("[handlers]CreateUser.Validate")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	userID, httpCode, err := h.uc.Create(ctx, payload)
	if err != nil {
		return c.JSON(httpCode, response.NewResponseError(
			httpCode,
			response.MsgFailed,
			err.Error()),
		)
	}

	return c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, response.MsgSuccess, map[string]int64{"id": userID}))
}

func (h *handlers) UserDetail(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
	defer cancel()

	userID := c.Get("identity").(*middleware.CustomClaims).UserID

	data, code, err := h.uc.Detail(ctx, userID)
	if err != nil {
		return c.JSON(code, response.NewResponseError(code, response.MsgFailed, err.Error()))
	}

	return c.JSON(code, response.NewResponse(code, response.MsgSuccess, data))
}

func (h *handlers) UpdateUser(c echo.Context) error {
	var (
		ctx, cancel = context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
		request     dtos.UpdateUserRequest
	)
	defer cancel()

	request.UserID = c.Get("identity").(entities.Identity).UserID

	if err := c.Bind(&request); err != nil {
		h.log.Z().Err(err).Msg("[handlers]UpdateUser.Bind")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	if err := request.Validate(); err != nil {
		h.log.Z().Err(err).Msg("[handlers]UpdateUser.Validate")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	if err := h.uc.PartialUpdate(ctx, request); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusInternalServerError, response.MsgFailed, err.Error()))
	}

	return c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, response.MsgSuccess, nil))
}

func (h *handlers) UpdateUserStatus(c echo.Context) error {
	var (
		ctx, cancel = context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
		payload     dtos.UpdateUserStatusPayload
	)
	defer cancel()

	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.log.Z().Err(err).Msg("[handlers]UpdateUser.ParseParam")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	switch c.QueryParam("is_active") {
	case BooleanTextFalse:
		payload.IsActive = false
	case BooleanTextTrue:
		payload.IsActive = true
	default:
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, ErrInvalidIsActive.Error()))
	}

	argsUpdateUserStatus := entities.UpdateUserStatus{
		UserID:   userID,
		IsActive: payload.IsActive,
	}

	err = h.uc.UpdateStatus(ctx, argsUpdateUserStatus)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusInternalServerError, response.MsgFailed, err.Error()))
	}

	return c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, response.MsgSuccess, nil))
}
