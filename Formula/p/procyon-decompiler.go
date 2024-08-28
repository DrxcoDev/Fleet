
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// procyon-decompilerFormula represents a formula in Go.
type procyon-decompilerFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg procyon-decompilerFormula) Print() {
    fmt.Printf("Name: procyon-decompiler\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := procyon-decompilerFormula{
        Description:  "Modern decompiler for Java 5 and beyond",
        Homepage:     "https://github.com/mstrobel/procyon",
        URL:          "https://github.com/mstrobel/procyon/releases/download/v0.6.0/procyon-decompiler-0.6.0.jar",
        Sha256:       "6d383fdfc3e0cd6e98271bc33299c2553943305e40c4cc48b127ab4157dff604",
        Dependencies: []string{"openjdk@21"},
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

    fmt.Printf("Installing formula: %s\n", "procyon-decompiler")
    if err := pkg.Installprocyon-decompiler(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg procyon-decompilerFormula) Installprocyon-decompiler() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "procyon-decompiler-0.6.0.jar"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "procyon-decompiler-0.6.0"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}