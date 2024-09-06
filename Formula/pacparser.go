package main

// Pacparser - Library to parse proxy auto-config (PAC) files
// Homepage: https://github.com/manugarg/pacparser

import (
	"fmt"
	
	"os/exec"
)

func installPacparser() {
	// Método 1: Descargar y extraer .tar.gz
	pacparser_tar_url := "https://github.com/manugarg/pacparser/archive/refs/tags/v1.4.5.tar.gz"
	pacparser_cmd_tar := exec.Command("curl", "-L", pacparser_tar_url, "-o", "package.tar.gz")
	err := pacparser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pacparser_zip_url := "https://github.com/manugarg/pacparser/archive/refs/tags/v1.4.5.zip"
	pacparser_cmd_zip := exec.Command("curl", "-L", pacparser_zip_url, "-o", "package.zip")
	err = pacparser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pacparser_bin_url := "https://github.com/manugarg/pacparser/archive/refs/tags/v1.4.5.bin"
	pacparser_cmd_bin := exec.Command("curl", "-L", pacparser_bin_url, "-o", "binary.bin")
	err = pacparser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pacparser_src_url := "https://github.com/manugarg/pacparser/archive/refs/tags/v1.4.5.src.tar.gz"
	pacparser_cmd_src := exec.Command("curl", "-L", pacparser_src_url, "-o", "source.tar.gz")
	err = pacparser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pacparser_cmd_direct := exec.Command("./binary")
	err = pacparser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}