variable "access_key" {
    description = "Value of AWS Administrator user access key"
    type = string
    sensitive = true
}

variable "secret_key" {
    description = "Value of AWS Administrator user secret key"
    type = string
    sensitive = true
}

variable "region" {
    description = "Value of the aws region"
    type = string
    default = "ap-south-1"
}

variable "key_name" {
    description = "Name of pem file to be created"
    type = string
}

variable "instance_type" {
    description = "Value of a valid aws instance type"
    default = "t2.micro"
}

variable "ami_id" {
    description = "Value of a valid AMI id for the selected AWS zone"
    type = string
}

variable "instance_name" {
    description = "Value of the Name tag for the EC2 instance"
    type = string
    default = "EC2AppServerInstance"
}