package main

// Espflash - Serial flasher utility for Espressif SoCs and modules based on esptool.py
// Homepage: https://github.com/esp-rs/espflash

import (
	"fmt"
	
	"os/exec"
)

func installEspflash() {
	// Método 1: Descargar y extraer .tar.gz
	espflash_tar_url := "https://github.com/esp-rs/espflash/archive/refs/tags/v3.1.1.tar.gz"
	espflash_cmd_tar := exec.Command("curl", "-L", espflash_tar_url, "-o", "package.tar.gz")
	err := espflash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	espflash_zip_url := "https://github.com/esp-rs/espflash/archive/refs/tags/v3.1.1.zip"
	espflash_cmd_zip := exec.Command("curl", "-L", espflash_zip_url, "-o", "package.zip")
	err = espflash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	espflash_bin_url := "https://github.com/esp-rs/espflash/archive/refs/tags/v3.1.1.bin"
	espflash_cmd_bin := exec.Command("curl", "-L", espflash_bin_url, "-o", "binary.bin")
	err = espflash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	espflash_src_url := "https://github.com/esp-rs/espflash/archive/refs/tags/v3.1.1.src.tar.gz"
	espflash_cmd_src := exec.Command("curl", "-L", espflash_src_url, "-o", "source.tar.gz")
	err = espflash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	espflash_cmd_direct := exec.Command("./binary")
	err = espflash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}