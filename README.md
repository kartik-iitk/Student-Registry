# Student Registry
## Task 5 - Programming Club - Secrectary Recruitment 2022-23

```Made By - Kartik Anant Kulkarni (210493)```

+ Database: *MySQL*
+ API creation: *GoLang*
+ References to create MySQL Database: [https://www.youtube.com/watch?v=BDShK0w1zAg&ab_channel=PnTutorialsbyPradnyankurNikam](https://www.youtube.com/watch?v=BDShK0w1zAg&ab_channel=PnTutorialsbyPradnyankurNikam), [
https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-20-04](
https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-20-04)
+ Prerequisite: Database Creation, User Creation.
+ Implementation:
    + Student Table is created with Roll No, Name, Branch and User ID.
    + `studentcontrollers.go` implements following api handlers:
        + get all students
        + get student by roll no.
        + create student entry
        + update student data for specific roll no.
        + delete student with specific roll no.
    + Roll no. was used as a filtering operations.
+ Sample calls to API:
   + To Create record  
    ```
    POST API
    http://localhost:8090/create
    Body:
    {
    "ID"  : 1
    "rollNo" : 100
    "name" : "Test",
    "branch" : "Test Branch",
    "userId" : 100
    }
    ```

   + To Get record
    ```
    GET API (all)
    http://localhost:8090/get
    ```

   + GET API (specific)
   ```
    http://localhost:8090/get/{rollNo}
    ```

   + To Update record
   ```
    PUT API
    http://localhost:8090/update/{rollNo}
    Body:
    {
    "ID"  : 1
    "rollNo" : 100
    "name" : "Test1",
    "branch" : "Test Branch",
    "userId" : 100
    }
    ```
    
   + To Delete record 
   ```
    DELETE API
    http://localhost:8090/delete/{rollNo}
    ```
