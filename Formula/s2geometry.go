package main

// S2geometry - Computational geometry and spatial indexing on the sphere
// Homepage: https://github.com/google/s2geometry

import (
	"fmt"
	
	"os/exec"
)

func installS2geometry() {
	// Método 1: Descargar y extraer .tar.gz
	s2geometry_tar_url := "https://github.com/google/s2geometry/archive/refs/tags/v0.11.1.tar.gz"
	s2geometry_cmd_tar := exec.Command("curl", "-L", s2geometry_tar_url, "-o", "package.tar.gz")
	err := s2geometry_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	s2geometry_zip_url := "https://github.com/google/s2geometry/archive/refs/tags/v0.11.1.zip"
	s2geometry_cmd_zip := exec.Command("curl", "-L", s2geometry_zip_url, "-o", "package.zip")
	err = s2geometry_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	s2geometry_bin_url := "https://github.com/google/s2geometry/archive/refs/tags/v0.11.1.bin"
	s2geometry_cmd_bin := exec.Command("curl", "-L", s2geometry_bin_url, "-o", "binary.bin")
	err = s2geometry_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	s2geometry_src_url := "https://github.com/google/s2geometry/archive/refs/tags/v0.11.1.src.tar.gz"
	s2geometry_cmd_src := exec.Command("curl", "-L", s2geometry_src_url, "-o", "source.tar.gz")
	err = s2geometry_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	s2geometry_cmd_direct := exec.Command("./binary")
	err = s2geometry_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: glog")
	exec.Command("latte", "install", "glog").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}