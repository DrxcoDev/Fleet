package main

// Mrbayes - Bayesian inference of phylogenies and evolutionary models
// Homepage: https://nbisweden.github.io/MrBayes/

import (
	"fmt"
	
	"os/exec"
)

func installMrbayes() {
	// Método 1: Descargar y extraer .tar.gz
	mrbayes_tar_url := "https://github.com/NBISweden/MrBayes/archive/refs/tags/v3.2.7a.tar.gz"
	mrbayes_cmd_tar := exec.Command("curl", "-L", mrbayes_tar_url, "-o", "package.tar.gz")
	err := mrbayes_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mrbayes_zip_url := "https://github.com/NBISweden/MrBayes/archive/refs/tags/v3.2.7a.zip"
	mrbayes_cmd_zip := exec.Command("curl", "-L", mrbayes_zip_url, "-o", "package.zip")
	err = mrbayes_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mrbayes_bin_url := "https://github.com/NBISweden/MrBayes/archive/refs/tags/v3.2.7a.bin"
	mrbayes_cmd_bin := exec.Command("curl", "-L", mrbayes_bin_url, "-o", "binary.bin")
	err = mrbayes_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mrbayes_src_url := "https://github.com/NBISweden/MrBayes/archive/refs/tags/v3.2.7a.src.tar.gz"
	mrbayes_cmd_src := exec.Command("curl", "-L", mrbayes_src_url, "-o", "source.tar.gz")
	err = mrbayes_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mrbayes_cmd_direct := exec.Command("./binary")
	err = mrbayes_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: beagle")
	exec.Command("latte", "install", "beagle").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
}