package main

// Erlang - Programming language for highly scalable real-time systems
// Homepage: https://www.erlang.org/

import (
	"fmt"
	
	"os/exec"
)

func installErlang() {
	// Método 1: Descargar y extraer .tar.gz
	erlang_tar_url := "https://github.com/erlang/otp/releases/download/OTP-26.2.5/otp_src_26.2.5.tar.gz"
	erlang_cmd_tar := exec.Command("curl", "-L", erlang_tar_url, "-o", "package.tar.gz")
	err := erlang_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	erlang_zip_url := "https://github.com/erlang/otp/releases/download/OTP-26.2.5/otp_src_26.2.5.zip"
	erlang_cmd_zip := exec.Command("curl", "-L", erlang_zip_url, "-o", "package.zip")
	err = erlang_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	erlang_bin_url := "https://github.com/erlang/otp/releases/download/OTP-26.2.5/otp_src_26.2.5.bin"
	erlang_cmd_bin := exec.Command("curl", "-L", erlang_bin_url, "-o", "binary.bin")
	err = erlang_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	erlang_src_url := "https://github.com/erlang/otp/releases/download/OTP-26.2.5/otp_src_26.2.5.src.tar.gz"
	erlang_cmd_src := exec.Command("curl", "-L", erlang_src_url, "-o", "source.tar.gz")
	err = erlang_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	erlang_cmd_direct := exec.Command("./binary")
	err = erlang_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: unixodbc")
	exec.Command("latte", "install", "unixodbc").Run()
	fmt.Println("Instalando dependencia: wxwidgets")
	exec.Command("latte", "install", "wxwidgets").Run()
}