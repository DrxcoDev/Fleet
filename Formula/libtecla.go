package main

// Libtecla - Command-line editing facilities similar to the tcsh shell
// Homepage: https://sites.astro.caltech.edu/~mcs/tecla/

import (
	"fmt"
	
	"os/exec"
)

func installLibtecla() {
	// Método 1: Descargar y extraer .tar.gz
	libtecla_tar_url := "https://sites.astro.caltech.edu/~mcs/tecla/libtecla-1.6.3.tar.gz"
	libtecla_cmd_tar := exec.Command("curl", "-L", libtecla_tar_url, "-o", "package.tar.gz")
	err := libtecla_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtecla_zip_url := "https://sites.astro.caltech.edu/~mcs/tecla/libtecla-1.6.3.zip"
	libtecla_cmd_zip := exec.Command("curl", "-L", libtecla_zip_url, "-o", "package.zip")
	err = libtecla_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtecla_bin_url := "https://sites.astro.caltech.edu/~mcs/tecla/libtecla-1.6.3.bin"
	libtecla_cmd_bin := exec.Command("curl", "-L", libtecla_bin_url, "-o", "binary.bin")
	err = libtecla_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtecla_src_url := "https://sites.astro.caltech.edu/~mcs/tecla/libtecla-1.6.3.src.tar.gz"
	libtecla_cmd_src := exec.Command("curl", "-L", libtecla_src_url, "-o", "source.tar.gz")
	err = libtecla_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtecla_cmd_direct := exec.Command("./binary")
	err = libtecla_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}