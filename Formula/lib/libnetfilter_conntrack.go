package main

import (
	"fmt"
	"log"
	"os/exec"
)

// libnetfilter_conntrackFormula represents a formula in Go.
type libnetfilter_conntrackFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg libnetfilter_conntrackFormula) Print() {
	fmt.Printf("Name: libnetfilter_conntrack\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := libnetfilter_conntrackFormula{
		Description:  "Library providing an API to the in-kernel connection tracking state table",
		Homepage:     "https://www.netfilter.org/projects/libnetfilter_conntrack/",
		URL:          "https://www.netfilter.org/projects/libnetfilter_conntrack/downloads.html",
		Sha256:       "c8a9efde4edf4bbd45308e4f1492b5efaff4dede7b4c9bf0f54f14f50699ba66",
		Dependencies: []string{"pkg-config", "libmnl", "libnfnetlink"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installlibnetfilter_conntrack(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg libnetfilter_conntrackFormula) Installlibnetfilter_conntrack() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "downloads.html"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "downloads"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
