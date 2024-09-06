package main

// Mxnet - Flexible and efficient library for deep learning
// Homepage: https://mxnet.apache.org

import (
	"fmt"
	
	"os/exec"
)

func installMxnet() {
	// Método 1: Descargar y extraer .tar.gz
	mxnet_tar_url := "https://www.apache.org/dyn/closer.lua?path=mxnet/1.9.1/apache-mxnet-src-1.9.1-incubating.tar.gz"
	mxnet_cmd_tar := exec.Command("curl", "-L", mxnet_tar_url, "-o", "package.tar.gz")
	err := mxnet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mxnet_zip_url := "https://www.apache.org/dyn/closer.lua?path=mxnet/1.9.1/apache-mxnet-src-1.9.1-incubating.zip"
	mxnet_cmd_zip := exec.Command("curl", "-L", mxnet_zip_url, "-o", "package.zip")
	err = mxnet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mxnet_bin_url := "https://www.apache.org/dyn/closer.lua?path=mxnet/1.9.1/apache-mxnet-src-1.9.1-incubating.bin"
	mxnet_cmd_bin := exec.Command("curl", "-L", mxnet_bin_url, "-o", "binary.bin")
	err = mxnet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mxnet_src_url := "https://www.apache.org/dyn/closer.lua?path=mxnet/1.9.1/apache-mxnet-src-1.9.1-incubating.src.tar.gz"
	mxnet_cmd_src := exec.Command("curl", "-L", mxnet_src_url, "-o", "source.tar.gz")
	err = mxnet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mxnet_cmd_direct := exec.Command("./binary")
	err = mxnet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: opencv")
	exec.Command("latte", "install", "opencv").Run()
}