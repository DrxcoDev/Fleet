package main

// Libusbmuxd - USB multiplexor library for iOS devices
// Homepage: https://www.libimobiledevice.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibusbmuxd() {
	// Método 1: Descargar y extraer .tar.gz
	libusbmuxd_tar_url := "https://github.com/libimobiledevice/libusbmuxd/releases/download/2.1.0/libusbmuxd-2.1.0.tar.bz2"
	libusbmuxd_cmd_tar := exec.Command("curl", "-L", libusbmuxd_tar_url, "-o", "package.tar.gz")
	err := libusbmuxd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libusbmuxd_zip_url := "https://github.com/libimobiledevice/libusbmuxd/releases/download/2.1.0/libusbmuxd-2.1.0.tar.bz2"
	libusbmuxd_cmd_zip := exec.Command("curl", "-L", libusbmuxd_zip_url, "-o", "package.zip")
	err = libusbmuxd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libusbmuxd_bin_url := "https://github.com/libimobiledevice/libusbmuxd/releases/download/2.1.0/libusbmuxd-2.1.0.tar.bz2"
	libusbmuxd_cmd_bin := exec.Command("curl", "-L", libusbmuxd_bin_url, "-o", "binary.bin")
	err = libusbmuxd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libusbmuxd_src_url := "https://github.com/libimobiledevice/libusbmuxd/releases/download/2.1.0/libusbmuxd-2.1.0.tar.bz2"
	libusbmuxd_cmd_src := exec.Command("curl", "-L", libusbmuxd_src_url, "-o", "source.tar.gz")
	err = libusbmuxd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libusbmuxd_cmd_direct := exec.Command("./binary")
	err = libusbmuxd_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libimobiledevice-glue")
	exec.Command("latte", "install", "libimobiledevice-glue").Run()
	fmt.Println("Instalando dependencia: libplist")
	exec.Command("latte", "install", "libplist").Run()
}