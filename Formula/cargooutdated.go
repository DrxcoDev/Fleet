package main

// CargoOutdated - Cargo subcommand for displaying when Rust dependencies are out of date
// Homepage: https://github.com/kbknapp/cargo-outdated

import (
	"fmt"
	
	"os/exec"
)

func installCargoOutdated() {
	// Método 1: Descargar y extraer .tar.gz
	cargooutdated_tar_url := "https://static.crates.io/crates/cargo-outdated/cargo-outdated-0.15.0.crate"
	cargooutdated_cmd_tar := exec.Command("curl", "-L", cargooutdated_tar_url, "-o", "package.tar.gz")
	err := cargooutdated_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargooutdated_zip_url := "https://static.crates.io/crates/cargo-outdated/cargo-outdated-0.15.0.crate"
	cargooutdated_cmd_zip := exec.Command("curl", "-L", cargooutdated_zip_url, "-o", "package.zip")
	err = cargooutdated_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargooutdated_bin_url := "https://static.crates.io/crates/cargo-outdated/cargo-outdated-0.15.0.crate"
	cargooutdated_cmd_bin := exec.Command("curl", "-L", cargooutdated_bin_url, "-o", "binary.bin")
	err = cargooutdated_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargooutdated_src_url := "https://static.crates.io/crates/cargo-outdated/cargo-outdated-0.15.0.crate"
	cargooutdated_cmd_src := exec.Command("curl", "-L", cargooutdated_src_url, "-o", "source.tar.gz")
	err = cargooutdated_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargooutdated_cmd_direct := exec.Command("./binary")
	err = cargooutdated_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
	fmt.Println("Instalando dependencia: libgit2@1.7")
	exec.Command("latte", "install", "libgit2@1.7").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}