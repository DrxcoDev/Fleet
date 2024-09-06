package main

// Nali - Tool for querying IP geographic information and CDN provider
// Homepage: https://github.com/zu1k/nali

import (
	"fmt"
	
	"os/exec"
)

func installNali() {
	// Método 1: Descargar y extraer .tar.gz
	nali_tar_url := "https://github.com/zu1k/nali/archive/refs/tags/v0.8.1.tar.gz"
	nali_cmd_tar := exec.Command("curl", "-L", nali_tar_url, "-o", "package.tar.gz")
	err := nali_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nali_zip_url := "https://github.com/zu1k/nali/archive/refs/tags/v0.8.1.zip"
	nali_cmd_zip := exec.Command("curl", "-L", nali_zip_url, "-o", "package.zip")
	err = nali_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nali_bin_url := "https://github.com/zu1k/nali/archive/refs/tags/v0.8.1.bin"
	nali_cmd_bin := exec.Command("curl", "-L", nali_bin_url, "-o", "binary.bin")
	err = nali_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nali_src_url := "https://github.com/zu1k/nali/archive/refs/tags/v0.8.1.src.tar.gz"
	nali_cmd_src := exec.Command("curl", "-L", nali_src_url, "-o", "source.tar.gz")
	err = nali_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nali_cmd_direct := exec.Command("./binary")
	err = nali_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}