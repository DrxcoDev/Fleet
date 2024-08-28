
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// spidermonkey@91Formula represents a formula in Go.
type spidermonkey@91Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg spidermonkey@91Formula) Print() {
    fmt.Printf("Name: spidermonkey@91\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := spidermonkey@91Formula{
        Description:  "JavaScript-C Engine",
        Homepage:     "https://spidermonkey.dev",
        URL:          "https://archive.mozilla.org/pub/firefox/releases/91.13.0esr/source/firefox-91.13.0esr.source.tar.xz",
        Sha256:       "ae8f67dc53fa55e1d06ee912bd04ee344c6d742210c401bc1dfae9410d198eaf",
        Dependencies: []string{"autoconf@2.13", "pkg-config", "python@3.9", "rust", "icu4c", "nspr", "readline"},
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

    fmt.Printf("Installing formula: %s\n", "spidermonkey@91")
    if err := pkg.Installspidermonkey@91(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg spidermonkey@91Formula) Installspidermonkey@91() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "firefox-91.13.0esr.source.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "firefox-91.13.0esr.source.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}