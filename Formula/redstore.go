package main

// Redstore - Lightweight RDF triplestore powered by Redland
// Homepage: https://www.aelius.com/njh/redstore/

import (
	"fmt"
	
	"os/exec"
)

func installRedstore() {
	// Método 1: Descargar y extraer .tar.gz
	redstore_tar_url := "https://www.aelius.com/njh/redstore/redstore-0.5.4.tar.gz"
	redstore_cmd_tar := exec.Command("curl", "-L", redstore_tar_url, "-o", "package.tar.gz")
	err := redstore_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redstore_zip_url := "https://www.aelius.com/njh/redstore/redstore-0.5.4.zip"
	redstore_cmd_zip := exec.Command("curl", "-L", redstore_zip_url, "-o", "package.zip")
	err = redstore_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redstore_bin_url := "https://www.aelius.com/njh/redstore/redstore-0.5.4.bin"
	redstore_cmd_bin := exec.Command("curl", "-L", redstore_bin_url, "-o", "binary.bin")
	err = redstore_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redstore_src_url := "https://www.aelius.com/njh/redstore/redstore-0.5.4.src.tar.gz"
	redstore_cmd_src := exec.Command("curl", "-L", redstore_src_url, "-o", "source.tar.gz")
	err = redstore_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redstore_cmd_direct := exec.Command("./binary")
	err = redstore_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: raptor")
	exec.Command("latte", "install", "raptor").Run()
	fmt.Println("Instalando dependencia: rasqal")
	exec.Command("latte", "install", "rasqal").Run()
	fmt.Println("Instalando dependencia: redland")
	exec.Command("latte", "install", "redland").Run()
}