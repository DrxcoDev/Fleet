package main

// Torsocks - Use SOCKS-friendly applications with Tor
// Homepage: https://gitlab.torproject.org/tpo/core/torsocks

import (
	"fmt"
	
	"os/exec"
)

func installTorsocks() {
	// Método 1: Descargar y extraer .tar.gz
	torsocks_tar_url := "https://gitlab.torproject.org/tpo/core/torsocks/-/archive/v2.4.0/torsocks-v2.4.0.tar.bz2"
	torsocks_cmd_tar := exec.Command("curl", "-L", torsocks_tar_url, "-o", "package.tar.gz")
	err := torsocks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	torsocks_zip_url := "https://gitlab.torproject.org/tpo/core/torsocks/-/archive/v2.4.0/torsocks-v2.4.0.tar.bz2"
	torsocks_cmd_zip := exec.Command("curl", "-L", torsocks_zip_url, "-o", "package.zip")
	err = torsocks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	torsocks_bin_url := "https://gitlab.torproject.org/tpo/core/torsocks/-/archive/v2.4.0/torsocks-v2.4.0.tar.bz2"
	torsocks_cmd_bin := exec.Command("curl", "-L", torsocks_bin_url, "-o", "binary.bin")
	err = torsocks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	torsocks_src_url := "https://gitlab.torproject.org/tpo/core/torsocks/-/archive/v2.4.0/torsocks-v2.4.0.tar.bz2"
	torsocks_cmd_src := exec.Command("curl", "-L", torsocks_src_url, "-o", "source.tar.gz")
	err = torsocks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	torsocks_cmd_direct := exec.Command("./binary")
	err = torsocks_cmd_direct.Run()
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
}