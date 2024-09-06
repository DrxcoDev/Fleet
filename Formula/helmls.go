package main

// HelmLs - Language server for Helm
// Homepage: https://github.com/mrjosh/helm-ls

import (
	"fmt"
	
	"os/exec"
)

func installHelmLs() {
	// Método 1: Descargar y extraer .tar.gz
	helmls_tar_url := "https://github.com/mrjosh/helm-ls/archive/refs/tags/v0.1.0.tar.gz"
	helmls_cmd_tar := exec.Command("curl", "-L", helmls_tar_url, "-o", "package.tar.gz")
	err := helmls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	helmls_zip_url := "https://github.com/mrjosh/helm-ls/archive/refs/tags/v0.1.0.zip"
	helmls_cmd_zip := exec.Command("curl", "-L", helmls_zip_url, "-o", "package.zip")
	err = helmls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	helmls_bin_url := "https://github.com/mrjosh/helm-ls/archive/refs/tags/v0.1.0.bin"
	helmls_cmd_bin := exec.Command("curl", "-L", helmls_bin_url, "-o", "binary.bin")
	err = helmls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	helmls_src_url := "https://github.com/mrjosh/helm-ls/archive/refs/tags/v0.1.0.src.tar.gz"
	helmls_cmd_src := exec.Command("curl", "-L", helmls_src_url, "-o", "source.tar.gz")
	err = helmls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	helmls_cmd_direct := exec.Command("./binary")
	err = helmls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}