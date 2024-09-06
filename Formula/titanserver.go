package main

// TitanServer - Distributed graph database
// Homepage: https://thinkaurelius.github.io/titan/

import (
	"fmt"
	
	"os/exec"
)

func installTitanServer() {
	// Método 1: Descargar y extraer .tar.gz
	titanserver_tar_url := "http://s3.thinkaurelius.com/downloads/titan/titan-1.0.0-hadoop1.zip"
	titanserver_cmd_tar := exec.Command("curl", "-L", titanserver_tar_url, "-o", "package.tar.gz")
	err := titanserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	titanserver_zip_url := "http://s3.thinkaurelius.com/downloads/titan/titan-1.0.0-hadoop1.zip"
	titanserver_cmd_zip := exec.Command("curl", "-L", titanserver_zip_url, "-o", "package.zip")
	err = titanserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	titanserver_bin_url := "http://s3.thinkaurelius.com/downloads/titan/titan-1.0.0-hadoop1.zip"
	titanserver_cmd_bin := exec.Command("curl", "-L", titanserver_bin_url, "-o", "binary.bin")
	err = titanserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	titanserver_src_url := "http://s3.thinkaurelius.com/downloads/titan/titan-1.0.0-hadoop1.zip"
	titanserver_cmd_src := exec.Command("curl", "-L", titanserver_src_url, "-o", "source.tar.gz")
	err = titanserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	titanserver_cmd_direct := exec.Command("./binary")
	err = titanserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}