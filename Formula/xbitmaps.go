package main

// Xbitmaps - Bitmap images used by multiple X11 applications
// Homepage: https://xcb.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installXbitmaps() {
	// Método 1: Descargar y extraer .tar.gz
	xbitmaps_tar_url := "https://xorg.freedesktop.org/archive/individual/data/xbitmaps-1.1.3.tar.xz"
	xbitmaps_cmd_tar := exec.Command("curl", "-L", xbitmaps_tar_url, "-o", "package.tar.gz")
	err := xbitmaps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xbitmaps_zip_url := "https://xorg.freedesktop.org/archive/individual/data/xbitmaps-1.1.3.tar.xz"
	xbitmaps_cmd_zip := exec.Command("curl", "-L", xbitmaps_zip_url, "-o", "package.zip")
	err = xbitmaps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xbitmaps_bin_url := "https://xorg.freedesktop.org/archive/individual/data/xbitmaps-1.1.3.tar.xz"
	xbitmaps_cmd_bin := exec.Command("curl", "-L", xbitmaps_bin_url, "-o", "binary.bin")
	err = xbitmaps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xbitmaps_src_url := "https://xorg.freedesktop.org/archive/individual/data/xbitmaps-1.1.3.tar.xz"
	xbitmaps_cmd_src := exec.Command("curl", "-L", xbitmaps_src_url, "-o", "source.tar.gz")
	err = xbitmaps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xbitmaps_cmd_direct := exec.Command("./binary")
	err = xbitmaps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: util-macros")
	exec.Command("latte", "install", "util-macros").Run()
}