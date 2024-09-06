package main

// Zssh - Interactive file transfers over SSH
// Homepage: https://zssh.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installZssh() {
	// Método 1: Descargar y extraer .tar.gz
	zssh_tar_url := "https://downloads.sourceforge.net/project/zssh/zssh/1.5/zssh-1.5c.tgz"
	zssh_cmd_tar := exec.Command("curl", "-L", zssh_tar_url, "-o", "package.tar.gz")
	err := zssh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zssh_zip_url := "https://downloads.sourceforge.net/project/zssh/zssh/1.5/zssh-1.5c.tgz"
	zssh_cmd_zip := exec.Command("curl", "-L", zssh_zip_url, "-o", "package.zip")
	err = zssh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zssh_bin_url := "https://downloads.sourceforge.net/project/zssh/zssh/1.5/zssh-1.5c.tgz"
	zssh_cmd_bin := exec.Command("curl", "-L", zssh_bin_url, "-o", "binary.bin")
	err = zssh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zssh_src_url := "https://downloads.sourceforge.net/project/zssh/zssh/1.5/zssh-1.5c.tgz"
	zssh_cmd_src := exec.Command("curl", "-L", zssh_src_url, "-o", "source.tar.gz")
	err = zssh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zssh_cmd_direct := exec.Command("./binary")
	err = zssh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: lrzsz")
	exec.Command("latte", "install", "lrzsz").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}