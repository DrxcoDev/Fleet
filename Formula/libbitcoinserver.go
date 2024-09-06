package main

// LibbitcoinServer - Bitcoin Full Node and Query Server
// Homepage: https://github.com/libbitcoin/libbitcoin-server

import (
	"fmt"
	
	"os/exec"
)

func installLibbitcoinServer() {
	// Método 1: Descargar y extraer .tar.gz
	libbitcoinserver_tar_url := "https://github.com/libbitcoin/libbitcoin-server/archive/refs/tags/v3.8.0.tar.gz"
	libbitcoinserver_cmd_tar := exec.Command("curl", "-L", libbitcoinserver_tar_url, "-o", "package.tar.gz")
	err := libbitcoinserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbitcoinserver_zip_url := "https://github.com/libbitcoin/libbitcoin-server/archive/refs/tags/v3.8.0.zip"
	libbitcoinserver_cmd_zip := exec.Command("curl", "-L", libbitcoinserver_zip_url, "-o", "package.zip")
	err = libbitcoinserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbitcoinserver_bin_url := "https://github.com/libbitcoin/libbitcoin-server/archive/refs/tags/v3.8.0.bin"
	libbitcoinserver_cmd_bin := exec.Command("curl", "-L", libbitcoinserver_bin_url, "-o", "binary.bin")
	err = libbitcoinserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbitcoinserver_src_url := "https://github.com/libbitcoin/libbitcoin-server/archive/refs/tags/v3.8.0.src.tar.gz"
	libbitcoinserver_cmd_src := exec.Command("curl", "-L", libbitcoinserver_src_url, "-o", "source.tar.gz")
	err = libbitcoinserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbitcoinserver_cmd_direct := exec.Command("./binary")
	err = libbitcoinserver_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: boost@1.76")
	exec.Command("latte", "install", "boost@1.76").Run()
	fmt.Println("Instalando dependencia: libbitcoin-node")
	exec.Command("latte", "install", "libbitcoin-node").Run()
	fmt.Println("Instalando dependencia: libbitcoin-protocol")
	exec.Command("latte", "install", "libbitcoin-protocol").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
}