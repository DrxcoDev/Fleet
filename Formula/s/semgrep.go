
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// semgrepFormula represents a formula in Go.
type semgrepFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg semgrepFormula) Print() {
    fmt.Printf("Name: semgrep\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := semgrepFormula{
        Description:  "Easily detect and prevent bugs and anti-patterns in your codebase",
        Homepage:     "https://semgrep.dev",
        URL:          "https://files.pythonhosted.org/packages/0e/af/9f2de5bd32549a1b705af7a7c054af3878816a1267cb389c03cc4f342a51/zipp-3.20.0.tar.gz",
        Sha256:       "0145e43d89664cfe1a2e533adc75adafed82fe2da404b4bbb6b026c0157bdb31",
        Dependencies: []string{"autoconf", "cmake", "coreutils", "dune", "ocaml", "opam", "pipenv", "pkg-config", "rust", "certifi", "gmp", "libev", "pcre", "pcre2", "python@3.12", "sqlite", "tree-sitter"},
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

    fmt.Printf("Installing formula: %s\n", "semgrep")
    if err := pkg.Installsemgrep(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg semgrepFormula) Installsemgrep() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "zipp-3.20.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "zipp-3.20.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
