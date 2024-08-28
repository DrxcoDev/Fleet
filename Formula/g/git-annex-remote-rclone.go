
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// git-annex-remote-rcloneFormula represents a formula in Go.
type git-annex-remote-rcloneFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg git-annex-remote-rcloneFormula) Print() {
    fmt.Printf("Name: git-annex-remote-rclone\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := git-annex-remote-rcloneFormula{
        Description:  "Use rclone supported cloud storage with git-annex",
        Homepage:     "https://github.com/git-annex-remote-rclone/git-annex-remote-rclone",
        URL:          "https://github.com/git-annex-remote-rclone/git-annex-remote-rclone/archive/refs/tags/v0.8.tar.gz",
        Sha256:       "60ec135de845b97d8eafb3de716f93fda72c9c69a99c58e6b9669eec71006cfc",
        Dependencies: []string{"git-annex", "rclone"},
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

    fmt.Printf("Installing formula: %s\n", "git-annex-remote-rclone")
    if err := pkg.Installgit-annex-remote-rclone(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg git-annex-remote-rcloneFormula) Installgit-annex-remote-rclone() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v0.8.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v0.8.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
