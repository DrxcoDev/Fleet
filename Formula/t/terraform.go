package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Definición de la estructura Formula
type Formula struct {
	Name        string
	Description string
	Homepage    string
	URL         string
	Sha256      string
	License     string
	Install     func() error
	Test        func() error
}

func (f *Formula) InstallPackage() error {
	fmt.Printf("Installing %s...\n", f.Name)
	if err := f.Install(); err != nil {
		return fmt.Errorf("installation failed: %v", err)
	}
	return nil
}

func (f *Formula) TestPackage() error {
	fmt.Printf("Testing %s...\n", f.Name)
	if err := f.Test(); err != nil {
		return fmt.Errorf("testing failed: %v", err)
	}
	return nil
}

var terraform = &Formula{
	Name:        "Terraform",
	Description: "Terraform Infrastructure as Code Tool",
	Homepage:    "https://www.terraform.io/",
	URL:         "https://releases.hashicorp.com/terraform/1.6.2/terraform_1.6.2_linux_amd64.zip",
	Sha256:      "fa83d76815e3b28e58e64f6d0689e4c53e6f33e7b0b61e9fa99b8c7d6c4a7c2b",
	License:     "MPL-2.0",
	Install: func() error {
		fmt.Println("Downloading Terraform...")
		cmd := exec.Command("curl", "-LO", "https://releases.hashicorp.com/terraform/1.6.2/terraform_1.6.2_linux_amd64.zip")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Terraform...")
		cmd = exec.Command("unzip", "terraform_1.6.2_linux_amd64.zip")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing Terraform...")
		cmd = exec.Command("mv", "terraform", "/usr/local/bin/terraform")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing Terraform...")
		cmd := exec.Command("terraform", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := terraform.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := terraform.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Terraform installed and tested successfully!")
}