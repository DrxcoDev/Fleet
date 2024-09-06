package main

// Snappy - Compression/decompression library aiming for high speed
// Homepage: https://google.github.io/snappy/

import (
	"fmt"
	
	"os/exec"
)

func installSnappy() {
	// Método 1: Descargar y extraer .tar.gz
	snappy_tar_url := "https://github.com/google/snappy/archive/refs/tags/1.2.1.tar.gz"
	snappy_cmd_tar := exec.Command("curl", "-L", snappy_tar_url, "-o", "package.tar.gz")
	err := snappy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snappy_zip_url := "https://github.com/google/snappy/archive/refs/tags/1.2.1.zip"
	snappy_cmd_zip := exec.Command("curl", "-L", snappy_zip_url, "-o", "package.zip")
	err = snappy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snappy_bin_url := "https://github.com/google/snappy/archive/refs/tags/1.2.1.bin"
	snappy_cmd_bin := exec.Command("curl", "-L", snappy_bin_url, "-o", "binary.bin")
	err = snappy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snappy_src_url := "https://github.com/google/snappy/archive/refs/tags/1.2.1.src.tar.gz"
	snappy_cmd_src := exec.Command("curl", "-L", snappy_src_url, "-o", "source.tar.gz")
	err = snappy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snappy_cmd_direct := exec.Command("./binary")
	err = snappy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}