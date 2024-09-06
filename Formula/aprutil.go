package main

// AprUtil - Companion library to apr, the Apache Portable Runtime library
// Homepage: https://apr.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installAprUtil() {
	// Método 1: Descargar y extraer .tar.gz
	aprutil_tar_url := "https://www.apache.org/dyn/closer.lua?path=apr/apr-util-1.6.3.tar.bz2"
	aprutil_cmd_tar := exec.Command("curl", "-L", aprutil_tar_url, "-o", "package.tar.gz")
	err := aprutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aprutil_zip_url := "https://www.apache.org/dyn/closer.lua?path=apr/apr-util-1.6.3.tar.bz2"
	aprutil_cmd_zip := exec.Command("curl", "-L", aprutil_zip_url, "-o", "package.zip")
	err = aprutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aprutil_bin_url := "https://www.apache.org/dyn/closer.lua?path=apr/apr-util-1.6.3.tar.bz2"
	aprutil_cmd_bin := exec.Command("curl", "-L", aprutil_bin_url, "-o", "binary.bin")
	err = aprutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aprutil_src_url := "https://www.apache.org/dyn/closer.lua?path=apr/apr-util-1.6.3.tar.bz2"
	aprutil_cmd_src := exec.Command("curl", "-L", aprutil_src_url, "-o", "source.tar.gz")
	err = aprutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aprutil_cmd_direct := exec.Command("./binary")
	err = aprutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: apr")
	exec.Command("latte", "install", "apr").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: mawk")
	exec.Command("latte", "install", "mawk").Run()
	fmt.Println("Instalando dependencia: unixodbc")
	exec.Command("latte", "install", "unixodbc").Run()
}