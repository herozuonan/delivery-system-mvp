package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ChatRequest struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	Message   string `json:"message"`
	Stage     string `json:"stage,omitempty"`
	Goal      string `json:"goal,omitempty"`
}

type ChatResponse struct {
	Reply     string `json:"reply"`
	SessionID string `json:"session_id"`
	TraceID   string `json:"trace_id"`
	Stage     string `json:"stage,omitempty"`
	Goal      string `json:"goal,omitempty"`
}

var (
	reqTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "aics",
			Name:      "http_requests_total",
			Help:      "Total HTTP requests",
		},
		[]string{"path", "method", "code"},
	)
	reqDur = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "aics",
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request duration in seconds",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

func main() {
	prometheus.MustRegister(reqTotal, reqDur)

	addr := getenv("ADDR", ":8080")
	mode := strings.ToLower(getenv("AI_MODE", "mock"))

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", instrument("/healthz", healthzHandler))
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/v1/chat", instrument("/v1/chat", chatHandler(mode)))
	mux.HandleFunc("/v1/chat/stream", instrument("/v1/chat/stream", chatStreamHandler(mode)))
	mux.HandleFunc("/", instrument("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("ai-cs-api"))
	}))

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      0, // streaming
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("ai-cs-api listening on %s (AI_MODE=%s)", addr, mode)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %v", err)
	}
}

func instrument(path string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &statusWriter{ResponseWriter: w, status: 200}
		h(rw, r)
		code := fmt.Sprintf("%d", rw.status)
		reqTotal.WithLabelValues(path, r.Method, code).Inc()
		reqDur.WithLabelValues(path, r.Method).Observe(time.Since(start).Seconds())
	}
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// Flush ensures streaming endpoints (SSE) keep working even when wrapped.
func (w *statusWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func healthzHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("ok"))
}

func chatHandler(mode string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req ChatRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&req); err != nil {
			http.Error(w, "invalid json: "+err.Error(), http.StatusBadRequest)
			return
		}

		if strings.TrimSpace(req.Message) == "" {
			http.Error(w, "message required", http.StatusBadRequest)
			return
		}
		if req.SessionID == "" {
			req.SessionID = "s-" + shortID()
		}

		traceID := "t-" + shortID()
		reply := generateReply(r.Context(), mode, req)

		resp := ChatResponse{
			Reply:     reply,
			SessionID: req.SessionID,
			TraceID:   traceID,
			Stage:     req.Stage,
			Goal:      req.Goal,
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(resp)
	}
}

func chatStreamHandler(mode string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// SSE is typically GET.
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		q := r.URL.Query()
		req := ChatRequest{
			UserID:    q.Get("user_id"),
			SessionID: q.Get("session_id"),
			Message:   q.Get("message"),
			Stage:     q.Get("stage"),
			Goal:      q.Get("goal"),
		}
		if strings.TrimSpace(req.Message) == "" {
			http.Error(w, "message query param required", http.StatusBadRequest)
			return
		}
		if req.SessionID == "" {
			req.SessionID = "s-" + shortID()
		}

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming unsupported", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no")

		traceID := "t-" + shortID()
		startEvent(w, flusher, map[string]any{
			"trace_id":   traceID,
			"session_id": req.SessionID,
			"stage":      req.Stage,
			"goal":       req.Goal,
		})

		reply := generateReply(r.Context(), mode, req)
		// Cheap “token-like” streaming for MVP.
		chunks := chunkString(reply, 16)
		for _, c := range chunks {
			if err := writeSSE(w, "delta", map[string]any{"text": c}); err != nil {
				return
			}
			flusher.Flush()
			select {
			case <-r.Context().Done():
				return
			case <-time.After(30 * time.Millisecond):
			}
		}

		_ = writeSSE(w, "done", map[string]any{"trace_id": traceID})
		flusher.Flush()
	}
}

func startEvent(w http.ResponseWriter, flusher http.Flusher, payload any) {
	_ = writeSSE(w, "start", payload)
	flusher.Flush()
}

func writeSSE(w http.ResponseWriter, event string, payload any) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	// SSE framing
	_, err = fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event, string(b))
	return err
}

func generateReply(ctx context.Context, mode string, req ChatRequest) string {
	_ = ctx
	// MVP: mock mode only; we will plug in real LLM later behind Model Gateway.
	s := strings.TrimSpace(req.Stage)
	g := strings.TrimSpace(req.Goal)
	prefix := ""
	if s != "" || g != "" {
		prefix = fmt.Sprintf("[stage=%s goal=%s] ", s, g)
	}
	if mode == "mock" {
		return prefix + "收到：" + req.Message + "（MVP mock 回复）"
	}
	return prefix + "收到：" + req.Message
}

func shortID() string {
	var b [6]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

func chunkString(s string, n int) []string {
	if n <= 0 || len(s) <= n {
		return []string{s}
	}
	out := make([]string, 0, (len(s)+n-1)/n)
	for len(s) > 0 {
		if len(s) <= n {
			out = append(out, s)
			break
		}
		out = append(out, s[:n])
		s = s[n:]
	}
	return out
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
