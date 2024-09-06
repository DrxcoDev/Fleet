package main

// Dotdrop - Save your dotfiles once, deploy them everywhere
// Homepage: https://github.com/deadc0de6/dotdrop

import (
	"fmt"
	
	"os/exec"
)

func installDotdrop() {
	// Método 1: Descargar y extraer .tar.gz
	dotdrop_tar_url := "https://files.pythonhosted.org/packages/78/bd/2684d3616a0838675e7355480bef03295ee12074f1d5745a37a18542c1cc/dotdrop-1.14.2.tar.gz"
	dotdrop_cmd_tar := exec.Command("curl", "-L", dotdrop_tar_url, "-o", "package.tar.gz")
	err := dotdrop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dotdrop_zip_url := "https://files.pythonhosted.org/packages/78/bd/2684d3616a0838675e7355480bef03295ee12074f1d5745a37a18542c1cc/dotdrop-1.14.2.zip"
	dotdrop_cmd_zip := exec.Command("curl", "-L", dotdrop_zip_url, "-o", "package.zip")
	err = dotdrop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dotdrop_bin_url := "https://files.pythonhosted.org/packages/78/bd/2684d3616a0838675e7355480bef03295ee12074f1d5745a37a18542c1cc/dotdrop-1.14.2.bin"
	dotdrop_cmd_bin := exec.Command("curl", "-L", dotdrop_bin_url, "-o", "binary.bin")
	err = dotdrop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dotdrop_src_url := "https://files.pythonhosted.org/packages/78/bd/2684d3616a0838675e7355480bef03295ee12074f1d5745a37a18542c1cc/dotdrop-1.14.2.src.tar.gz"
	dotdrop_cmd_src := exec.Command("curl", "-L", dotdrop_src_url, "-o", "source.tar.gz")
	err = dotdrop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dotdrop_cmd_direct := exec.Command("./binary")
	err = dotdrop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libmagic")
	exec.Command("latte", "install", "libmagic").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}