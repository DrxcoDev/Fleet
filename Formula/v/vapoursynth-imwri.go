
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// vapoursynth-imwriFormula represents a formula in Go.
type vapoursynth-imwriFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg vapoursynth-imwriFormula) Print() {
    fmt.Printf("Name: vapoursynth-imwri\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := vapoursynth-imwriFormula{
        Description:  "VapourSynth filters - ImageMagick HDRI writer/reader",
        Homepage:     "https://github.com/vapoursynth/vs-imwri",
        URL:          "https://github.com/vapoursynth/vs-imwri/archive/refs/tags/R2.tar.gz",
        Sha256:       "ddcbf221b69cd9a1dd5bae38d04f951bd1096897a9af4192b90a72ffad84a468",
        Dependencies: []string{"meson", "ninja", "pkg-config", "imagemagick", "vapoursynth", "jpeg-xl", "libheif", "libtiff"},
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

    fmt.Printf("Installing formula: %s\n", "vapoursynth-imwri")
    if err := pkg.Installvapoursynth-imwri(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg vapoursynth-imwriFormula) Installvapoursynth-imwri() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "R2.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "R2.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}