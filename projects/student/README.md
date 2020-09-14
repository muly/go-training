# Training

## Deployment

* First build the program to check for any build errors like show below.
* This will generate an executable file in the current directory(for say the project directory).

```cmd
    go build -o studentsapi cmd/main.go
```

* Now run the build file like shown below

```cmd
    ./studentsapi
```

* Deployment is successful if the below output is shown in the terminal

```cmd
    2020/09/11 00:52:18 Listening on port 8080
```

## REQUEST

* For post and get requests  
<http://localhost:8080/students>

* For get request for one student by id  
* For changes to a student, put request by id  
* For delete request by id  
* For patch request by id  
<http://localhost:8080/students/[id_of_student]>

## API

### Handlers

1. HandlePost
2. HandleGet
3. HandleGetById
4. HandlePut
5. HandlePatch
6. HandleDelete

## CRUD

* This is generic crud code
* Try not to edit. only generic code in this file
* For any customizations done in seperate methods should be in other go files  
* Methods in crud are explained below

### HTTP Methods

1. POST

    * Post posts the given list of records into the database collection
    * The request body contains one or more number of student data
    * All individual student data is checked for existence in database

2. GET

    * This can be run in two different functions
        * Get
            * Get gets all the records in the database
        * GetById
            * GetById gets all the records based on the id of the student provided
            * if id is blank in the input, returns generic error
            * if record is not found, returns not found error

3. PUT

    * Put updates the record
    * If unique fields which being id is missing in the parameters, return error
    * Matches the record based on the ud and updates the field with what is provided in the input struct

4. PATCH

    * Patch updates the record with only provided fields

5. DELETE

    * Delete deletes the record
    * If id is blank in the input, returns generic error
    * If id is not found in the database, returns not found error
    * Matches the record based on the id and deletes the record

below is the postman collection with examples for all the above mentioned methods. Note: you have to update the uuid in these, with the id that you are testing with.
- `training 2020.postman_collection.json`

## Sample Post
## Sample Get

## Troubleshooting

* For any errors at the time of build, first resolve them and then build again.
* To resolve build errors, the termial locates the error location (file, line_numnber, column_number) like below.
* Once all the errors are resolved, then build the program again.

```cmd
    # github.com/muly/go-training/projects/student/api
    api/handlers.go:26:44: undefined: students
    api/handlers.go:31:28: undefined: students
```
