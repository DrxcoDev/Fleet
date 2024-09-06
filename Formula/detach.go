package main

// Detach - Execute given command in detached process
// Homepage: http://inglorion.net/software/detach/

import (
	"fmt"
	
	"os/exec"
)

func installDetach() {
	// Método 1: Descargar y extraer .tar.gz
	detach_tar_url := "http://inglorion.net/download/detach-0.2.3.tar.bz2"
	detach_cmd_tar := exec.Command("curl", "-L", detach_tar_url, "-o", "package.tar.gz")
	err := detach_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	detach_zip_url := "http://inglorion.net/download/detach-0.2.3.tar.bz2"
	detach_cmd_zip := exec.Command("curl", "-L", detach_zip_url, "-o", "package.zip")
	err = detach_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	detach_bin_url := "http://inglorion.net/download/detach-0.2.3.tar.bz2"
	detach_cmd_bin := exec.Command("curl", "-L", detach_bin_url, "-o", "binary.bin")
	err = detach_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	detach_src_url := "http://inglorion.net/download/detach-0.2.3.tar.bz2"
	detach_cmd_src := exec.Command("curl", "-L", detach_src_url, "-o", "source.tar.gz")
	err = detach_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	detach_cmd_direct := exec.Command("./binary")
	err = detach_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}