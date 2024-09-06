package main

// ReconNg - Web Reconnaissance Framework
// Homepage: https://github.com/lanmaster53/recon-ng

import (
	"fmt"
	
	"os/exec"
)

func installReconNg() {
	// Método 1: Descargar y extraer .tar.gz
	reconng_tar_url := "https://github.com/lanmaster53/recon-ng/archive/refs/tags/v5.1.2.tar.gz"
	reconng_cmd_tar := exec.Command("curl", "-L", reconng_tar_url, "-o", "package.tar.gz")
	err := reconng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reconng_zip_url := "https://github.com/lanmaster53/recon-ng/archive/refs/tags/v5.1.2.zip"
	reconng_cmd_zip := exec.Command("curl", "-L", reconng_zip_url, "-o", "package.zip")
	err = reconng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reconng_bin_url := "https://github.com/lanmaster53/recon-ng/archive/refs/tags/v5.1.2.bin"
	reconng_cmd_bin := exec.Command("curl", "-L", reconng_bin_url, "-o", "binary.bin")
	err = reconng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reconng_src_url := "https://github.com/lanmaster53/recon-ng/archive/refs/tags/v5.1.2.src.tar.gz"
	reconng_cmd_src := exec.Command("curl", "-L", reconng_src_url, "-o", "source.tar.gz")
	err = reconng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reconng_cmd_direct := exec.Command("./binary")
	err = reconng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}