package main

// LibbitcoinExplorer - Bitcoin command-line tool
// Homepage: https://github.com/libbitcoin/libbitcoin-explorer

import (
	"fmt"
	
	"os/exec"
)

func installLibbitcoinExplorer() {
	// Método 1: Descargar y extraer .tar.gz
	libbitcoinexplorer_tar_url := "https://github.com/libbitcoin/libbitcoin-explorer/archive/refs/tags/v3.8.0.tar.gz"
	libbitcoinexplorer_cmd_tar := exec.Command("curl", "-L", libbitcoinexplorer_tar_url, "-o", "package.tar.gz")
	err := libbitcoinexplorer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbitcoinexplorer_zip_url := "https://github.com/libbitcoin/libbitcoin-explorer/archive/refs/tags/v3.8.0.zip"
	libbitcoinexplorer_cmd_zip := exec.Command("curl", "-L", libbitcoinexplorer_zip_url, "-o", "package.zip")
	err = libbitcoinexplorer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbitcoinexplorer_bin_url := "https://github.com/libbitcoin/libbitcoin-explorer/archive/refs/tags/v3.8.0.bin"
	libbitcoinexplorer_cmd_bin := exec.Command("curl", "-L", libbitcoinexplorer_bin_url, "-o", "binary.bin")
	err = libbitcoinexplorer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbitcoinexplorer_src_url := "https://github.com/libbitcoin/libbitcoin-explorer/archive/refs/tags/v3.8.0.src.tar.gz"
	libbitcoinexplorer_cmd_src := exec.Command("curl", "-L", libbitcoinexplorer_src_url, "-o", "source.tar.gz")
	err = libbitcoinexplorer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbitcoinexplorer_cmd_direct := exec.Command("./binary")
	err = libbitcoinexplorer_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libbitcoin-client")
	exec.Command("latte", "install", "libbitcoin-client").Run()
	fmt.Println("Instalando dependencia: libbitcoin-network")
	exec.Command("latte", "install", "libbitcoin-network").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
}