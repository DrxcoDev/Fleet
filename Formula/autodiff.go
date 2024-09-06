package main

// Autodiff - Automatic differentiation made easier for C++
// Homepage: https://autodiff.github.io

import (
	"fmt"
	
	"os/exec"
)

func installAutodiff() {
	// Método 1: Descargar y extraer .tar.gz
	autodiff_tar_url := "https://github.com/autodiff/autodiff/archive/refs/tags/v1.1.2.tar.gz"
	autodiff_cmd_tar := exec.Command("curl", "-L", autodiff_tar_url, "-o", "package.tar.gz")
	err := autodiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autodiff_zip_url := "https://github.com/autodiff/autodiff/archive/refs/tags/v1.1.2.zip"
	autodiff_cmd_zip := exec.Command("curl", "-L", autodiff_zip_url, "-o", "package.zip")
	err = autodiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autodiff_bin_url := "https://github.com/autodiff/autodiff/archive/refs/tags/v1.1.2.bin"
	autodiff_cmd_bin := exec.Command("curl", "-L", autodiff_bin_url, "-o", "binary.bin")
	err = autodiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autodiff_src_url := "https://github.com/autodiff/autodiff/archive/refs/tags/v1.1.2.src.tar.gz"
	autodiff_cmd_src := exec.Command("curl", "-L", autodiff_src_url, "-o", "source.tar.gz")
	err = autodiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autodiff_cmd_direct := exec.Command("./binary")
	err = autodiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: pybind11")
	exec.Command("latte", "install", "pybind11").Run()
}