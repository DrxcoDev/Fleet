
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// sdl2_imageFormula represents a formula in Go.
type sdl2_imageFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg sdl2_imageFormula) Print() {
    fmt.Printf("Name: sdl2_image\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := sdl2_imageFormula{
        Description:  "Library for loading images as SDL surfaces and textures",
        Homepage:     "https://github.com/libsdl-org/SDL_image",
        URL:          "https://github.com/libsdl-org/SDL_image.git",
        Sha256:       "f4508612c1f8603d490e9947ccc09a2954a482d76f94962d6bfc18577616db52",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "jpeg-turbo", "jpeg-xl", "libavif", "libpng", "libtiff", "sdl2", "webp"},
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

    fmt.Printf("Installing formula: %s\n", "sdl2_image")
    if err := pkg.Installsdl2_image(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg sdl2_imageFormula) Installsdl2_image() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "SDL_image.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "SDL_image"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}