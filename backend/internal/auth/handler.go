package auth

import (
	"encoding/json"
	"net/http"

	"idinex-go/internal/shared"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		shared.WriteError(
			w,
			http.StatusMethodNotAllowed,
			"method_not_allowed",
			"Method not allowed.",
		)
		return
	}

	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(
			w,
			http.StatusBadRequest,
			"invalid_request",
			"Invalid request body.",
		)
		return
	}

	if err := ValidateRegisterRequest(req); err != nil {
		shared.WriteError(
			w,
			http.StatusBadRequest,
			"invalid_request",
			err.Error(),
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
