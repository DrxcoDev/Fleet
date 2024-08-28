
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// sagittarius-schemeFormula represents a formula in Go.
type sagittarius-schemeFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg sagittarius-schemeFormula) Print() {
    fmt.Printf("Name: sagittarius-scheme\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := sagittarius-schemeFormula{
        Description:  "Free Scheme implementation supporting R6RS and R7RS",
        Homepage:     "https://bitbucket.org/ktakashi/sagittarius-scheme/wiki/Home",
        URL:          "https://bitbucket.org/ktakashi/sagittarius-scheme/downloads/sagittarius-0.9.11.tar.gz",
        Sha256:       "d14d120ab7276dccad71129c24ccac4760dc929e801e471f7d3d439801972bdb",
        Dependencies: []string{"cmake", "pkg-config", "bdw-gc", "openssl@3", "unixodbc"},
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

    fmt.Printf("Installing formula: %s\n", "sagittarius-scheme")
    if err := pkg.Installsagittarius-scheme(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg sagittarius-schemeFormula) Installsagittarius-scheme() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "sagittarius-0.9.11.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "sagittarius-0.9.11.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}