package main

import (
	"fmt"
	"log"
	"os/exec"
)

type abicompliancecheckerFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg abicompliancecheckerFormula) Print() {
	fmt.Printf("Name: abi-compliance-checker\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := abi - compliance - checkerFormula{
		Description:  "Tool for checking backward API/ABI compatibility of a C/C++ library",
		Homepage:     "https://lvc.github.io/abi-compliance-checker/",
		URL:          "https://github.com/lvc/abi-compliance-checker/archive/refs/tags/2.3.tar.gz",
		Sha256:       "06af34b7632a01e00b3d6d5ad826d4102e7a840e32b4a0a0bc2a58c3fc799cef",
		Dependencies: []string{"gcc", "universal-ctags"},
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

	fmt.Printf("Installing formula: %s\n", "abi-compliance-checker")
	if err := pkg.Installabicompliancechecker(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
	cmd := exec.Command("brew", "list", name)
	err := cmd.Run()
	return err == nil
}

func (pkg abicompliancecheckerFormula) Installabicompliancechecker() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "2.3.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "2.3.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
