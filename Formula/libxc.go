package main

// Libxc - Library of exchange and correlation functionals for codes
// Homepage: https://tddft.org/programs/libxc/

import (
	"fmt"
	
	"os/exec"
)

func installLibxc() {
	// Método 1: Descargar y extraer .tar.gz
	libxc_tar_url := "https://gitlab.com/libxc/libxc/-/archive/6.2.2/libxc-6.2.2.tar.bz2"
	libxc_cmd_tar := exec.Command("curl", "-L", libxc_tar_url, "-o", "package.tar.gz")
	err := libxc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxc_zip_url := "https://gitlab.com/libxc/libxc/-/archive/6.2.2/libxc-6.2.2.tar.bz2"
	libxc_cmd_zip := exec.Command("curl", "-L", libxc_zip_url, "-o", "package.zip")
	err = libxc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxc_bin_url := "https://gitlab.com/libxc/libxc/-/archive/6.2.2/libxc-6.2.2.tar.bz2"
	libxc_cmd_bin := exec.Command("curl", "-L", libxc_bin_url, "-o", "binary.bin")
	err = libxc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxc_src_url := "https://gitlab.com/libxc/libxc/-/archive/6.2.2/libxc-6.2.2.tar.bz2"
	libxc_cmd_src := exec.Command("curl", "-L", libxc_src_url, "-o", "source.tar.gz")
	err = libxc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxc_cmd_direct := exec.Command("./binary")
	err = libxc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}