
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// vowpal-wabbitFormula represents a formula in Go.
type vowpal-wabbitFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg vowpal-wabbitFormula) Print() {
    fmt.Printf("Name: vowpal-wabbit\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := vowpal-wabbitFormula{
        Description:  "Online learning algorithm",
        Homepage:     "https://github.com/VowpalWabbit/vowpal_wabbit",
        URL:          "https://github.com/VowpalWabbit/vowpal_wabbit/archive/refs/tags/9.10.0.tar.gz",
        Sha256:       "24bc424f2e333c4995e596f13ce7b4bda399467dc711098d4fb399fc954bf6bf",
        Dependencies: []string{"cmake", "rapidjson", "spdlog", "boost", "eigen", "fmt", "sse2neon"},
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

    fmt.Printf("Installing formula: %s\n", "vowpal-wabbit")
    if err := pkg.Installvowpal-wabbit(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg vowpal-wabbitFormula) Installvowpal-wabbit() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "9.10.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "9.10.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}