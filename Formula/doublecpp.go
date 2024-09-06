package main

// Doublecpp - Double dispatch in C++
// Homepage: https://doublecpp.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDoublecpp() {
	// Método 1: Descargar y extraer .tar.gz
	doublecpp_tar_url := "https://downloads.sourceforge.net/project/doublecpp/doublecpp/0.6.3/doublecpp-0.6.3.tar.gz"
	doublecpp_cmd_tar := exec.Command("curl", "-L", doublecpp_tar_url, "-o", "package.tar.gz")
	err := doublecpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	doublecpp_zip_url := "https://downloads.sourceforge.net/project/doublecpp/doublecpp/0.6.3/doublecpp-0.6.3.zip"
	doublecpp_cmd_zip := exec.Command("curl", "-L", doublecpp_zip_url, "-o", "package.zip")
	err = doublecpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	doublecpp_bin_url := "https://downloads.sourceforge.net/project/doublecpp/doublecpp/0.6.3/doublecpp-0.6.3.bin"
	doublecpp_cmd_bin := exec.Command("curl", "-L", doublecpp_bin_url, "-o", "binary.bin")
	err = doublecpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	doublecpp_src_url := "https://downloads.sourceforge.net/project/doublecpp/doublecpp/0.6.3/doublecpp-0.6.3.src.tar.gz"
	doublecpp_cmd_src := exec.Command("curl", "-L", doublecpp_src_url, "-o", "source.tar.gz")
	err = doublecpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	doublecpp_cmd_direct := exec.Command("./binary")
	err = doublecpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}