package main

// GnomeRecipes - Formula for GNOME recipes
// Homepage: https://wiki.gnome.org/Apps/Recipes

import (
	"fmt"
	
	"os/exec"
)

func installGnomeRecipes() {
	// Método 1: Descargar y extraer .tar.gz
	gnomerecipes_tar_url := "https://gitlab.gnome.org/GNOME/recipes.git"
	gnomerecipes_cmd_tar := exec.Command("curl", "-L", gnomerecipes_tar_url, "-o", "package.tar.gz")
	err := gnomerecipes_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnomerecipes_zip_url := "https://gitlab.gnome.org/GNOME/recipes.git"
	gnomerecipes_cmd_zip := exec.Command("curl", "-L", gnomerecipes_zip_url, "-o", "package.zip")
	err = gnomerecipes_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnomerecipes_bin_url := "https://gitlab.gnome.org/GNOME/recipes.git"
	gnomerecipes_cmd_bin := exec.Command("curl", "-L", gnomerecipes_bin_url, "-o", "binary.bin")
	err = gnomerecipes_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnomerecipes_src_url := "https://gitlab.gnome.org/GNOME/recipes.git"
	gnomerecipes_cmd_src := exec.Command("curl", "-L", gnomerecipes_src_url, "-o", "source.tar.gz")
	err = gnomerecipes_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnomerecipes_cmd_direct := exec.Command("./binary")
	err = gnomerecipes_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: itstool")
	exec.Command("latte", "install", "itstool").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
	exec.Command("latte", "install", "adwaita-icon-theme").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gnome-autoar")
	exec.Command("latte", "install", "gnome-autoar").Run()
	fmt.Println("Instalando dependencia: gspell")
	exec.Command("latte", "install", "gspell").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: libcanberra")
	exec.Command("latte", "install", "libcanberra").Run()
	fmt.Println("Instalando dependencia: librest")
	exec.Command("latte", "install", "librest").Run()
	fmt.Println("Instalando dependencia: libsoup@2")
	exec.Command("latte", "install", "libsoup@2").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: enchant")
	exec.Command("latte", "install", "enchant").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
}