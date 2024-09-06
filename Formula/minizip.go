package main

// Minizip - C library for zip/unzip via zLib
// Homepage: https://www.winimage.com/zLibDll/minizip.html

import (
	"fmt"
	
	"os/exec"
)

func installMinizip() {
	// Método 1: Descargar y extraer .tar.gz
	minizip_tar_url := "https://zlib.net/zlib-1.3.1.tar.gz"
	minizip_cmd_tar := exec.Command("curl", "-L", minizip_tar_url, "-o", "package.tar.gz")
	err := minizip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minizip_zip_url := "https://zlib.net/zlib-1.3.1.zip"
	minizip_cmd_zip := exec.Command("curl", "-L", minizip_zip_url, "-o", "package.zip")
	err = minizip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minizip_bin_url := "https://zlib.net/zlib-1.3.1.bin"
	minizip_cmd_bin := exec.Command("curl", "-L", minizip_bin_url, "-o", "binary.bin")
	err = minizip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minizip_src_url := "https://zlib.net/zlib-1.3.1.src.tar.gz"
	minizip_cmd_src := exec.Command("curl", "-L", minizip_src_url, "-o", "source.tar.gz")
	err = minizip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minizip_cmd_direct := exec.Command("./binary")
	err = minizip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}