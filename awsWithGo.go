// package main

// import (
//     "fmt"
//     "github.com/aws/aws-sdk-go/aws"
//     "github.com/aws/aws-sdk-go/aws/session"
//     "github.com/aws/aws-sdk-go/service/ec2"
// )

// func main() {
//     // Create a new AWS session
//     sess, err := session.NewSession(&aws.Config{
//         Region: aws.String("ap-south-1"), // Corrected AWS region format
//     })
//     if err != nil {
//         fmt.Println("Error creating session:", err)
//         return
//     }

//     // Create a new EC2 service client
//     svc := ec2.New(sess)

//     // Specify the parameters for the instance
//     params := &ec2.RunInstancesInput{
//         ImageId:      aws.String("ami-0ad37af6d0ff85233"), // Specify the AMI ID
//         InstanceType: aws.String("t2.micro"),              // Specify the instance type
//         MinCount:     aws.Int64(1),
//         MaxCount:     aws.Int64(1),
//     }

//     // Run the instance
//     resp, err := svc.RunInstances(params)
//     if err != nil {
//         fmt.Println("Error launching instance:", err)
//         return
//     }

//     // Output the instance ID
//     instanceID := *resp.Instances[0].InstanceId
//     fmt.Println("Launched instance with ID:", instanceID)
// }

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	// "time"
)

func main() {
	// Create a new AWS session

// session.NewSession() this is a function from the AWS SDK  library for go
    // it is used to create a new session
   // it will returns two values 'sess' a new AWS session and 'err' an error(if any)
  // It takes a configuration struct as an argument. In this case, 
 // it sets the AWS region to "ap-northeast-1".
 // It returns two values, sess (a new AWS session) and err (an error, if any).
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"), // Corrected AWS region format
	})

	//If there is an error during session creation, print an error message and exit the program.

	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}
//	Create a new EC2 service client using the AWS session (sess) created in the previous step.
	// Create a new EC2 service client
	svc := ec2.New(sess)

	// Specify the parameters for the instance
	// Create a RunInstancesInput struct with parameters such as the AMI ID, instance type,
	//  minimum and maximum count of instances.
	params := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-0ad37af6d0ff85233"), // Specify the AMI ID
		InstanceType: aws.String("t2.micro"),              // Specify the instance type
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	// Run the instance
	resp, err := svc.RunInstances(params)
	if err != nil {
		fmt.Println("Error launching instance:", err)
		return
	}
	// Use the EC2 service client (svc) to launch an EC2 instance with the specified parameters.
	// resp will contain information about the launched instances.
	// Output the instance ID

	instanceID := *resp.Instances[0].InstanceId
	// Extract the instance ID from the response and print it to the console.

	fmt.Println("Launched instance with ID:", instanceID)

	// Wait for the instance to be in a running state

	// Use the WaitUntilInstanceRunning method to wait until 
	// the launched instance is in a running state.

	fmt.Println("Waiting for the instance to be in a running state...")

	err = svc.WaitUntilInstanceRunning(&ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	})
	
	if err != nil {
		fmt.Println("Error waiting for instance to be running:", err)
		return
	}

}