
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// mpvFormula represents a formula in Go.
type mpvFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg mpvFormula) Print() {
    fmt.Printf("Name: mpv\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := mpvFormula{
        Description:  "Media player based on MPlayer and mplayer2",
        Homepage:     "https://mpv.io",
        URL:          "https://github.com/mpv-player/mpv/archive/refs/tags/v0.38.0.tar.gz",
        Sha256:       "3a1e0b8cc2f36392714bd044fae7066f54b353421c32c782a6f67a9736e4ffdc",
        Dependencies: []string{"docutils", "meson", "pkg-config", "ffmpeg", "jpeg-turbo", "libarchive", "libass", "libbluray", "libplacebo", "little-cms2", "luajit", "mujs", "rubberband", "uchardet", "vapoursynth", "vulkan-loader", "yt-dlp", "zimg", "molten-vk", "alsa-lib", "libdrm", "libva", "libvdpau", "libx11", "libxext", "libxkbcommon", "libxpresent", "libxrandr", "libxscrnsaver", "libxv", "mesa", "pulseaudio", "wayland"},
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

    fmt.Printf("Installing formula: %s\n", "mpv")
    if err := pkg.Installmpv(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg mpvFormula) Installmpv() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v0.38.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v0.38.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
