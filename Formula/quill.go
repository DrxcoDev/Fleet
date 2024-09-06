package main

// Quill - C++17 Asynchronous Low Latency Logging Library
// Homepage: https://github.com/odygrd/quill

import (
	"fmt"
	
	"os/exec"
)

func installQuill() {
	// Método 1: Descargar y extraer .tar.gz
	quill_tar_url := "https://github.com/odygrd/quill/archive/refs/tags/v7.0.0.tar.gz"
	quill_cmd_tar := exec.Command("curl", "-L", quill_tar_url, "-o", "package.tar.gz")
	err := quill_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	quill_zip_url := "https://github.com/odygrd/quill/archive/refs/tags/v7.0.0.zip"
	quill_cmd_zip := exec.Command("curl", "-L", quill_zip_url, "-o", "package.zip")
	err = quill_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	quill_bin_url := "https://github.com/odygrd/quill/archive/refs/tags/v7.0.0.bin"
	quill_cmd_bin := exec.Command("curl", "-L", quill_bin_url, "-o", "binary.bin")
	err = quill_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	quill_src_url := "https://github.com/odygrd/quill/archive/refs/tags/v7.0.0.src.tar.gz"
	quill_cmd_src := exec.Command("curl", "-L", quill_src_url, "-o", "source.tar.gz")
	err = quill_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	quill_cmd_direct := exec.Command("./binary")
	err = quill_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}