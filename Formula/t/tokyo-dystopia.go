
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// tokyo-dystopiaFormula represents a formula in Go.
type tokyo-dystopiaFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg tokyo-dystopiaFormula) Print() {
    fmt.Printf("Name: tokyo-dystopia\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := tokyo-dystopiaFormula{
        Description:  "Lightweight full-text search system",
        Homepage:     "https://dbmx.net/tokyodystopia/",
        URL:          "https://dbmx.net/tokyodystopia/tokyodystopia-0.9.15.tar.gz",
        Sha256:       "1babf4df457924711f3ad2fc7e8f9f797950f53de9f1b740290ec1c112ca83e8",
        Dependencies: []string{"tokyo-cabinet"},
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

    fmt.Printf("Installing formula: %s\n", "tokyo-dystopia")
    if err := pkg.Installtokyo-dystopia(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg tokyo-dystopiaFormula) Installtokyo-dystopia() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "tokyodystopia-0.9.15.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "tokyodystopia-0.9.15.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}