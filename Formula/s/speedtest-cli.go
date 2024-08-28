
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// speedtest-cliFormula represents a formula in Go.
type speedtest-cliFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg speedtest-cliFormula) Print() {
    fmt.Printf("Name: speedtest-cli\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := speedtest-cliFormula{
        Description:  "Command-line interface for https://speedtest.net bandwidth tests",
        Homepage:     "https://github.com/sivel/speedtest-cli",
        URL:          "https://github.com/sivel/speedtest-cli/commit/22210ca35228f0bbcef75a7c14587c4ecb875ab4.patch?full_index=1",
        Sha256:       "d0456eb9fded20fb1580dbc6e3bc451a10c3fbcd3441efea66035aa848440c09",
        Dependencies: []string{"python@3.12"},
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

    fmt.Printf("Installing formula: %s\n", "speedtest-cli")
    if err := pkg.Installspeedtest-cli(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg speedtest-cliFormula) Installspeedtest-cli() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "22210ca35228f0bbcef75a7c14587c4ecb875ab4.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "22210ca35228f0bbcef75a7c14587c4ecb875ab4"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}