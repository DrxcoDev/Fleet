package main

// Rtaudio - API for realtime audio input/output
// Homepage: https://www.music.mcgill.ca/~gary/rtaudio/

import (
	"fmt"
	
	"os/exec"
)

func installRtaudio() {
	// Método 1: Descargar y extraer .tar.gz
	rtaudio_tar_url := "https://www.music.mcgill.ca/~gary/rtaudio/release/rtaudio-6.0.1.tar.gz"
	rtaudio_cmd_tar := exec.Command("curl", "-L", rtaudio_tar_url, "-o", "package.tar.gz")
	err := rtaudio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rtaudio_zip_url := "https://www.music.mcgill.ca/~gary/rtaudio/release/rtaudio-6.0.1.zip"
	rtaudio_cmd_zip := exec.Command("curl", "-L", rtaudio_zip_url, "-o", "package.zip")
	err = rtaudio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rtaudio_bin_url := "https://www.music.mcgill.ca/~gary/rtaudio/release/rtaudio-6.0.1.bin"
	rtaudio_cmd_bin := exec.Command("curl", "-L", rtaudio_bin_url, "-o", "binary.bin")
	err = rtaudio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rtaudio_src_url := "https://www.music.mcgill.ca/~gary/rtaudio/release/rtaudio-6.0.1.src.tar.gz"
	rtaudio_cmd_src := exec.Command("curl", "-L", rtaudio_src_url, "-o", "source.tar.gz")
	err = rtaudio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rtaudio_cmd_direct := exec.Command("./binary")
	err = rtaudio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}