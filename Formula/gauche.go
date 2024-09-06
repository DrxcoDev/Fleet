package main

// Gauche - R7RS Scheme implementation, developed to be a handy script interpreter
// Homepage: https://practical-scheme.net/gauche/

import (
	"fmt"
	
	"os/exec"
)

func installGauche() {
	// Método 1: Descargar y extraer .tar.gz
	gauche_tar_url := "https://github.com/shirok/Gauche/releases/download/release0_9_15/Gauche-0.9.15.tgz"
	gauche_cmd_tar := exec.Command("curl", "-L", gauche_tar_url, "-o", "package.tar.gz")
	err := gauche_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gauche_zip_url := "https://github.com/shirok/Gauche/releases/download/release0_9_15/Gauche-0.9.15.tgz"
	gauche_cmd_zip := exec.Command("curl", "-L", gauche_zip_url, "-o", "package.zip")
	err = gauche_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gauche_bin_url := "https://github.com/shirok/Gauche/releases/download/release0_9_15/Gauche-0.9.15.tgz"
	gauche_cmd_bin := exec.Command("curl", "-L", gauche_bin_url, "-o", "binary.bin")
	err = gauche_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gauche_src_url := "https://github.com/shirok/Gauche/releases/download/release0_9_15/Gauche-0.9.15.tgz"
	gauche_cmd_src := exec.Command("curl", "-L", gauche_src_url, "-o", "source.tar.gz")
	err = gauche_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gauche_cmd_direct := exec.Command("./binary")
	err = gauche_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ca-certificates")
	exec.Command("latte", "install", "ca-certificates").Run()
	fmt.Println("Instalando dependencia: mbedtls")
	exec.Command("latte", "install", "mbedtls").Run()
}