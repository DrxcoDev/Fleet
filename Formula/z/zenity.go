
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// zenityFormula represents a formula in Go.
type zenityFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg zenityFormula) Print() {
    fmt.Printf("Name: zenity\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := zenityFormula{
        Description:  "GTK+ dialog boxes for the command-line",
        Homepage:     "https://wiki.gnome.org/Projects/Zenity",
        URL:          "https://download.gnome.org/sources/zenity/4.0/zenity-4.0.2.tar.xz",
        Sha256:       "3b1e00581cd9fb4b0fa0506018c864ba34cb5a4a5db6fd67dcd06b60c7ea88d0",
        Dependencies: []string{"gettext", "help2man", "itstool", "meson", "ninja", "pkg-config", "glib", "gtk4", "libadwaita", "pango", "gettext"},
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

    fmt.Printf("Installing formula: %s\n", "zenity")
    if err := pkg.Installzenity(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg zenityFormula) Installzenity() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "zenity-4.0.2.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "zenity-4.0.2.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}