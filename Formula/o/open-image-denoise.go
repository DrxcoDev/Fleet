
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// open-image-denoiseFormula represents a formula in Go.
type open-image-denoiseFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg open-image-denoiseFormula) Print() {
    fmt.Printf("Name: open-image-denoise\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := open-image-denoiseFormula{
        Description:  "High-performance denoising library for ray tracing",
        Homepage:     "https://openimagedenoise.github.io",
        URL:          "https://github.com/RenderKit/oidn/commit/e5e52d335c58365b6cbd91f9a8a6f9ee9a085bf5.patch?full_index=1",
        Sha256:       "e5e42bb52b9790bbce3c8f82413986d5a23d389e1488965b738810b0d9fb0d2a",
        Dependencies: []string{"cmake", "ispc", "python@3.12", "tbb"},
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

    fmt.Printf("Installing formula: %s\n", "open-image-denoise")
    if err := pkg.Installopen-image-denoise(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg open-image-denoiseFormula) Installopen-image-denoise() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "e5e52d335c58365b6cbd91f9a8a6f9ee9a085bf5.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "e5e52d335c58365b6cbd91f9a8a6f9ee9a085bf5"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}