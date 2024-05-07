package check

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/MatthewAraujo/health-checker-website/types"
	"github.com/MatthewAraujo/health-checker-website/utils"
	"github.com/go-playground/validator"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) *http.ServeMux {
	router.HandleFunc("/check", checkHandler)
	return router
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	var payload types.Payload
	if err := utils.ParseJSON(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("validation error: %s", errors))
		return
	}

	if payload.Domain == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("domain is required"))
		return
	}

	d := payload.Domain
	if err := Check(d); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Domain is reachable"})

}

func Check(d string) error {
	p := "80"
	address := d + ":" + p
	timeout := time.Duration(5 * time.Second)
	_, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return fmt.Errorf("%s is not reachable", address)
	}

	return nil
}
