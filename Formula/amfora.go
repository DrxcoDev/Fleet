package main

// Amfora - Fancy terminal browser for the Gemini protocol
// Homepage: https://github.com/makew0rld/amfora

import (
	"fmt"
	
	"os/exec"
)

func installAmfora() {
	// Método 1: Descargar y extraer .tar.gz
	amfora_tar_url := "https://github.com/makew0rld/amfora/archive/refs/tags/v1.10.0.tar.gz"
	amfora_cmd_tar := exec.Command("curl", "-L", amfora_tar_url, "-o", "package.tar.gz")
	err := amfora_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	amfora_zip_url := "https://github.com/makew0rld/amfora/archive/refs/tags/v1.10.0.zip"
	amfora_cmd_zip := exec.Command("curl", "-L", amfora_zip_url, "-o", "package.zip")
	err = amfora_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	amfora_bin_url := "https://github.com/makew0rld/amfora/archive/refs/tags/v1.10.0.bin"
	amfora_cmd_bin := exec.Command("curl", "-L", amfora_bin_url, "-o", "binary.bin")
	err = amfora_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	amfora_src_url := "https://github.com/makew0rld/amfora/archive/refs/tags/v1.10.0.src.tar.gz"
	amfora_cmd_src := exec.Command("curl", "-L", amfora_src_url, "-o", "source.tar.gz")
	err = amfora_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	amfora_cmd_direct := exec.Command("./binary")
	err = amfora_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}