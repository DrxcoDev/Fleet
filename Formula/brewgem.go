package main

// BrewGem - Install RubyGems as Homebrew formulae
// Homepage: https://github.com/sportngin/brew-gem

import (
	"fmt"
	
	"os/exec"
)

func installBrewGem() {
	// Método 1: Descargar y extraer .tar.gz
	brewgem_tar_url := "https://github.com/sportngin/brew-gem/archive/refs/tags/v1.2.0.tar.gz"
	brewgem_cmd_tar := exec.Command("curl", "-L", brewgem_tar_url, "-o", "package.tar.gz")
	err := brewgem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brewgem_zip_url := "https://github.com/sportngin/brew-gem/archive/refs/tags/v1.2.0.zip"
	brewgem_cmd_zip := exec.Command("curl", "-L", brewgem_zip_url, "-o", "package.zip")
	err = brewgem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brewgem_bin_url := "https://github.com/sportngin/brew-gem/archive/refs/tags/v1.2.0.bin"
	brewgem_cmd_bin := exec.Command("curl", "-L", brewgem_bin_url, "-o", "binary.bin")
	err = brewgem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brewgem_src_url := "https://github.com/sportngin/brew-gem/archive/refs/tags/v1.2.0.src.tar.gz"
	brewgem_cmd_src := exec.Command("curl", "-L", brewgem_src_url, "-o", "source.tar.gz")
	err = brewgem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brewgem_cmd_direct := exec.Command("./binary")
	err = brewgem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}