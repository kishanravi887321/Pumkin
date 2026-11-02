package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	defaultAddress  = ":8180"
	downloadTimeout = 30 * time.Minute
)

func main() {
	address := os.Getenv("CUTTER_GO_ADDRESS")
	if address == "" {
		address = defaultAddress
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/mp3", mp3Handler)

	server := &http.Server{
		Addr:              address,
		Handler:           cors(mux),
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Printf("cutter_go listening at %s", address)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func healthHandler(response http.ResponseWriter, _ *http.Request) {
	writeJSON(response, http.StatusOK, `{"status":"ok"}`)
}

func mp3Handler(response http.ResponseWriter, request *http.Request) {
	sourceURL := request.URL.Query().Get("url")
	if err := validateYouTubeURL(sourceURL); err != nil {
		writeJSON(response, http.StatusBadRequest, fmt.Sprintf(`{"error":%q}`, err.Error()))
		return
	}

	ytDLP, err := findExecutable("yt-dlp", "yt-dlp.exe")
	if err != nil {
		writeJSON(response, http.StatusInternalServerError, `{"error":"yt-dlp is not installed on the server"}`)
		return
	}
	ffmpeg, err := findExecutable("ffmpeg", "ffmpeg.exe")
	if err != nil {
		writeJSON(response, http.StatusInternalServerError, `{"error":"FFmpeg is not installed on the server"}`)
		return
	}

	workdir, err := os.MkdirTemp("", "cutter-go-mp3-*")
	if err != nil {
		writeJSON(response, http.StatusInternalServerError, `{"error":"could not create temporary download directory"}`)
		return
	}
	defer func() {
		if cleanupErr := os.RemoveAll(workdir); cleanupErr != nil {
			log.Printf("temporary cleanup failed for %s: %v", workdir, cleanupErr)
		} else {
			log.Printf("temporary directory removed: %s", workdir)
		}
	}()

	outputTemplate := filepath.Join(workdir, "%(title).200s [%(id)s].%(ext)s")
	ctx, cancel := context.WithTimeout(request.Context(), downloadTimeout)
	defer cancel()

	args := []string{
		"--newline",
		"--no-playlist",
		"--format", "bestaudio/best",
		"--output", outputTemplate,
		"--postprocessor-args", "ExtractAudio+ffmpeg_o:-loglevel error",
		"--extract-audio",
		"--audio-format", "mp3",
		"--audio-quality", "192K",
		"--ffmpeg-location", filepath.Dir(ffmpeg),
		sourceURL,
	}
	log.Printf("MP3 download started: %s", sourceURL)
	command := exec.CommandContext(ctx, ytDLP, args...)
	output, err := command.CombinedOutput()
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			writeJSON(response, http.StatusGatewayTimeout, `{"error":"MP3 download timed out"}`)
			return
		}
		log.Printf("yt-dlp failed: %s", strings.TrimSpace(string(output)))
		writeJSON(response, http.StatusInternalServerError, `{"error":"MP3 download or conversion failed"}`)
		return
	}

	mp3Path, err := findMP3(workdir)
	if err != nil {
		writeJSON(response, http.StatusInternalServerError, `{"error":"MP3 file was not created"}`)
		return
	}

	file, err := os.Open(mp3Path)
	if err != nil {
		writeJSON(response, http.StatusInternalServerError, `{"error":"MP3 file could not be opened"}`)
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		writeJSON(response, http.StatusInternalServerError, `{"error":"MP3 file metadata could not be read"}`)
		return
	}

	response.Header().Set("Content-Type", "audio/mpeg")
	response.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(mp3Path)))
	response.Header().Set("Content-Length", fmt.Sprint(info.Size()))
	http.ServeContent(response, request, filepath.Base(mp3Path), info.ModTime(), file)
	log.Printf("MP3 response completed: %s (%d bytes)", filepath.Base(mp3Path), info.Size())
}

func validateYouTubeURL(rawURL string) error {
	if rawURL == "" {
		return errors.New("query parameter 'url' is required")
	}
	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil || parsed.Scheme != "https" || parsed.Hostname() == "" {
		return errors.New("a valid HTTPS YouTube URL is required")
	}
	host := strings.ToLower(parsed.Hostname())
	if host != "youtube.com" && host != "www.youtube.com" && host != "m.youtube.com" && host != "youtu.be" {
		return errors.New("only YouTube URLs are supported")
	}
	return nil
}

func findMP3(workdir string) (string, error) {
	entries, err := os.ReadDir(workdir)
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.EqualFold(filepath.Ext(entry.Name()), ".mp3") {
			return filepath.Join(workdir, entry.Name()), nil
		}
	}
	return "", os.ErrNotExist
}

func findExecutable(names ...string) (string, error) {
	for _, name := range names {
		if path, err := exec.LookPath(name); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("%s not found", strings.Join(names, " or "))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if request.Method == http.MethodOptions {
			response.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(response, request)
	})
}

func writeJSON(response http.ResponseWriter, status int, body string) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	_, _ = response.Write([]byte(body))
}
