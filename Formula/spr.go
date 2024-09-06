package main

// Spr - Submit pull requests for individual, amendable, rebaseable commits to GitHub
// Homepage: https://github.com/spacedentist/spr

import (
	"fmt"
	
	"os/exec"
)

func installSpr() {
	// Método 1: Descargar y extraer .tar.gz
	spr_tar_url := "https://github.com/spacedentist/spr/archive/refs/tags/v1.3.5.tar.gz"
	spr_cmd_tar := exec.Command("curl", "-L", spr_tar_url, "-o", "package.tar.gz")
	err := spr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spr_zip_url := "https://github.com/spacedentist/spr/archive/refs/tags/v1.3.5.zip"
	spr_cmd_zip := exec.Command("curl", "-L", spr_zip_url, "-o", "package.zip")
	err = spr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spr_bin_url := "https://github.com/spacedentist/spr/archive/refs/tags/v1.3.5.bin"
	spr_cmd_bin := exec.Command("curl", "-L", spr_bin_url, "-o", "binary.bin")
	err = spr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spr_src_url := "https://github.com/spacedentist/spr/archive/refs/tags/v1.3.5.src.tar.gz"
	spr_cmd_src := exec.Command("curl", "-L", spr_src_url, "-o", "source.tar.gz")
	err = spr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spr_cmd_direct := exec.Command("./binary")
	err = spr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}