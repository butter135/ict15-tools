package middleware

import (
	"net/http"
	"net/url"
	"encoding/json"
	"strings"
)

func NormalizeInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if strings.HasPrefix(ct, "application/json") {
			// JSONを読み取ってmapにパース
			var body map[string]string
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			// mapからFormに変換
			form := url.Values{}
			for k, v := range body {
				form.Set(k, v)
			}

			// r.Form / r.PostForm に設定
			r.PostForm = form
		} else {
			// 通常フォームならそのままパース
			_ = r.ParseForm()
		}

		next.ServeHTTP(w, r)
	})
}
