package main

import (
	"fmt"
	"log"
	"os/exec"
)

// cbmcFormula represents a formula in Go.
type cbmcFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg cbmcFormula) Print() {
	fmt.Printf("Name: cbmc\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := cbmcFormula{
		Description:  "C Bounded Model Checker",
		Homepage:     "https://www.cprover.org/cbmc/",
		URL:          "https://github.com/diffblue/cbmc.git",
		Sha256:       "f0a0d46330bb1eb42845215cd701359603fb06d881e936b70df190c1d70e5f15",
		Dependencies: []string{"cmake", "maven", "openjdk@21", "rust"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installcbmc(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg cbmcFormula) Installcbmc() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "cbmc.git"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "cbmc"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
