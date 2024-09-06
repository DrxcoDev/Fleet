package main

// Qjackctl - Simple Qt application to control the JACK sound server daemon
// Homepage: https://qjackctl.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installQjackctl() {
	// Método 1: Descargar y extraer .tar.gz
	qjackctl_tar_url := "https://downloads.sourceforge.net/project/qjackctl/qjackctl/1.0.1/qjackctl-1.0.1.tar.gz"
	qjackctl_cmd_tar := exec.Command("curl", "-L", qjackctl_tar_url, "-o", "package.tar.gz")
	err := qjackctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qjackctl_zip_url := "https://downloads.sourceforge.net/project/qjackctl/qjackctl/1.0.1/qjackctl-1.0.1.zip"
	qjackctl_cmd_zip := exec.Command("curl", "-L", qjackctl_zip_url, "-o", "package.zip")
	err = qjackctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qjackctl_bin_url := "https://downloads.sourceforge.net/project/qjackctl/qjackctl/1.0.1/qjackctl-1.0.1.bin"
	qjackctl_cmd_bin := exec.Command("curl", "-L", qjackctl_bin_url, "-o", "binary.bin")
	err = qjackctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qjackctl_src_url := "https://downloads.sourceforge.net/project/qjackctl/qjackctl/1.0.1/qjackctl-1.0.1.src.tar.gz"
	qjackctl_cmd_src := exec.Command("curl", "-L", qjackctl_src_url, "-o", "source.tar.gz")
	err = qjackctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qjackctl_cmd_direct := exec.Command("./binary")
	err = qjackctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jack")
	exec.Command("latte", "install", "jack").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}