package main

// Capnp - Data interchange format and capability-based RPC system
// Homepage: https://capnproto.org/

import (
	"fmt"
	
	"os/exec"
)

func installCapnp() {
	// Método 1: Descargar y extraer .tar.gz
	capnp_tar_url := "https://capnproto.org/capnproto-c++-1.0.2.tar.gz"
	capnp_cmd_tar := exec.Command("curl", "-L", capnp_tar_url, "-o", "package.tar.gz")
	err := capnp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	capnp_zip_url := "https://capnproto.org/capnproto-c++-1.0.2.zip"
	capnp_cmd_zip := exec.Command("curl", "-L", capnp_zip_url, "-o", "package.zip")
	err = capnp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	capnp_bin_url := "https://capnproto.org/capnproto-c++-1.0.2.bin"
	capnp_cmd_bin := exec.Command("curl", "-L", capnp_bin_url, "-o", "binary.bin")
	err = capnp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	capnp_src_url := "https://capnproto.org/capnproto-c++-1.0.2.src.tar.gz"
	capnp_cmd_src := exec.Command("curl", "-L", capnp_src_url, "-o", "source.tar.gz")
	err = capnp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	capnp_cmd_direct := exec.Command("./binary")
	err = capnp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}