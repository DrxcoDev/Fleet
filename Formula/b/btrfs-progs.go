
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// btrfs-progsFormula represents a formula in Go.
type btrfs-progsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg btrfs-progsFormula) Print() {
    fmt.Printf("Name: btrfs-progs\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := btrfs-progsFormula{
        Description:  "Userspace utilities to manage btrfs filesystems",
        Homepage:     "https://btrfs.readthedocs.io/en/latest/",
        URL:          "https://mirrors.edge.kernel.org/pub/linux/kernel/people/kdave/btrfs-progs/",
        Sha256:       "daa071b3db25e8a8f23e945383d81837773ef2a53c5fcdbf3845f0a036b4e56d",
        Dependencies: []string{"pkg-config", "python-setuptools", "python@3.12", "sphinx-doc", "e2fsprogs", "lzo", "systemd", "util-linux", "zlib", "zstd"},
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

    fmt.Printf("Installing formula: %s\n", "btrfs-progs")
    if err := pkg.Installbtrfs-progs(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg btrfs-progsFormula) Installbtrfs-progs() error {
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
