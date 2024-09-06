package main

// Libdazzle - GNOME companion library to GObject and Gtk+
// Homepage: https://gitlab.gnome.org/GNOME/libdazzle

import (
	"fmt"
	
	"os/exec"
)

func installLibdazzle() {
	// Método 1: Descargar y extraer .tar.gz
	libdazzle_tar_url := "https://download.gnome.org/sources/libdazzle/3.44/libdazzle-3.44.0.tar.xz"
	libdazzle_cmd_tar := exec.Command("curl", "-L", libdazzle_tar_url, "-o", "package.tar.gz")
	err := libdazzle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdazzle_zip_url := "https://download.gnome.org/sources/libdazzle/3.44/libdazzle-3.44.0.tar.xz"
	libdazzle_cmd_zip := exec.Command("curl", "-L", libdazzle_zip_url, "-o", "package.zip")
	err = libdazzle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdazzle_bin_url := "https://download.gnome.org/sources/libdazzle/3.44/libdazzle-3.44.0.tar.xz"
	libdazzle_cmd_bin := exec.Command("curl", "-L", libdazzle_bin_url, "-o", "binary.bin")
	err = libdazzle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdazzle_src_url := "https://download.gnome.org/sources/libdazzle/3.44/libdazzle-3.44.0.tar.xz"
	libdazzle_cmd_src := exec.Command("curl", "-L", libdazzle_src_url, "-o", "source.tar.gz")
	err = libdazzle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdazzle_cmd_direct := exec.Command("./binary")
	err = libdazzle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
}