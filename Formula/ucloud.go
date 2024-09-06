package main

// Ucloud - Official tool for managing UCloud services
// Homepage: https://www.ucloud.cn

import (
	"fmt"
	
	"os/exec"
)

func installUcloud() {
	// Método 1: Descargar y extraer .tar.gz
	ucloud_tar_url := "https://github.com/ucloud/ucloud-cli/archive/refs/tags/v0.2.0.tar.gz"
	ucloud_cmd_tar := exec.Command("curl", "-L", ucloud_tar_url, "-o", "package.tar.gz")
	err := ucloud_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ucloud_zip_url := "https://github.com/ucloud/ucloud-cli/archive/refs/tags/v0.2.0.zip"
	ucloud_cmd_zip := exec.Command("curl", "-L", ucloud_zip_url, "-o", "package.zip")
	err = ucloud_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ucloud_bin_url := "https://github.com/ucloud/ucloud-cli/archive/refs/tags/v0.2.0.bin"
	ucloud_cmd_bin := exec.Command("curl", "-L", ucloud_bin_url, "-o", "binary.bin")
	err = ucloud_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ucloud_src_url := "https://github.com/ucloud/ucloud-cli/archive/refs/tags/v0.2.0.src.tar.gz"
	ucloud_cmd_src := exec.Command("curl", "-L", ucloud_src_url, "-o", "source.tar.gz")
	err = ucloud_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ucloud_cmd_direct := exec.Command("./binary")
	err = ucloud_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}