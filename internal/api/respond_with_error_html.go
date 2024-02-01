package api

import (
	"fmt"
	"net/http"
)

func respondWithErrorHTML(w http.ResponseWriter, code int, errorMsg string) {
	errorHTML := fmt.Sprintf(`
        <div class="error">
            <p>%s</p>
        </div>
    `, errorMsg)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	w.Write([]byte(errorHTML))
}
