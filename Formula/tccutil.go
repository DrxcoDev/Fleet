package main

// Tccutil - Utility to modify the macOS Accessibility Database (TCC.db)
// Homepage: https://github.com/jacobsalmela/tccutil

import (
	"fmt"
	
	"os/exec"
)

func installTccutil() {
	// Método 1: Descargar y extraer .tar.gz
	tccutil_tar_url := "https://github.com/jacobsalmela/tccutil/archive/refs/tags/v1.4.0.tar.gz"
	tccutil_cmd_tar := exec.Command("curl", "-L", tccutil_tar_url, "-o", "package.tar.gz")
	err := tccutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tccutil_zip_url := "https://github.com/jacobsalmela/tccutil/archive/refs/tags/v1.4.0.zip"
	tccutil_cmd_zip := exec.Command("curl", "-L", tccutil_zip_url, "-o", "package.zip")
	err = tccutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tccutil_bin_url := "https://github.com/jacobsalmela/tccutil/archive/refs/tags/v1.4.0.bin"
	tccutil_cmd_bin := exec.Command("curl", "-L", tccutil_bin_url, "-o", "binary.bin")
	err = tccutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tccutil_src_url := "https://github.com/jacobsalmela/tccutil/archive/refs/tags/v1.4.0.src.tar.gz"
	tccutil_cmd_src := exec.Command("curl", "-L", tccutil_src_url, "-o", "source.tar.gz")
	err = tccutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tccutil_cmd_direct := exec.Command("./binary")
	err = tccutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}