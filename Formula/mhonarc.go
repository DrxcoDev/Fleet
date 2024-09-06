package main

// Mhonarc - Mail-to-HTML converter
// Homepage: https://www.mhonarc.org/

import (
	"fmt"
	
	"os/exec"
)

func installMhonarc() {
	// Método 1: Descargar y extraer .tar.gz
	mhonarc_tar_url := "https://www.mhonarc.org/release/MHonArc/tar/MHonArc-2.6.19.tar.bz2"
	mhonarc_cmd_tar := exec.Command("curl", "-L", mhonarc_tar_url, "-o", "package.tar.gz")
	err := mhonarc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mhonarc_zip_url := "https://www.mhonarc.org/release/MHonArc/tar/MHonArc-2.6.19.tar.bz2"
	mhonarc_cmd_zip := exec.Command("curl", "-L", mhonarc_zip_url, "-o", "package.zip")
	err = mhonarc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mhonarc_bin_url := "https://www.mhonarc.org/release/MHonArc/tar/MHonArc-2.6.19.tar.bz2"
	mhonarc_cmd_bin := exec.Command("curl", "-L", mhonarc_bin_url, "-o", "binary.bin")
	err = mhonarc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mhonarc_src_url := "https://www.mhonarc.org/release/MHonArc/tar/MHonArc-2.6.19.tar.bz2"
	mhonarc_cmd_src := exec.Command("curl", "-L", mhonarc_src_url, "-o", "source.tar.gz")
	err = mhonarc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mhonarc_cmd_direct := exec.Command("./binary")
	err = mhonarc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: perl")
	exec.Command("latte", "install", "perl").Run()
}