package main

// Newsboat - RSS/Atom feed reader for text terminals
// Homepage: https://newsboat.org/

import (
	"fmt"
	
	"os/exec"
)

func installNewsboat() {
	// Método 1: Descargar y extraer .tar.gz
	newsboat_tar_url := "https://newsboat.org/releases/2.36/newsboat-2.36.tar.xz"
	newsboat_cmd_tar := exec.Command("curl", "-L", newsboat_tar_url, "-o", "package.tar.gz")
	err := newsboat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	newsboat_zip_url := "https://newsboat.org/releases/2.36/newsboat-2.36.tar.xz"
	newsboat_cmd_zip := exec.Command("curl", "-L", newsboat_zip_url, "-o", "package.zip")
	err = newsboat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	newsboat_bin_url := "https://newsboat.org/releases/2.36/newsboat-2.36.tar.xz"
	newsboat_cmd_bin := exec.Command("curl", "-L", newsboat_bin_url, "-o", "binary.bin")
	err = newsboat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	newsboat_src_url := "https://newsboat.org/releases/2.36/newsboat-2.36.tar.xz"
	newsboat_cmd_src := exec.Command("curl", "-L", newsboat_src_url, "-o", "source.tar.gz")
	err = newsboat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	newsboat_cmd_direct := exec.Command("./binary")
	err = newsboat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
	exec.Command("latte", "install", "asciidoctor").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: make")
	exec.Command("latte", "install", "make").Run()
}