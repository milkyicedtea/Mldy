package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

func DownloadYtdlp(destPath string) error {
	// Determine the URL based on platform
	url := "https://github.com/yt-dlp/yt-dlp/releases/latest/download/"
	if runtime.GOOS == "windows" {
		url += "yt-dlp.exe"
	} else {
		url += "yt-dlp"
	}

	// Download the file using Go's http client (no external dependencies)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download yt-dlp: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status when downloading yt-dlp: %s", resp.Status)
	}

	// Create the output file
	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer out.Close()

	// Copy the content to the output file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write yt-dlp to file: %w", err)
	}

	// Make the file executable on non-Windows platforms
	if runtime.GOOS != "windows" {
		if err := os.Chmod(destPath, 0755); err != nil {
			return fmt.Errorf("failed to make yt-dlp executable: %w", err)
		}
	}

	log.Println("Successfully downloaded yt-dlp to", destPath)
	return nil
}

func DownloadFFmpeg(appDir, destPath string) error {
	// FFmpeg download is more complex as it comes in an archive
	var url string
	var archiveName string
	var executablePath string

	// Determine the URL and paths based on platform
	if runtime.GOOS == "windows" {
		// Example URL - you might need to update this with a more current version
		url = "https://github.com/BtbN/FFmpeg-Builds/releases/download/latest/ffmpeg-master-latest-win64-gpl.zip"
		archiveName = filepath.Join(appDir, "ffmpeg.zip")
		executablePath = "ffmpeg-master-latest-win64-gpl/bin/ffmpeg.exe"
	} else if runtime.GOOS == "darwin" {
		url = "https://evermeet.cx/ffmpeg/getrelease/zip"
		archiveName = filepath.Join(appDir, "ffmpeg.zip")
		executablePath = "ffmpeg"
	} else {
		// Linux - you might need a different approach for Linux
		url = "https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz"
		archiveName = filepath.Join(appDir, "ffmpeg.tar.xz")
		executablePath = "ffmpeg-*-amd64-static/ffmpeg"
	}

	// Download the archive
	cmd := exec.Command("curl", "-L", url, "-o", archiveName)
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: 0x08000000,
		}
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	// Extract the archive
	log.Println("Successfully downloaded ffmpeg to", destPath)
	log.Println("Extracting ffmpeg archive...")
	if strings.HasSuffix(archiveName, ".zip") {
		// Extract ZIP
		cmd = exec.Command("powershell", "-Command",
			"Expand-Archive", "-Path", archiveName, "-DestinationPath", appDir, "-Force")
		if runtime.GOOS == "windows" {
			cmd.SysProcAttr = &syscall.SysProcAttr{
				HideWindow:    true,
				CreationFlags: 0x08000000,
			}
		}
		if err := cmd.Run(); err != nil {
			return err
		}
	} else if strings.HasSuffix(archiveName, ".tar.xz") {
		// Extract tar.xz on Linux
		cmd = exec.Command("tar", "-xf", archiveName, "-C", appDir)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	// Find the extracted ffmpeg executable
	var ffmpegPath string
	if runtime.GOOS == "windows" {
		ffmpegPath = filepath.Join(appDir, executablePath)
	} else {
		// For Linux/Mac, might need to find the file
		matches, err := filepath.Glob(filepath.Join(appDir, executablePath))
		if err != nil || len(matches) == 0 {
			return fmt.Errorf("could not find extracted ffmpeg executable")
		}
		ffmpegPath = matches[0]
	}

	// Copy the executable to the destination
	input, err := os.Open(ffmpegPath)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer output.Close()

	if _, err = io.Copy(output, input); err != nil {
		return err
	}

	// Make executable on non-Windows platforms
	if runtime.GOOS != "windows" {
		if err := os.Chmod(destPath, 0755); err != nil {
			return err
		}
	}

	// Clean up the archive and extracted files
	os.Remove(archiveName)

	return nil
}

func GetDownloadsFolder() string {
	if runtime.GOOS == "windows" {
		// Use Windows API to get the true Downloads directory path
		return getWindowsDownloadsFolder()
	}

	// Fallback for other operating systems
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error getting home directory:", err)
		return ""
	}
	return filepath.Join(homeDir, "Downloads")
}

func getWindowsDownloadsFolder() string {
	// Windows Known Folder ID for Downloads
	// {374DE290-123F-4565-9164-39C4925E467B}
	folderIDDownloads := syscall.GUID{
		Data1: 0x374DE290,
		Data2: 0x123F,
		Data3: 0x4565,
		Data4: [8]byte{0x91, 0x64, 0x39, 0xC4, 0x92, 0x5E, 0x46, 0x7B},
	}

	// load shell32.dll
	shell32 := syscall.NewLazyDLL("shell32.dll")
	//if err != nil {
	//	log.Println("Error loading shell32.dll:", err)
	//	return filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
	//}
	//defer func(shell32 *syscall.LazyDLL) {
	//	err := shell32.Release()
	//	if err != nil {
	//		log.Println("Error releasing shell32:", err)
	//	}
	//}(shell32)

	// get SHGetKnownFolderPath function
	proc := shell32.NewProc("SHGetKnownFolderPath")
	//if err != nil {
	//	log.Println("Error finding SHGetKnownFolderPath:", err)
	//	return filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
	//}

	var path uintptr
	// call SHGetKnownFolderPath
	r, _, _ := proc.Call(
		uintptr(unsafe.Pointer(&folderIDDownloads)), // rfid
		0,                                           // dwFlags
		0,                                           // hToken
		uintptr(unsafe.Pointer(&path)),              // ppszPath
	)

	if r != 0 {
		return filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
	}

	// convert wide char to string
	// path is returned as a wide char string
	downloadsPath := syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(path))[:])

	// free the allocated memory
	ole32 := syscall.NewLazyDLL("ole32.dll")
	//if err == nil {
	//	defer func(ole32 *syscall.DLL) {
	//		err := ole32.Release()
	//		if err != nil {
	//			log.Println("Error releasing ole32:", err)
	//		}
	//	}(ole32)
	proc = ole32.NewProc("CoTaskMemFree")
	proc.Call(path)

	return downloadsPath
}
