Title
The purpose of this project is to focus on solving three different algorithm problems and developing a user management system with CRUD operations. 
Additionally, the aim is to evaluate the ability to create appropriate test scenarios, adhere to coding standards suitable for software development processes,and document the projects with English explanations.
Overview

Q1: 
This algorithm aims to sort a list of words in descending order based on the number of "a" letters each word contains. 
If two words have the same number of "a" letters, the length of the words is considered, prioritizing the longer word. 
If the lengths are also equal, the original order of the list is preserved. The algorithm is written in Go and has been thoroughly tested for both stability and accuracy.

Q2:
This solution calculates all perfect squares up to a given integer input. 
The algorithm starts from the square of 2 (2 * 2 = 4) and recursively calculates the squares of each subsequent integer until the specified limit is exceeded. 
It is important to note that 2 is used as the starting point and only its square (4) is included in the output. 
The user input is thoroughly validated, and meaningful error messages are provided for invalid inputs. 
The use of recursion enhances both the simplicity and efficiency of the algorithm. 
However, the behavior of using "2" as the starting point should be kept in mind when interpreting the results. The solution has been tested extensively for correctness and clarity.

Q3:
This algorithm identifies the most frequently occurring element in a list. 
The frequency of each element is calculated using a map structure, and the element with the highest frequency is returned. 
If multiple elements have the same frequency, the algorithm prioritizes the first element encountered in the list. 
The solution has been tested with various data sets and evaluated for both performance and accuracy.

Q4:
This system is designed to perform essential CRUD operations such as listing, adding, updating, and deleting users. 
The backend is built with Go and adheres to RESTful API standards, while SQLite is used for persistent storage of user data. 
On the frontend, a modern and user-friendly interface is created using Next.js and Tailwind CSS, providing both aesthetic and functional user experiences. 
Axios is utilized for seamless communication between the frontend and backend, ensuring all CRUD operations are integrated smoothly. 
The system is supported by a comprehensive testing framework and has been optimized to efficiently manage users.

Setup

Q1:
Go to the sw-q1 folder and use the following commands to run the q1.go file. 
The algorithm will sort the given words based on the number of "a" letters they contain.

cd sw-q1
go run q1.go

Q2:
Go to the sw-q2 folder and use the following commands to run the q2.go file. 
The algorithm will generate perfect squares in order for a given number. 
You can also run the go test command to verify the correctness of the solution.

cd sw-q2
go run q2.go
go test

Q3:
Go to the sw-q3 folder and use the following commands to run the q3.go file. 
The algorithm will find the most frequently occurring element in an array.

cd sw-q3
go run q3.go

Q4:
Running the Backend:
Go to the sw-q4 folder. Run the go mod tidy command to install Go dependencies. 
Then, use the go run main.go command to start the backend application. 
The backend application will run at http://localhost:8080/users.


cd sw-q4
go mod tidy
go run main.go

## Routing and CORS Configuration

### Routing
The backend uses the [gorilla/mux](https://github.com/gorilla/mux) library to define the following API endpoints:

- **List all users**: GET /users
- **Retrieve a specific user by ID**: GET /users/{id}
- **Create a new user**: POST /users
- **Update an existing user**: PUT /users/{id}
- **Delete a user**: DELETE /users/{id}

Example routing code:


router := mux.NewRouter()
router.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")


Running the Frontend:
Go to the sw-q4/frontend folder. Run the npm install command to install Node.js dependencies. 
Then, use the npm run dev command to start the frontend application. 
The frontend interface will run at http://localhost:3000.

cd sw-q4/frontend
npm install
npm run dev


The project is available at: https://github.com/0xEfeBa/fill-labs
For communication, you can reach me at: efeba2024@gmail.com






