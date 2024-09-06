package main

// Libxp - X Print Client Library
// Homepage: https://gitlab.freedesktop.org/xorg/lib/libxp

import (
	"fmt"
	
	"os/exec"
)

func installLibxp() {
	// Método 1: Descargar y extraer .tar.gz
	libxp_tar_url := "https://gitlab.freedesktop.org/xorg/lib/libxp/-/archive/libXp-1.0.4/libxp-libXp-1.0.4.tar.bz2"
	libxp_cmd_tar := exec.Command("curl", "-L", libxp_tar_url, "-o", "package.tar.gz")
	err := libxp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxp_zip_url := "https://gitlab.freedesktop.org/xorg/lib/libxp/-/archive/libXp-1.0.4/libxp-libXp-1.0.4.tar.bz2"
	libxp_cmd_zip := exec.Command("curl", "-L", libxp_zip_url, "-o", "package.zip")
	err = libxp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxp_bin_url := "https://gitlab.freedesktop.org/xorg/lib/libxp/-/archive/libXp-1.0.4/libxp-libXp-1.0.4.tar.bz2"
	libxp_cmd_bin := exec.Command("curl", "-L", libxp_bin_url, "-o", "binary.bin")
	err = libxp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxp_src_url := "https://gitlab.freedesktop.org/xorg/lib/libxp/-/archive/libXp-1.0.4/libxp-libXp-1.0.4.tar.bz2"
	libxp_cmd_src := exec.Command("curl", "-L", libxp_src_url, "-o", "source.tar.gz")
	err = libxp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxp_cmd_direct := exec.Command("./binary")
	err = libxp_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: util-macros")
	exec.Command("latte", "install", "util-macros").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
}