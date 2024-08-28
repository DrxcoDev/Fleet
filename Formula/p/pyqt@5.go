
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// pyqt@5Formula represents a formula in Go.
type pyqt@5Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg pyqt@5Formula) Print() {
    fmt.Printf("Name: pyqt@5\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := pyqt@5Formula{
        Description:  "Python bindings for v5 of Qt",
        Homepage:     "https://www.riverbankcomputing.com/software/pyqt/intro",
        URL:          "https://files.pythonhosted.org/packages/cf/4b/ca01d875eff114ba5221ce9311912fbbc142b7bb4cbc4435e04f4f1f73cb/PyQtWebEngine-5.15.6.tar.gz",
        Sha256:       "ae241ef2a61c782939c58b52c2aea53ad99b30f3934c8358d5e0a6ebb3fd0721",
        Dependencies: []string{"pyqt-builder", "python@3.12", "qt@5"},
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

    fmt.Printf("Installing formula: %s\n", "pyqt@5")
    if err := pkg.Installpyqt@5(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg pyqt@5Formula) Installpyqt@5() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "PyQtWebEngine-5.15.6.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "PyQtWebEngine-5.15.6.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}