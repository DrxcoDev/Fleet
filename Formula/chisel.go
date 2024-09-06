package main

// Chisel - Collection of LLDB commands to assist debugging iOS apps
// Homepage: https://github.com/facebook/chisel

import (
	"fmt"
	
	"os/exec"
)

func installChisel() {
	// Método 1: Descargar y extraer .tar.gz
	chisel_tar_url := "https://github.com/facebook/chisel/archive/refs/tags/2.0.1.tar.gz"
	chisel_cmd_tar := exec.Command("curl", "-L", chisel_tar_url, "-o", "package.tar.gz")
	err := chisel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chisel_zip_url := "https://github.com/facebook/chisel/archive/refs/tags/2.0.1.zip"
	chisel_cmd_zip := exec.Command("curl", "-L", chisel_zip_url, "-o", "package.zip")
	err = chisel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chisel_bin_url := "https://github.com/facebook/chisel/archive/refs/tags/2.0.1.bin"
	chisel_cmd_bin := exec.Command("curl", "-L", chisel_bin_url, "-o", "binary.bin")
	err = chisel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chisel_src_url := "https://github.com/facebook/chisel/archive/refs/tags/2.0.1.src.tar.gz"
	chisel_cmd_src := exec.Command("curl", "-L", chisel_src_url, "-o", "source.tar.gz")
	err = chisel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chisel_cmd_direct := exec.Command("./binary")
	err = chisel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}