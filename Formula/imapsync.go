package main

// Imapsync - Migrate or backup IMAP mail accounts
// Homepage: https://imapsync.lamiral.info/

import (
	"fmt"
	
	"os/exec"
)

func installImapsync() {
	// Método 1: Descargar y extraer .tar.gz
	imapsync_tar_url := "https://imapsync.lamiral.info/dist2/imapsync-2.290.tgz"
	imapsync_cmd_tar := exec.Command("curl", "-L", imapsync_tar_url, "-o", "package.tar.gz")
	err := imapsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imapsync_zip_url := "https://imapsync.lamiral.info/dist2/imapsync-2.290.tgz"
	imapsync_cmd_zip := exec.Command("curl", "-L", imapsync_zip_url, "-o", "package.zip")
	err = imapsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imapsync_bin_url := "https://imapsync.lamiral.info/dist2/imapsync-2.290.tgz"
	imapsync_cmd_bin := exec.Command("curl", "-L", imapsync_bin_url, "-o", "binary.bin")
	err = imapsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imapsync_src_url := "https://imapsync.lamiral.info/dist2/imapsync-2.290.tgz"
	imapsync_cmd_src := exec.Command("curl", "-L", imapsync_src_url, "-o", "source.tar.gz")
	err = imapsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imapsync_cmd_direct := exec.Command("./binary")
	err = imapsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pod2man")
	exec.Command("latte", "install", "pod2man").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}