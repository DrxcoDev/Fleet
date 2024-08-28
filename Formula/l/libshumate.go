
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libshumateFormula represents a formula in Go.
type libshumateFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libshumateFormula) Print() {
    fmt.Printf("Name: libshumate\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libshumateFormula{
        Description:  "Shumate is a GTK toolkit providing widgets for embedded maps",
        Homepage:     "https://gitlab.gnome.org/GNOME/libshumate",
        URL:          "https://download.gnome.org/sources/libshumate/1.2/libshumate-1.2.3.tar.xz",
        Sha256:       "ae4fbb8002a69796f35a85f5733689d78fd5b807a76cb30f1981f7689cf52c33",
        Dependencies: []string{"gettext", "gobject-introspection", "meson", "ninja", "pkg-config", "vala", "cairo", "cmake", "gdk-pixbuf", "gi-docgen", "glib", "graphene", "gtk4", "json-glib", "libsoup", "pango", "protobuf-c", "sqlite", "gettext"},
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

    fmt.Printf("Installing formula: %s\n", "libshumate")
    if err := pkg.Installlibshumate(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libshumateFormula) Installlibshumate() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libshumate-1.2.3.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libshumate-1.2.3.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}