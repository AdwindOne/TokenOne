package api

import (
	"net/http"

	"github.com/tokenone/tokenone-backend/internal/core/wallets"
	"github.com/tokenone/tokenone-backend/pkg/utils"
)

// Response structures
type MnemonicResponse struct {
	Mnemonic string `json:"mnemonic"`
}

type ImportRequest struct {
	Mnemonic string `json:"mnemonic"`
}

type ImportResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// CreateWalletHandler handles the creation of a new wallet.
// POST /api/v1/wallets/create
func CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	mnemonic, err := wallets.GenerateMnemonic()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate mnemonic")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, MnemonicResponse{Mnemonic: mnemonic})
}

// ImportWalletHandler handles the import of an existing wallet using a mnemonic.
// POST /api/v1/wallets/import
func ImportWalletHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	var req ImportRequest
	if err := utils.DecodeJSONBody(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	if req.Mnemonic == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Mnemonic cannot be empty")
		return
	}

	isValid := wallets.ValidateMnemonic(req.Mnemonic)
	if !isValid {
		utils.RespondWithJSON(w, http.StatusOK, ImportResponse{Success: false, Message: "Invalid mnemonic"})
		return
	}

	// In a real application, we would store or further process the validated mnemonic here.
	// For now, just confirm validation.
	utils.RespondWithJSON(w, http.StatusOK, ImportResponse{Success: true, Message: "Mnemonic is valid and imported (placeholder)"})
}

// ConfigureRoutes sets up the API routes.
// This function will be called from main.go
func ConfigureRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/wallets/create", CreateWalletHandler)
	mux.HandleFunc("/api/v1/wallets/import", ImportWalletHandler)
}
