package main

// Amber - Crystal web framework. Bare metal performance, productivity and happiness
// Homepage: https://amberframework.org/

import (
	"fmt"
	
	"os/exec"
)

func installAmber() {
	// Método 1: Descargar y extraer .tar.gz
	amber_tar_url := "https://github.com/amberframework/amber/archive/refs/tags/v1.4.1.tar.gz"
	amber_cmd_tar := exec.Command("curl", "-L", amber_tar_url, "-o", "package.tar.gz")
	err := amber_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	amber_zip_url := "https://github.com/amberframework/amber/archive/refs/tags/v1.4.1.zip"
	amber_cmd_zip := exec.Command("curl", "-L", amber_zip_url, "-o", "package.zip")
	err = amber_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	amber_bin_url := "https://github.com/amberframework/amber/archive/refs/tags/v1.4.1.bin"
	amber_cmd_bin := exec.Command("curl", "-L", amber_bin_url, "-o", "binary.bin")
	err = amber_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	amber_src_url := "https://github.com/amberframework/amber/archive/refs/tags/v1.4.1.src.tar.gz"
	amber_cmd_src := exec.Command("curl", "-L", amber_src_url, "-o", "source.tar.gz")
	err = amber_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	amber_cmd_direct := exec.Command("./binary")
	err = amber_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: crystal")
	exec.Command("latte", "install", "crystal").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}