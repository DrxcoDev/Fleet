package main

// Libvmaf - Perceptual video quality assessment based on multi-method fusion
// Homepage: https://github.com/Netflix/vmaf

import (
	"fmt"
	
	"os/exec"
)

func installLibvmaf() {
	// Método 1: Descargar y extraer .tar.gz
	libvmaf_tar_url := "https://github.com/Netflix/vmaf/archive/refs/tags/v3.0.0.tar.gz"
	libvmaf_cmd_tar := exec.Command("curl", "-L", libvmaf_tar_url, "-o", "package.tar.gz")
	err := libvmaf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvmaf_zip_url := "https://github.com/Netflix/vmaf/archive/refs/tags/v3.0.0.zip"
	libvmaf_cmd_zip := exec.Command("curl", "-L", libvmaf_zip_url, "-o", "package.zip")
	err = libvmaf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvmaf_bin_url := "https://github.com/Netflix/vmaf/archive/refs/tags/v3.0.0.bin"
	libvmaf_cmd_bin := exec.Command("curl", "-L", libvmaf_bin_url, "-o", "binary.bin")
	err = libvmaf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvmaf_src_url := "https://github.com/Netflix/vmaf/archive/refs/tags/v3.0.0.src.tar.gz"
	libvmaf_cmd_src := exec.Command("curl", "-L", libvmaf_src_url, "-o", "source.tar.gz")
	err = libvmaf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvmaf_cmd_direct := exec.Command("./binary")
	err = libvmaf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: nasm")
	exec.Command("latte", "install", "nasm").Run()
}