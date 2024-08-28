
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libpeas@1Formula represents a formula in Go.
type libpeas@1Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libpeas@1Formula) Print() {
    fmt.Printf("Name: libpeas@1\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libpeas@1Formula{
        Description:  "GObject plugin library",
        Homepage:     "https://wiki.gnome.org/Projects/Libpeas",
        URL:          "https://download.gnome.org/sources/libpeas/1.36/libpeas-1.36.0.tar.xz",
        Sha256:       "b44e4a9f261b68e1ae0f92a08a962d53b7b56f7c90acdc664b062a87cc2d541c",
        Dependencies: []string{"meson", "ninja", "pkg-config", "vala", "glib", "gobject-introspection", "gtk+3", "pygobject3", "python@3.12", "gettext"},
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

    fmt.Printf("Installing formula: %s\n", "libpeas@1")
    if err := pkg.Installlibpeas@1(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libpeas@1Formula) Installlibpeas@1() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libpeas-1.36.0.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libpeas-1.36.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}