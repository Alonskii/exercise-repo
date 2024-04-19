This is a description of the different accessment questions.  

Question_One demonstrates how to swap the values of two integers in Go.

How to Use
Ensure you have Go installed on your machine.
Clone or download the script.
Open a terminal and navigate to the directory containing the script.
Run the script using the go run command: 

go run main.go 

QUESTION_TWO: 

Sum Even Numbers

The script calculates the sum of even numbers in a given slice of integers. It defines a function sumEvenNumbers that iterates through the provided slice and adds each even number to the total sum. The main function demonstrates its usage by computing the sum of even numbers in a sample slice and printing the result.

QUESTION_THREE: 

Logger Script

The script demonstrates a simple logging functionality in Go. It defines two types of loggers: FileLogger, which logs messages to a file, and ConsoleLogger, which logs messages to the console. The script includes an example of how to use both loggers simultaneously. To use the logging functionality, simply call the Log method on the desired logger with the message to be logged. Ensure that you have the necessary permissions to create and write to files if using the FileLogger. 

QUESTION_FOUR: 

The script calculates the sum of elements in a large array by dividing it into smaller sub-arrays and utilizing concurrent goroutines for parallel computation. The number of goroutines used can be adjusted to optimize performance. To use, specify your large array and the desired number of goroutines. The script then distributes the workload among goroutines, calculates the sum of sub-arrays concurrently, and aggregates the results to output the total sum.

QUESTION_FIVE: 

Random Number Generator and Squares Calculator

The script generates random integers and calculates their squares concurrently using goroutines in Go. The producer function generates 5 random numbers with a 500-millisecond delay between each, sending them to a channel. Meanwhile, the consumer function reads these integers from the channel and computes their squares, printing the results. 

QUESTION_SIX: 

The script creates a basic HTTP server that listens on port 8080 and responds with "Hello, World!" to any incoming requests to the root URL ("/"). To run the server, execute the script, and it will print a message indicating that it's listening on port 8080. You can then access the server by navigating to http://localhost:8080 in your web browser.

QUESTION_SEVEN: 

Simple Todo API
This is a basic Todo API, using Gorilla Mux for routing. The API allows users to perform CRUD operations on tasks. It supports fetching all tasks, creating a new task, updating an existing task, and deleting a task. Tasks are stored in memory and served over HTTP. The API listens on port 8000 by default.

QUESTION_EIGHT: 

The script contains tests for the sumEvenNumbers function, which calculates the sum of even numbers in a given slice. Test cases cover scenarios such as all even numbers, all odd numbers, mixed numbers, an empty slice, single elements, and negative numbers. Each test case ensures the function behaves as expected, reporting any discrepancies between actual and expected results. Additionally, there's a test to ensure the main function runs without errors. 

QUESTION_NINE: 

The script implements unit tests for a simple HTTP server using Go's net/http package. The tests ensure that the server responds correctly with a "Hello, World!" message for GET requests and returns the appropriate status codes for other HTTP methods like POST, PUT, and DELETE. The TestHelloServer function validates the response body and status code for a GET request, while TestHelloServerWithDifferentMethods tests various HTTP methods. This script helps maintain the functionality and integrity of the server's response handling.
