package main

// ApmBashCompletion - Completion for Atom Package Manager
// Homepage: https://github.com/vigo/apm-bash-completion

import (
	"fmt"
	
	"os/exec"
)

func installApmBashCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	apmbashcompletion_tar_url := "https://github.com/vigo/apm-bash-completion/archive/refs/tags/1.0.0.tar.gz"
	apmbashcompletion_cmd_tar := exec.Command("curl", "-L", apmbashcompletion_tar_url, "-o", "package.tar.gz")
	err := apmbashcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apmbashcompletion_zip_url := "https://github.com/vigo/apm-bash-completion/archive/refs/tags/1.0.0.zip"
	apmbashcompletion_cmd_zip := exec.Command("curl", "-L", apmbashcompletion_zip_url, "-o", "package.zip")
	err = apmbashcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apmbashcompletion_bin_url := "https://github.com/vigo/apm-bash-completion/archive/refs/tags/1.0.0.bin"
	apmbashcompletion_cmd_bin := exec.Command("curl", "-L", apmbashcompletion_bin_url, "-o", "binary.bin")
	err = apmbashcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apmbashcompletion_src_url := "https://github.com/vigo/apm-bash-completion/archive/refs/tags/1.0.0.src.tar.gz"
	apmbashcompletion_cmd_src := exec.Command("curl", "-L", apmbashcompletion_src_url, "-o", "source.tar.gz")
	err = apmbashcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apmbashcompletion_cmd_direct := exec.Command("./binary")
	err = apmbashcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}