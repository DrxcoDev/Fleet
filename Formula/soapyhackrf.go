package main

// Soapyhackrf - SoapySDR HackRF module
// Homepage: https://github.com/pothosware/SoapyHackRF/wiki

import (
	"fmt"
	
	"os/exec"
)

func installSoapyhackrf() {
	// Método 1: Descargar y extraer .tar.gz
	soapyhackrf_tar_url := "https://github.com/pothosware/SoapyHackRF/archive/refs/tags/soapy-hackrf-0.3.4.tar.gz"
	soapyhackrf_cmd_tar := exec.Command("curl", "-L", soapyhackrf_tar_url, "-o", "package.tar.gz")
	err := soapyhackrf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	soapyhackrf_zip_url := "https://github.com/pothosware/SoapyHackRF/archive/refs/tags/soapy-hackrf-0.3.4.zip"
	soapyhackrf_cmd_zip := exec.Command("curl", "-L", soapyhackrf_zip_url, "-o", "package.zip")
	err = soapyhackrf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	soapyhackrf_bin_url := "https://github.com/pothosware/SoapyHackRF/archive/refs/tags/soapy-hackrf-0.3.4.bin"
	soapyhackrf_cmd_bin := exec.Command("curl", "-L", soapyhackrf_bin_url, "-o", "binary.bin")
	err = soapyhackrf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	soapyhackrf_src_url := "https://github.com/pothosware/SoapyHackRF/archive/refs/tags/soapy-hackrf-0.3.4.src.tar.gz"
	soapyhackrf_cmd_src := exec.Command("curl", "-L", soapyhackrf_src_url, "-o", "source.tar.gz")
	err = soapyhackrf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	soapyhackrf_cmd_direct := exec.Command("./binary")
	err = soapyhackrf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: hackrf")
	exec.Command("latte", "install", "hackrf").Run()
	fmt.Println("Instalando dependencia: soapysdr")
	exec.Command("latte", "install", "soapysdr").Run()
}