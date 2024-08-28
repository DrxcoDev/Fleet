
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// python@3.10Formula represents a formula in Go.
type python@3.10Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg python@3.10Formula) Print() {
    fmt.Printf("Name: python@3.10\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := python@3.10Formula{
        Description:  "Interpreted, interactive, object-oriented programming language",
        Homepage:     "https://www.python.org/",
        URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/a1618a5005d0b01d63b720321806820a03432f1a/python/3.10-distutils-scheme.diff",
        Sha256:       "d1a29b3c9ecf8aecd65e1e54efc42fb1422b2f5d05cba0c747178f4ef8a69683",
        Dependencies: []string{"pkg-config", "gdbm", "mpdecimal", "openssl@3", "readline", "sqlite", "xz", "libnsl"},
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

    fmt.Printf("Installing formula: %s\n", "python@3.10")
    if err := pkg.Installpython@3.10(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg python@3.10Formula) Installpython@3.10() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "3.10-distutils-scheme.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "3.10-distutils-scheme"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}