package main

// GhzWeb - Web interface for ghz
// Homepage: https://ghz.sh

import (
	"fmt"
	
	"os/exec"
)

func installGhzWeb() {
	// Método 1: Descargar y extraer .tar.gz
	ghzweb_tar_url := "https://github.com/bojand/ghz/archive/refs/tags/v0.120.0.tar.gz"
	ghzweb_cmd_tar := exec.Command("curl", "-L", ghzweb_tar_url, "-o", "package.tar.gz")
	err := ghzweb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghzweb_zip_url := "https://github.com/bojand/ghz/archive/refs/tags/v0.120.0.zip"
	ghzweb_cmd_zip := exec.Command("curl", "-L", ghzweb_zip_url, "-o", "package.zip")
	err = ghzweb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghzweb_bin_url := "https://github.com/bojand/ghz/archive/refs/tags/v0.120.0.bin"
	ghzweb_cmd_bin := exec.Command("curl", "-L", ghzweb_bin_url, "-o", "binary.bin")
	err = ghzweb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghzweb_src_url := "https://github.com/bojand/ghz/archive/refs/tags/v0.120.0.src.tar.gz"
	ghzweb_cmd_src := exec.Command("curl", "-L", ghzweb_src_url, "-o", "source.tar.gz")
	err = ghzweb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghzweb_cmd_direct := exec.Command("./binary")
	err = ghzweb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}