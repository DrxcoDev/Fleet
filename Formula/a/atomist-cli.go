
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// atomist-cliFormula represents a formula in Go.
type atomist-cliFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg atomist-cliFormula) Print() {
    fmt.Printf("Name: atomist-cli\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := atomist-cliFormula{
        Description:  "Unified command-line tool for interacting with Atomist services",
        Homepage:     "https://github.com/atomist/cli",
        URL:          "https://registry.npmjs.org/@atomist/cli/-/cli-1.8.0.tgz",
        Sha256:       "a8a420e052281702ab92e1b510422b86556e366f8c6b149fbae2ef3864466708",
        Dependencies: []string{"node", "macos-term-size"},
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

    fmt.Printf("Installing formula: %s\n", "atomist-cli")
    if err := pkg.Installatomist-cli(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg atomist-cliFormula) Installatomist-cli() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "cli-1.8.0.tgz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "cli-1.8.0"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
