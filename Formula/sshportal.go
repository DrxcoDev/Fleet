package main

// Sshportal - SSH & Telnet bastion server
// Homepage: https://v1.manfred.life/sshportal/

import (
	"fmt"
	
	"os/exec"
)

func installSshportal() {
	// Método 1: Descargar y extraer .tar.gz
	sshportal_tar_url := "https://github.com/moul/sshportal/archive/refs/tags/v1.19.5.tar.gz"
	sshportal_cmd_tar := exec.Command("curl", "-L", sshportal_tar_url, "-o", "package.tar.gz")
	err := sshportal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sshportal_zip_url := "https://github.com/moul/sshportal/archive/refs/tags/v1.19.5.zip"
	sshportal_cmd_zip := exec.Command("curl", "-L", sshportal_zip_url, "-o", "package.zip")
	err = sshportal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sshportal_bin_url := "https://github.com/moul/sshportal/archive/refs/tags/v1.19.5.bin"
	sshportal_cmd_bin := exec.Command("curl", "-L", sshportal_bin_url, "-o", "binary.bin")
	err = sshportal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sshportal_src_url := "https://github.com/moul/sshportal/archive/refs/tags/v1.19.5.src.tar.gz"
	sshportal_cmd_src := exec.Command("curl", "-L", sshportal_src_url, "-o", "source.tar.gz")
	err = sshportal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sshportal_cmd_direct := exec.Command("./binary")
	err = sshportal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}