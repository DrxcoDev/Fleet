
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// epoll-shimFormula represents a formula in Go.
type epoll-shimFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg epoll-shimFormula) Print() {
    fmt.Printf("Name: epoll-shim\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := epoll-shimFormula{
        Description:  "Small epoll implementation using kqueue",
        Homepage:     "https://github.com/jiixyj/epoll-shim",
        URL:          "https://github.com/jiixyj/epoll-shim/archive/refs/tags/v0.0.20240608.tar.gz",
        Sha256:       "ad0c5b28e749ebd570f59dd630db71ce9d45435a664e44b1e2c290602d66c0db",
        Dependencies: []string{"cmake"},
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

    fmt.Printf("Installing formula: %s\n", "epoll-shim")
    if err := pkg.Installepoll-shim(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg epoll-shimFormula) Installepoll-shim() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v0.0.20240608.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v0.0.20240608.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
