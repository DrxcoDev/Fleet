
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// librsvgFormula represents a formula in Go.
type librsvgFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg librsvgFormula) Print() {
    fmt.Printf("Name: librsvg\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := librsvgFormula{
        Description:  "Library to render SVG files using Cairo",
        Homepage:     "https://wiki.gnome.org/Projects/LibRsvg",
        URL:          "https://download.gnome.org/sources/librsvg/2.58/librsvg-2.58.3.tar.xz",
        Sha256:       "1692a55a60d93415e41ab9fcab6597fa54080ab9ac61d7fd63ab6dab338f97bf",
        Dependencies: []string{"gobject-introspection", "pkg-config", "rust", "cairo", "gdk-pixbuf", "glib", "pango", "fontconfig", "freetype", "gettext", "harfbuzz", "libpng"},
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

    fmt.Printf("Installing formula: %s\n", "librsvg")
    if err := pkg.Installlibrsvg(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg librsvgFormula) Installlibrsvg() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "librsvg-2.58.3.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "librsvg-2.58.3.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}