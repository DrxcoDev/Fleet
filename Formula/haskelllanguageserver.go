package main

// HaskellLanguageServer - Integration point for ghcide and haskell-ide-engine. One IDE to rule them all
// Homepage: https://github.com/haskell/haskell-language-server

import (
	"fmt"
	
	"os/exec"
)

func installHaskellLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	haskelllanguageserver_tar_url := "https://github.com/haskell/haskell-language-server/releases/download/2.9.0.1/haskell-language-server-2.9.0.1-src.tar.gz"
	haskelllanguageserver_cmd_tar := exec.Command("curl", "-L", haskelllanguageserver_tar_url, "-o", "package.tar.gz")
	err := haskelllanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	haskelllanguageserver_zip_url := "https://github.com/haskell/haskell-language-server/releases/download/2.9.0.1/haskell-language-server-2.9.0.1-src.zip"
	haskelllanguageserver_cmd_zip := exec.Command("curl", "-L", haskelllanguageserver_zip_url, "-o", "package.zip")
	err = haskelllanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	haskelllanguageserver_bin_url := "https://github.com/haskell/haskell-language-server/releases/download/2.9.0.1/haskell-language-server-2.9.0.1-src.bin"
	haskelllanguageserver_cmd_bin := exec.Command("curl", "-L", haskelllanguageserver_bin_url, "-o", "binary.bin")
	err = haskelllanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	haskelllanguageserver_src_url := "https://github.com/haskell/haskell-language-server/releases/download/2.9.0.1/haskell-language-server-2.9.0.1-src.src.tar.gz"
	haskelllanguageserver_cmd_src := exec.Command("curl", "-L", haskelllanguageserver_src_url, "-o", "source.tar.gz")
	err = haskelllanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	haskelllanguageserver_cmd_direct := exec.Command("./binary")
	err = haskelllanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
	fmt.Println("Instalando dependencia: ghc@9.6")
	exec.Command("latte", "install", "ghc@9.6").Run()
	fmt.Println("Instalando dependencia: ghc@9.8")
	exec.Command("latte", "install", "ghc@9.8").Run()
}