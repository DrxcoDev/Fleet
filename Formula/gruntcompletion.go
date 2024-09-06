package main

// GruntCompletion - Bash and Zsh completion for Grunt
// Homepage: https://gruntjs.com/

import (
	"fmt"
	
	"os/exec"
)

func installGruntCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	gruntcompletion_tar_url := "https://github.com/gruntjs/grunt-cli/archive/refs/tags/v1.5.0.tar.gz"
	gruntcompletion_cmd_tar := exec.Command("curl", "-L", gruntcompletion_tar_url, "-o", "package.tar.gz")
	err := gruntcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gruntcompletion_zip_url := "https://github.com/gruntjs/grunt-cli/archive/refs/tags/v1.5.0.zip"
	gruntcompletion_cmd_zip := exec.Command("curl", "-L", gruntcompletion_zip_url, "-o", "package.zip")
	err = gruntcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gruntcompletion_bin_url := "https://github.com/gruntjs/grunt-cli/archive/refs/tags/v1.5.0.bin"
	gruntcompletion_cmd_bin := exec.Command("curl", "-L", gruntcompletion_bin_url, "-o", "binary.bin")
	err = gruntcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gruntcompletion_src_url := "https://github.com/gruntjs/grunt-cli/archive/refs/tags/v1.5.0.src.tar.gz"
	gruntcompletion_cmd_src := exec.Command("curl", "-L", gruntcompletion_src_url, "-o", "source.tar.gz")
	err = gruntcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gruntcompletion_cmd_direct := exec.Command("./binary")
	err = gruntcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}