package main

// Urlview - URL extractor/launcher
// Homepage: https://packages.debian.org/sid/misc/urlview

import (
	"fmt"
	
	"os/exec"
)

func installUrlview() {
	// Método 1: Descargar y extraer .tar.gz
	urlview_tar_url := "https://deb.debian.org/debian/pool/main/u/urlview/urlview_0.9.orig.tar.gz"
	urlview_cmd_tar := exec.Command("curl", "-L", urlview_tar_url, "-o", "package.tar.gz")
	err := urlview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	urlview_zip_url := "https://deb.debian.org/debian/pool/main/u/urlview/urlview_0.9.orig.zip"
	urlview_cmd_zip := exec.Command("curl", "-L", urlview_zip_url, "-o", "package.zip")
	err = urlview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	urlview_bin_url := "https://deb.debian.org/debian/pool/main/u/urlview/urlview_0.9.orig.bin"
	urlview_cmd_bin := exec.Command("curl", "-L", urlview_bin_url, "-o", "binary.bin")
	err = urlview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	urlview_src_url := "https://deb.debian.org/debian/pool/main/u/urlview/urlview_0.9.orig.src.tar.gz"
	urlview_cmd_src := exec.Command("curl", "-L", urlview_src_url, "-o", "source.tar.gz")
	err = urlview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	urlview_cmd_direct := exec.Command("./binary")
	err = urlview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}