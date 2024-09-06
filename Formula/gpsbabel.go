package main

// Gpsbabel - Converts/uploads GPS waypoints, tracks, and routes
// Homepage: https://www.gpsbabel.org/

import (
	"fmt"
	
	"os/exec"
)

func installGpsbabel() {
	// Método 1: Descargar y extraer .tar.gz
	gpsbabel_tar_url := "https://github.com/GPSBabel/gpsbabel/archive/refs/tags/gpsbabel_1_9_0.tar.gz"
	gpsbabel_cmd_tar := exec.Command("curl", "-L", gpsbabel_tar_url, "-o", "package.tar.gz")
	err := gpsbabel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpsbabel_zip_url := "https://github.com/GPSBabel/gpsbabel/archive/refs/tags/gpsbabel_1_9_0.zip"
	gpsbabel_cmd_zip := exec.Command("curl", "-L", gpsbabel_zip_url, "-o", "package.zip")
	err = gpsbabel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpsbabel_bin_url := "https://github.com/GPSBabel/gpsbabel/archive/refs/tags/gpsbabel_1_9_0.bin"
	gpsbabel_cmd_bin := exec.Command("curl", "-L", gpsbabel_bin_url, "-o", "binary.bin")
	err = gpsbabel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpsbabel_src_url := "https://github.com/GPSBabel/gpsbabel/archive/refs/tags/gpsbabel_1_9_0.src.tar.gz"
	gpsbabel_cmd_src := exec.Command("curl", "-L", gpsbabel_src_url, "-o", "source.tar.gz")
	err = gpsbabel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpsbabel_cmd_direct := exec.Command("./binary")
	err = gpsbabel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: shapelib")
	exec.Command("latte", "install", "shapelib").Run()
}