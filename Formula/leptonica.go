package main

// Leptonica - Image processing and image analysis library
// Homepage: http://www.leptonica.org/

import (
	"fmt"
	
	"os/exec"
)

func installLeptonica() {
	// Método 1: Descargar y extraer .tar.gz
	leptonica_tar_url := "https://github.com/DanBloomberg/leptonica/releases/download/1.84.1/leptonica-1.84.1.tar.gz"
	leptonica_cmd_tar := exec.Command("curl", "-L", leptonica_tar_url, "-o", "package.tar.gz")
	err := leptonica_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leptonica_zip_url := "https://github.com/DanBloomberg/leptonica/releases/download/1.84.1/leptonica-1.84.1.zip"
	leptonica_cmd_zip := exec.Command("curl", "-L", leptonica_zip_url, "-o", "package.zip")
	err = leptonica_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leptonica_bin_url := "https://github.com/DanBloomberg/leptonica/releases/download/1.84.1/leptonica-1.84.1.bin"
	leptonica_cmd_bin := exec.Command("curl", "-L", leptonica_bin_url, "-o", "binary.bin")
	err = leptonica_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leptonica_src_url := "https://github.com/DanBloomberg/leptonica/releases/download/1.84.1/leptonica-1.84.1.src.tar.gz"
	leptonica_cmd_src := exec.Command("curl", "-L", leptonica_src_url, "-o", "source.tar.gz")
	err = leptonica_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leptonica_cmd_direct := exec.Command("./binary")
	err = leptonica_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: openjpeg")
	exec.Command("latte", "install", "openjpeg").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
}