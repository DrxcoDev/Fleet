
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// rdkitFormula represents a formula in Go.
type rdkitFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg rdkitFormula) Print() {
    fmt.Printf("Name: rdkit\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := rdkitFormula{
        Description:  "Open-source chemoinformatics library",
        Homepage:     "https://rdkit.org/",
        URL:          "https://github.com/rdkit/rdkit/archive/refs/tags/Release_2024_03_5.tar.gz",
        Sha256:       "daa5c0cf45a634f85ea6035150a1a4e1710c0596d688ce92d6f903c3e08e7f25",
        Dependencies: []string{"catch2", "cmake", "pkg-config", "swig", "boost", "boost-python3", "cairo", "eigen", "freetype", "numpy", "postgresql@14", "py3cairo", "python@3.12"},
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

    fmt.Printf("Installing formula: %s\n", "rdkit")
    if err := pkg.Installrdkit(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg rdkitFormula) Installrdkit() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "Release_2024_03_5.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "Release_2024_03_5.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
