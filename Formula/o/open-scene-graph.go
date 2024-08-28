
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// open-scene-graphFormula represents a formula in Go.
type open-scene-graphFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg open-scene-graphFormula) Print() {
    fmt.Printf("Name: open-scene-graph\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := open-scene-graphFormula{
        Description:  "3D graphics toolkit",
        Homepage:     "https://github.com/openscenegraph/OpenSceneGraph",
        URL:          "https://github.com/openscenegraph/OpenSceneGraph/commit/21f5a0adfb57dc4c28b696e93beface45de28194.patch?full_index=1",
        Sha256:       "43c4367454e8de65443937a3509f96d4d273b50431b0a4fde16607c88183b247",
        Dependencies: []string{"cmake", "doxygen", "graphviz", "pkg-config", "fontconfig", "freetype", "jpeg-turbo", "sdl2", "cairo", "giflib", "glib", "libpng", "librsvg", "libx11", "libxinerama", "libxrandr", "mesa"},
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

    fmt.Printf("Installing formula: %s\n", "open-scene-graph")
    if err := pkg.Installopen-scene-graph(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg open-scene-graphFormula) Installopen-scene-graph() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "21f5a0adfb57dc4c28b696e93beface45de28194.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "21f5a0adfb57dc4c28b696e93beface45de28194"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}