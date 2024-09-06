package main

// Elm - Functional programming language for building browser-based GUIs
// Homepage: https://elm-lang.org

import (
	"fmt"
	
	"os/exec"
)

func installElm() {
	// Método 1: Descargar y extraer .tar.gz
	elm_tar_url := "https://github.com/elm/compiler/archive/refs/tags/0.19.1.tar.gz"
	elm_cmd_tar := exec.Command("curl", "-L", elm_tar_url, "-o", "package.tar.gz")
	err := elm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	elm_zip_url := "https://github.com/elm/compiler/archive/refs/tags/0.19.1.zip"
	elm_cmd_zip := exec.Command("curl", "-L", elm_zip_url, "-o", "package.zip")
	err = elm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	elm_bin_url := "https://github.com/elm/compiler/archive/refs/tags/0.19.1.bin"
	elm_cmd_bin := exec.Command("curl", "-L", elm_bin_url, "-o", "binary.bin")
	err = elm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	elm_src_url := "https://github.com/elm/compiler/archive/refs/tags/0.19.1.src.tar.gz"
	elm_cmd_src := exec.Command("curl", "-L", elm_src_url, "-o", "source.tar.gz")
	err = elm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	elm_cmd_direct := exec.Command("./binary")
	err = elm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
}