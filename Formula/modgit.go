package main

// Modgit - Tool for git repo deploy with filters. Used for magento development
// Homepage: https://github.com/jreinke/modgit

import (
	"fmt"
	
	"os/exec"
)

func installModgit() {
	// Método 1: Descargar y extraer .tar.gz
	modgit_tar_url := "https://github.com/jreinke/modgit/archive/refs/tags/v1.1.0.tar.gz"
	modgit_cmd_tar := exec.Command("curl", "-L", modgit_tar_url, "-o", "package.tar.gz")
	err := modgit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	modgit_zip_url := "https://github.com/jreinke/modgit/archive/refs/tags/v1.1.0.zip"
	modgit_cmd_zip := exec.Command("curl", "-L", modgit_zip_url, "-o", "package.zip")
	err = modgit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	modgit_bin_url := "https://github.com/jreinke/modgit/archive/refs/tags/v1.1.0.bin"
	modgit_cmd_bin := exec.Command("curl", "-L", modgit_bin_url, "-o", "binary.bin")
	err = modgit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	modgit_src_url := "https://github.com/jreinke/modgit/archive/refs/tags/v1.1.0.src.tar.gz"
	modgit_cmd_src := exec.Command("curl", "-L", modgit_src_url, "-o", "source.tar.gz")
	err = modgit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	modgit_cmd_direct := exec.Command("./binary")
	err = modgit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}