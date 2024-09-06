package main

// Rtags - Source code cross-referencer like ctags with a clang frontend
// Homepage: https://github.com/Andersbakken/rtags

import (
	"fmt"
	
	"os/exec"
)

func installRtags() {
	// Método 1: Descargar y extraer .tar.gz
	rtags_tar_url := "https://github.com/Andersbakken/rtags.git"
	rtags_cmd_tar := exec.Command("curl", "-L", rtags_tar_url, "-o", "package.tar.gz")
	err := rtags_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rtags_zip_url := "https://github.com/Andersbakken/rtags.git"
	rtags_cmd_zip := exec.Command("curl", "-L", rtags_zip_url, "-o", "package.zip")
	err = rtags_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rtags_bin_url := "https://github.com/Andersbakken/rtags.git"
	rtags_cmd_bin := exec.Command("curl", "-L", rtags_bin_url, "-o", "binary.bin")
	err = rtags_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rtags_src_url := "https://github.com/Andersbakken/rtags.git"
	rtags_cmd_src := exec.Command("curl", "-L", rtags_src_url, "-o", "source.tar.gz")
	err = rtags_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rtags_cmd_direct := exec.Command("./binary")
	err = rtags_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: emacs")
	exec.Command("latte", "install", "emacs").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}