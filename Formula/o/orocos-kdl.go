
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// orocos-kdlFormula represents a formula in Go.
type orocos-kdlFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg orocos-kdlFormula) Print() {
    fmt.Printf("Name: orocos-kdl\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := orocos-kdlFormula{
        Description:  "Orocos Kinematics and Dynamics C++ library",
        Homepage:     "https://orocos.org/",
        URL:          "https://github.com/orocos/orocos_kinematics_dynamics/commit/ef39a4fd5cfb1400b2e6e034b1a99b8ad91192cf.patch?full_index=1",
        Sha256:       "b2ac2ff5d5d3285e7dfb4fbfc81364b1abc808cdd7d22415e446bfbdca189edd",
        Dependencies: []string{"cmake", "eigen"},
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

    fmt.Printf("Installing formula: %s\n", "orocos-kdl")
    if err := pkg.Installorocos-kdl(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg orocos-kdlFormula) Installorocos-kdl() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "ef39a4fd5cfb1400b2e6e034b1a99b8ad91192cf.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "ef39a4fd5cfb1400b2e6e034b1a99b8ad91192cf"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}