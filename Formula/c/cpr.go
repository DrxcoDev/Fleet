package main

import (
	"fmt"
	"log"
	"os/exec"
)

// cprFormula represents a formula in Go.
type cprFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg cprFormula) Print() {
	fmt.Printf("Name: cpr\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := cprFormula{
		Description:  "C++ Requests, a spiritual port of Python Requests",
		Homepage:     "https://docs.libcpr.org/",
		URL:          "https://github.com/libcpr/cpr/archive/refs/tags/1.10.5.tar.gz",
		Sha256:       "c5238ed4fc1a27b5b501e68ba01255e01fadd98398e67c709ee0c72e58c144aa",
		Dependencies: []string{"cmake", "openssl@3"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installcpr(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg cprFormula) Installcpr() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "1.10.5.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "1.10.5.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
