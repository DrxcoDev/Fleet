package main

// QtPerconaServer - Qt SQL Database Driver
// Homepage: https://www.qt.io/

import (
	"fmt"
	
	"os/exec"
)

func installQtPerconaServer() {
	// Método 1: Descargar y extraer .tar.gz
	qtperconaserver_tar_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtperconaserver_cmd_tar := exec.Command("curl", "-L", qtperconaserver_tar_url, "-o", "package.tar.gz")
	err := qtperconaserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qtperconaserver_zip_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtperconaserver_cmd_zip := exec.Command("curl", "-L", qtperconaserver_zip_url, "-o", "package.zip")
	err = qtperconaserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qtperconaserver_bin_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtperconaserver_cmd_bin := exec.Command("curl", "-L", qtperconaserver_bin_url, "-o", "binary.bin")
	err = qtperconaserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qtperconaserver_src_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtperconaserver_cmd_src := exec.Command("curl", "-L", qtperconaserver_src_url, "-o", "source.tar.gz")
	err = qtperconaserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qtperconaserver_cmd_direct := exec.Command("./binary")
	err = qtperconaserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: percona-server")
	exec.Command("latte", "install", "percona-server").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
}