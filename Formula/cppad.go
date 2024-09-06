package main

// Cppad - Differentiation of C++ Algorithms
// Homepage: https://www.coin-or.org/CppAD

import (
	"fmt"
	
	"os/exec"
)

func installCppad() {
	// Método 1: Descargar y extraer .tar.gz
	cppad_tar_url := "https://github.com/coin-or/CppAD/archive/refs/tags/20240000.7.tar.gz"
	cppad_cmd_tar := exec.Command("curl", "-L", cppad_tar_url, "-o", "package.tar.gz")
	err := cppad_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cppad_zip_url := "https://github.com/coin-or/CppAD/archive/refs/tags/20240000.7.zip"
	cppad_cmd_zip := exec.Command("curl", "-L", cppad_zip_url, "-o", "package.zip")
	err = cppad_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cppad_bin_url := "https://github.com/coin-or/CppAD/archive/refs/tags/20240000.7.bin"
	cppad_cmd_bin := exec.Command("curl", "-L", cppad_bin_url, "-o", "binary.bin")
	err = cppad_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cppad_src_url := "https://github.com/coin-or/CppAD/archive/refs/tags/20240000.7.src.tar.gz"
	cppad_cmd_src := exec.Command("curl", "-L", cppad_src_url, "-o", "source.tar.gz")
	err = cppad_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cppad_cmd_direct := exec.Command("./binary")
	err = cppad_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}