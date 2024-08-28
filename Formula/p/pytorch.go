
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// pytorchFormula represents a formula in Go.
type pytorchFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg pytorchFormula) Print() {
    fmt.Printf("Name: pytorch\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := pytorchFormula{
        Description:  "Tensors and dynamic neural networks",
        Homepage:     "https://pytorch.org/",
        URL:          "https://github.com/pytorch/pytorch/commit/38d9bb5abcc31ba97927a5399b88afe2cf60bf64.patch?full_index=1",
        Sha256:       "c9bf84d154e5f3f9b67a68d25765f32a08b9deb3127254971ed6351231eba228",
        Dependencies: []string{"cmake", "ninja", "python@3.12", "abseil", "eigen", "libuv", "libyaml", "numpy", "openblas", "protobuf", "pybind11", "sleef", "libomp"},
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

    fmt.Printf("Installing formula: %s\n", "pytorch")
    if err := pkg.Installpytorch(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg pytorchFormula) Installpytorch() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "38d9bb5abcc31ba97927a5399b88afe2cf60bf64.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "38d9bb5abcc31ba97927a5399b88afe2cf60bf64"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
