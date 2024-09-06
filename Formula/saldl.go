package main

// Saldl - CLI downloader optimized for speed and early preview
// Homepage: https://saldl.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installSaldl() {
	// Método 1: Descargar y extraer .tar.gz
	saldl_tar_url := "https://github.com/saldl/saldl/archive/refs/tags/v41.tar.gz"
	saldl_cmd_tar := exec.Command("curl", "-L", saldl_tar_url, "-o", "package.tar.gz")
	err := saldl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	saldl_zip_url := "https://github.com/saldl/saldl/archive/refs/tags/v41.zip"
	saldl_cmd_zip := exec.Command("curl", "-L", saldl_zip_url, "-o", "package.zip")
	err = saldl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	saldl_bin_url := "https://github.com/saldl/saldl/archive/refs/tags/v41.bin"
	saldl_cmd_bin := exec.Command("curl", "-L", saldl_bin_url, "-o", "binary.bin")
	err = saldl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	saldl_src_url := "https://github.com/saldl/saldl/archive/refs/tags/v41.src.tar.gz"
	saldl_cmd_src := exec.Command("curl", "-L", saldl_src_url, "-o", "source.tar.gz")
	err = saldl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	saldl_cmd_direct := exec.Command("./binary")
	err = saldl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
	exec.Command("latte", "install", "asciidoc").Run()
	fmt.Println("Instalando dependencia: docbook-xsl")
	exec.Command("latte", "install", "docbook-xsl").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
}