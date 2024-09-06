package main

// TomeePlume - Apache TomEE Plume
// Homepage: https://tomee.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installTomeePlume() {
	// Método 1: Descargar y extraer .tar.gz
	tomeeplume_tar_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-plume.tar.gz"
	tomeeplume_cmd_tar := exec.Command("curl", "-L", tomeeplume_tar_url, "-o", "package.tar.gz")
	err := tomeeplume_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tomeeplume_zip_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-plume.zip"
	tomeeplume_cmd_zip := exec.Command("curl", "-L", tomeeplume_zip_url, "-o", "package.zip")
	err = tomeeplume_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tomeeplume_bin_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-plume.bin"
	tomeeplume_cmd_bin := exec.Command("curl", "-L", tomeeplume_bin_url, "-o", "binary.bin")
	err = tomeeplume_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tomeeplume_src_url := "https://www.apache.org/dyn/closer.lua?path=tomee/tomee-9.1.3/apache-tomee-9.1.3-plume.src.tar.gz"
	tomeeplume_cmd_src := exec.Command("curl", "-L", tomeeplume_src_url, "-o", "source.tar.gz")
	err = tomeeplume_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tomeeplume_cmd_direct := exec.Command("./binary")
	err = tomeeplume_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}