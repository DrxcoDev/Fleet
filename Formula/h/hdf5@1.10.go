
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// hdf5@1.10Formula represents a formula in Go.
type hdf5@1.10Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg hdf5@1.10Formula) Print() {
    fmt.Printf("Name: hdf5@1.10\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := hdf5@1.10Formula{
        Description:  "File format designed to store large amounts of data",
        Homepage:     "https://www.hdfgroup.org/HDF5",
        URL:          "https://support.hdfgroup.org/ftp/HDF5/releases/hdf5-1.10/hdf5-1.10.11/src/hdf5-1.10.11.tar.bz2",
        Sha256:       "706516f685fc2d8a3d2c1e6565b36cbc885445f2c0f901a2bd735f3229151590",
        Dependencies: []string{"autoconf", "automake", "libtool", "gcc", "libaec"},
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

    fmt.Printf("Installing formula: %s\n", "hdf5@1.10")
    if err := pkg.Installhdf5@1.10(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg hdf5@1.10Formula) Installhdf5@1.10() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "hdf5-1.10.11.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "hdf5-1.10.11.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}