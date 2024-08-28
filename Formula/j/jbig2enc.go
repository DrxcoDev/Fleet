
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// jbig2encFormula represents a formula in Go.
type jbig2encFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg jbig2encFormula) Print() {
    fmt.Printf("Name: jbig2enc\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := jbig2encFormula{
        Description:  "JBIG2 encoder (for monochrome documents)",
        Homepage:     "https://github.com/agl/jbig2enc",
        URL:          "https://github.com/agl/jbig2enc/commit/d211d8c9c65fbc103594580484a3b7fa0249e160.patch?full_index=1",
        Sha256:       "a1e7b44b9ea28d32d034718fb10022961dcec32b74beda56575f84416081bd43",
        Dependencies: []string{"autoconf", "automake", "libtool", "leptonica", "giflib", "jpeg-turbo", "libpng", "libtiff", "webp"},
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

    fmt.Printf("Installing formula: %s\n", "jbig2enc")
    if err := pkg.Installjbig2enc(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg jbig2encFormula) Installjbig2enc() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "d211d8c9c65fbc103594580484a3b7fa0249e160.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "d211d8c9c65fbc103594580484a3b7fa0249e160"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
