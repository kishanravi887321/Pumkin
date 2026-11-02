package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const videoURL = "https://m.youtube.com/watch?v=asG7cwxi1sA&list=RD5Bme69vWxkQ&index=2"

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed to determine repository root: %v\n", err)
		os.Exit(1)
	}

	logFile, err := os.OpenFile(filepath.Join(root, "out.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		fmt.Printf("failed to create out.log: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	logger := log.New(io.MultiWriter(os.Stdout, logFile), "", log.LstdFlags)
	logger.Printf("download started; output directory: %s", root)
	logger.Printf("source URL: %s", videoURL)

	downloaderPath, err := findYTDLP()
	if err != nil {
		logger.Printf("yt-dlp is required but was not found: %v", err)
		logger.Println("Install it with: python -m pip install --upgrade yt-dlp")
		os.Exit(1)
	}
	logger.Printf("yt-dlp executable: %s", downloaderPath)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	outputTemplate := filepath.Join(root, "%(title)s.%(ext)s")
	command := exec.CommandContext(
		ctx,
		downloaderPath,
		"--newline",
		"--no-playlist",
		"--format", "bestvideo*+bestaudio/best",
		"--merge-output-format", "mp4",
		"--output", outputTemplate,
		videoURL,
	)
	command.Stdout = logger.Writer()
	command.Stderr = logger.Writer()

	if err := command.Run(); err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			logger.Println("download failed: 30-minute timeout exceeded")
		} else {
			logger.Printf("download failed: %v", err)
		}
		os.Exit(1)
	}

	logger.Println("download completed successfully")
}

func findYTDLP() (string, error) {
	if path, err := exec.LookPath("yt-dlp"); err == nil {
		return path, nil
	}

	if path, err := exec.LookPath("yt-dlp.exe"); err == nil {
		return path, nil
	}

	return "", errors.New("yt-dlp was not found on PATH")
}
