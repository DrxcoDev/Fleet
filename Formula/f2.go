package main

// F2 - Command-line batch renaming tool
// Homepage: https://github.com/ayoisaiah/f2

import (
	"fmt"
	
	"os/exec"
)

func installF2() {
	// Método 1: Descargar y extraer .tar.gz
	f2_tar_url := "https://github.com/ayoisaiah/f2/archive/refs/tags/v1.9.1.tar.gz"
	f2_cmd_tar := exec.Command("curl", "-L", f2_tar_url, "-o", "package.tar.gz")
	err := f2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	f2_zip_url := "https://github.com/ayoisaiah/f2/archive/refs/tags/v1.9.1.zip"
	f2_cmd_zip := exec.Command("curl", "-L", f2_zip_url, "-o", "package.zip")
	err = f2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	f2_bin_url := "https://github.com/ayoisaiah/f2/archive/refs/tags/v1.9.1.bin"
	f2_cmd_bin := exec.Command("curl", "-L", f2_bin_url, "-o", "binary.bin")
	err = f2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	f2_src_url := "https://github.com/ayoisaiah/f2/archive/refs/tags/v1.9.1.src.tar.gz"
	f2_cmd_src := exec.Command("curl", "-L", f2_src_url, "-o", "source.tar.gz")
	err = f2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	f2_cmd_direct := exec.Command("./binary")
	err = f2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}