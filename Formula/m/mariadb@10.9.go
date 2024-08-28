
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// mariadb@10.9Formula represents a formula in Go.
type mariadb@10.9Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg mariadb@10.9Formula) Print() {
    fmt.Printf("Name: mariadb@10.9\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := mariadb@10.9Formula{
        Description:  "Drop-in replacement for MySQL",
        Homepage:     "https://mariadb.org/",
        URL:          "https://github.com/MariaDB/server/commit/cd5808eb8da13c5626d4bdeb452cef6ada29cb1d.patch?full_index=1",
        Sha256:       "4d288f82f56c61278aefecba8a90d214810b754e234f40b338e8cc809e0369e9",
        Dependencies: []string{"bison", "cmake", "fmt", "pkg-config", "groonga", "openssl@3", "pcre2", "zstd", "linux-pam", "readline"},
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

    fmt.Printf("Installing formula: %s\n", "mariadb@10.9")
    if err := pkg.Installmariadb@10.9(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg mariadb@10.9Formula) Installmariadb@10.9() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "cd5808eb8da13c5626d4bdeb452cef6ada29cb1d.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "cd5808eb8da13c5626d4bdeb452cef6ada29cb1d"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}