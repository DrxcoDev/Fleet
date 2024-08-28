
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// twtxtFormula represents a formula in Go.
type twtxtFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg twtxtFormula) Print() {
    fmt.Printf("Name: twtxt\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := twtxtFormula{
        Description:  "Decentralised, minimalist microblogging service for hackers",
        Homepage:     "https://github.com/buckket/twtxt",
        URL:          "https://github.com/buckket/twtxt/commit/12bdd3670bff339fd27a7cd71c8ec64086b4ae5b.patch?full_index=1",
        Sha256:       "e206e7d18040d2b6c0d93ef2d7e4770c3e24448621bc6b5e0f206e193c6298ad",
        Dependencies: []string{"python@3.12"},
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

    fmt.Printf("Installing formula: %s\n", "twtxt")
    if err := pkg.Installtwtxt(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg twtxtFormula) Installtwtxt() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "12bdd3670bff339fd27a7cd71c8ec64086b4ae5b.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "12bdd3670bff339fd27a7cd71c8ec64086b4ae5b"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
