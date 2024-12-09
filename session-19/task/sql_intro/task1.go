package sql_intro

import (
	"fmt"
	"os"
	"os/exec"
)

func RunContainer() {
	// Remove any leftover container
	RemoveContainer()
	var db_pass string
	fmt.Print("Please enter password for postgresql. Default value is \"securepass\": ")
	fmt.Scanln(&db_pass)
	if db_pass == "" {
		db_pass = "securepass"
	}
	pass_var := fmt.Sprintf("POSTGRES_PASSWORD=%s", db_pass)
	cmd := exec.Command("docker", "container", "run",
		"--name", "school-db", "-e",
		"POSTGRES_USER=student", "-e", pass_var,
		"-p", "5432:5432", "-d",
		"postgres")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Container started successfully")
	}

}
func RemoveContainer() {
	cmd := exec.Command("docker", "container", "rm", "school-db", "--force")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Container removed successfully")
	}
}

func RunTask1() {

}
