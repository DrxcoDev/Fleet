package main

// Nsync - C library that exports various synchronization primitives
// Homepage: https://github.com/google/nsync

import (
	"fmt"
	
	"os/exec"
)

func installNsync() {
	// Método 1: Descargar y extraer .tar.gz
	nsync_tar_url := "https://github.com/google/nsync/archive/refs/tags/1.29.2.tar.gz"
	nsync_cmd_tar := exec.Command("curl", "-L", nsync_tar_url, "-o", "package.tar.gz")
	err := nsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nsync_zip_url := "https://github.com/google/nsync/archive/refs/tags/1.29.2.zip"
	nsync_cmd_zip := exec.Command("curl", "-L", nsync_zip_url, "-o", "package.zip")
	err = nsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nsync_bin_url := "https://github.com/google/nsync/archive/refs/tags/1.29.2.bin"
	nsync_cmd_bin := exec.Command("curl", "-L", nsync_bin_url, "-o", "binary.bin")
	err = nsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nsync_src_url := "https://github.com/google/nsync/archive/refs/tags/1.29.2.src.tar.gz"
	nsync_cmd_src := exec.Command("curl", "-L", nsync_src_url, "-o", "source.tar.gz")
	err = nsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nsync_cmd_direct := exec.Command("./binary")
	err = nsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}