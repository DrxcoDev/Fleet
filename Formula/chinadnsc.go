package main

// ChinadnsC - Port of ChinaDNS to C: fix irregularities with DNS in China
// Homepage: https://github.com/shadowsocks/ChinaDNS

import (
	"fmt"
	
	"os/exec"
)

func installChinadnsC() {
	// Método 1: Descargar y extraer .tar.gz
	chinadnsc_tar_url := "https://github.com/shadowsocks/ChinaDNS/releases/download/1.3.2/chinadns-1.3.2.tar.gz"
	chinadnsc_cmd_tar := exec.Command("curl", "-L", chinadnsc_tar_url, "-o", "package.tar.gz")
	err := chinadnsc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chinadnsc_zip_url := "https://github.com/shadowsocks/ChinaDNS/releases/download/1.3.2/chinadns-1.3.2.zip"
	chinadnsc_cmd_zip := exec.Command("curl", "-L", chinadnsc_zip_url, "-o", "package.zip")
	err = chinadnsc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chinadnsc_bin_url := "https://github.com/shadowsocks/ChinaDNS/releases/download/1.3.2/chinadns-1.3.2.bin"
	chinadnsc_cmd_bin := exec.Command("curl", "-L", chinadnsc_bin_url, "-o", "binary.bin")
	err = chinadnsc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chinadnsc_src_url := "https://github.com/shadowsocks/ChinaDNS/releases/download/1.3.2/chinadns-1.3.2.src.tar.gz"
	chinadnsc_cmd_src := exec.Command("curl", "-L", chinadnsc_src_url, "-o", "source.tar.gz")
	err = chinadnsc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chinadnsc_cmd_direct := exec.Command("./binary")
	err = chinadnsc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}