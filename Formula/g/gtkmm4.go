package main

import (
	"fmt"
	"log"
	"os/exec"
)

// gtkmm4Formula represents a formula in Go.
type gtkmm4Formula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg gtkmm4Formula) Print() {
	fmt.Printf("Name: gtkmm4\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := gtkmm4Formula{
		Description:  "C++ interfaces for GTK+ and GNOME",
		Homepage:     "https://www.gtkmm.org/",
		URL:          "https://download.gnome.org/sources/gtkmm/4.14/gtkmm-4.14.0.tar.xz",
		Sha256:       "96c1cfa59f83f8a91af058e5fcc1b5f888875718ae2974b4ec512918919a91ee",
		Dependencies: []string{"meson", "ninja", "pkg-config", "cairo", "cairomm", "gdk-pixbuf", "glib", "glibmm", "graphene", "gtk4", "libsigc++", "pangomm"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installgtkmm4(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg gtkmm4Formula) Installgtkmm4() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "gtkmm-4.14.0.tar.xz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "gtkmm-4.14.0.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
