package main

import (
	"fmt"
	"log"
	"os/exec"
)

// gangliaFormula represents a formula in Go.
type gangliaFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg gangliaFormula) Print() {
	fmt.Printf("Name: ganglia\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := gangliaFormula{
		Description:  "Scalable distributed monitoring system",
		Homepage:     "https://ganglia.sourceforge.net/",
		URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/03cf8088210822aa2c1ab544ed58ea04c897d9c4/libtool/configure-pre-0.4.2.418-big_sur.diff",
		Sha256:       "83af02f2aa2b746bb7225872cab29a253264be49db0ecebb12f841562d9a2923",
		Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "apr", "confuse", "pcre", "rrdtool"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installganglia(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg gangliaFormula) Installganglia() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "configure-pre-0.4.2.418-big_sur.diff"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "configure-pre-0.4.2.418-big_sur"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
