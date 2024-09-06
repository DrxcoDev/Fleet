package main

// Kitex - Golang RPC framework for microservices
// Homepage: https://github.com/cloudwego/kitex

import (
	"fmt"
	
	"os/exec"
)

func installKitex() {
	// Método 1: Descargar y extraer .tar.gz
	kitex_tar_url := "https://github.com/cloudwego/kitex/archive/refs/tags/v0.10.3.tar.gz"
	kitex_cmd_tar := exec.Command("curl", "-L", kitex_tar_url, "-o", "package.tar.gz")
	err := kitex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kitex_zip_url := "https://github.com/cloudwego/kitex/archive/refs/tags/v0.10.3.zip"
	kitex_cmd_zip := exec.Command("curl", "-L", kitex_zip_url, "-o", "package.zip")
	err = kitex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kitex_bin_url := "https://github.com/cloudwego/kitex/archive/refs/tags/v0.10.3.bin"
	kitex_cmd_bin := exec.Command("curl", "-L", kitex_bin_url, "-o", "binary.bin")
	err = kitex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kitex_src_url := "https://github.com/cloudwego/kitex/archive/refs/tags/v0.10.3.src.tar.gz"
	kitex_cmd_src := exec.Command("curl", "-L", kitex_src_url, "-o", "source.tar.gz")
	err = kitex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kitex_cmd_direct := exec.Command("./binary")
	err = kitex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: thriftgo")
	exec.Command("latte", "install", "thriftgo").Run()
}