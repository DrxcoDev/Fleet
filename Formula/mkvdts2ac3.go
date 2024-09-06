package main

// Mkvdts2ac3 - Convert DTS audio to AC3 within a matroska file
// Homepage: https://github.com/JakeWharton/mkvdts2ac3

import (
	"fmt"
	
	"os/exec"
)

func installMkvdts2ac3() {
	// Método 1: Descargar y extraer .tar.gz
	mkvdts2ac3_tar_url := "https://github.com/JakeWharton/mkvdts2ac3/archive/refs/tags/1.6.0.tar.gz"
	mkvdts2ac3_cmd_tar := exec.Command("curl", "-L", mkvdts2ac3_tar_url, "-o", "package.tar.gz")
	err := mkvdts2ac3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mkvdts2ac3_zip_url := "https://github.com/JakeWharton/mkvdts2ac3/archive/refs/tags/1.6.0.zip"
	mkvdts2ac3_cmd_zip := exec.Command("curl", "-L", mkvdts2ac3_zip_url, "-o", "package.zip")
	err = mkvdts2ac3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mkvdts2ac3_bin_url := "https://github.com/JakeWharton/mkvdts2ac3/archive/refs/tags/1.6.0.bin"
	mkvdts2ac3_cmd_bin := exec.Command("curl", "-L", mkvdts2ac3_bin_url, "-o", "binary.bin")
	err = mkvdts2ac3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mkvdts2ac3_src_url := "https://github.com/JakeWharton/mkvdts2ac3/archive/refs/tags/1.6.0.src.tar.gz"
	mkvdts2ac3_cmd_src := exec.Command("curl", "-L", mkvdts2ac3_src_url, "-o", "source.tar.gz")
	err = mkvdts2ac3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mkvdts2ac3_cmd_direct := exec.Command("./binary")
	err = mkvdts2ac3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: mkvtoolnix")
	exec.Command("latte", "install", "mkvtoolnix").Run()
}