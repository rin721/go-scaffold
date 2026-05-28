package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rei0721/go-scaffold/internal/modules/user/service"
	"github.com/rei0721/go-scaffold/pkg/iam"
	"github.com/rei0721/go-scaffold/pkg/logger"
	apperrors "github.com/rei0721/go-scaffold/types/errors"
	"github.com/rei0721/go-scaffold/types/result"
)

const principalContextKey = "user.principal"

type UserHandler struct {
	service service.UserService
	logger  logger.Logger
}

type createUserRequest struct {
	Username    string   `json:"username" binding:"required"`
	Email       string   `json:"email" binding:"required"`
	Password    string   `json:"password" binding:"required"`
	DisplayName string   `json:"displayName"`
	Status      string   `json:"status"`
	Roles       []string `json:"roles"`
}

type updateUserRequest struct {
	Username    *string   `json:"username"`
	Email       *string   `json:"email"`
	Password    *string   `json:"password"`
	DisplayName *string   `json:"displayName"`
	Status      *string   `json:"status"`
	Roles       *[]string `json:"roles"`
}

type loginRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type createRoleRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type updateRoleRequest struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Permissions *[]string `json:"permissions"`
}

type createPermissionRequest struct {
	Code        string `json:"code" binding:"required"`
	Description string `json:"description"`
}

type updatePermissionRequest struct {
	Code        *string `json:"code"`
	Description *string `json:"description"`
}

func NewUserHandler(service service.UserService, logger logger.Logger) *UserHandler {
	return &UserHandler{service: service, logger: logger}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		result.BadRequest(c, err.Error())
		return
	}
	token, err := h.service.Register(c.Request.Context(), service.CreateUserInput{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		DisplayName: req.DisplayName,
		Status:      req.Status,
	})
	if err != nil {
		h.writeError(c, err)
		return
	}
	c.JSON(http.StatusCreated, result.Success(token))
}

func (h *UserHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		result.BadRequest(c, err.Error())
		return
	}
	token, err := h.service.Login(c.Request.Context(), service.LoginInput{
		Identifier: req.Identifier,
		Password:   req.Password,
	})
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, token)
}

func (h *UserHandler) Me(c *gin.Context) {
	principal, ok := principalFromGin(c)
	if !ok {
		result.Unauthorized(c, "authentication required")
		return
	}
	result.OK(c, principal)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		result.BadRequest(c, err.Error())
		return
	}
	user, err := h.service.CreateUser(c.Request.Context(), service.CreateUserInput{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		DisplayName: req.DisplayName,
		Status:      req.Status,
		Roles:       req.Roles,
	})
	if err != nil {
		h.writeError(c, err)
		return
	}
	c.JSON(http.StatusCreated, result.Success(user))
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers(c.Request.Context())
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	user, err := h.service.GetUser(c.Request.Context(), id)
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		result.BadRequest(c, err.Error())
		return
	}
	user, err := h.service.UpdateUser(c.Request.Context(), id, service.UpdateUserInput{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		DisplayName: req.DisplayName,
		Status:      req.Status,
		Roles:       req.Roles,
	})
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	if err := h.service.DeleteUser(c.Request.Context(), id); err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, gin.H{"deleted": true})
}

func (h *UserHandler) AssignRoleToUser(c *gin.Context) {
	userID, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	roleID, ok := parseIDParam(c, "roleID")
	if !ok {
		return
	}
	user, err := h.service.AssignRoleToUser(c.Request.Context(), userID, roleID)
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, user)
}

func (h *UserHandler) RemoveRoleFromUser(c *gin.Context) {
	userID, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	roleID, ok := parseIDParam(c, "roleID")
	if !ok {
		return
	}
	user, err := h.service.RemoveRoleFromUser(c.Request.Context(), userID, roleID)
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, user)
}

func (h *UserHandler) CreateRole(c *gin.Context) {
	var req createRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		result.BadRequest(c, err.Error())
		return
	}
	role, err := h.service.CreateRole(c.Request.Context(), service.CreateRoleInput{
		Name:        req.Name,
		Description: req.Description,
		Permissions: req.Permissions,
	})
	if err != nil {
		h.writeError(c, err)
		return
	}
	c.JSON(http.StatusCreated, result.Success(role))
}

func (h *UserHandler) ListRoles(c *gin.Context) {
	roles, err := h.service.ListRoles(c.Request.Context())
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, roles)
}

func (h *UserHandler) GetRole(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	role, err := h.service.GetRole(c.Request.Context(), id)
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, role)
}

func (h *UserHandler) UpdateRole(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	var req updateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		result.BadRequest(c, err.Error())
		return
	}
	role, err := h.service.UpdateRole(c.Request.Context(), id, service.UpdateRoleInput{
		Name:        req.Name,
		Description: req.Description,
		Permissions: req.Permissions,
	})
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, role)
}

func (h *UserHandler) DeleteRole(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	if err := h.service.DeleteRole(c.Request.Context(), id); err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, gin.H{"deleted": true})
}

func (h *UserHandler) AssignPermissionToRole(c *gin.Context) {
	roleID, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	permissionID, ok := parseIDParam(c, "permissionID")
	if !ok {
		return
	}
	role, err := h.service.AssignPermissionToRole(c.Request.Context(), roleID, permissionID)
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, role)
}

func (h *UserHandler) RemovePermissionFromRole(c *gin.Context) {
	roleID, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	permissionID, ok := parseIDParam(c, "permissionID")
	if !ok {
		return
	}
	role, err := h.service.RemovePermissionFromRole(c.Request.Context(), roleID, permissionID)
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, role)
}

func (h *UserHandler) CreatePermission(c *gin.Context) {
	var req createPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		result.BadRequest(c, err.Error())
		return
	}
	permission, err := h.service.CreatePermission(c.Request.Context(), service.CreatePermissionInput{
		Code:        req.Code,
		Description: req.Description,
	})
	if err != nil {
		h.writeError(c, err)
		return
	}
	c.JSON(http.StatusCreated, result.Success(permission))
}

func (h *UserHandler) ListPermissions(c *gin.Context) {
	permissions, err := h.service.ListPermissions(c.Request.Context())
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, permissions)
}

func (h *UserHandler) GetPermission(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	permission, err := h.service.GetPermission(c.Request.Context(), id)
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, permission)
}

func (h *UserHandler) UpdatePermission(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	var req updatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		result.BadRequest(c, err.Error())
		return
	}
	permission, err := h.service.UpdatePermission(c.Request.Context(), id, service.UpdatePermissionInput{
		Code:        req.Code,
		Description: req.Description,
	})
	if err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, permission)
}

func (h *UserHandler) DeletePermission(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	if err := h.service.DeletePermission(c.Request.Context(), id); err != nil {
		h.writeError(c, err)
		return
	}
	result.OK(c, gin.H{"deleted": true})
}

func (h *UserHandler) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := bearerToken(c.GetHeader("Authorization"))
		if !ok {
			result.Unauthorized(c, "authentication required")
			c.Abort()
			return
		}
		principal, err := h.service.AuthenticateToken(c.Request.Context(), token)
		if err != nil {
			h.writeError(c, err)
			c.Abort()
			return
		}
		c.Set(principalContextKey, principal)
		iamPrincipal := iam.Principal{
			ID:    strconv.FormatUint(uint64(principal.User.ID), 10),
			Name:  principal.User.Username,
			Roles: append([]string(nil), principal.Roles...),
			Attributes: map[string]string{
				"email":  principal.User.Email,
				"status": principal.User.Status,
			},
		}
		c.Request = c.Request.WithContext(iam.ContextWithPrincipal(c.Request.Context(), iamPrincipal))
		c.Next()
	}
}

func (h *UserHandler) RequirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		principal, ok := principalFromGin(c)
		if !ok {
			result.Unauthorized(c, "authentication required")
			c.Abort()
			return
		}
		if err := h.service.Authorize(c.Request.Context(), *principal, permission); err != nil {
			h.writeError(c, err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func bearerToken(header string) (string, bool) {
	header = strings.TrimSpace(header)
	if header == "" {
		return "", false
	}
	parts := strings.Fields(header)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
		return "", false
	}
	return parts[1], true
}

func principalFromGin(c *gin.Context) (*service.AuthPrincipal, bool) {
	value, ok := c.Get(principalContextKey)
	if !ok {
		return nil, false
	}
	principal, ok := value.(*service.AuthPrincipal)
	return principal, ok
}

func parseIDParam(c *gin.Context, name string) (uint, bool) {
	value, err := strconv.ParseUint(c.Param(name), 10, 64)
	if err != nil || value == 0 {
		result.BadRequest(c, "invalid id")
		return 0, false
	}
	return uint(value), true
}

func (h *UserHandler) writeError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrUsernameRequired), errors.Is(err, service.ErrInvalidUsername):
		writeUserError(c, http.StatusBadRequest, apperrors.ErrInvalidUsername, err.Error())
	case errors.Is(err, service.ErrEmailRequired), errors.Is(err, service.ErrInvalidEmail):
		writeUserError(c, http.StatusBadRequest, apperrors.ErrInvalidEmail, err.Error())
	case errors.Is(err, service.ErrPasswordRequired), errors.Is(err, service.ErrInvalidPassword):
		writeUserError(c, http.StatusBadRequest, apperrors.ErrInvalidPassword, err.Error())
	case errors.Is(err, service.ErrDisplayNameTooLong), errors.Is(err, service.ErrInvalidStatus),
		errors.Is(err, service.ErrRoleNameRequired), errors.Is(err, service.ErrInvalidRoleName),
		errors.Is(err, service.ErrPermissionRequired), errors.Is(err, service.ErrInvalidPermission):
		result.BadRequest(c, err.Error())
	case errors.Is(err, service.ErrDuplicateUsername):
		writeUserError(c, http.StatusUnprocessableEntity, apperrors.ErrDuplicateUsername, err.Error())
	case errors.Is(err, service.ErrDuplicateEmail):
		writeUserError(c, http.StatusUnprocessableEntity, apperrors.ErrDuplicateEmail, err.Error())
	case errors.Is(err, service.ErrDuplicateRole), errors.Is(err, service.ErrDuplicatePermission),
		errors.Is(err, service.ErrProtectedRole), errors.Is(err, service.ErrProtectedPermission):
		writeUserError(c, http.StatusUnprocessableEntity, apperrors.ErrBusinessLogic, err.Error())
	case errors.Is(err, service.ErrInvalidCredentials), errors.Is(err, service.ErrUserDisabled),
		errors.Is(err, service.ErrInvalidToken), errors.Is(err, service.ErrExpiredToken):
		result.Unauthorized(c, err.Error())
	case errors.Is(err, service.ErrPermissionDenied):
		result.Forbidden(c, err.Error())
	case errors.Is(err, service.ErrUserNotFound):
		writeUserError(c, http.StatusNotFound, apperrors.ErrUserNotFound, err.Error())
	case errors.Is(err, service.ErrRoleNotFound), errors.Is(err, service.ErrPermissionNotFound):
		result.NotFound(c, err.Error())
	default:
		if h.logger != nil {
			h.logger.Error("user request failed", "error", err)
		}
		result.InternalError(c, "internal server error")
	}
}

func writeUserError(c *gin.Context, status int, code int, message string) {
	c.JSON(status, result.ErrorWithTrace(code, message, result.GetTraceID(c)))
}
