
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// ffmpeg@2.8Formula represents a formula in Go.
type ffmpeg@2.8Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg ffmpeg@2.8Formula) Print() {
    fmt.Printf("Name: ffmpeg@2.8\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := ffmpeg@2.8Formula{
        Description:  "Play, record, convert, and stream audio and video",
        Homepage:     "https://ffmpeg.org/",
        URL:          "https://ffmpeg.org/download.html",
        Sha256:       "a547e14ec86a6927096e887ac5f7308cdfdd3f4e0d3341afb3367f2e86e0e82a",
        Dependencies: []string{"pkg-config", "texi2html", "yasm", "fontconfig", "freetype", "frei0r", "lame", "libass", "libogg", "libvo-aacenc", "libvorbis", "libvpx", "opencore-amr", "openssl@3", "opus", "rtmpdump", "sdl12-compat", "snappy", "speex", "theora", "x264", "x265", "xvid", "xz", "alsa-lib"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        if !isFormulaInstalled(dep) {
            fmt.Printf("Installing dependency: %s\n", dep)
            cmd := exec.Command("brew", "install", dep)
            if err := cmd.Run(); err != nil {
                log.Fatalf("Error installing dependency %s: %v", dep, err)
            }
        } else {
            fmt.Printf("Dependency %s is already installed.\n", dep)
        }
    }

    fmt.Printf("Installing formula: %s\n", "ffmpeg@2.8")
    if err := pkg.Installffmpeg@2.8(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg ffmpeg@2.8Formula) Installffmpeg@2.8() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "download.html"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "download"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}