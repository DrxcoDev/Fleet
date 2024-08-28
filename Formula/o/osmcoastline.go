
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// osmcoastlineFormula represents a formula in Go.
type osmcoastlineFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg osmcoastlineFormula) Print() {
    fmt.Printf("Name: osmcoastline\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := osmcoastlineFormula{
        Description:  "Extracts coastline data from OpenStreetMap planet file",
        Homepage:     "https://osmcode.org/osmcoastline/",
        URL:          "https://github.com/osmcode/osmcoastline/commit/67cc33161069f65e315acae952492ab5ee07af15.patch?full_index=1",
        Sha256:       "31b89e33b22ccdfe289a5da67480f9791bdd4f410c6a7831f0c1e007c4258e68",
        Dependencies: []string{"cmake", "libosmium", "expat", "gdal", "geos", "libspatialite", "lz4"},
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

    fmt.Printf("Installing formula: %s\n", "osmcoastline")
    if err := pkg.Installosmcoastline(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg osmcoastlineFormula) Installosmcoastline() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "67cc33161069f65e315acae952492ab5ee07af15.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "67cc33161069f65e315acae952492ab5ee07af15"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
