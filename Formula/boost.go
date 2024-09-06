package main

// Boost - Collection of portable C++ source libraries
// Homepage: https://www.boost.org/

import (
	"fmt"
	
	"os/exec"
)

func installBoost() {
	// Método 1: Descargar y extraer .tar.gz
	boost_tar_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boost_cmd_tar := exec.Command("curl", "-L", boost_tar_url, "-o", "package.tar.gz")
	err := boost_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boost_zip_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boost_cmd_zip := exec.Command("curl", "-L", boost_zip_url, "-o", "package.zip")
	err = boost_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boost_bin_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boost_cmd_bin := exec.Command("curl", "-L", boost_bin_url, "-o", "binary.bin")
	err = boost_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boost_src_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boost_cmd_src := exec.Command("curl", "-L", boost_src_url, "-o", "source.tar.gz")
	err = boost_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boost_cmd_direct := exec.Command("./binary")
	err = boost_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}