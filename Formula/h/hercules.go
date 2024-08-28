
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// herculesFormula represents a formula in Go.
type herculesFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg herculesFormula) Print() {
    fmt.Printf("Name: hercules\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := herculesFormula{
        Description:  "System/370, ESA/390 and z/Architecture Emulator",
        Homepage:     "https://sdl-hercules-390.github.io/html/",
        URL:          "https://github.com/SDL-Hercules-390/telnet/archive/729f0b688c1426018112c1e509f207fb5f266efa.tar.gz",
        Sha256:       "222bc9c5b56056b3fa4afdf4dd78ab1c87673c26c725309b1b3a6fd3e0e88d51",
        Dependencies: []string{"autoconf", "automake", "cmake", "gnu-sed", "libtool"},
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

    fmt.Printf("Installing formula: %s\n", "hercules")
    if err := pkg.Installhercules(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg herculesFormula) Installhercules() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "729f0b688c1426018112c1e509f207fb5f266efa.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "729f0b688c1426018112c1e509f207fb5f266efa.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
