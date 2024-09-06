package main

// Align - Text column alignment filter
// Homepage: https://kinzler.com/me/align/

import (
	"fmt"
	
	"os/exec"
)

func installAlign() {
	// Método 1: Descargar y extraer .tar.gz
	align_tar_url := "https://kinzler.com/me/align/align-1.7.5.tgz"
	align_cmd_tar := exec.Command("curl", "-L", align_tar_url, "-o", "package.tar.gz")
	err := align_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	align_zip_url := "https://kinzler.com/me/align/align-1.7.5.tgz"
	align_cmd_zip := exec.Command("curl", "-L", align_zip_url, "-o", "package.zip")
	err = align_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	align_bin_url := "https://kinzler.com/me/align/align-1.7.5.tgz"
	align_cmd_bin := exec.Command("curl", "-L", align_bin_url, "-o", "binary.bin")
	err = align_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	align_src_url := "https://kinzler.com/me/align/align-1.7.5.tgz"
	align_cmd_src := exec.Command("curl", "-L", align_src_url, "-o", "source.tar.gz")
	err = align_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	align_cmd_direct := exec.Command("./binary")
	err = align_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}