package main

// Flake - FLAC audio encoder
// Homepage: https://flake-enc.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installFlake() {
	// Método 1: Descargar y extraer .tar.gz
	flake_tar_url := "https://downloads.sourceforge.net/project/flake-enc/flake/0.11/flake-0.11.tar.bz2"
	flake_cmd_tar := exec.Command("curl", "-L", flake_tar_url, "-o", "package.tar.gz")
	err := flake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flake_zip_url := "https://downloads.sourceforge.net/project/flake-enc/flake/0.11/flake-0.11.tar.bz2"
	flake_cmd_zip := exec.Command("curl", "-L", flake_zip_url, "-o", "package.zip")
	err = flake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flake_bin_url := "https://downloads.sourceforge.net/project/flake-enc/flake/0.11/flake-0.11.tar.bz2"
	flake_cmd_bin := exec.Command("curl", "-L", flake_bin_url, "-o", "binary.bin")
	err = flake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flake_src_url := "https://downloads.sourceforge.net/project/flake-enc/flake/0.11/flake-0.11.tar.bz2"
	flake_cmd_src := exec.Command("curl", "-L", flake_src_url, "-o", "source.tar.gz")
	err = flake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flake_cmd_direct := exec.Command("./binary")
	err = flake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}