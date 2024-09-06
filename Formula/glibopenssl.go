package main

// GlibOpenssl - OpenSSL GIO module for glib
// Homepage: https://launchpad.net/glib-networking

import (
	"fmt"
	
	"os/exec"
)

func installGlibOpenssl() {
	// Método 1: Descargar y extraer .tar.gz
	glibopenssl_tar_url := "https://download.gnome.org/sources/glib-openssl/2.50/glib-openssl-2.50.8.tar.xz"
	glibopenssl_cmd_tar := exec.Command("curl", "-L", glibopenssl_tar_url, "-o", "package.tar.gz")
	err := glibopenssl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glibopenssl_zip_url := "https://download.gnome.org/sources/glib-openssl/2.50/glib-openssl-2.50.8.tar.xz"
	glibopenssl_cmd_zip := exec.Command("curl", "-L", glibopenssl_zip_url, "-o", "package.zip")
	err = glibopenssl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glibopenssl_bin_url := "https://download.gnome.org/sources/glib-openssl/2.50/glib-openssl-2.50.8.tar.xz"
	glibopenssl_cmd_bin := exec.Command("curl", "-L", glibopenssl_bin_url, "-o", "binary.bin")
	err = glibopenssl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glibopenssl_src_url := "https://download.gnome.org/sources/glib-openssl/2.50/glib-openssl-2.50.8.tar.xz"
	glibopenssl_cmd_src := exec.Command("curl", "-L", glibopenssl_src_url, "-o", "source.tar.gz")
	err = glibopenssl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glibopenssl_cmd_direct := exec.Command("./binary")
	err = glibopenssl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: openssl@1.1")
	exec.Command("latte", "install", "openssl@1.1").Run()
}