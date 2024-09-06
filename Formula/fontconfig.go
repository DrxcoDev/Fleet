package main

// Fontconfig - XML-based font configuration API for X Windows
// Homepage: https://wiki.freedesktop.org/www/Software/fontconfig/

import (
	"fmt"
	
	"os/exec"
)

func installFontconfig() {
	// Método 1: Descargar y extraer .tar.gz
	fontconfig_tar_url := "https://www.freedesktop.org/software/fontconfig/release/fontconfig-2.15.0.tar.xz"
	fontconfig_cmd_tar := exec.Command("curl", "-L", fontconfig_tar_url, "-o", "package.tar.gz")
	err := fontconfig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fontconfig_zip_url := "https://www.freedesktop.org/software/fontconfig/release/fontconfig-2.15.0.tar.xz"
	fontconfig_cmd_zip := exec.Command("curl", "-L", fontconfig_zip_url, "-o", "package.zip")
	err = fontconfig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fontconfig_bin_url := "https://www.freedesktop.org/software/fontconfig/release/fontconfig-2.15.0.tar.xz"
	fontconfig_cmd_bin := exec.Command("curl", "-L", fontconfig_bin_url, "-o", "binary.bin")
	err = fontconfig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fontconfig_src_url := "https://www.freedesktop.org/software/fontconfig/release/fontconfig-2.15.0.tar.xz"
	fontconfig_cmd_src := exec.Command("curl", "-L", fontconfig_src_url, "-o", "source.tar.gz")
	err = fontconfig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fontconfig_cmd_direct := exec.Command("./binary")
	err = fontconfig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}