package main

// Freeswitch - Telephony platform to route various communication protocols
// Homepage: https://freeswitch.org

import (
	"fmt"
	
	"os/exec"
)

func installFreeswitch() {
	// Método 1: Descargar y extraer .tar.gz
	freeswitch_tar_url := "https://github.com/signalwire/freeswitch.git"
	freeswitch_cmd_tar := exec.Command("curl", "-L", freeswitch_tar_url, "-o", "package.tar.gz")
	err := freeswitch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freeswitch_zip_url := "https://github.com/signalwire/freeswitch.git"
	freeswitch_cmd_zip := exec.Command("curl", "-L", freeswitch_zip_url, "-o", "package.zip")
	err = freeswitch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freeswitch_bin_url := "https://github.com/signalwire/freeswitch.git"
	freeswitch_cmd_bin := exec.Command("curl", "-L", freeswitch_bin_url, "-o", "binary.bin")
	err = freeswitch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freeswitch_src_url := "https://github.com/signalwire/freeswitch.git"
	freeswitch_cmd_src := exec.Command("curl", "-L", freeswitch_src_url, "-o", "source.tar.gz")
	err = freeswitch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freeswitch_cmd_direct := exec.Command("./binary")
	err = freeswitch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: yasm")
	exec.Command("latte", "install", "yasm").Run()
	fmt.Println("Instalando dependencia: ffmpeg@5")
	exec.Command("latte", "install", "ffmpeg@5").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: ldns")
	exec.Command("latte", "install", "ldns").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: opencore-amr")
	exec.Command("latte", "install", "opencore-amr").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: sofia-sip")
	exec.Command("latte", "install", "sofia-sip").Run()
	fmt.Println("Instalando dependencia: speex")
	exec.Command("latte", "install", "speex").Run()
	fmt.Println("Instalando dependencia: speexdsp")
	exec.Command("latte", "install", "speexdsp").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}