package main

// Pywhat - Identify anything: emails, IP addresses, and more
// Homepage: https://github.com/bee-san/pyWhat

import (
	"fmt"
	
	"os/exec"
)

func installPywhat() {
	// Método 1: Descargar y extraer .tar.gz
	pywhat_tar_url := "https://files.pythonhosted.org/packages/ae/31/57bb23df3d3474c1e0a0ae207f8571e763018fa064823310a76758eaef81/pywhat-5.1.0.tar.gz"
	pywhat_cmd_tar := exec.Command("curl", "-L", pywhat_tar_url, "-o", "package.tar.gz")
	err := pywhat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pywhat_zip_url := "https://files.pythonhosted.org/packages/ae/31/57bb23df3d3474c1e0a0ae207f8571e763018fa064823310a76758eaef81/pywhat-5.1.0.zip"
	pywhat_cmd_zip := exec.Command("curl", "-L", pywhat_zip_url, "-o", "package.zip")
	err = pywhat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pywhat_bin_url := "https://files.pythonhosted.org/packages/ae/31/57bb23df3d3474c1e0a0ae207f8571e763018fa064823310a76758eaef81/pywhat-5.1.0.bin"
	pywhat_cmd_bin := exec.Command("curl", "-L", pywhat_bin_url, "-o", "binary.bin")
	err = pywhat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pywhat_src_url := "https://files.pythonhosted.org/packages/ae/31/57bb23df3d3474c1e0a0ae207f8571e763018fa064823310a76758eaef81/pywhat-5.1.0.src.tar.gz"
	pywhat_cmd_src := exec.Command("curl", "-L", pywhat_src_url, "-o", "source.tar.gz")
	err = pywhat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pywhat_cmd_direct := exec.Command("./binary")
	err = pywhat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}