
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// seven-kingdomsFormula represents a formula in Go.
type seven-kingdomsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg seven-kingdomsFormula) Print() {
    fmt.Printf("Name: seven-kingdoms\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := seven-kingdomsFormula{
        Description:  "Real-time strategy game developed by Trevor Chan of Enlight Software",
        Homepage:     "https://7kfans.com",
        URL:          "https://downloads.sourceforge.net/project/skfans/7KAA%202.15.6/7kaa-2.15.6.tar.gz",
        Sha256:       "b3ba3202ca789169da6b0eb0173e686b97866e37ac12f0f10429e747c4ce0751",
        Dependencies: []string{"pkg-config", "enet", "gettext", "sdl2", "gcc", "openal-soft"},
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

    fmt.Printf("Installing formula: %s\n", "seven-kingdoms")
    if err := pkg.Installseven-kingdoms(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg seven-kingdomsFormula) Installseven-kingdoms() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "7kaa-2.15.6.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "7kaa-2.15.6.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}