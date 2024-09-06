package main

// Speexdsp - Speex audio processing library
// Homepage: https://github.com/xiph/speexdsp

import (
	"fmt"
	
	"os/exec"
)

func installSpeexdsp() {
	// Método 1: Descargar y extraer .tar.gz
	speexdsp_tar_url := "https://github.com/xiph/speexdsp/archive/refs/tags/SpeexDSP-1.2.1.tar.gz"
	speexdsp_cmd_tar := exec.Command("curl", "-L", speexdsp_tar_url, "-o", "package.tar.gz")
	err := speexdsp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	speexdsp_zip_url := "https://github.com/xiph/speexdsp/archive/refs/tags/SpeexDSP-1.2.1.zip"
	speexdsp_cmd_zip := exec.Command("curl", "-L", speexdsp_zip_url, "-o", "package.zip")
	err = speexdsp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	speexdsp_bin_url := "https://github.com/xiph/speexdsp/archive/refs/tags/SpeexDSP-1.2.1.bin"
	speexdsp_cmd_bin := exec.Command("curl", "-L", speexdsp_bin_url, "-o", "binary.bin")
	err = speexdsp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	speexdsp_src_url := "https://github.com/xiph/speexdsp/archive/refs/tags/SpeexDSP-1.2.1.src.tar.gz"
	speexdsp_cmd_src := exec.Command("curl", "-L", speexdsp_src_url, "-o", "source.tar.gz")
	err = speexdsp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	speexdsp_cmd_direct := exec.Command("./binary")
	err = speexdsp_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}