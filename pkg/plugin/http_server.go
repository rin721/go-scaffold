package plugin

import (
	"encoding/json"
	"net/http"
)

// NewHTTPServer exposes a Plugin through the standard JSON HTTP transport.
func NewHTTPServer(plugin Plugin) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != HTTPInvokePath {
			writeHTTPPluginError(w, http.StatusNotFound, "plugin endpoint not found")
			return
		}
		if r.Method != http.MethodPost {
			writeHTTPPluginError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		if plugin == nil {
			writeHTTPPluginError(w, http.StatusInternalServerError, ErrNilPlugin.Error())
			return
		}

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeHTTPPluginError(w, http.StatusBadRequest, err.Error())
			return
		}
		if err := req.Validate(); err != nil {
			writeHTTPPluginError(w, http.StatusBadRequest, err.Error())
			return
		}
		if req.Plugin == "" {
			req.Plugin = plugin.Metadata().Name
		}

		resp, err := plugin.Invoke(r.Context(), req)
		if resp == nil {
			resp = &Response{}
		}
		if err != nil {
			resp.Error = err.Error()
			writeHTTPPluginResponse(w, http.StatusInternalServerError, resp)
			return
		}
		writeHTTPPluginResponse(w, http.StatusOK, resp)
	})
}

func writeHTTPPluginError(w http.ResponseWriter, status int, message string) {
	writeHTTPPluginResponse(w, status, &Response{Error: message})
}

func writeHTTPPluginResponse(w http.ResponseWriter, status int, resp *Response) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}
