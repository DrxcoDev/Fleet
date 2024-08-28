
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// mikutterFormula represents a formula in Go.
type mikutterFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg mikutterFormula) Print() {
    fmt.Printf("Name: mikutter\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := mikutterFormula{
        Description:  "Extensible Twitter client",
        Homepage:     "https://mikutter.hachune.net/",
        URL:          "https://rubygems.org/downloads/typed-array-0.1.2.gem",
        Sha256:       "891fa1de2cdccad5f9e03936569c3c15d413d8c6658e2edfe439d9386d169b62",
        Dependencies: []string{"at-spi2-core", "cairo", "fontconfig", "freetype", "gdk-pixbuf", "glib", "gobject-introspection", "gtk+3", "harfbuzz", "pango", "ruby", "gettext", "terminal-notifier"},
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

    fmt.Printf("Installing formula: %s\n", "mikutter")
    if err := pkg.Installmikutter(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg mikutterFormula) Installmikutter() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "typed-array-0.1.2.gem"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "typed-array-0.1.2"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
