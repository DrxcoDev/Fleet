package main

// GhcAT94 - Glorious Glasgow Haskell Compilation System
// Homepage: https://haskell.org/ghc/

import (
	"fmt"
	
	"os/exec"
)

func installGhcAT94() {
	// Método 1: Descargar y extraer .tar.gz
	ghcat94_tar_url := "https://downloads.haskell.org/~ghc/9.4.8/ghc-9.4.8-src.tar.xz"
	ghcat94_cmd_tar := exec.Command("curl", "-L", ghcat94_tar_url, "-o", "package.tar.gz")
	err := ghcat94_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghcat94_zip_url := "https://downloads.haskell.org/~ghc/9.4.8/ghc-9.4.8-src.tar.xz"
	ghcat94_cmd_zip := exec.Command("curl", "-L", ghcat94_zip_url, "-o", "package.zip")
	err = ghcat94_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghcat94_bin_url := "https://downloads.haskell.org/~ghc/9.4.8/ghc-9.4.8-src.tar.xz"
	ghcat94_cmd_bin := exec.Command("curl", "-L", ghcat94_bin_url, "-o", "binary.bin")
	err = ghcat94_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghcat94_src_url := "https://downloads.haskell.org/~ghc/9.4.8/ghc-9.4.8-src.tar.xz"
	ghcat94_cmd_src := exec.Command("curl", "-L", ghcat94_src_url, "-o", "source.tar.gz")
	err = ghcat94_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghcat94_cmd_direct := exec.Command("./binary")
	err = ghcat94_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
}