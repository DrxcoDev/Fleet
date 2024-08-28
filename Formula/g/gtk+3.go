
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gtk+3Formula represents a formula in Go.
type gtk+3Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gtk+3Formula) Print() {
    fmt.Printf("Name: gtk+3\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gtk+3Formula{
        Description:  "Toolkit for creating graphical user interfaces",
        Homepage:     "https://gtk.org/",
        URL:          "https://download.gnome.org/sources/gtk+/3.24/gtk+-3.24.43.tar.xz",
        Sha256:       "880cff997f0d940867b79dee4f87f22c8fa03adc377c2f793f85c72818a8cad6",
        Dependencies: []string{"docbook", "docbook-xsl", "gettext", "gobject-introspection", "meson", "ninja", "pkg-config", "at-spi2-core", "cairo", "fribidi", "gdk-pixbuf", "glib", "gsettings-desktop-schemas", "harfbuzz", "hicolor-icon-theme", "libepoxy", "pango", "gettext", "cmake", "fontconfig", "iso-codes", "libx11", "libxdamage", "libxext", "libxfixes", "libxi", "libxinerama", "libxkbcommon", "libxrandr", "wayland", "wayland-protocols", "xorgproto"},
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

    fmt.Printf("Installing formula: %s\n", "gtk+3")
    if err := pkg.Installgtk+3(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gtk+3Formula) Installgtk+3() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gtk+-3.24.43.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gtk+-3.24.43.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
