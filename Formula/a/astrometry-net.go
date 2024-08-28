
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// astrometry-netFormula represents a formula in Go.
type astrometry-netFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg astrometry-netFormula) Print() {
    fmt.Printf("Name: astrometry-net\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := astrometry-netFormula{
        Description:  "Automatic identification of astronomical images",
        Homepage:     "https://github.com/dstndstn/astrometry.net",
        URL:          "https://github.com/dstndstn/astrometry.net/commit/7991b0d20280e8fc6a9b17d5669312eee69e4c43.patch?full_index=1",
        Sha256:       "4fa0feec1bd4159749789c5f115dd0e99e6a92e6032b90da273560a9064419d5",
        Dependencies: []string{"pkg-config", "python-setuptools", "swig", "cairo", "cfitsio", "gsl", "jpeg-turbo", "libpng", "netpbm", "numpy", "python@3.12", "wcslib"},
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

    fmt.Printf("Installing formula: %s\n", "astrometry-net")
    if err := pkg.Installastrometry-net(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg astrometry-netFormula) Installastrometry-net() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "7991b0d20280e8fc6a9b17d5669312eee69e4c43.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "7991b0d20280e8fc6a9b17d5669312eee69e4c43"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}