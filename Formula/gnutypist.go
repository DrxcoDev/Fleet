package main

// GnuTypist - GNU typing tutor
// Homepage: https://www.gnu.org/software/gtypist/

import (
	"fmt"
	
	"os/exec"
)

func installGnuTypist() {
	// Método 1: Descargar y extraer .tar.gz
	gnutypist_tar_url := "https://ftp.gnu.org/gnu/gtypist/gtypist-2.9.5.tar.xz"
	gnutypist_cmd_tar := exec.Command("curl", "-L", gnutypist_tar_url, "-o", "package.tar.gz")
	err := gnutypist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnutypist_zip_url := "https://ftp.gnu.org/gnu/gtypist/gtypist-2.9.5.tar.xz"
	gnutypist_cmd_zip := exec.Command("curl", "-L", gnutypist_zip_url, "-o", "package.zip")
	err = gnutypist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnutypist_bin_url := "https://ftp.gnu.org/gnu/gtypist/gtypist-2.9.5.tar.xz"
	gnutypist_cmd_bin := exec.Command("curl", "-L", gnutypist_bin_url, "-o", "binary.bin")
	err = gnutypist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnutypist_src_url := "https://ftp.gnu.org/gnu/gtypist/gtypist-2.9.5.tar.xz"
	gnutypist_cmd_src := exec.Command("curl", "-L", gnutypist_src_url, "-o", "source.tar.gz")
	err = gnutypist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnutypist_cmd_direct := exec.Command("./binary")
	err = gnutypist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}