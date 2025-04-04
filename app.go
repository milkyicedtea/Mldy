package main

import (
	"Wails/utils"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// App struct
type App struct {
	ctx               context.Context
	ytdlpExecutable   string
	ffmpegExecutable  string
	dependenciesReady bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.dependenciesReady = false
	defer func() {
		a.dependenciesReady = true
		log.Println("Dependencies ready:", a.dependenciesReady)
	}()

	// Set up a default path where we'll store yt-dlp
	userDir, err := os.UserConfigDir()
	if err != nil {
		log.Println("Error getting user config dir:", err)
		return
	}

	appDir := filepath.Join(userDir, "Mldy")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		log.Println("Error creating app directory:", err)
		return
	}

	ytdlpPath := filepath.Join(appDir, "yt-dlp")
	if runtime.GOOS == "windows" {
		ytdlpPath += ".exe"
	}

	// Check if yt-dlp already exists
	if _, err := os.Stat(ytdlpPath); os.IsNotExist(err) {
		// Download yt-dlp manually with proper window hiding
		err = utils.DownloadYtdlp(ytdlpPath)
		if err != nil {
			log.Println("Error downloading yt-dlp:", err)
			return
		}
	}
	a.ytdlpExecutable = ytdlpPath

	ffmpegPath := filepath.Join(appDir, "ffmpeg")
	if runtime.GOOS == "windows" {
		ffmpegPath += ".exe"
	}

	// Check if FFmpeg already exists
	if _, err := os.Stat(ffmpegPath); os.IsNotExist(err) {
		err = utils.DownloadFFmpeg(appDir, ffmpegPath)
		if err != nil {
			log.Println("Error downloading FFmpeg:", err)
		}
	}
	log.Printf("a.ffmpegExecutable is going to be: %s\n", ffmpegPath)
	a.ffmpegExecutable = ffmpegPath

	log.Println("Using yt-dlp at:", ytdlpPath)
	log.Println("Using FFmpeg at:", ffmpegPath)
}

type VideoRequest struct {
	Url string `json:"url"`
}

func (a *App) AreDependenciesReady() bool {
	return a.dependenciesReady
}

func (a *App) Download(video VideoRequest) (string, error) {
	log.Printf("Downloading video %s", video.Url)

	if a.ytdlpExecutable == "" {
		return "", fmt.Errorf("yt-dlp not available")
	}

	metadataCmd := exec.Command(a.ytdlpExecutable,
		"--no-download",
		"--print-json",
		"--skip-download",
		video.Url,
	)

	if runtime.GOOS == "windows" {
		metadataCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, CreationFlags: 0x08000000}
	}

	metadataOutput, err := metadataCmd.Output()
	if err != nil {
		log.Println("Error fetching video metadata:", err)
		return "", nil
	}
	// log.Printf("Metadata output: %+v\n", string(metadataOutput))

	var videoInfo struct {
		Id       string `json:"id"`
		Title    string `json:"title"`
		Uploader string `json:"uploader"`
	}
	if err := json.Unmarshal(metadataOutput, &videoInfo); err != nil {
		log.Println("Error parsing video metadata:", err)
		return "", nil
	}

	log.Println("Metadata fetched - Title:", videoInfo.Title, "Uploader:", videoInfo.Uploader, "Id:", videoInfo.Id)

	filename := videoInfo.Title + ".mp3"
	filename = strings.Map(func(r rune) rune {
		if strings.ContainsRune(`<>:"/\|?*`, r) {
			return '-'
		}
		return r
	}, filename)

	targetDir := filepath.Join(utils.GetDownloadsFolder(), "Mldy Downloads")

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		err = os.MkdirAll(targetDir, 0755)
		if err != nil {
			log.Println("Error creating download directory:", err)
			return "", err
		}
	} else if err != nil {
		log.Println("Error checking download directory:", err)
		return "", err
	}

	outPath := filepath.Join(targetDir, filename)
	log.Printf("Downloading video %s to %s", video.Url, outPath)

	//log.Println("ytdlp cmd")
	ytdlCmd := exec.Command(a.ytdlpExecutable,
		"--format", "bestaudio[ext!=webm]",
		"--no-cache-dir",
		"--output", "-",
		video.Url,
	)

	if runtime.GOOS == "windows" {
		ytdlCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, CreationFlags: 0x08000000}
	}

	ytdlOut, err := ytdlCmd.StdoutPipe()
	if err != nil {
		log.Println("Error creating yt-dlp pipe:", err)
		return "", nil
	}

	ytdlErr, err := ytdlCmd.StderrPipe()
	if err != nil {
		log.Println("Error creating yt-dlp pipe:", err)
		return "", nil
	}

	//log.Println("ffmpeg cmd")
	log.Printf("a.ffmpegExecutable: %s\n", a.ffmpegExecutable)
	ffmpegCmd := exec.Command(a.ffmpegExecutable,
		"-i", "pipe:0",
		"-i", fmt.Sprintf("https://img.youtube.com/vi/%s/maxresdefault.jpg", videoInfo.Id),
		"-map", "0:a",
		"-map", "1:v",
		"-ab", "320k",
		"-ar", "48000",
		"-metadata", "title="+videoInfo.Title,
		"-metadata", "artist="+videoInfo.Uploader,
		"-c:v", "mjpeg",
		"-id3v2_version", "3",
		"-metadata:s:v", "title=Album cover",
		"-metadata:s:v", "comment=Cover (front)",
		"-disposition:v", "attached_pic",
		"-f", "mp3",
		"-y",
		outPath,
	)
	//log.Println("ffmpeg cmd finished")

	if runtime.GOOS == "windows" {
		ffmpegCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, CreationFlags: 0x08000000}
	}

	ffmpegCmd.Stdin = ytdlOut
	ffmpegErr, err := ffmpegCmd.StderrPipe()
	if err != nil {
		log.Println("Error creating ffmpeg error pipe:", err)
		return "", err
	}

	// start yt-dlp
	if err := ytdlCmd.Start(); err != nil {
		log.Println("Error starting yt-dlp:", err)
	}

	// start FFmpeg
	if err := ffmpegCmd.Start(); err != nil {
		log.Println("Error running FFmpeg:", err)
	}

	// capture errors from yt-dlp
	go func() {
		errOutput, _ := io.ReadAll(ytdlErr)
		if len(errOutput) > 0 {
			log.Println("yt-dlp error output:", string(errOutput))
		}
	}()

	// capture errors from ffmpeg
	go func() {
		errOutput, _ := io.ReadAll(ffmpegErr)
		if len(errOutput) > 0 {
			log.Println("FFmpeg error output:", string(errOutput))
		}
	}()

	if err := ffmpegCmd.Wait(); err != nil {
		log.Println("Error waiting for FFmpeg:", err)
		return "", err
	}

	// wait for ytdlp to finish
	if err := ytdlCmd.Wait(); err != nil {
		log.Println("yt-dlp finished with error:", err)
	}

	return outPath, nil
}
