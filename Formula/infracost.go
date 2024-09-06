package main

// Infracost - Cost estimates for Terraform
// Homepage: https://www.infracost.io/docs/

import (
	"fmt"
	
	"os/exec"
)

func installInfracost() {
	// Método 1: Descargar y extraer .tar.gz
	infracost_tar_url := "https://github.com/infracost/infracost/archive/refs/tags/v0.10.39.tar.gz"
	infracost_cmd_tar := exec.Command("curl", "-L", infracost_tar_url, "-o", "package.tar.gz")
	err := infracost_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	infracost_zip_url := "https://github.com/infracost/infracost/archive/refs/tags/v0.10.39.zip"
	infracost_cmd_zip := exec.Command("curl", "-L", infracost_zip_url, "-o", "package.zip")
	err = infracost_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	infracost_bin_url := "https://github.com/infracost/infracost/archive/refs/tags/v0.10.39.bin"
	infracost_cmd_bin := exec.Command("curl", "-L", infracost_bin_url, "-o", "binary.bin")
	err = infracost_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	infracost_src_url := "https://github.com/infracost/infracost/archive/refs/tags/v0.10.39.src.tar.gz"
	infracost_cmd_src := exec.Command("curl", "-L", infracost_src_url, "-o", "source.tar.gz")
	err = infracost_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	infracost_cmd_direct := exec.Command("./binary")
	err = infracost_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}