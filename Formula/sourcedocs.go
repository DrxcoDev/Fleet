package main

// Sourcedocs - Generate Markdown files from inline source code documentation
// Homepage: https://github.com/SourceDocs/SourceDocs

import (
	"fmt"
	
	"os/exec"
)

func installSourcedocs() {
	// Método 1: Descargar y extraer .tar.gz
	sourcedocs_tar_url := "https://github.com/SourceDocs/SourceDocs/archive/refs/tags/2.0.1.tar.gz"
	sourcedocs_cmd_tar := exec.Command("curl", "-L", sourcedocs_tar_url, "-o", "package.tar.gz")
	err := sourcedocs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sourcedocs_zip_url := "https://github.com/SourceDocs/SourceDocs/archive/refs/tags/2.0.1.zip"
	sourcedocs_cmd_zip := exec.Command("curl", "-L", sourcedocs_zip_url, "-o", "package.zip")
	err = sourcedocs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sourcedocs_bin_url := "https://github.com/SourceDocs/SourceDocs/archive/refs/tags/2.0.1.bin"
	sourcedocs_cmd_bin := exec.Command("curl", "-L", sourcedocs_bin_url, "-o", "binary.bin")
	err = sourcedocs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sourcedocs_src_url := "https://github.com/SourceDocs/SourceDocs/archive/refs/tags/2.0.1.src.tar.gz"
	sourcedocs_cmd_src := exec.Command("curl", "-L", sourcedocs_src_url, "-o", "source.tar.gz")
	err = sourcedocs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sourcedocs_cmd_direct := exec.Command("./binary")
	err = sourcedocs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}