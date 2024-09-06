package main

// AwsAuth - Allows you to programmatically authenticate into AWS accounts through IAM roles
// Homepage: https://github.com/iamarkadyt/aws-auth

import (
	"fmt"
	
	"os/exec"
)

func installAwsAuth() {
	// Método 1: Descargar y extraer .tar.gz
	awsauth_tar_url := "https://registry.npmjs.org/@iamarkadyt/aws-auth/-/aws-auth-2.2.4.tgz"
	awsauth_cmd_tar := exec.Command("curl", "-L", awsauth_tar_url, "-o", "package.tar.gz")
	err := awsauth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsauth_zip_url := "https://registry.npmjs.org/@iamarkadyt/aws-auth/-/aws-auth-2.2.4.tgz"
	awsauth_cmd_zip := exec.Command("curl", "-L", awsauth_zip_url, "-o", "package.zip")
	err = awsauth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsauth_bin_url := "https://registry.npmjs.org/@iamarkadyt/aws-auth/-/aws-auth-2.2.4.tgz"
	awsauth_cmd_bin := exec.Command("curl", "-L", awsauth_bin_url, "-o", "binary.bin")
	err = awsauth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsauth_src_url := "https://registry.npmjs.org/@iamarkadyt/aws-auth/-/aws-auth-2.2.4.tgz"
	awsauth_cmd_src := exec.Command("curl", "-L", awsauth_src_url, "-o", "source.tar.gz")
	err = awsauth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsauth_cmd_direct := exec.Command("./binary")
	err = awsauth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}