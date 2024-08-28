
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gtranslatorFormula represents a formula in Go.
type gtranslatorFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gtranslatorFormula) Print() {
    fmt.Printf("Name: gtranslator\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gtranslatorFormula{
        Description:  "GNOME gettext PO file editor",
        Homepage:     "https://wiki.gnome.org/Design/Apps/Translator",
        URL:          "https://download.gnome.org/sources/gtranslator/46/gtranslator-46.1.tar.xz",
        Sha256:       "c41a8150fa0f64e40c47a0c00f121a8005656d9b678671faac04c6389018234e",
        Dependencies: []string{"desktop-file-utils", "itstool", "meson", "ninja", "pkg-config", "adwaita-icon-theme", "cairo", "gettext", "glib", "gspell", "gtk4", "gtksourceview5", "json-glib", "libadwaita", "libgda", "libsoup", "libspelling", "pango"},
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

    fmt.Printf("Installing formula: %s\n", "gtranslator")
    if err := pkg.Installgtranslator(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gtranslatorFormula) Installgtranslator() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gtranslator-46.1.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gtranslator-46.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
