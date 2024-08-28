
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// glibmm@2.66Formula represents a formula in Go.
type glibmm@2.66Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg glibmm@2.66Formula) Print() {
    fmt.Printf("Name: glibmm@2.66\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := glibmm@2.66Formula{
        Description:  "C++ interface to glib",
        Homepage:     "https://www.gtkmm.org/",
        URL:          "https://download.gnome.org/sources/glibmm/2.66/",
        Sha256:       "e73a3eb58812790f1057663483cb3047cb1c33f80497faa11393ddf8e411c20b",
        Dependencies: []string{"meson", "ninja", "pkg-config", "glib", "libsigc++@2"},
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

    fmt.Printf("Installing formula: %s\n", "glibmm@2.66")
    if err := pkg.Installglibmm@2.66(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg glibmm@2.66Formula) Installglibmm@2.66() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := ""
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := ""
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
