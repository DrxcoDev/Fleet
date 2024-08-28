
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gtksourceviewmm3Formula represents a formula in Go.
type gtksourceviewmm3Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gtksourceviewmm3Formula) Print() {
    fmt.Printf("Name: gtksourceviewmm3\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gtksourceviewmm3Formula{
        Description:  "C++ bindings for gtksourceview3",
        Homepage:     "https://gitlab.gnome.org/GNOME/gtksourceviewmm",
        URL:          "https://download.gnome.org/sources/gtksourceviewmm/3.18/gtksourceviewmm-3.18.0.tar.xz",
        Sha256:       "b6489e67344895dad8ef93a694a47c49fec20dc0735989355b11489cd85fbccd",
        Dependencies: []string{"pkg-config", "atkmm@2.28", "cairomm@1.14", "glib", "glibmm@2.66", "gtk+3", "gtkmm3", "gtksourceview3", "libsigc++@2", "pangomm@2.46", "at-spi2-core", "cairo", "gdk-pixbuf", "gettext", "harfbuzz", "pango"},
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

    fmt.Printf("Installing formula: %s\n", "gtksourceviewmm3")
    if err := pkg.Installgtksourceviewmm3(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gtksourceviewmm3Formula) Installgtksourceviewmm3() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gtksourceviewmm-3.18.0.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gtksourceviewmm-3.18.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
