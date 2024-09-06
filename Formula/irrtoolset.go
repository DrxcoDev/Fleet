package main

// Irrtoolset - Tools to work with Internet routing policies
// Homepage: https://github.com/irrtoolset/irrtoolset

import (
	"fmt"
	
	"os/exec"
)

func installIrrtoolset() {
	// Método 1: Descargar y extraer .tar.gz
	irrtoolset_tar_url := "https://github.com/irrtoolset/irrtoolset/archive/refs/tags/release-5.1.3.tar.gz"
	irrtoolset_cmd_tar := exec.Command("curl", "-L", irrtoolset_tar_url, "-o", "package.tar.gz")
	err := irrtoolset_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	irrtoolset_zip_url := "https://github.com/irrtoolset/irrtoolset/archive/refs/tags/release-5.1.3.zip"
	irrtoolset_cmd_zip := exec.Command("curl", "-L", irrtoolset_zip_url, "-o", "package.zip")
	err = irrtoolset_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	irrtoolset_bin_url := "https://github.com/irrtoolset/irrtoolset/archive/refs/tags/release-5.1.3.bin"
	irrtoolset_cmd_bin := exec.Command("curl", "-L", irrtoolset_bin_url, "-o", "binary.bin")
	err = irrtoolset_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	irrtoolset_src_url := "https://github.com/irrtoolset/irrtoolset/archive/refs/tags/release-5.1.3.src.tar.gz"
	irrtoolset_cmd_src := exec.Command("curl", "-L", irrtoolset_src_url, "-o", "source.tar.gz")
	err = irrtoolset_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	irrtoolset_cmd_direct := exec.Command("./binary")
	err = irrtoolset_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}