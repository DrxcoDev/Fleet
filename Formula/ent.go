package main

// Ent - Pseudorandom number sequence test program
// Homepage: https://www.fourmilab.ch/random/

import (
	"fmt"
	
	"os/exec"
)

func installEnt() {
	// Método 1: Descargar y extraer .tar.gz
	ent_tar_url := "https://github.com/psm14/ent/archive/refs/tags/1.0.tar.gz"
	ent_cmd_tar := exec.Command("curl", "-L", ent_tar_url, "-o", "package.tar.gz")
	err := ent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ent_zip_url := "https://github.com/psm14/ent/archive/refs/tags/1.0.zip"
	ent_cmd_zip := exec.Command("curl", "-L", ent_zip_url, "-o", "package.zip")
	err = ent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ent_bin_url := "https://github.com/psm14/ent/archive/refs/tags/1.0.bin"
	ent_cmd_bin := exec.Command("curl", "-L", ent_bin_url, "-o", "binary.bin")
	err = ent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ent_src_url := "https://github.com/psm14/ent/archive/refs/tags/1.0.src.tar.gz"
	ent_cmd_src := exec.Command("curl", "-L", ent_src_url, "-o", "source.tar.gz")
	err = ent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ent_cmd_direct := exec.Command("./binary")
	err = ent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}