package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("fetching data...")
	os.Chdir("../data")
	cmd := exec.Command("./data", "-fetch")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println("pre processing...")
	cmd = exec.Command("./data", "-preprocess")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println("one hot encoding...")
	cmd = exec.Command("./data", "-onehot")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println("converting data to csv...")
	os.Chdir("../pycsv")
	cmd = exec.Command("python3", "../pycsv/to_csv.py")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println("training the model...")
	os.Chdir("../regressor")
	cmd = exec.Command("python3", "../regressor/note-exp.py")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
