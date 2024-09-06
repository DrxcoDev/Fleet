package main

// Chainsaw - Rapidly Search and Hunt through Windows Forensic Artefacts
// Homepage: https://github.com/WithSecureLabs/chainsaw

import (
	"fmt"
	
	"os/exec"
)

func installChainsaw() {
	// Método 1: Descargar y extraer .tar.gz
	chainsaw_tar_url := "https://github.com/WithSecureLabs/chainsaw/archive/refs/tags/v2.10.0.tar.gz"
	chainsaw_cmd_tar := exec.Command("curl", "-L", chainsaw_tar_url, "-o", "package.tar.gz")
	err := chainsaw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chainsaw_zip_url := "https://github.com/WithSecureLabs/chainsaw/archive/refs/tags/v2.10.0.zip"
	chainsaw_cmd_zip := exec.Command("curl", "-L", chainsaw_zip_url, "-o", "package.zip")
	err = chainsaw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chainsaw_bin_url := "https://github.com/WithSecureLabs/chainsaw/archive/refs/tags/v2.10.0.bin"
	chainsaw_cmd_bin := exec.Command("curl", "-L", chainsaw_bin_url, "-o", "binary.bin")
	err = chainsaw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chainsaw_src_url := "https://github.com/WithSecureLabs/chainsaw/archive/refs/tags/v2.10.0.src.tar.gz"
	chainsaw_cmd_src := exec.Command("curl", "-L", chainsaw_src_url, "-o", "source.tar.gz")
	err = chainsaw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chainsaw_cmd_direct := exec.Command("./binary")
	err = chainsaw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}