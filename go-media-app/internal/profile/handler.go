package profile

import (
	"fmt"
	"net/http"
	"go-chat-app/pkg/logger"
	"go-chat-app/internal/s3utils"
)

func GetProfilePhotoHandler(log *logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		userId := mux.Vars(r)["userId"]

		profilePhotoURL, err := s3utils.GetProfilePhotoURL(userId)
		if err != nil {
			log.Errorf("Error al obtener foto de perfil: %v", err)
			http.Error(w, fmt.Sprintf("Error al obtener la foto de perfil: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"photoUrl": "%s"}`, profilePhotoURL)
	}
}
