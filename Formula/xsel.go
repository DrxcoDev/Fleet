package main

// Xsel - Command-line program for getting and setting the contents of the X selection
// Homepage: https://www.vergenet.net/~conrad/software/xsel/

import (
	"fmt"
	
	"os/exec"
)

func installXsel() {
	// Método 1: Descargar y extraer .tar.gz
	xsel_tar_url := "https://github.com/kfish/xsel/archive/refs/tags/1.2.1.tar.gz"
	xsel_cmd_tar := exec.Command("curl", "-L", xsel_tar_url, "-o", "package.tar.gz")
	err := xsel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xsel_zip_url := "https://github.com/kfish/xsel/archive/refs/tags/1.2.1.zip"
	xsel_cmd_zip := exec.Command("curl", "-L", xsel_zip_url, "-o", "package.zip")
	err = xsel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xsel_bin_url := "https://github.com/kfish/xsel/archive/refs/tags/1.2.1.bin"
	xsel_cmd_bin := exec.Command("curl", "-L", xsel_bin_url, "-o", "binary.bin")
	err = xsel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xsel_src_url := "https://github.com/kfish/xsel/archive/refs/tags/1.2.1.src.tar.gz"
	xsel_cmd_src := exec.Command("curl", "-L", xsel_src_url, "-o", "source.tar.gz")
	err = xsel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xsel_cmd_direct := exec.Command("./binary")
	err = xsel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
}