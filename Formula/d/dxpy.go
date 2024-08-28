package main

import (
	"fmt"
	"log"
	"os/exec"
)

// dxpyFormula represents a formula in Go.
type dxpyFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg dxpyFormula) Print() {
	fmt.Printf("Name: dxpy\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := dxpyFormula{
		Description:  "DNAnexus toolkit utilities and platform API bindings for Python",
		Homepage:     "https://github.com/dnanexus/dx-toolkit",
		URL:          "https://files.pythonhosted.org/packages/20/07/2a94288afc0f6c9434d6709c5320ee21eaedb2f463ede25ed9cf6feff330/websocket-client-1.7.0.tar.gz",
		Sha256:       "10e511ea3a8c744631d3bd77e61eb17ed09304c413ad42cf6ddfa4c7787e8fe6",
		Dependencies: []string{"certifi", "cryptography", "python@3.12", "readline"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installdxpy(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg dxpyFormula) Installdxpy() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "websocket-client-1.7.0.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "websocket-client-1.7.0.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
