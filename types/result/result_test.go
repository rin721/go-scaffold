package result

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	apperrors "github.com/rei0721/go-scaffold/types/errors"
)

type responseBody struct {
	Code       int             `json:"code"`
	Message    string          `json:"message"`
	Data       json.RawMessage `json:"data,omitempty"`
	TraceID    string          `json:"traceId,omitempty"`
	ServerTime int64           `json:"serverTime"`
}

func TestSuccessErrorAndPaginationContracts(t *testing.T) {
	success := Success("ok")
	if success.Code != 0 {
		t.Fatalf("Success code = %d, want 0", success.Code)
	}
	if success.Message != "success" {
		t.Fatalf("Success message = %q, want success", success.Message)
	}
	if success.Data != "ok" {
		t.Fatalf("Success data = %q, want ok", success.Data)
	}
	if success.ServerTime == 0 {
		t.Fatal("Success server time should be set")
	}

	errResult := Error(apperrors.ErrInvalidParams, "bad request")
	if errResult.Code != apperrors.ErrInvalidParams {
		t.Fatalf("Error code = %d, want %d", errResult.Code, apperrors.ErrInvalidParams)
	}
	if errResult.TraceID != "" {
		t.Fatalf("Error trace id = %q, want empty", errResult.TraceID)
	}

	traceResult := ErrorWithTrace(apperrors.ErrInternalServer, "internal error", "trace-1")
	if traceResult.TraceID != "trace-1" {
		t.Fatalf("ErrorWithTrace trace id = %q, want trace-1", traceResult.TraceID)
	}

	page := NewPageResult([]string{"a", "b"}, 2, 10, 21)
	if page.Pagination.TotalPages != 3 {
		t.Fatalf("TotalPages = %d, want 3", page.Pagination.TotalPages)
	}
	if page.Pagination.Page != 2 || page.Pagination.PageSize != 10 || page.Pagination.Total != 21 {
		t.Fatalf("unexpected pagination metadata: %+v", page.Pagination)
	}
}

func TestGinResponseHelpersContract(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		call       func(*gin.Context)
		wantStatus int
		wantCode   int
	}{
		{
			name:       "bad request",
			call:       func(c *gin.Context) { BadRequest(c, "bad request") },
			wantStatus: http.StatusBadRequest,
			wantCode:   apperrors.ErrInvalidParams,
		},
		{
			name:       "unauthorized",
			call:       func(c *gin.Context) { Unauthorized(c, "unauthorized") },
			wantStatus: http.StatusUnauthorized,
			wantCode:   apperrors.ErrUnauthorized,
		},
		{
			name:       "forbidden",
			call:       func(c *gin.Context) { Forbidden(c, "forbidden") },
			wantStatus: http.StatusForbidden,
			wantCode:   apperrors.ErrPermissionDenied,
		},
		{
			name:       "not found",
			call:       func(c *gin.Context) { NotFound(c, "not found") },
			wantStatus: http.StatusNotFound,
			wantCode:   apperrors.ErrResourceNotFound,
		},
		{
			name:       "internal error",
			call:       func(c *gin.Context) { InternalError(c, "internal error") },
			wantStatus: http.StatusInternalServerError,
			wantCode:   apperrors.ErrInternalServer,
		},
		{
			name:       "generic fail",
			call:       func(c *gin.Context) { Fail(c, http.StatusBadRequest, "bad request") },
			wantStatus: http.StatusBadRequest,
			wantCode:   apperrors.ErrInvalidParams,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Set("trace_id", "trace-1")

			tt.call(ctx)

			if recorder.Code != tt.wantStatus {
				t.Fatalf("HTTP status = %d, want %d", recorder.Code, tt.wantStatus)
			}

			var body responseBody
			if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
				t.Fatalf("decode response: %v", err)
			}
			if body.Code != tt.wantCode {
				t.Fatalf("body code = %d, want %d", body.Code, tt.wantCode)
			}
			if body.TraceID != "trace-1" {
				t.Fatalf("trace id = %q, want trace-1", body.TraceID)
			}
			if body.ServerTime == 0 {
				t.Fatal("server time should be set")
			}
		})
	}
}

func TestGetTraceIDContract(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	if got := GetTraceID(ctx); got != "" {
		t.Fatalf("GetTraceID without value = %q, want empty", got)
	}

	ctx.Set("trace_id", 123)
	if got := GetTraceID(ctx); got != "" {
		t.Fatalf("GetTraceID with non-string value = %q, want empty", got)
	}

	ctx.Set("trace_id", "trace-1")
	if got := GetTraceID(ctx); got != "trace-1" {
		t.Fatalf("GetTraceID = %q, want trace-1", got)
	}
}
