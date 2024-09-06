package main

// LunarDate - Chinese lunar date library
// Homepage: https://github.com/yetist/lunar-date

import (
	"fmt"
	
	"os/exec"
)

func installLunarDate() {
	// Método 1: Descargar y extraer .tar.gz
	lunardate_tar_url := "https://github.com/yetist/lunar-date/releases/download/v3.0.1/lunar-date-3.0.1.tar.xz"
	lunardate_cmd_tar := exec.Command("curl", "-L", lunardate_tar_url, "-o", "package.tar.gz")
	err := lunardate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lunardate_zip_url := "https://github.com/yetist/lunar-date/releases/download/v3.0.1/lunar-date-3.0.1.tar.xz"
	lunardate_cmd_zip := exec.Command("curl", "-L", lunardate_zip_url, "-o", "package.zip")
	err = lunardate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lunardate_bin_url := "https://github.com/yetist/lunar-date/releases/download/v3.0.1/lunar-date-3.0.1.tar.xz"
	lunardate_cmd_bin := exec.Command("curl", "-L", lunardate_bin_url, "-o", "binary.bin")
	err = lunardate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lunardate_src_url := "https://github.com/yetist/lunar-date/releases/download/v3.0.1/lunar-date-3.0.1.tar.xz"
	lunardate_cmd_src := exec.Command("curl", "-L", lunardate_src_url, "-o", "source.tar.gz")
	err = lunardate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lunardate_cmd_direct := exec.Command("./binary")
	err = lunardate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}