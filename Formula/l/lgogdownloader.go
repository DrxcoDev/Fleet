
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// lgogdownloaderFormula represents a formula in Go.
type lgogdownloaderFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg lgogdownloaderFormula) Print() {
    fmt.Printf("Name: lgogdownloader\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := lgogdownloaderFormula{
        Description:  "Unofficial downloader for GOG.com games",
        Homepage:     "https://sites.google.com/site/gogdownloader/",
        URL:          "https://github.com/Sude-/lgogdownloader/releases/download/v3.15/lgogdownloader-3.15.tar.gz",
        Sha256:       "e86b00e2db846abc67a68d99e04b029fb4ec8018f934f7bd1321fbcc5425e9e4",
        Dependencies: []string{"cmake", "help2man", "pkg-config", "boost", "htmlcxx", "jsoncpp", "rhash", "tidy-html5", "tinyxml2"},
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

    fmt.Printf("Installing formula: %s\n", "lgogdownloader")
    if err := pkg.Installlgogdownloader(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg lgogdownloaderFormula) Installlgogdownloader() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "lgogdownloader-3.15.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "lgogdownloader-3.15.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
