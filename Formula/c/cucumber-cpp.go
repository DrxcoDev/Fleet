
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// cucumber-cppFormula represents a formula in Go.
type cucumber-cppFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg cucumber-cppFormula) Print() {
    fmt.Printf("Name: cucumber-cpp\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := cucumber-cppFormula{
        Description:  "Support for writing Cucumber step definitions in C++",
        Homepage:     "https://cucumber.io",
        URL:          "https://github.com/cucumber/cucumber-cpp.git",
        Sha256:       "3c9e95d394b01d3a55aee2d7b2add5de066e47843a7c98f58a9a30c58346fff0",
        Dependencies: []string{"cmake", "nlohmann-json", "qt@5", "ruby", "asio", "tclap"},
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

    fmt.Printf("Installing formula: %s\n", "cucumber-cpp")
    if err := pkg.Installcucumber-cpp(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg cucumber-cppFormula) Installcucumber-cpp() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "cucumber-cpp.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "cucumber-cpp"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}