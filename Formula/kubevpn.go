package main

// Kubevpn - Offers a Cloud-Native Dev Environment that connects to your K8s cluster network
// Homepage: https://www.kubevpn.cn

import (
	"fmt"
	
	"os/exec"
)

func installKubevpn() {
	// Método 1: Descargar y extraer .tar.gz
	kubevpn_tar_url := "https://github.com/kubenetworks/kubevpn/archive/refs/tags/v2.2.17.tar.gz"
	kubevpn_cmd_tar := exec.Command("curl", "-L", kubevpn_tar_url, "-o", "package.tar.gz")
	err := kubevpn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubevpn_zip_url := "https://github.com/kubenetworks/kubevpn/archive/refs/tags/v2.2.17.zip"
	kubevpn_cmd_zip := exec.Command("curl", "-L", kubevpn_zip_url, "-o", "package.zip")
	err = kubevpn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubevpn_bin_url := "https://github.com/kubenetworks/kubevpn/archive/refs/tags/v2.2.17.bin"
	kubevpn_cmd_bin := exec.Command("curl", "-L", kubevpn_bin_url, "-o", "binary.bin")
	err = kubevpn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubevpn_src_url := "https://github.com/kubenetworks/kubevpn/archive/refs/tags/v2.2.17.src.tar.gz"
	kubevpn_cmd_src := exec.Command("curl", "-L", kubevpn_src_url, "-o", "source.tar.gz")
	err = kubevpn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubevpn_cmd_direct := exec.Command("./binary")
	err = kubevpn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}