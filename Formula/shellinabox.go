package main

// Shellinabox - Export command-line tools to web based terminal emulator
// Homepage: https://github.com/shellinabox/shellinabox

import (
	"fmt"
	
	"os/exec"
)

func installShellinabox() {
	// Método 1: Descargar y extraer .tar.gz
	shellinabox_tar_url := "https://github.com/shellinabox/shellinabox/archive/refs/tags/v2.20.tar.gz"
	shellinabox_cmd_tar := exec.Command("curl", "-L", shellinabox_tar_url, "-o", "package.tar.gz")
	err := shellinabox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shellinabox_zip_url := "https://github.com/shellinabox/shellinabox/archive/refs/tags/v2.20.zip"
	shellinabox_cmd_zip := exec.Command("curl", "-L", shellinabox_zip_url, "-o", "package.zip")
	err = shellinabox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shellinabox_bin_url := "https://github.com/shellinabox/shellinabox/archive/refs/tags/v2.20.bin"
	shellinabox_cmd_bin := exec.Command("curl", "-L", shellinabox_bin_url, "-o", "binary.bin")
	err = shellinabox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shellinabox_src_url := "https://github.com/shellinabox/shellinabox/archive/refs/tags/v2.20.src.tar.gz"
	shellinabox_cmd_src := exec.Command("curl", "-L", shellinabox_src_url, "-o", "source.tar.gz")
	err = shellinabox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shellinabox_cmd_direct := exec.Command("./binary")
	err = shellinabox_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}