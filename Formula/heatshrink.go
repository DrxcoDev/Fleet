package main

// Heatshrink - Data compression library for embedded/real-time systems
// Homepage: https://github.com/atomicobject/heatshrink

import (
	"fmt"
	
	"os/exec"
)

func installHeatshrink() {
	// Método 1: Descargar y extraer .tar.gz
	heatshrink_tar_url := "https://github.com/atomicobject/heatshrink/archive/refs/tags/v0.4.1.tar.gz"
	heatshrink_cmd_tar := exec.Command("curl", "-L", heatshrink_tar_url, "-o", "package.tar.gz")
	err := heatshrink_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	heatshrink_zip_url := "https://github.com/atomicobject/heatshrink/archive/refs/tags/v0.4.1.zip"
	heatshrink_cmd_zip := exec.Command("curl", "-L", heatshrink_zip_url, "-o", "package.zip")
	err = heatshrink_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	heatshrink_bin_url := "https://github.com/atomicobject/heatshrink/archive/refs/tags/v0.4.1.bin"
	heatshrink_cmd_bin := exec.Command("curl", "-L", heatshrink_bin_url, "-o", "binary.bin")
	err = heatshrink_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	heatshrink_src_url := "https://github.com/atomicobject/heatshrink/archive/refs/tags/v0.4.1.src.tar.gz"
	heatshrink_cmd_src := exec.Command("curl", "-L", heatshrink_src_url, "-o", "source.tar.gz")
	err = heatshrink_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	heatshrink_cmd_direct := exec.Command("./binary")
	err = heatshrink_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}