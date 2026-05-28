package service

import (
	"context"
	"errors"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rei0721/go-scaffold/internal/modules/user/model"
	"github.com/rei0721/go-scaffold/internal/modules/user/repository"
	authapi "github.com/rei0721/go-scaffold/pkg/auth"
	passwordcrypto "github.com/rei0721/go-scaffold/pkg/crypto"
	"github.com/rei0721/go-scaffold/pkg/database"
	rbacapi "github.com/rei0721/go-scaffold/pkg/rbac"
	"gorm.io/gorm"
)

var (
	ErrMissingDatabase     = errors.New("database is nil")
	ErrUsernameRequired    = errors.New("username is required")
	ErrInvalidUsername     = errors.New("username must be 3-64 characters and contain only letters, numbers, dot, underscore or dash")
	ErrEmailRequired       = errors.New("email is required")
	ErrInvalidEmail        = errors.New("email is invalid")
	ErrPasswordRequired    = errors.New("password is required")
	ErrInvalidPassword     = errors.New("password is invalid")
	ErrDisplayNameTooLong  = errors.New("display name is too long")
	ErrInvalidStatus       = errors.New("user status is invalid")
	ErrDuplicateUsername   = errors.New("username already exists")
	ErrDuplicateEmail      = errors.New("email already exists")
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrUserDisabled        = errors.New("user is disabled")
	ErrRoleNameRequired    = errors.New("role name is required")
	ErrInvalidRoleName     = errors.New("role name must be 2-64 characters and contain only letters, numbers, dot, underscore or dash")
	ErrDuplicateRole       = errors.New("role already exists")
	ErrRoleNotFound        = errors.New("role not found")
	ErrProtectedRole       = errors.New("built-in role cannot be deleted")
	ErrPermissionRequired  = errors.New("permission code is required")
	ErrInvalidPermission   = errors.New("permission code must use resource:action format")
	ErrDuplicatePermission = errors.New("permission already exists")
	ErrPermissionNotFound  = errors.New("permission not found")
	ErrProtectedPermission = errors.New("built-in permission cannot be deleted")
	ErrPermissionDenied    = errors.New("permission denied")
)

const (
	PermissionUsersCreate           = "users:create"
	PermissionUsersRead             = "users:read"
	PermissionUsersUpdate           = "users:update"
	PermissionUsersDelete           = "users:delete"
	PermissionUsersAssignRoles      = "users:assign_roles"
	PermissionRolesCreate           = "roles:create"
	PermissionRolesRead             = "roles:read"
	PermissionRolesUpdate           = "roles:update"
	PermissionRolesDelete           = "roles:delete"
	PermissionRolesAssignPermission = "roles:assign_permissions"
	PermissionPermissionsCreate     = "permissions:create"
	PermissionPermissionsRead       = "permissions:read"
	PermissionPermissionsUpdate     = "permissions:update"
	PermissionPermissionsDelete     = "permissions:delete"
	PermissionAll                   = "*:*"
)

var (
	usernamePattern   = regexp.MustCompile(`^[A-Za-z0-9_.-]{3,64}$`)
	roleNamePattern   = regexp.MustCompile(`^[A-Za-z0-9_.-]{2,64}$`)
	permissionPattern = regexp.MustCompile(`^[A-Za-z0-9_.*-]+:[A-Za-z0-9_.*-]+$`)
)

type UserDTO struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DisplayName string    `json:"displayName,omitempty"`
	Status      string    `json:"status"`
	Roles       []RoleDTO `json:"roles,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type RoleDTO struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Permissions []PermissionDTO `json:"permissions,omitempty"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

type PermissionDTO struct {
	ID          uint      `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type AuthToken struct {
	Token     string    `json:"token"`
	TokenType string    `json:"tokenType"`
	ExpiresAt time.Time `json:"expiresAt"`
	User      UserDTO   `json:"user"`
}

type AuthPrincipal struct {
	User        UserDTO  `json:"user"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type CreateUserInput struct {
	Username    string
	Email       string
	Password    string
	DisplayName string
	Status      string
	Roles       []string
}

type UpdateUserInput struct {
	Username    *string
	Email       *string
	Password    *string
	DisplayName *string
	Status      *string
	Roles       *[]string
}

type LoginInput struct {
	Identifier string
	Password   string
}

type CreateRoleInput struct {
	Name        string
	Description string
	Permissions []string
}

type UpdateRoleInput struct {
	Name        *string
	Description *string
	Permissions *[]string
}

type CreatePermissionInput struct {
	Code        string
	Description string
}

type UpdatePermissionInput struct {
	Code        *string
	Description *string
}

type RBACSeed struct {
	Roles           []RBACRoleSeed
	Permissions     []RBACPermissionSeed
	RolePermissions []RBACRolePermissionSeed
}

type RBACRoleSeed struct {
	Name        string
	Description string
}

type RBACPermissionSeed struct {
	Code        string
	Description string
}

type RBACRolePermissionSeed struct {
	Role        string
	Permissions []string
}

type UserService interface {
	Register(ctx context.Context, input CreateUserInput) (*AuthToken, error)
	Login(ctx context.Context, input LoginInput) (*AuthToken, error)
	AuthenticateToken(ctx context.Context, token string) (*AuthPrincipal, error)
	Authorize(ctx context.Context, principal AuthPrincipal, permission string) error
	ApplyRBACSeed(ctx context.Context, seed RBACSeed) error

	CreateUser(ctx context.Context, input CreateUserInput) (*UserDTO, error)
	ListUsers(ctx context.Context) ([]UserDTO, error)
	GetUser(ctx context.Context, id uint) (*UserDTO, error)
	UpdateUser(ctx context.Context, id uint, input UpdateUserInput) (*UserDTO, error)
	DeleteUser(ctx context.Context, id uint) error
	AssignRoleToUser(ctx context.Context, userID uint, roleID uint) (*UserDTO, error)
	RemoveRoleFromUser(ctx context.Context, userID uint, roleID uint) (*UserDTO, error)

	CreateRole(ctx context.Context, input CreateRoleInput) (*RoleDTO, error)
	ListRoles(ctx context.Context) ([]RoleDTO, error)
	GetRole(ctx context.Context, id uint) (*RoleDTO, error)
	UpdateRole(ctx context.Context, id uint, input UpdateRoleInput) (*RoleDTO, error)
	DeleteRole(ctx context.Context, id uint) error
	AssignPermissionToRole(ctx context.Context, roleID uint, permissionID uint) (*RoleDTO, error)
	RemovePermissionFromRole(ctx context.Context, roleID uint, permissionID uint) (*RoleDTO, error)

	CreatePermission(ctx context.Context, input CreatePermissionInput) (*PermissionDTO, error)
	ListPermissions(ctx context.Context) ([]PermissionDTO, error)
	GetPermission(ctx context.Context, id uint) (*PermissionDTO, error)
	UpdatePermission(ctx context.Context, id uint, input UpdatePermissionInput) (*PermissionDTO, error)
	DeletePermission(ctx context.Context, id uint) error
}

type userService struct {
	db     database.Database
	repo   repository.Repository
	crypto passwordcrypto.Crypto
	tokens authapi.TokenService
	rbac   rbacapi.Authorizer
}

func NewUserService(db database.Database, repo repository.Repository, crypto passwordcrypto.Crypto, tokens authapi.TokenService, authorizer rbacapi.Authorizer) UserService {
	if authorizer == nil {
		authorizer = rbacapi.NewDefaultAuthorizer()
	}
	return &userService{db: db, repo: repo, crypto: crypto, tokens: tokens, rbac: authorizer}
}

func (s *userService) Register(ctx context.Context, input CreateUserInput) (*AuthToken, error) {
	user, err := s.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}
	return s.issueTokenForUser(ctx, user.ID)
}

func (s *userService) Login(ctx context.Context, input LoginInput) (*AuthToken, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	identifier := strings.TrimSpace(input.Identifier)
	if identifier == "" || strings.TrimSpace(input.Password) == "" {
		return nil, ErrInvalidCredentials
	}

	user, err := s.findUserByIdentifier(ctx, s.db.DB(), identifier)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}
	if user.Status != model.UserStatusActive {
		return nil, ErrUserDisabled
	}
	if err := s.crypto.VerifyPassword(user.PasswordHash, input.Password); err != nil {
		return nil, ErrInvalidCredentials
	}
	return s.issueTokenForUser(ctx, user.ID)
}

func (s *userService) AuthenticateToken(ctx context.Context, token string) (*AuthPrincipal, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	claims, err := s.tokens.Verify(ctx, strings.TrimSpace(token))
	if err != nil {
		return nil, normalizeTokenError(err)
	}
	userID, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil || userID == 0 {
		return nil, ErrInvalidToken
	}
	user, err := s.repo.FindUserByID(ctx, s.db.DB(), uint(userID))
	if err != nil {
		return nil, normalizeUserNotFound(err)
	}
	if user.Status != model.UserStatusActive {
		return nil, ErrUserDisabled
	}
	dto, err := s.userDTO(ctx, s.db.DB(), *user)
	if err != nil {
		return nil, err
	}
	permissions, err := s.repo.ListPermissionsForUser(ctx, s.db.DB(), user.ID)
	if err != nil {
		return nil, err
	}
	principal := &AuthPrincipal{User: *dto}
	for _, role := range dto.Roles {
		principal.Roles = append(principal.Roles, role.Name)
	}
	for _, permission := range permissions {
		principal.Permissions = append(principal.Permissions, permission.Code)
	}
	return principal, nil
}

func (s *userService) Authorize(ctx context.Context, principal AuthPrincipal, permission string) error {
	if err := s.ensureReady(); err != nil {
		return err
	}
	policies, err := s.repo.ListRolePermissionPolicies(ctx, s.db.DB())
	if err != nil {
		return err
	}
	allowed, err := s.rbac.Authorize(ctx, rbacapi.Principal{
		ID:    strconv.FormatUint(uint64(principal.User.ID), 10),
		Roles: principal.Roles,
	}, strings.TrimSpace(permission), toRBACPolicies(policies))
	if err != nil {
		return err
	}
	if allowed {
		return nil
	}
	return ErrPermissionDenied
}

func (s *userService) ApplyRBACSeed(ctx context.Context, seed RBACSeed) error {
	if err := s.ensureReady(); err != nil {
		return err
	}
	return s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		if err := s.ensureDefaultRBAC(ctx, tx); err != nil {
			return err
		}
		if err := s.applyRoleSeeds(ctx, tx, seed.Roles); err != nil {
			return err
		}
		if err := s.applyPermissionSeeds(ctx, tx, seed.Permissions); err != nil {
			return err
		}
		return s.applyRolePermissionSeeds(ctx, tx, seed.RolePermissions)
	})
}

func (s *userService) CreateUser(ctx context.Context, input CreateUserInput) (*UserDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	var created model.User
	err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		userCount, err := s.repo.CountUsers(ctx, tx)
		if err != nil {
			return err
		}
		if err := s.ensureDefaultRBAC(ctx, tx); err != nil {
			return err
		}
		user, err := s.createUserInTx(ctx, tx, input)
		if err != nil {
			return err
		}
		roleNames := normalizedNames(input.Roles)
		if len(roleNames) == 0 {
			roleNames = []string{model.RoleUser}
			if userCount == 0 {
				roleNames = []string{model.RoleAdmin}
			}
		}
		if err := s.replaceUserRolesByName(ctx, tx, user.ID, roleNames); err != nil {
			return err
		}
		created = *user
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.userDTO(ctx, s.db.DB(), created)
}

func (s *userService) ListUsers(ctx context.Context) ([]UserDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	users, err := s.repo.ListUsers(ctx, s.db.DB())
	if err != nil {
		return nil, err
	}
	return s.userDTOs(ctx, s.db.DB(), users)
}

func (s *userService) GetUser(ctx context.Context, id uint) (*UserDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	user, err := s.repo.FindUserByID(ctx, s.db.DB(), id)
	if err != nil {
		return nil, normalizeUserNotFound(err)
	}
	return s.userDTO(ctx, s.db.DB(), *user)
}

func (s *userService) UpdateUser(ctx context.Context, id uint, input UpdateUserInput) (*UserDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	var updated model.User
	err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		user, err := s.repo.FindUserByID(ctx, tx, id)
		if err != nil {
			return normalizeUserNotFound(err)
		}
		if input.Username != nil {
			username, err := normalizeUsername(*input.Username)
			if err != nil {
				return err
			}
			if username != user.Username {
				if _, err := s.repo.FindUserByUsername(ctx, tx, username); err == nil {
					return ErrDuplicateUsername
				} else if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}
			}
			user.Username = username
		}
		if input.Email != nil {
			email, err := normalizeEmail(*input.Email)
			if err != nil {
				return err
			}
			if email != user.Email {
				if _, err := s.repo.FindUserByEmail(ctx, tx, email); err == nil {
					return ErrDuplicateEmail
				} else if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}
			}
			user.Email = email
		}
		if input.Password != nil {
			password := strings.TrimSpace(*input.Password)
			if password == "" {
				return ErrPasswordRequired
			}
			hash, err := s.crypto.HashPassword(password)
			if err != nil {
				return ErrInvalidPassword
			}
			user.PasswordHash = hash
		}
		if input.DisplayName != nil {
			displayName := strings.TrimSpace(*input.DisplayName)
			if len(displayName) > 100 {
				return ErrDisplayNameTooLong
			}
			user.DisplayName = displayName
		}
		if input.Status != nil {
			status, err := normalizeStatus(*input.Status)
			if err != nil {
				return err
			}
			user.Status = status
		}
		if err := s.repo.UpdateUser(ctx, tx, user); err != nil {
			return err
		}
		if input.Roles != nil {
			if err := s.replaceUserRolesByName(ctx, tx, user.ID, normalizedNames(*input.Roles)); err != nil {
				return err
			}
		}
		updated = *user
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.userDTO(ctx, s.db.DB(), updated)
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	if err := s.ensureReady(); err != nil {
		return err
	}
	return s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		if _, err := s.repo.FindUserByID(ctx, tx, id); err != nil {
			return normalizeUserNotFound(err)
		}
		return s.repo.DeleteUser(ctx, tx, id)
	})
}

func (s *userService) AssignRoleToUser(ctx context.Context, userID uint, roleID uint) (*UserDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	var user model.User
	err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		current, err := s.repo.FindUserByID(ctx, tx, userID)
		if err != nil {
			return normalizeUserNotFound(err)
		}
		if _, err := s.repo.FindRoleByID(ctx, tx, roleID); err != nil {
			return normalizeRoleNotFound(err)
		}
		if err := s.repo.AssignRoleToUser(ctx, tx, userID, roleID); err != nil {
			return err
		}
		user = *current
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.userDTO(ctx, s.db.DB(), user)
}

func (s *userService) RemoveRoleFromUser(ctx context.Context, userID uint, roleID uint) (*UserDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	var user model.User
	err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		current, err := s.repo.FindUserByID(ctx, tx, userID)
		if err != nil {
			return normalizeUserNotFound(err)
		}
		if _, err := s.repo.FindRoleByID(ctx, tx, roleID); err != nil {
			return normalizeRoleNotFound(err)
		}
		if err := s.repo.RemoveRoleFromUser(ctx, tx, userID, roleID); err != nil {
			return err
		}
		user = *current
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.userDTO(ctx, s.db.DB(), user)
}

func (s *userService) CreateRole(ctx context.Context, input CreateRoleInput) (*RoleDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	var role model.Role
	err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		if err := s.ensureDefaultRBAC(ctx, tx); err != nil {
			return err
		}
		name, err := normalizeRoleName(input.Name)
		if err != nil {
			return err
		}
		if _, err := s.repo.FindRoleByName(ctx, tx, name); err == nil {
			return ErrDuplicateRole
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		role = model.Role{Name: name, Description: strings.TrimSpace(input.Description)}
		if err := s.repo.CreateRole(ctx, tx, &role); err != nil {
			return err
		}
		return s.replaceRolePermissionsByCode(ctx, tx, role.ID, normalizedNames(input.Permissions))
	})
	if err != nil {
		return nil, err
	}
	return s.roleDTO(ctx, s.db.DB(), role)
}

func (s *userService) ListRoles(ctx context.Context) ([]RoleDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	roles, err := s.repo.ListRoles(ctx, s.db.DB())
	if err != nil {
		return nil, err
	}
	dtos := make([]RoleDTO, 0, len(roles))
	for _, role := range roles {
		dto, err := s.roleDTO(ctx, s.db.DB(), role)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, *dto)
	}
	return dtos, nil
}

func (s *userService) GetRole(ctx context.Context, id uint) (*RoleDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	role, err := s.repo.FindRoleByID(ctx, s.db.DB(), id)
	if err != nil {
		return nil, normalizeRoleNotFound(err)
	}
	return s.roleDTO(ctx, s.db.DB(), *role)
}

func (s *userService) UpdateRole(ctx context.Context, id uint, input UpdateRoleInput) (*RoleDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	var updated model.Role
	err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		role, err := s.repo.FindRoleByID(ctx, tx, id)
		if err != nil {
			return normalizeRoleNotFound(err)
		}
		if input.Name != nil {
			name, err := normalizeRoleName(*input.Name)
			if err != nil {
				return err
			}
			if name != role.Name {
				if _, err := s.repo.FindRoleByName(ctx, tx, name); err == nil {
					return ErrDuplicateRole
				} else if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}
			}
			role.Name = name
		}
		if input.Description != nil {
			role.Description = strings.TrimSpace(*input.Description)
		}
		if err := s.repo.UpdateRole(ctx, tx, role); err != nil {
			return err
		}
		if input.Permissions != nil {
			if err := s.replaceRolePermissionsByCode(ctx, tx, role.ID, normalizedNames(*input.Permissions)); err != nil {
				return err
			}
		}
		updated = *role
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.roleDTO(ctx, s.db.DB(), updated)
}

func (s *userService) DeleteRole(ctx context.Context, id uint) error {
	if err := s.ensureReady(); err != nil {
		return err
	}
	return s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		role, err := s.repo.FindRoleByID(ctx, tx, id)
		if err != nil {
			return normalizeRoleNotFound(err)
		}
		if role.Name == model.RoleAdmin || role.Name == model.RoleUser {
			return ErrProtectedRole
		}
		return s.repo.DeleteRole(ctx, tx, id)
	})
}

func (s *userService) AssignPermissionToRole(ctx context.Context, roleID uint, permissionID uint) (*RoleDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	var role model.Role
	err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		current, err := s.repo.FindRoleByID(ctx, tx, roleID)
		if err != nil {
			return normalizeRoleNotFound(err)
		}
		if _, err := s.repo.FindPermissionByID(ctx, tx, permissionID); err != nil {
			return normalizePermissionNotFound(err)
		}
		if err := s.repo.AssignPermissionToRole(ctx, tx, roleID, permissionID); err != nil {
			return err
		}
		role = *current
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.roleDTO(ctx, s.db.DB(), role)
}

func (s *userService) RemovePermissionFromRole(ctx context.Context, roleID uint, permissionID uint) (*RoleDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	var role model.Role
	err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		current, err := s.repo.FindRoleByID(ctx, tx, roleID)
		if err != nil {
			return normalizeRoleNotFound(err)
		}
		if _, err := s.repo.FindPermissionByID(ctx, tx, permissionID); err != nil {
			return normalizePermissionNotFound(err)
		}
		if err := s.repo.RemovePermissionFromRole(ctx, tx, roleID, permissionID); err != nil {
			return err
		}
		role = *current
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.roleDTO(ctx, s.db.DB(), role)
}

func (s *userService) CreatePermission(ctx context.Context, input CreatePermissionInput) (*PermissionDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	code, err := normalizePermissionCode(input.Code)
	if err != nil {
		return nil, err
	}
	var permission model.Permission
	err = s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		if _, err := s.repo.FindPermissionByCode(ctx, tx, code); err == nil {
			return ErrDuplicatePermission
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		permission = model.Permission{Code: code, Description: strings.TrimSpace(input.Description)}
		return s.repo.CreatePermission(ctx, tx, &permission)
	})
	if err != nil {
		return nil, err
	}
	return permissionDTO(permission), nil
}

func (s *userService) ListPermissions(ctx context.Context) ([]PermissionDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	permissions, err := s.repo.ListPermissions(ctx, s.db.DB())
	if err != nil {
		return nil, err
	}
	return permissionDTOs(permissions), nil
}

func (s *userService) GetPermission(ctx context.Context, id uint) (*PermissionDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	permission, err := s.repo.FindPermissionByID(ctx, s.db.DB(), id)
	if err != nil {
		return nil, normalizePermissionNotFound(err)
	}
	return permissionDTO(*permission), nil
}

func (s *userService) UpdatePermission(ctx context.Context, id uint, input UpdatePermissionInput) (*PermissionDTO, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}
	var updated model.Permission
	err := s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		permission, err := s.repo.FindPermissionByID(ctx, tx, id)
		if err != nil {
			return normalizePermissionNotFound(err)
		}
		if input.Code != nil {
			code, err := normalizePermissionCode(*input.Code)
			if err != nil {
				return err
			}
			if code != permission.Code {
				if _, err := s.repo.FindPermissionByCode(ctx, tx, code); err == nil {
					return ErrDuplicatePermission
				} else if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}
			}
			permission.Code = code
		}
		if input.Description != nil {
			permission.Description = strings.TrimSpace(*input.Description)
		}
		if err := s.repo.UpdatePermission(ctx, tx, permission); err != nil {
			return err
		}
		updated = *permission
		return nil
	})
	if err != nil {
		return nil, err
	}
	return permissionDTO(updated), nil
}

func (s *userService) DeletePermission(ctx context.Context, id uint) error {
	if err := s.ensureReady(); err != nil {
		return err
	}
	return s.db.WithTx(ctx, func(ctx context.Context, tx *gorm.DB) error {
		permission, err := s.repo.FindPermissionByID(ctx, tx, id)
		if err != nil {
			return normalizePermissionNotFound(err)
		}
		if permission.Code == PermissionAll {
			return ErrProtectedPermission
		}
		return s.repo.DeletePermission(ctx, tx, id)
	})
}

func (s *userService) ensureReady() error {
	if s.db == nil {
		return ErrMissingDatabase
	}
	if s.repo == nil || s.crypto == nil || s.tokens == nil || s.rbac == nil {
		return errors.New("user service dependency is nil")
	}
	return nil
}

func (s *userService) createUserInTx(ctx context.Context, tx *gorm.DB, input CreateUserInput) (*model.User, error) {
	username, err := normalizeUsername(input.Username)
	if err != nil {
		return nil, err
	}
	email, err := normalizeEmail(input.Email)
	if err != nil {
		return nil, err
	}
	password := strings.TrimSpace(input.Password)
	if password == "" {
		return nil, ErrPasswordRequired
	}
	displayName := strings.TrimSpace(input.DisplayName)
	if len(displayName) > 100 {
		return nil, ErrDisplayNameTooLong
	}
	status, err := normalizeStatus(input.Status)
	if err != nil {
		return nil, err
	}
	if _, err := s.repo.FindUserByUsername(ctx, tx, username); err == nil {
		return nil, ErrDuplicateUsername
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if _, err := s.repo.FindUserByEmail(ctx, tx, email); err == nil {
		return nil, ErrDuplicateEmail
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	hash, err := s.crypto.HashPassword(password)
	if err != nil {
		return nil, ErrInvalidPassword
	}
	user := &model.User{
		Username:     username,
		Email:        email,
		PasswordHash: hash,
		DisplayName:  displayName,
		Status:       status,
	}
	if err := s.repo.CreateUser(ctx, tx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) issueTokenForUser(ctx context.Context, id uint) (*AuthToken, error) {
	user, err := s.repo.FindUserByID(ctx, s.db.DB(), id)
	if err != nil {
		return nil, normalizeUserNotFound(err)
	}
	dto, err := s.userDTO(ctx, s.db.DB(), *user)
	if err != nil {
		return nil, err
	}
	token, err := s.tokens.Issue(ctx, authapi.Claims{
		Subject:  strconv.FormatUint(uint64(user.ID), 10),
		Username: user.Username,
	})
	if err != nil {
		return nil, err
	}
	return &AuthToken{Token: token.Value, TokenType: token.Type, ExpiresAt: token.ExpiresAt, User: *dto}, nil
}

func (s *userService) findUserByIdentifier(ctx context.Context, db *gorm.DB, identifier string) (*model.User, error) {
	var user *model.User
	var err error
	if strings.Contains(identifier, "@") {
		user, err = s.repo.FindUserByEmail(ctx, db, strings.ToLower(identifier))
	} else {
		user, err = s.repo.FindUserByUsername(ctx, db, identifier)
	}
	if err != nil {
		return nil, normalizeUserNotFound(err)
	}
	return user, nil
}

func (s *userService) ensureDefaultRBAC(ctx context.Context, tx *gorm.DB) error {
	adminRole, err := s.findOrCreateRole(ctx, tx, model.RoleAdmin, "Built-in administrator role")
	if err != nil {
		return err
	}
	if _, err := s.findOrCreateRole(ctx, tx, model.RoleUser, "Built-in default user role"); err != nil {
		return err
	}
	adminPermission, err := s.findOrCreatePermission(ctx, tx, PermissionAll, "All permissions")
	if err != nil {
		return err
	}
	return s.repo.AssignPermissionToRole(ctx, tx, adminRole.ID, adminPermission.ID)
}

func (s *userService) findOrCreateRole(ctx context.Context, tx *gorm.DB, name string, description string) (*model.Role, error) {
	role, err := s.repo.FindRoleByName(ctx, tx, name)
	if err == nil {
		return role, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	role = &model.Role{Name: name, Description: description}
	if err := s.repo.CreateRole(ctx, tx, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *userService) findOrCreatePermission(ctx context.Context, tx *gorm.DB, code string, description string) (*model.Permission, error) {
	permission, err := s.repo.FindPermissionByCode(ctx, tx, code)
	if err == nil {
		return permission, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	permission = &model.Permission{Code: code, Description: description}
	if err := s.repo.CreatePermission(ctx, tx, permission); err != nil {
		return nil, err
	}
	return permission, nil
}

func (s *userService) applyRoleSeeds(ctx context.Context, tx *gorm.DB, roles []RBACRoleSeed) error {
	seen := map[string]struct{}{}
	for _, item := range roles {
		name, err := normalizeRoleName(item.Name)
		if err != nil {
			return err
		}
		if _, ok := seen[name]; ok {
			continue
		}
		seen[name] = struct{}{}
		if _, err := s.findOrCreateRole(ctx, tx, name, strings.TrimSpace(item.Description)); err != nil {
			return err
		}
	}
	return nil
}

func (s *userService) applyPermissionSeeds(ctx context.Context, tx *gorm.DB, permissions []RBACPermissionSeed) error {
	seen := map[string]struct{}{}
	for _, item := range permissions {
		code, err := normalizePermissionCode(item.Code)
		if err != nil {
			return err
		}
		if _, ok := seen[code]; ok {
			continue
		}
		seen[code] = struct{}{}
		if _, err := s.findOrCreatePermission(ctx, tx, code, strings.TrimSpace(item.Description)); err != nil {
			return err
		}
	}
	return nil
}

func (s *userService) applyRolePermissionSeeds(ctx context.Context, tx *gorm.DB, grants []RBACRolePermissionSeed) error {
	for _, grant := range grants {
		roleName, err := normalizeRoleName(grant.Role)
		if err != nil {
			return err
		}
		role, err := s.repo.FindRoleByName(ctx, tx, roleName)
		if err != nil {
			return normalizeRoleNotFound(err)
		}
		seen := map[string]struct{}{}
		for _, value := range grant.Permissions {
			code, err := normalizePermissionCode(value)
			if err != nil {
				return err
			}
			if _, ok := seen[code]; ok {
				continue
			}
			seen[code] = struct{}{}
			permission, err := s.repo.FindPermissionByCode(ctx, tx, code)
			if err != nil {
				return normalizePermissionNotFound(err)
			}
			if err := s.repo.AssignPermissionToRole(ctx, tx, role.ID, permission.ID); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *userService) replaceUserRolesByName(ctx context.Context, tx *gorm.DB, userID uint, roleNames []string) error {
	roleIDs := make([]uint, 0, len(roleNames))
	for _, roleName := range roleNames {
		role, err := s.repo.FindRoleByName(ctx, tx, roleName)
		if err != nil {
			return normalizeRoleNotFound(err)
		}
		roleIDs = append(roleIDs, role.ID)
	}
	return s.repo.ReplaceUserRoles(ctx, tx, userID, roleIDs)
}

func (s *userService) replaceRolePermissionsByCode(ctx context.Context, tx *gorm.DB, roleID uint, codes []string) error {
	permissionIDs := make([]uint, 0, len(codes))
	for _, code := range codes {
		permission, err := s.repo.FindPermissionByCode(ctx, tx, code)
		if err != nil {
			return normalizePermissionNotFound(err)
		}
		permissionIDs = append(permissionIDs, permission.ID)
	}
	return s.repo.ReplaceRolePermissions(ctx, tx, roleID, permissionIDs)
}

func (s *userService) userDTO(ctx context.Context, db *gorm.DB, user model.User) (*UserDTO, error) {
	roles, err := s.repo.ListRolesForUser(ctx, db, user.ID)
	if err != nil {
		return nil, err
	}
	dto := UserDTO{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		Status:      user.Status,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	for _, role := range roles {
		dto.Roles = append(dto.Roles, RoleDTO{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
		})
	}
	return &dto, nil
}

func (s *userService) userDTOs(ctx context.Context, db *gorm.DB, users []model.User) ([]UserDTO, error) {
	dtos := make([]UserDTO, 0, len(users))
	for _, user := range users {
		dto, err := s.userDTO(ctx, db, user)
		if err != nil {
			return nil, err
		}
		dtos = append(dtos, *dto)
	}
	return dtos, nil
}

func (s *userService) roleDTO(ctx context.Context, db *gorm.DB, role model.Role) (*RoleDTO, error) {
	permissions, err := s.repo.ListPermissionsForRole(ctx, db, role.ID)
	if err != nil {
		return nil, err
	}
	dto := RoleDTO{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
		Permissions: permissionDTOs(permissions),
	}
	return &dto, nil
}

func permissionDTO(permission model.Permission) *PermissionDTO {
	return &PermissionDTO{
		ID:          permission.ID,
		Code:        permission.Code,
		Description: permission.Description,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}
}

func permissionDTOs(permissions []model.Permission) []PermissionDTO {
	dtos := make([]PermissionDTO, 0, len(permissions))
	for _, permission := range permissions {
		dtos = append(dtos, *permissionDTO(permission))
	}
	return dtos
}

func normalizeUsername(value string) (string, error) {
	username := strings.TrimSpace(value)
	if username == "" {
		return "", ErrUsernameRequired
	}
	if !usernamePattern.MatchString(username) {
		return "", ErrInvalidUsername
	}
	return username, nil
}

func normalizeEmail(value string) (string, error) {
	email := strings.ToLower(strings.TrimSpace(value))
	if email == "" {
		return "", ErrEmailRequired
	}
	parsed, err := mail.ParseAddress(email)
	if err != nil || parsed.Address != email || strings.Contains(email, " ") {
		return "", ErrInvalidEmail
	}
	return email, nil
}

func normalizeStatus(value string) (string, error) {
	status := strings.ToLower(strings.TrimSpace(value))
	if status == "" {
		return model.UserStatusActive, nil
	}
	if status != model.UserStatusActive && status != model.UserStatusDisabled {
		return "", ErrInvalidStatus
	}
	return status, nil
}

func normalizeRoleName(value string) (string, error) {
	name := strings.ToLower(strings.TrimSpace(value))
	if name == "" {
		return "", ErrRoleNameRequired
	}
	if !roleNamePattern.MatchString(name) {
		return "", ErrInvalidRoleName
	}
	return name, nil
}

func normalizePermissionCode(value string) (string, error) {
	code := strings.ToLower(strings.TrimSpace(value))
	if code == "" {
		return "", ErrPermissionRequired
	}
	if !permissionPattern.MatchString(code) {
		return "", ErrInvalidPermission
	}
	return code, nil
}

func normalizedNames(values []string) []string {
	seen := map[string]struct{}{}
	var out []string
	for _, value := range values {
		name := strings.ToLower(strings.TrimSpace(value))
		if name == "" {
			continue
		}
		if _, ok := seen[name]; ok {
			continue
		}
		seen[name] = struct{}{}
		out = append(out, name)
	}
	return out
}

func normalizeUserNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrUserNotFound
	}
	return err
}

func normalizeTokenError(err error) error {
	if errors.Is(err, authapi.ErrExpiredToken) {
		return ErrExpiredToken
	}
	if errors.Is(err, authapi.ErrInvalidToken) {
		return ErrInvalidToken
	}
	return err
}

func toRBACPolicies(policies []repository.RolePermissionPolicy) []rbacapi.Policy {
	out := make([]rbacapi.Policy, 0, len(policies))
	for _, policy := range policies {
		out = append(out, rbacapi.Policy{
			Role:       policy.RoleName,
			Permission: policy.PermissionCode,
		})
	}
	return out
}

func normalizeRoleNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRoleNotFound
	}
	return err
}

func normalizePermissionNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrPermissionNotFound
	}
	return err
}
