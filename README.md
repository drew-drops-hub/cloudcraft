# CloudCraft

CloudCraft is a CLI tool that simplifies the process of creating EC2 instances on AWS using Terraform. It provides an interactive prompt to gather necessary configuration details and then uses Terraform to provision the infrastructure.

## Prerequisites

Before using CloudCraft, ensure you have the following installed on your system:

- Go (version 1.22 or later)
- Terraform (version 0.12 or later)

## Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/drew-drops-hub/cloudcraft.git
    cd cloudcraft
    ```
2. Build the application:
    ```bash
    go build
    ```

## Usage

1. Run the CloudCraft executable:
    ```bash
    ./cloudcraft
    ```
2. Follow the prompts to enter your AWS credentials and EC2 instance details.

3. Review the configuration summary and confirm to proceed.

4. CloudCraft will use Terraform to create your EC2 instance.

## Configuration

CloudCraft will prompt you for the following information:

- AWS Access Key
- AWS Secret Key
- AWS Region
- EC2 Instance Type
- AMI ID
- Instance Name
- Key Pair Name

## Important Notes

- Ensure that Terraform is installed and available in your system PATH.
- AWS credentials entered are not stored and are only used for the current session.
- The Terraform state file will be created in the same directory. Be cautious about sharing this file as it may contain sensitive information.

## Contributing

Contributions to CloudCraft are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.