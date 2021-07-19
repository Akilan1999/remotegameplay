package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
  err := LaplaceConfig()
  if err != nil {
  	fmt.Println(err)
  }
  _, err = LaplaceServer(23424)
	if err != nil {
		fmt.Println(err)
	}
}

// LaplaceConfig Sets up the laplace configuration
func LaplaceConfig() error {
	// Setting up environment variable Laplace
	curDir := os.Getenv("PWD")
	os.Setenv("LAPLACE", curDir)

	cmd := exec.Command("./laplace","-setconfig")

	// USE THE FOLLOWING TO DEBUG
	//cmdReader, err := cmd.StdoutPipe()
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
	//	return
	//}
	//
	//// the following is used to print output of the command
	//// as it makes progress...
	//scanner := bufio.NewScanner(cmdReader)
	//go func() {
	//	for scanner.Scan() {
	//		fmt.Printf("%s\n", scanner.Text())
	//		//
	//		// TODO:
	//		// send output to server
	//	}
	//}()
	//
	//if err := cmd.Start(); err != nil {
	//	return err
	//}

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// LaplaceServer Executes the laplace server
func LaplaceServer(port int)(*exec.Cmd, error) {
	cmd := exec.Command("./laplace","-tls","-addr","0.0.0.0:" + fmt.Sprint(port))

	// USE THE FOLLOWING TO DEBUG
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return nil, err
	}

	// the following is used to print output of the command
	// as it makes progress...
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
			//
			// TODO:
			// send output to server
		}
	}()

	if err = cmd.Run(); err != nil {
		return nil, err
	}

	return cmd, nil
}