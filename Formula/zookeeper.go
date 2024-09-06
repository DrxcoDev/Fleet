package main

// Zookeeper - Centralized server for distributed coordination of services
// Homepage: https://zookeeper.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installZookeeper() {
	// Método 1: Descargar y extraer .tar.gz
	zookeeper_tar_url := "https://www.apache.org/dyn/closer.lua?path=zookeeper/zookeeper-3.9.2/apache-zookeeper-3.9.2.tar.gz"
	zookeeper_cmd_tar := exec.Command("curl", "-L", zookeeper_tar_url, "-o", "package.tar.gz")
	err := zookeeper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zookeeper_zip_url := "https://www.apache.org/dyn/closer.lua?path=zookeeper/zookeeper-3.9.2/apache-zookeeper-3.9.2.zip"
	zookeeper_cmd_zip := exec.Command("curl", "-L", zookeeper_zip_url, "-o", "package.zip")
	err = zookeeper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zookeeper_bin_url := "https://www.apache.org/dyn/closer.lua?path=zookeeper/zookeeper-3.9.2/apache-zookeeper-3.9.2.bin"
	zookeeper_cmd_bin := exec.Command("curl", "-L", zookeeper_bin_url, "-o", "binary.bin")
	err = zookeeper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zookeeper_src_url := "https://www.apache.org/dyn/closer.lua?path=zookeeper/zookeeper-3.9.2/apache-zookeeper-3.9.2.src.tar.gz"
	zookeeper_cmd_src := exec.Command("curl", "-L", zookeeper_src_url, "-o", "source.tar.gz")
	err = zookeeper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zookeeper_cmd_direct := exec.Command("./binary")
	err = zookeeper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: cppunit")
	exec.Command("latte", "install", "cppunit").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: maven")
	exec.Command("latte", "install", "maven").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}