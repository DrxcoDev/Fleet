package main

// Uggconv - Universal Game Genie code converter
// Homepage: https://wyrmcorp.com/software/uggconv/index.shtml

import (
	"fmt"
	
	"os/exec"
)

func installUggconv() {
	// Método 1: Descargar y extraer .tar.gz
	uggconv_tar_url := "https://wyrmcorp.com/software/uggconv/uggconv-1.0.tar.gz"
	uggconv_cmd_tar := exec.Command("curl", "-L", uggconv_tar_url, "-o", "package.tar.gz")
	err := uggconv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uggconv_zip_url := "https://wyrmcorp.com/software/uggconv/uggconv-1.0.zip"
	uggconv_cmd_zip := exec.Command("curl", "-L", uggconv_zip_url, "-o", "package.zip")
	err = uggconv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uggconv_bin_url := "https://wyrmcorp.com/software/uggconv/uggconv-1.0.bin"
	uggconv_cmd_bin := exec.Command("curl", "-L", uggconv_bin_url, "-o", "binary.bin")
	err = uggconv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uggconv_src_url := "https://wyrmcorp.com/software/uggconv/uggconv-1.0.src.tar.gz"
	uggconv_cmd_src := exec.Command("curl", "-L", uggconv_src_url, "-o", "source.tar.gz")
	err = uggconv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uggconv_cmd_direct := exec.Command("./binary")
	err = uggconv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}