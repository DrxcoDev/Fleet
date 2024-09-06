package main

// AdwaitaIconTheme - Icons for the GNOME project
// Homepage: https://developer.gnome.org

import (
	"fmt"
	
	"os/exec"
)

func installAdwaitaIconTheme() {
	// Método 1: Descargar y extraer .tar.gz
	adwaitaicontheme_tar_url := "https://download.gnome.org/sources/adwaita-icon-theme/46/adwaita-icon-theme-46.2.tar.xz"
	adwaitaicontheme_cmd_tar := exec.Command("curl", "-L", adwaitaicontheme_tar_url, "-o", "package.tar.gz")
	err := adwaitaicontheme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	adwaitaicontheme_zip_url := "https://download.gnome.org/sources/adwaita-icon-theme/46/adwaita-icon-theme-46.2.tar.xz"
	adwaitaicontheme_cmd_zip := exec.Command("curl", "-L", adwaitaicontheme_zip_url, "-o", "package.zip")
	err = adwaitaicontheme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	adwaitaicontheme_bin_url := "https://download.gnome.org/sources/adwaita-icon-theme/46/adwaita-icon-theme-46.2.tar.xz"
	adwaitaicontheme_cmd_bin := exec.Command("curl", "-L", adwaitaicontheme_bin_url, "-o", "binary.bin")
	err = adwaitaicontheme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	adwaitaicontheme_src_url := "https://download.gnome.org/sources/adwaita-icon-theme/46/adwaita-icon-theme-46.2.tar.xz"
	adwaitaicontheme_cmd_src := exec.Command("curl", "-L", adwaitaicontheme_src_url, "-o", "source.tar.gz")
	err = adwaitaicontheme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	adwaitaicontheme_cmd_direct := exec.Command("./binary")
	err = adwaitaicontheme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gtk4")
	exec.Command("latte", "install", "gtk4").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: librsvg")
	exec.Command("latte", "install", "librsvg").Run()
}