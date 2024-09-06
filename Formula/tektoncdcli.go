package main

// TektoncdCli - CLI for interacting with TektonCD
// Homepage: https://github.com/tektoncd/cli

import (
	"fmt"
	
	"os/exec"
)

func installTektoncdCli() {
	// Método 1: Descargar y extraer .tar.gz
	tektoncdcli_tar_url := "https://github.com/tektoncd/cli/archive/refs/tags/v0.38.0.tar.gz"
	tektoncdcli_cmd_tar := exec.Command("curl", "-L", tektoncdcli_tar_url, "-o", "package.tar.gz")
	err := tektoncdcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tektoncdcli_zip_url := "https://github.com/tektoncd/cli/archive/refs/tags/v0.38.0.zip"
	tektoncdcli_cmd_zip := exec.Command("curl", "-L", tektoncdcli_zip_url, "-o", "package.zip")
	err = tektoncdcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tektoncdcli_bin_url := "https://github.com/tektoncd/cli/archive/refs/tags/v0.38.0.bin"
	tektoncdcli_cmd_bin := exec.Command("curl", "-L", tektoncdcli_bin_url, "-o", "binary.bin")
	err = tektoncdcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tektoncdcli_src_url := "https://github.com/tektoncd/cli/archive/refs/tags/v0.38.0.src.tar.gz"
	tektoncdcli_cmd_src := exec.Command("curl", "-L", tektoncdcli_src_url, "-o", "source.tar.gz")
	err = tektoncdcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tektoncdcli_cmd_direct := exec.Command("./binary")
	err = tektoncdcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}