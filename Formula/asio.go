package main

// Asio - Cross-platform C++ Library for asynchronous programming
// Homepage: https://think-async.com/Asio

import (
	"fmt"
	
	"os/exec"
)

func installAsio() {
	// Método 1: Descargar y extraer .tar.gz
	asio_tar_url := "https://downloads.sourceforge.net/project/asio/asio/1.30.2%20%28Stable%29/asio-1.30.2.tar.bz2"
	asio_cmd_tar := exec.Command("curl", "-L", asio_tar_url, "-o", "package.tar.gz")
	err := asio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asio_zip_url := "https://downloads.sourceforge.net/project/asio/asio/1.30.2%20%28Stable%29/asio-1.30.2.tar.bz2"
	asio_cmd_zip := exec.Command("curl", "-L", asio_zip_url, "-o", "package.zip")
	err = asio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asio_bin_url := "https://downloads.sourceforge.net/project/asio/asio/1.30.2%20%28Stable%29/asio-1.30.2.tar.bz2"
	asio_cmd_bin := exec.Command("curl", "-L", asio_bin_url, "-o", "binary.bin")
	err = asio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asio_src_url := "https://downloads.sourceforge.net/project/asio/asio/1.30.2%20%28Stable%29/asio-1.30.2.tar.bz2"
	asio_cmd_src := exec.Command("curl", "-L", asio_src_url, "-o", "source.tar.gz")
	err = asio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asio_cmd_direct := exec.Command("./binary")
	err = asio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}