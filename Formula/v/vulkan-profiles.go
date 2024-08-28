
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// vulkan-profilesFormula represents a formula in Go.
type vulkan-profilesFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg vulkan-profilesFormula) Print() {
    fmt.Printf("Name: vulkan-profiles\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := vulkan-profilesFormula{
        Description:  "Tools for Vulkan profiles",
        Homepage:     "https://github.com/KhronosGroup/Vulkan-Profiles",
        URL:          "https://github.com/KhronosGroup/Vulkan-Profiles/archive/refs/tags/v1.3.294.tar.gz",
        Sha256:       "ea897d3797bef8d85d4a760ddf1f25c28b9d2e6806a650ee1391664d37ee835d",
        Dependencies: []string{"cmake", "pkg-config", "python@3.12", "vulkan-tools", "jsoncpp", "valijson", "vulkan-headers", "vulkan-loader", "vulkan-utility-libraries", "molten-vk", "mesa"},
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

    fmt.Printf("Installing formula: %s\n", "vulkan-profiles")
    if err := pkg.Installvulkan-profiles(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg vulkan-profilesFormula) Installvulkan-profiles() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v1.3.294.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v1.3.294.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}