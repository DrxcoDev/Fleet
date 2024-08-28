
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// google-authenticator-libpamFormula represents a formula in Go.
type google-authenticator-libpamFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg google-authenticator-libpamFormula) Print() {
    fmt.Printf("Name: google-authenticator-libpam\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := google-authenticator-libpamFormula{
        Description:  "PAM module for two-factor authentication",
        Homepage:     "https://github.com/google/google-authenticator-libpam",
        URL:          "https://github.com/google/google-authenticator-libpam/archive/refs/tags/1.10.tar.gz",
        Sha256:       "9d9a5c7680d94e3ff83b8c38eec4450c933e15d9135b7e270508f500b0ebe920",
        Dependencies: []string{"autoconf", "automake", "libtool", "qrencode", "linux-pam"},
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

    fmt.Printf("Installing formula: %s\n", "google-authenticator-libpam")
    if err := pkg.Installgoogle-authenticator-libpam(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg google-authenticator-libpamFormula) Installgoogle-authenticator-libpam() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "1.10.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "1.10.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}