package main

// Rfcstrip - Strips headers and footers from RFCs and Internet-Drafts
// Homepage: https://github.com/mbj4668/rfcstrip

import (
	"fmt"
	
	"os/exec"
)

func installRfcstrip() {
	// Método 1: Descargar y extraer .tar.gz
	rfcstrip_tar_url := "https://github.com/mbj4668/rfcstrip/archive/refs/tags/1.3.tar.gz"
	rfcstrip_cmd_tar := exec.Command("curl", "-L", rfcstrip_tar_url, "-o", "package.tar.gz")
	err := rfcstrip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rfcstrip_zip_url := "https://github.com/mbj4668/rfcstrip/archive/refs/tags/1.3.zip"
	rfcstrip_cmd_zip := exec.Command("curl", "-L", rfcstrip_zip_url, "-o", "package.zip")
	err = rfcstrip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rfcstrip_bin_url := "https://github.com/mbj4668/rfcstrip/archive/refs/tags/1.3.bin"
	rfcstrip_cmd_bin := exec.Command("curl", "-L", rfcstrip_bin_url, "-o", "binary.bin")
	err = rfcstrip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rfcstrip_src_url := "https://github.com/mbj4668/rfcstrip/archive/refs/tags/1.3.src.tar.gz"
	rfcstrip_cmd_src := exec.Command("curl", "-L", rfcstrip_src_url, "-o", "source.tar.gz")
	err = rfcstrip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rfcstrip_cmd_direct := exec.Command("./binary")
	err = rfcstrip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}