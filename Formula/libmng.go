package main

// Libmng - MNG/JNG reference library
// Homepage: https://sourceforge.net/projects/libmng/

import (
	"fmt"
	
	"os/exec"
)

func installLibmng() {
	// Método 1: Descargar y extraer .tar.gz
	libmng_tar_url := "https://downloads.sourceforge.net/project/libmng/libmng-devel/2.0.3/libmng-2.0.3.tar.gz"
	libmng_cmd_tar := exec.Command("curl", "-L", libmng_tar_url, "-o", "package.tar.gz")
	err := libmng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmng_zip_url := "https://downloads.sourceforge.net/project/libmng/libmng-devel/2.0.3/libmng-2.0.3.zip"
	libmng_cmd_zip := exec.Command("curl", "-L", libmng_zip_url, "-o", "package.zip")
	err = libmng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmng_bin_url := "https://downloads.sourceforge.net/project/libmng/libmng-devel/2.0.3/libmng-2.0.3.bin"
	libmng_cmd_bin := exec.Command("curl", "-L", libmng_bin_url, "-o", "binary.bin")
	err = libmng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmng_src_url := "https://downloads.sourceforge.net/project/libmng/libmng-devel/2.0.3/libmng-2.0.3.src.tar.gz"
	libmng_cmd_src := exec.Command("curl", "-L", libmng_src_url, "-o", "source.tar.gz")
	err = libmng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmng_cmd_direct := exec.Command("./binary")
	err = libmng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
}