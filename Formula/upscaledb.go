package main

// Upscaledb - Database for embedded devices
// Homepage: https://upscaledb.com/

import (
	"fmt"
	
	"os/exec"
)

func installUpscaledb() {
	// Método 1: Descargar y extraer .tar.gz
	upscaledb_tar_url := "https://github.com/cruppstahl/upscaledb.git"
	upscaledb_cmd_tar := exec.Command("curl", "-L", upscaledb_tar_url, "-o", "package.tar.gz")
	err := upscaledb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	upscaledb_zip_url := "https://github.com/cruppstahl/upscaledb.git"
	upscaledb_cmd_zip := exec.Command("curl", "-L", upscaledb_zip_url, "-o", "package.zip")
	err = upscaledb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	upscaledb_bin_url := "https://github.com/cruppstahl/upscaledb.git"
	upscaledb_cmd_bin := exec.Command("curl", "-L", upscaledb_bin_url, "-o", "binary.bin")
	err = upscaledb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	upscaledb_src_url := "https://github.com/cruppstahl/upscaledb.git"
	upscaledb_cmd_src := exec.Command("curl", "-L", upscaledb_src_url, "-o", "source.tar.gz")
	err = upscaledb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	upscaledb_cmd_direct := exec.Command("./binary")
	err = upscaledb_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}