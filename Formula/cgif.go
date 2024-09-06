package main

// Cgif - GIF encoder written in C
// Homepage: https://github.com/dloebl/cgif

import (
	"fmt"
	
	"os/exec"
)

func installCgif() {
	// Método 1: Descargar y extraer .tar.gz
	cgif_tar_url := "https://github.com/dloebl/cgif/archive/refs/tags/v0.4.1.tar.gz"
	cgif_cmd_tar := exec.Command("curl", "-L", cgif_tar_url, "-o", "package.tar.gz")
	err := cgif_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cgif_zip_url := "https://github.com/dloebl/cgif/archive/refs/tags/v0.4.1.zip"
	cgif_cmd_zip := exec.Command("curl", "-L", cgif_zip_url, "-o", "package.zip")
	err = cgif_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cgif_bin_url := "https://github.com/dloebl/cgif/archive/refs/tags/v0.4.1.bin"
	cgif_cmd_bin := exec.Command("curl", "-L", cgif_bin_url, "-o", "binary.bin")
	err = cgif_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cgif_src_url := "https://github.com/dloebl/cgif/archive/refs/tags/v0.4.1.src.tar.gz"
	cgif_cmd_src := exec.Command("curl", "-L", cgif_src_url, "-o", "source.tar.gz")
	err = cgif_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cgif_cmd_direct := exec.Command("./binary")
	err = cgif_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
}