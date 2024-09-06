package main

// Rke - Rancher Kubernetes Engine, a Kubernetes installer that works everywhere
// Homepage: https://rke.docs.rancher.com/

import (
	"fmt"
	
	"os/exec"
)

func installRke() {
	// Método 1: Descargar y extraer .tar.gz
	rke_tar_url := "https://github.com/rancher/rke/archive/refs/tags/v1.6.1.tar.gz"
	rke_cmd_tar := exec.Command("curl", "-L", rke_tar_url, "-o", "package.tar.gz")
	err := rke_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rke_zip_url := "https://github.com/rancher/rke/archive/refs/tags/v1.6.1.zip"
	rke_cmd_zip := exec.Command("curl", "-L", rke_zip_url, "-o", "package.zip")
	err = rke_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rke_bin_url := "https://github.com/rancher/rke/archive/refs/tags/v1.6.1.bin"
	rke_cmd_bin := exec.Command("curl", "-L", rke_bin_url, "-o", "binary.bin")
	err = rke_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rke_src_url := "https://github.com/rancher/rke/archive/refs/tags/v1.6.1.src.tar.gz"
	rke_cmd_src := exec.Command("curl", "-L", rke_src_url, "-o", "source.tar.gz")
	err = rke_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rke_cmd_direct := exec.Command("./binary")
	err = rke_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}