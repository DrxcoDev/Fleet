package main

import (
	"fmt"
	"log"
	"os/exec"
)

// ideviceinstallerFormula represents a formula in Go.
type ideviceinstallerFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg ideviceinstallerFormula) Print() {
	fmt.Printf("Name: ideviceinstaller\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := ideviceinstallerFormula{
		Description:  "Tool for managing apps on iOS devices",
		Homepage:     "https://www.libimobiledevice.org/",
		URL:          "https://github.com/libimobiledevice/ideviceinstaller/releases/download/1.1.1/ideviceinstaller-1.1.1.tar.bz2",
		Sha256:       "85e4dd1751770e3068d164a48202b714923ac6acb344934f67ee2c9b6bd5b375",
		Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "libimobiledevice", "libplist", "libzip"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installideviceinstaller(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg ideviceinstallerFormula) Installideviceinstaller() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "ideviceinstaller-1.1.1.tar.bz2"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "ideviceinstaller-1.1.1.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
