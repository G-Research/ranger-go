# ranger-go

A Go client library for Apache Ranger.

This project provides a Go client for interacting with Apache Ranger's REST API. It includes functionality for managing policies, such as creating, updating, deleting, and fetching policies.

This currently supports the following Ranger services:

* Kafka
* Hive

## Features

- Get policies by service name or all policies.
- Get a specific policy by ID.
- Create new policies.
- Update existing policies.
- Delete policies.
- Includes unit tests with a mock Ranger server for testing.

## Prerequisites

- Go 1.24 or later
- Docker and Docker Compose (optional, for testing with a real Ranger instance)

## Installation

To install the Ranger Go client, you can use the following command:

```bash
go get github.com/g-research/ranger-go
```

Alternatively import into a Go project:

```go
import "github.com/g-research/ranger-go/ranger"
```

And then run `go get` with no arguments to download the package.

## Usage

Example: Get policies

```go
// Create a new Ranger client
client := ranger.NewClient("http://localhost:6080", "admin", "rangerR0cks!")

// Retrieve the dev_kafka policies
serviceName := "dev_kafka"
policies, err := client.GetPolicies(serviceName)
if err != nil {
    fmt.Println("Error fetching policies:", err)
    return
}

// List the policies
fmt.Println("Kafka policies:")
for _, policy := range policies {
    fmt.Println(policy.Name)
}
```

Example: Create a policy

```go
// Create a new Ranger client
client := ranger.NewClient("http://localhost:6080", "admin", "rangerR0cks!")

// Define a new policy
policy := ranger.Policy{
    Name:        "Test Policy 1",
    Description: "This is a test policy",
    IsEnabled:   true,
    Service:     "dev_kafka",
    PolicyType:  0,
    ServiceType: "kafka",
    Resources: ranger.Resources{
        Topic: &ranger.ResourceType{
            Values:      []string{"topic-*"},
            IsExcludes:  false,
            IsRecursive: false,
        },
    },
    IsAuditEnabled: true,
}

// Create the policy
createdPolicy, err := client.CreatePolicy(&policy)
if err != nil {
    fmt.Println("Error creating policy:", err)
} else {
    fmt.Println("Created policy with ID:", createdPolicy.ID)
}
```

Example: Delete a policy

```go
// Create a new Ranger client
client := ranger.NewClient("http://localhost:6080", "admin", "rangerR0cks!")

// Get existing policy
policyId := 65
policy, err := client.GetPolicy(policyId)

if err != nil {
	fmt.Printf("Error fetching policy: %v\n", err)
	return
}

// Update the description
policy.Description = "Updated description"

updatedPolicy, err := client.UpdatePolicy(policy)

if err != nil {
	fmt.Printf("Error updating policy: %v\n", err)
	return
}

fmt.Printf("New Policy description: %s\n", updatedPolicy.Description)
```

Example: Delete a policy

```go
// Create a new Ranger client
client := ranger.NewClient("http://localhost:6080", "admin", "rangerR0cks!")

// Delete the policy
policyId := 65
err := client.DeletePolicy(policyId)

if err != nil {
    fmt.Printf("Error deleting policy: %v\n", err)
    return
}

fmt.Printf("Successfully deleted policy with ID %d\n", policyId)
```