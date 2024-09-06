package main

// Libuv - Multi-platform support library with a focus on asynchronous I/O
// Homepage: https://libuv.org

import (
	"fmt"
	
	"os/exec"
)

func installLibuv() {
	// Método 1: Descargar y extraer .tar.gz
	libuv_tar_url := "https://github.com/libuv/libuv/archive/refs/tags/v1.48.0.tar.gz"
	libuv_cmd_tar := exec.Command("curl", "-L", libuv_tar_url, "-o", "package.tar.gz")
	err := libuv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libuv_zip_url := "https://github.com/libuv/libuv/archive/refs/tags/v1.48.0.zip"
	libuv_cmd_zip := exec.Command("curl", "-L", libuv_zip_url, "-o", "package.zip")
	err = libuv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libuv_bin_url := "https://github.com/libuv/libuv/archive/refs/tags/v1.48.0.bin"
	libuv_cmd_bin := exec.Command("curl", "-L", libuv_bin_url, "-o", "binary.bin")
	err = libuv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libuv_src_url := "https://github.com/libuv/libuv/archive/refs/tags/v1.48.0.src.tar.gz"
	libuv_cmd_src := exec.Command("curl", "-L", libuv_src_url, "-o", "source.tar.gz")
	err = libuv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libuv_cmd_direct := exec.Command("./binary")
	err = libuv_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
}