package main

// Mp4v2 - Read, create, and modify MP4 files
// Homepage: https://mp4v2.org

import (
	"fmt"
	
	"os/exec"
)

func installMp4v2() {
	// Método 1: Descargar y extraer .tar.gz
	mp4v2_tar_url := "https://github.com/enzo1982/mp4v2/releases/download/v2.1.3/mp4v2-2.1.3.tar.bz2"
	mp4v2_cmd_tar := exec.Command("curl", "-L", mp4v2_tar_url, "-o", "package.tar.gz")
	err := mp4v2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mp4v2_zip_url := "https://github.com/enzo1982/mp4v2/releases/download/v2.1.3/mp4v2-2.1.3.tar.bz2"
	mp4v2_cmd_zip := exec.Command("curl", "-L", mp4v2_zip_url, "-o", "package.zip")
	err = mp4v2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mp4v2_bin_url := "https://github.com/enzo1982/mp4v2/releases/download/v2.1.3/mp4v2-2.1.3.tar.bz2"
	mp4v2_cmd_bin := exec.Command("curl", "-L", mp4v2_bin_url, "-o", "binary.bin")
	err = mp4v2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mp4v2_src_url := "https://github.com/enzo1982/mp4v2/releases/download/v2.1.3/mp4v2-2.1.3.tar.bz2"
	mp4v2_cmd_src := exec.Command("curl", "-L", mp4v2_src_url, "-o", "source.tar.gz")
	err = mp4v2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mp4v2_cmd_direct := exec.Command("./binary")
	err = mp4v2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}