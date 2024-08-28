
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// aarch64-elf-gdbFormula represents a formula in Go.
type aarch64-elf-gdbFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg aarch64-elf-gdbFormula) Print() {
    fmt.Printf("Name: aarch64-elf-gdb\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := aarch64-elf-gdbFormula{
        Description:  "GNU debugger for aarch64-elf cross development",
        Homepage:     "https://www.gnu.org/software/gdb/",
        URL:          "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz",
        Sha256:       "393ec8e3717b1abf6fd13cc593b076697c41f6d3fe5f173051116ed119f19538",
        Dependencies: []string{"pkg-config", "aarch64-elf-gcc", "gmp", "mpfr", "python@3.12", "readline", "xz", "zstd", "texinfo"},
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

    fmt.Printf("Installing formula: %s\n", "aarch64-elf-gdb")
    if err := pkg.Installaarch64-elf-gdb(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg aarch64-elf-gdbFormula) Installaarch64-elf-gdb() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gdb-15.1.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gdb-15.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}