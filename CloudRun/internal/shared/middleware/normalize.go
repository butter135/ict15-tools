package middleware

import (
	"net/http"
	"encoding/json"
	"bytes"
	"io"
)

func NormalizeInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			// formをパース
			if err := r.ParseForm(); err != nil {
				http.Error(w, "invalid form", http.StatusBadRequest)
				return
			}

			// フォームからJSONに変換
			data := map[string]string{
				"text":  r.FormValue("text"),
				"token": r.FormValue("token"),
			}
			jsonBytes, err := json.Marshal(data)
			if err != nil {
				http.Error(w, "cannot normalize input", http.StatusInternalServerError)
				return
			}

			// r.BodyをJSONに差し替える（↓重要）
			r.Body = io.NopCloser(bytes.NewReader(jsonBytes))
			r.Header.Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}
