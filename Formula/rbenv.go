package main

// Rbenv - Ruby version manager
// Homepage: https://rbenv.org

import (
	"fmt"
	
	"os/exec"
)

func installRbenv() {
	// Método 1: Descargar y extraer .tar.gz
	rbenv_tar_url := "https://github.com/rbenv/rbenv/archive/refs/tags/v1.3.0.tar.gz"
	rbenv_cmd_tar := exec.Command("curl", "-L", rbenv_tar_url, "-o", "package.tar.gz")
	err := rbenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbenv_zip_url := "https://github.com/rbenv/rbenv/archive/refs/tags/v1.3.0.zip"
	rbenv_cmd_zip := exec.Command("curl", "-L", rbenv_zip_url, "-o", "package.zip")
	err = rbenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbenv_bin_url := "https://github.com/rbenv/rbenv/archive/refs/tags/v1.3.0.bin"
	rbenv_cmd_bin := exec.Command("curl", "-L", rbenv_bin_url, "-o", "binary.bin")
	err = rbenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbenv_src_url := "https://github.com/rbenv/rbenv/archive/refs/tags/v1.3.0.src.tar.gz"
	rbenv_cmd_src := exec.Command("curl", "-L", rbenv_src_url, "-o", "source.tar.gz")
	err = rbenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbenv_cmd_direct := exec.Command("./binary")
	err = rbenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby-build")
	exec.Command("latte", "install", "ruby-build").Run()
}