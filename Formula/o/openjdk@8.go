
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// openjdk@8Formula represents a formula in Go.
type openjdk@8Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg openjdk@8Formula) Print() {
    fmt.Printf("Name: openjdk@8\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := openjdk@8Formula{
        Description:  "Development kit for the Java programming language",
        Homepage:     "https://openjdk.java.net/",
        URL:          "https://raw.githubusercontent.com/macports/macports-ports/04ad4a17332e391cd359271965d4c6dac87a7eb2/java/openjdk8/files/0006-Disable-C-11-warnings.patch",
        Sha256:       "127d9508b72005e849a6ada6adf04bd49a236731d769810e67793bdf2aa722fe",
        Dependencies: []string{"autoconf", "pkg-config", "freetype", "giflib", "gawk", "alsa-lib", "fontconfig", "libx11", "libxext", "libxi", "libxrandr", "libxrender", "libxt", "libxtst"},
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

    fmt.Printf("Installing formula: %s\n", "openjdk@8")
    if err := pkg.Installopenjdk@8(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg openjdk@8Formula) Installopenjdk@8() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "0006-Disable-C-11-warnings.patch"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "0006-Disable-C-11-warnings"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}