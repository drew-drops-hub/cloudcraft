package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Config struct {
	AccessKey    string
	SecretKey    string
	Region       string
	KeyName      string
	InstanceType string
	AmiID        string
	InstanceName string
}

func main() {
	config := getInputs()
	if confirmConfig(config) {
		runTerraform(config)
	} else {
		fmt.Println("Operation cancelled.")
	}
}

func getInputs() Config {
	reader := bufio.NewReader(os.Stdin)
	config := Config{}

	fmt.Print("Enter AWS Access Key: ")
	config.AccessKey, _ = reader.ReadString('\n')
	config.AccessKey = strings.TrimSpace(config.AccessKey)

	fmt.Print("Enter AWS Secret Key: ")
	config.SecretKey, _ = reader.ReadString('\n')
	config.SecretKey = strings.TrimSpace(config.SecretKey)

	fmt.Print("Enter AWS Region (default: ap-south-1): ")
	config.Region, _ = reader.ReadString('\n')
	config.Region = strings.TrimSpace(config.Region)
	if config.Region == "" {
		config.Region = "ap-south-1"
	}

	fmt.Print("Enter Key Name for PEM file: ")
	config.KeyName, _ = reader.ReadString('\n')
	config.KeyName = strings.TrimSpace(config.KeyName)

	fmt.Print("Enter Instance Type (default: t2.micro): ")
	config.InstanceType, _ = reader.ReadString('\n')
	config.InstanceType = strings.TrimSpace(config.InstanceType)
	if config.InstanceType == "" {
		config.InstanceType = "t2.micro"
	}

	fmt.Print("Enter AMI ID: ")
	config.AmiID, _ = reader.ReadString('\n')
	config.AmiID = strings.TrimSpace(config.AmiID)

	fmt.Print("Enter Instance Name (default: EC2AppServerInstance): ")
	config.InstanceName, _ = reader.ReadString('\n')
	config.InstanceName = strings.TrimSpace(config.InstanceName)
	if config.InstanceName == "" {
		config.InstanceName = "EC2AppServerInstance"
	}

	return config
}

func confirmConfig(config Config) bool {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Configuration", "Value"})

	table.Append([]string{"AWS Region", config.Region})
	table.Append([]string{"Key Name", config.KeyName})
	table.Append([]string{"Instance Type", config.InstanceType})
	table.Append([]string{"AMI ID", config.AmiID})
	table.Append([]string{"Instance Name", config.InstanceName})

	fmt.Println("\nConfiguration Summary:")
	table.Render()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nDo you want to proceed with this configuration? (y/n): ")
	response, _ := reader.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(response)) == "y"
}

func runTerraform(config Config) {
	args := []string{
		"apply",
		"-auto-approve",
		fmt.Sprintf("-var=access_key=%s", config.AccessKey),
		fmt.Sprintf("-var=secret_key=%s", config.SecretKey),
		fmt.Sprintf("-var=region=%s", config.Region),
		fmt.Sprintf("-var=key_name=%s", config.KeyName),
		fmt.Sprintf("-var=instance_type=%s", config.InstanceType),
		fmt.Sprintf("-var=ami_id=%s", config.AmiID),
		fmt.Sprintf("-var=instance_name=%s", config.InstanceName),
	}

	cmd := exec.Command("terraform", args...)
	cmd.Dir = filepath.Join("terraform", "aws", "node")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error applying Terraform configuration:", err)
		return
	}

	fmt.Println("EC2 instance created successfully!")
}
