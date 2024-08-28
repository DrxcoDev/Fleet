
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// vulkan-extensionlayerFormula represents a formula in Go.
type vulkan-extensionlayerFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg vulkan-extensionlayerFormula) Print() {
    fmt.Printf("Name: vulkan-extensionlayer\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := vulkan-extensionlayerFormula{
        Description:  "Layer providing Vulkan features when native support is unavailable",
        Homepage:     "https://github.com/KhronosGroup/Vulkan-ExtensionLayer",
        URL:          "https://github.com/KhronosGroup/Vulkan-ExtensionLayer/archive/refs/tags/v1.3.294.tar.gz",
        Sha256:       "981601afb0d386993abe15f7b7bf5783a7404f24f1e2d71640f0fdf6daa7234d",
        Dependencies: []string{"cmake", "googletest", "python@3.12", "vulkan-loader", "vulkan-tools", "glslang", "spirv-headers", "spirv-tools", "vulkan-headers", "vulkan-utility-libraries", "libxcb", "libxrandr", "mesa", "pkg-config", "wayland"},
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

    fmt.Printf("Installing formula: %s\n", "vulkan-extensionlayer")
    if err := pkg.Installvulkan-extensionlayer(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg vulkan-extensionlayerFormula) Installvulkan-extensionlayer() error {
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
