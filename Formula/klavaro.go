package main

// Klavaro - Free touch typing tutor program
// Homepage: https://klavaro.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installKlavaro() {
	// Método 1: Descargar y extraer .tar.gz
	klavaro_tar_url := "https://downloads.sourceforge.net/project/klavaro/klavaro-3.14.tar.bz2"
	klavaro_cmd_tar := exec.Command("curl", "-L", klavaro_tar_url, "-o", "package.tar.gz")
	err := klavaro_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	klavaro_zip_url := "https://downloads.sourceforge.net/project/klavaro/klavaro-3.14.tar.bz2"
	klavaro_cmd_zip := exec.Command("curl", "-L", klavaro_zip_url, "-o", "package.zip")
	err = klavaro_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	klavaro_bin_url := "https://downloads.sourceforge.net/project/klavaro/klavaro-3.14.tar.bz2"
	klavaro_cmd_bin := exec.Command("curl", "-L", klavaro_bin_url, "-o", "binary.bin")
	err = klavaro_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	klavaro_src_url := "https://downloads.sourceforge.net/project/klavaro/klavaro-3.14.tar.bz2"
	klavaro_cmd_src := exec.Command("curl", "-L", klavaro_src_url, "-o", "source.tar.gz")
	err = klavaro_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	klavaro_cmd_direct := exec.Command("./binary")
	err = klavaro_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
	exec.Command("latte", "install", "adwaita-icon-theme").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: gtkdatabox")
	exec.Command("latte", "install", "gtkdatabox").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: perl-xml-parser")
	exec.Command("latte", "install", "perl-xml-parser").Run()
}