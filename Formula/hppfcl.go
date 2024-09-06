package main

// HppFcl - Extension of the Flexible Collision Library
// Homepage: https://github.com/humanoid-path-planner/hpp-fcl

import (
	"fmt"
	
	"os/exec"
)

func installHppFcl() {
	// Método 1: Descargar y extraer .tar.gz
	hppfcl_tar_url := "https://github.com/humanoid-path-planner/hpp-fcl/releases/download/v2.4.5/hpp-fcl-2.4.5.tar.gz"
	hppfcl_cmd_tar := exec.Command("curl", "-L", hppfcl_tar_url, "-o", "package.tar.gz")
	err := hppfcl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hppfcl_zip_url := "https://github.com/humanoid-path-planner/hpp-fcl/releases/download/v2.4.5/hpp-fcl-2.4.5.zip"
	hppfcl_cmd_zip := exec.Command("curl", "-L", hppfcl_zip_url, "-o", "package.zip")
	err = hppfcl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hppfcl_bin_url := "https://github.com/humanoid-path-planner/hpp-fcl/releases/download/v2.4.5/hpp-fcl-2.4.5.bin"
	hppfcl_cmd_bin := exec.Command("curl", "-L", hppfcl_bin_url, "-o", "binary.bin")
	err = hppfcl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hppfcl_src_url := "https://github.com/humanoid-path-planner/hpp-fcl/releases/download/v2.4.5/hpp-fcl-2.4.5.src.tar.gz"
	hppfcl_cmd_src := exec.Command("curl", "-L", hppfcl_src_url, "-o", "source.tar.gz")
	err = hppfcl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hppfcl_cmd_direct := exec.Command("./binary")
	err = hppfcl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: assimp")
	exec.Command("latte", "install", "assimp").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: boost-python3")
	exec.Command("latte", "install", "boost-python3").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: eigenpy")
	exec.Command("latte", "install", "eigenpy").Run()
	fmt.Println("Instalando dependencia: octomap")
	exec.Command("latte", "install", "octomap").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}