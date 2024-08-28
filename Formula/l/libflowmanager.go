
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libflowmanagerFormula represents a formula in Go.
type libflowmanagerFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libflowmanagerFormula) Print() {
    fmt.Printf("Name: libflowmanager\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libflowmanagerFormula{
        Description:  "Flow-based measurement tasks with packet-based inputs",
        Homepage:     "https://github.com/LibtraceTeam/libflowmanager",
        URL:          "https://github.com/LibtraceTeam/libflowmanager/commit/a60a04a3b4a12faf48854b34908f9db0c4f080b0.patch?full_index=1",
        Sha256:       "15d93f863374eff428c69e6e1733bdc861c831714f8d7d7c1323ebf1b9ba9a4c",
        Dependencies: []string{"autoconf", "automake", "libtool", "libtrace"},
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

    fmt.Printf("Installing formula: %s\n", "libflowmanager")
    if err := pkg.Installlibflowmanager(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libflowmanagerFormula) Installlibflowmanager() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "a60a04a3b4a12faf48854b34908f9db0c4f080b0.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "a60a04a3b4a12faf48854b34908f9db0c4f080b0"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}