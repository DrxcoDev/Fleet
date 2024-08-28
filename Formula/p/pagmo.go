package main

import (
	"fmt"
	"log"
	"os/exec"
)

// pagmoFormula represents a formula in Go.
type pagmoFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg pagmoFormula) Print() {
	fmt.Printf("Name: pagmo\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := pagmoFormula{
		Description:  "Scientific library for massively parallel optimization",
		Homepage:     "https://esa.github.io/pagmo2/",
		URL:          "https://github.com/esa/pagmo2/archive/refs/tags/v2.19.1.tar.gz",
		Sha256:       "033a66f14e5276aa4b64bddc1114539bc2cc99eae1ce4ffd1e84a01413aaff88",
		Dependencies: []string{"cmake", "boost", "eigen", "nlopt", "tbb"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installpagmo(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg pagmoFormula) Installpagmo() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "v2.19.1.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "v2.19.1.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
