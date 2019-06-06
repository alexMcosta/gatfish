package main

func main() {
	// Collect the YAML file
	var resource Tags
	yaml := resource.Configure()

	processCentral(yaml)

	// Process the YAML file
	//cloud.ProcessCentral(yaml)

	// Output logs
}
