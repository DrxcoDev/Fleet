package main

// LuckyCommit - Customize your git commit hashes!
// Homepage: https://github.com/not-an-aardvark/lucky-commit

import (
	"fmt"
	
	"os/exec"
)

func installLuckyCommit() {
	// Método 1: Descargar y extraer .tar.gz
	luckycommit_tar_url := "https://github.com/not-an-aardvark/lucky-commit/archive/refs/tags/v2.2.3.tar.gz"
	luckycommit_cmd_tar := exec.Command("curl", "-L", luckycommit_tar_url, "-o", "package.tar.gz")
	err := luckycommit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	luckycommit_zip_url := "https://github.com/not-an-aardvark/lucky-commit/archive/refs/tags/v2.2.3.zip"
	luckycommit_cmd_zip := exec.Command("curl", "-L", luckycommit_zip_url, "-o", "package.zip")
	err = luckycommit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	luckycommit_bin_url := "https://github.com/not-an-aardvark/lucky-commit/archive/refs/tags/v2.2.3.bin"
	luckycommit_cmd_bin := exec.Command("curl", "-L", luckycommit_bin_url, "-o", "binary.bin")
	err = luckycommit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	luckycommit_src_url := "https://github.com/not-an-aardvark/lucky-commit/archive/refs/tags/v2.2.3.src.tar.gz"
	luckycommit_cmd_src := exec.Command("curl", "-L", luckycommit_src_url, "-o", "source.tar.gz")
	err = luckycommit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	luckycommit_cmd_direct := exec.Command("./binary")
	err = luckycommit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: opencl-icd-loader")
	exec.Command("latte", "install", "opencl-icd-loader").Run()
	fmt.Println("Instalando dependencia: pocl")
	exec.Command("latte", "install", "pocl").Run()
}