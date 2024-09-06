package main

// CargoFuzz - Command-line helpers for fuzzing
// Homepage: https://rust-fuzz.github.io/book/cargo-fuzz.html

import (
	"fmt"
	
	"os/exec"
)

func installCargoFuzz() {
	// Método 1: Descargar y extraer .tar.gz
	cargofuzz_tar_url := "https://github.com/rust-fuzz/cargo-fuzz/archive/refs/tags/0.12.0.tar.gz"
	cargofuzz_cmd_tar := exec.Command("curl", "-L", cargofuzz_tar_url, "-o", "package.tar.gz")
	err := cargofuzz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargofuzz_zip_url := "https://github.com/rust-fuzz/cargo-fuzz/archive/refs/tags/0.12.0.zip"
	cargofuzz_cmd_zip := exec.Command("curl", "-L", cargofuzz_zip_url, "-o", "package.zip")
	err = cargofuzz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargofuzz_bin_url := "https://github.com/rust-fuzz/cargo-fuzz/archive/refs/tags/0.12.0.bin"
	cargofuzz_cmd_bin := exec.Command("curl", "-L", cargofuzz_bin_url, "-o", "binary.bin")
	err = cargofuzz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargofuzz_src_url := "https://github.com/rust-fuzz/cargo-fuzz/archive/refs/tags/0.12.0.src.tar.gz"
	cargofuzz_cmd_src := exec.Command("curl", "-L", cargofuzz_src_url, "-o", "source.tar.gz")
	err = cargofuzz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargofuzz_cmd_direct := exec.Command("./binary")
	err = cargofuzz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
}