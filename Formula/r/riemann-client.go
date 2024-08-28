
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// riemann-clientFormula represents a formula in Go.
type riemann-clientFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg riemann-clientFormula) Print() {
    fmt.Printf("Name: riemann-client\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := riemann-clientFormula{
        Description:  "C client library for the Riemann monitoring system",
        Homepage:     "https://git.madhouse-project.org/algernon/riemann-c-client",
        URL:          "https://git.madhouse-project.org/algernon/riemann-c-client/archive/riemann-c-client-2.2.2.tar.gz",
        Sha256:       "fc7736635610401fa99e4911155329f537c997b57d218821538f0e6b8b212111",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "json-c", "openssl@3", "protobuf-c"},
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

    fmt.Printf("Installing formula: %s\n", "riemann-client")
    if err := pkg.Installriemann-client(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg riemann-clientFormula) Installriemann-client() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "riemann-c-client-2.2.2.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "riemann-c-client-2.2.2.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
