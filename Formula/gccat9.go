package main

// GccAT9 - GNU compiler collection
// Homepage: https://gcc.gnu.org/

import (
	"fmt"
	
	"os/exec"
)

func installGccAT9() {
	// Método 1: Descargar y extraer .tar.gz
	gccat9_tar_url := "https://ftp.gnu.org/gnu/gcc/gcc-9.5.0/gcc-9.5.0.tar.xz"
	gccat9_cmd_tar := exec.Command("curl", "-L", gccat9_tar_url, "-o", "package.tar.gz")
	err := gccat9_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gccat9_zip_url := "https://ftp.gnu.org/gnu/gcc/gcc-9.5.0/gcc-9.5.0.tar.xz"
	gccat9_cmd_zip := exec.Command("curl", "-L", gccat9_zip_url, "-o", "package.zip")
	err = gccat9_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gccat9_bin_url := "https://ftp.gnu.org/gnu/gcc/gcc-9.5.0/gcc-9.5.0.tar.xz"
	gccat9_cmd_bin := exec.Command("curl", "-L", gccat9_bin_url, "-o", "binary.bin")
	err = gccat9_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gccat9_src_url := "https://ftp.gnu.org/gnu/gcc/gcc-9.5.0/gcc-9.5.0.tar.xz"
	gccat9_cmd_src := exec.Command("curl", "-L", gccat9_src_url, "-o", "source.tar.gz")
	err = gccat9_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gccat9_cmd_direct := exec.Command("./binary")
	err = gccat9_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: isl")
	exec.Command("latte", "install", "isl").Run()
	fmt.Println("Instalando dependencia: libmpc")
	exec.Command("latte", "install", "libmpc").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
	fmt.Println("Instalando dependencia: binutils")
	exec.Command("latte", "install", "binutils").Run()
}