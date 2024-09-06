package main

// Binutils - GNU binary tools for native development
// Homepage: https://www.gnu.org/software/binutils/binutils.html

import (
	"fmt"
	
	"os/exec"
)

func installBinutils() {
	// Método 1: Descargar y extraer .tar.gz
	binutils_tar_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	binutils_cmd_tar := exec.Command("curl", "-L", binutils_tar_url, "-o", "package.tar.gz")
	err := binutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	binutils_zip_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	binutils_cmd_zip := exec.Command("curl", "-L", binutils_zip_url, "-o", "package.zip")
	err = binutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	binutils_bin_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	binutils_cmd_bin := exec.Command("curl", "-L", binutils_bin_url, "-o", "binary.bin")
	err = binutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	binutils_src_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	binutils_cmd_src := exec.Command("curl", "-L", binutils_src_url, "-o", "source.tar.gz")
	err = binutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	binutils_cmd_direct := exec.Command("./binary")
	err = binutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}