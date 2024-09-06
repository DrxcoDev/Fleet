package main

// Highlight - Convert source code to formatted text with syntax highlighting
// Homepage: http://andre-simon.de/doku/highlight/en/highlight.php

import (
	"fmt"
	
	"os/exec"
)

func installHighlight() {
	// Método 1: Descargar y extraer .tar.gz
	highlight_tar_url := "http://andre-simon.de/zip/highlight-4.13.tar.bz2"
	highlight_cmd_tar := exec.Command("curl", "-L", highlight_tar_url, "-o", "package.tar.gz")
	err := highlight_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	highlight_zip_url := "http://andre-simon.de/zip/highlight-4.13.tar.bz2"
	highlight_cmd_zip := exec.Command("curl", "-L", highlight_zip_url, "-o", "package.zip")
	err = highlight_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	highlight_bin_url := "http://andre-simon.de/zip/highlight-4.13.tar.bz2"
	highlight_cmd_bin := exec.Command("curl", "-L", highlight_bin_url, "-o", "binary.bin")
	err = highlight_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	highlight_src_url := "http://andre-simon.de/zip/highlight-4.13.tar.bz2"
	highlight_cmd_src := exec.Command("curl", "-L", highlight_src_url, "-o", "source.tar.gz")
	err = highlight_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	highlight_cmd_direct := exec.Command("./binary")
	err = highlight_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
}