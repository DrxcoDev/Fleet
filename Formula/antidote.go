package main

// Antidote - Plugin manager for zsh, inspired by antigen and antibody
// Homepage: https://getantidote.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installAntidote() {
	// Método 1: Descargar y extraer .tar.gz
	antidote_tar_url := "https://github.com/mattmc3/antidote/archive/refs/tags/v1.9.7.tar.gz"
	antidote_cmd_tar := exec.Command("curl", "-L", antidote_tar_url, "-o", "package.tar.gz")
	err := antidote_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	antidote_zip_url := "https://github.com/mattmc3/antidote/archive/refs/tags/v1.9.7.zip"
	antidote_cmd_zip := exec.Command("curl", "-L", antidote_zip_url, "-o", "package.zip")
	err = antidote_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	antidote_bin_url := "https://github.com/mattmc3/antidote/archive/refs/tags/v1.9.7.bin"
	antidote_cmd_bin := exec.Command("curl", "-L", antidote_bin_url, "-o", "binary.bin")
	err = antidote_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	antidote_src_url := "https://github.com/mattmc3/antidote/archive/refs/tags/v1.9.7.src.tar.gz"
	antidote_cmd_src := exec.Command("curl", "-L", antidote_src_url, "-o", "source.tar.gz")
	err = antidote_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	antidote_cmd_direct := exec.Command("./binary")
	err = antidote_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}