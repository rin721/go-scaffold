package httptransport

// 本测试文件固定 HTTP 路由、中间件顺序和错误响应契约，防止注释补全和后续重构改变外部可观察行为。

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

// DB 实现数据库测试桩的底层连接访问入口，按测试场景返回预置的 GORM 句柄。
func (db *fakeDatabase) DB() *gorm.DB {
	return nil
}

// Close 实现测试桩的资源关闭入口，用于验证生命周期调用而不释放外部资源。
func (db *fakeDatabase) Close() error {
	return nil
}

// Ping 实现数据库测试桩的健康检查入口，按测试需要返回成功或预设错误。
func (db *fakeDatabase) Ping(context.Context) error {
	return db.pingErr
}

// Reload 实现测试桩的配置重载入口，用于验证调用路径而不触发真实资源替换。
func (db *fakeDatabase) Reload(*database.Config) error {
	return nil
}

// WithTx 实现数据库测试桩的事务入口，用于把被测逻辑限制在可观察的回调调用内。
func (db *fakeDatabase) WithTx(context.Context, database.TxFunc) error {
	return nil
}

// WithTxOptions 实现数据库测试桩的事务入口，用于把被测逻辑限制在可观察的回调调用内。
func (db *fakeDatabase) WithTxOptions(context.Context, *database.TxOptions, database.TxFunc) error {
	return nil
}

type routerResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

// TestNewRouterHealthEndpoint 固定 HTTP 路由、中间件顺序和错误响应契约，确保后续注释补全或结构调整不改变该场景。
func TestNewRouterHealthEndpoint(t *testing.T) {
	router := newTestRouter(RouterDeps{})

	recorder, body := performRouterRequest(t, router, http.MethodGet, "/health")

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected /health status %d, got %d with body %s", http.StatusOK, recorder.Code, recorder.Body.String())
	}
	assertSuccessResponse(t, body)
	assertDataValue(t, body.Data, "status", "ok")
}

// TestNewRouterReadyEndpoint 固定 HTTP 路由、中间件顺序和错误响应契约，确保后续注释补全或结构调整不改变该场景。
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

// TestNewRouterDoesNotRegisterRemovedUserManagementRoutes 固定 HTTP 路由、中间件顺序和错误响应契约，确保后续注释补全或结构调整不改变该场景。
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

// newTestRouter 构造当前测试场景所需的最小依赖集合，避免测试直接耦合生产装配流程。
func newTestRouter(deps RouterDeps) *gin.Engine {
	gin.SetMode(gin.TestMode)
	return NewRouter(deps)
}

// performRouterRequest 执行测试 HTTP 请求并返回响应记录器，封装路由调用细节。
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

// assertSuccessResponse 校验测试响应或状态中的关键字段，使测试断言聚焦在对外契约而非重复解析细节。
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

// assertDataValue 校验测试响应或状态中的关键字段，使测试断言聚焦在对外契约而非重复解析细节。
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
