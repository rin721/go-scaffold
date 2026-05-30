package httptransport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rei0721/go-scaffold/pkg/database"
	apperrors "github.com/rei0721/go-scaffold/types/errors"
	"gorm.io/gorm"
)

type fakeDatabase struct {
	pingErr error
}

func (db *fakeDatabase) DB() *gorm.DB {
	return nil
}

func (db *fakeDatabase) Close() error {
	return nil
}

func (db *fakeDatabase) Ping(context.Context) error {
	return db.pingErr
}

func (db *fakeDatabase) Reload(*database.Config) error {
	return nil
}

func (db *fakeDatabase) WithTx(context.Context, database.TxFunc) error {
	return nil
}

func (db *fakeDatabase) WithTxOptions(context.Context, *database.TxOptions, database.TxFunc) error {
	return nil
}

type routerResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

func TestNewRouterHealthEndpoint(t *testing.T) {
	router := newTestRouter(RouterDeps{})

	recorder, body := performRouterRequest(t, router, http.MethodGet, "/health")

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected /health status %d, got %d with body %s", http.StatusOK, recorder.Code, recorder.Body.String())
	}
	assertSuccessResponse(t, body)
	assertDataValue(t, body.Data, "status", "ok")
}

func TestNewRouterReadyEndpoint(t *testing.T) {
	tests := []struct {
		name           string
		db             database.Database
		wantHTTPStatus int
		wantCode       int
		wantMessage    string
		wantStatus     string
		wantDBCheck    string
	}{
		{
			name:           "missing database",
			db:             nil,
			wantHTTPStatus: http.StatusServiceUnavailable,
			wantCode:       apperrors.ErrDatabaseError,
			wantMessage:    "not ready",
			wantStatus:     "not_ready",
			wantDBCheck:    "missing",
		},
		{
			name:           "ping failure",
			db:             &fakeDatabase{pingErr: errors.New("db offline")},
			wantHTTPStatus: http.StatusServiceUnavailable,
			wantCode:       apperrors.ErrDatabaseError,
			wantMessage:    "not ready",
			wantStatus:     "not_ready",
			wantDBCheck:    "db offline",
		},
		{
			name:           "ready",
			db:             &fakeDatabase{},
			wantHTTPStatus: http.StatusOK,
			wantCode:       0,
			wantMessage:    "success",
			wantStatus:     "ready",
			wantDBCheck:    "ok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newTestRouter(RouterDeps{Database: tt.db})

			recorder, body := performRouterRequest(t, router, http.MethodGet, "/ready")

			if recorder.Code != tt.wantHTTPStatus {
				t.Fatalf("expected /ready status %d, got %d with body %s", tt.wantHTTPStatus, recorder.Code, recorder.Body.String())
			}
			if body.Code != tt.wantCode {
				t.Fatalf("expected response code %d, got %d", tt.wantCode, body.Code)
			}
			if body.Message != tt.wantMessage {
				t.Fatalf("expected response message %q, got %q", tt.wantMessage, body.Message)
			}
			if body.Data == nil {
				t.Fatal("expected response data to be present")
			}
			assertDataValue(t, body.Data, "status", tt.wantStatus)
			checks, ok := body.Data["checks"].(map[string]any)
			if !ok {
				t.Fatalf("expected data.checks to be an object, got %#v", body.Data["checks"])
			}
			assertDataValue(t, checks, "database", tt.wantDBCheck)
		})
	}
}

func TestNewRouterDoesNotRegisterRemovedUserManagementRoutes(t *testing.T) {
	router := newTestRouter(RouterDeps{})

	for _, path := range []string{
		"/api/v1/auth/login",
		"/api/v1/auth/register",
		"/api/v1/users",
		"/api/v1/roles",
		"/api/v1/permissions",
	} {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, path, nil)

		router.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusNotFound {
			t.Fatalf("expected %s to be unregistered with status %d, got %d", path, http.StatusNotFound, recorder.Code)
		}
	}
}

func newTestRouter(deps RouterDeps) *gin.Engine {
	gin.SetMode(gin.TestMode)
	return NewRouter(deps)
}

func performRouterRequest(t *testing.T, router http.Handler, method string, path string) (*httptest.ResponseRecorder, routerResponse) {
	t.Helper()

	request := httptest.NewRequest(method, path, nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var body routerResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response body %q: %v", recorder.Body.String(), err)
	}
	return recorder, body
}

func assertSuccessResponse(t *testing.T, body routerResponse) {
	t.Helper()

	if body.Code != 0 {
		t.Fatalf("expected response code 0, got %d", body.Code)
	}
	if body.Message != "success" {
		t.Fatalf("expected response message success, got %q", body.Message)
	}
	if body.Data == nil {
		t.Fatal("expected response data to be present")
	}
}

func assertDataValue(t *testing.T, data map[string]any, key string, want string) {
	t.Helper()

	got, ok := data[key].(string)
	if !ok {
		t.Fatalf("expected data.%s to be a string, got %#v", key, data[key])
	}
	if got != want {
		t.Fatalf("expected data.%s %q, got %q", key, want, got)
	}
}
