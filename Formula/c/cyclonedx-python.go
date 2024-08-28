
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// cyclonedx-pythonFormula represents a formula in Go.
type cyclonedx-pythonFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg cyclonedx-pythonFormula) Print() {
    fmt.Printf("Name: cyclonedx-python\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := cyclonedx-pythonFormula{
        Description:  "Creates CycloneDX Software Bill of Materials (SBOM) from Python projects",
        Homepage:     "https://cyclonedx.org/",
        URL:          "https://files.pythonhosted.org/packages/b3/bf/cfe993a8acab0976a08cfa1a0a23cf9ce212b8c52cca40fbcca6e994acea/webcolors-24.6.0.tar.gz",
        Sha256:       "1d160d1de46b3e81e58d0a280d0c78b467dc80f47294b91b1ad8029d2cedb55b",
        Dependencies: []string{"rust", "python@3.12"},
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

    fmt.Printf("Installing formula: %s\n", "cyclonedx-python")
    if err := pkg.Installcyclonedx-python(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg cyclonedx-pythonFormula) Installcyclonedx-python() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "webcolors-24.6.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "webcolors-24.6.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}