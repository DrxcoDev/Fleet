package main

// Jpdfbookmarks - Create and edit bookmarks on existing PDF files
// Homepage: https://sourceforge.net/projects/jpdfbookmarks/

import (
	"fmt"
	
	"os/exec"
)

func installJpdfbookmarks() {
	// Método 1: Descargar y extraer .tar.gz
	jpdfbookmarks_tar_url := "https://downloads.sourceforge.net/project/jpdfbookmarks/JPdfBookmarks-2.5.2/jpdfbookmarks-2.5.2.tar.gz"
	jpdfbookmarks_cmd_tar := exec.Command("curl", "-L", jpdfbookmarks_tar_url, "-o", "package.tar.gz")
	err := jpdfbookmarks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jpdfbookmarks_zip_url := "https://downloads.sourceforge.net/project/jpdfbookmarks/JPdfBookmarks-2.5.2/jpdfbookmarks-2.5.2.zip"
	jpdfbookmarks_cmd_zip := exec.Command("curl", "-L", jpdfbookmarks_zip_url, "-o", "package.zip")
	err = jpdfbookmarks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jpdfbookmarks_bin_url := "https://downloads.sourceforge.net/project/jpdfbookmarks/JPdfBookmarks-2.5.2/jpdfbookmarks-2.5.2.bin"
	jpdfbookmarks_cmd_bin := exec.Command("curl", "-L", jpdfbookmarks_bin_url, "-o", "binary.bin")
	err = jpdfbookmarks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jpdfbookmarks_src_url := "https://downloads.sourceforge.net/project/jpdfbookmarks/JPdfBookmarks-2.5.2/jpdfbookmarks-2.5.2.src.tar.gz"
	jpdfbookmarks_cmd_src := exec.Command("curl", "-L", jpdfbookmarks_src_url, "-o", "source.tar.gz")
	err = jpdfbookmarks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jpdfbookmarks_cmd_direct := exec.Command("./binary")
	err = jpdfbookmarks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}