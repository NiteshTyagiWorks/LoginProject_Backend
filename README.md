# LoginProject_Backend

## Endpoints

### 1. Check User

- **Method:** GET
- **URL:** http://localhost:8000/v1/user/check/{email}
- **Response:**
    - 200 - "message": "success"
    - 502 - "message": "mongo: no documents in result"

### 2. Make User

- **Method:** POST
- **URL:** http://localhost:8000/v1/user/make
- **Request:**
    ```json
    {
        "email" : "example@gmail.com",
        "name" : "John Doe",
        "picture" : "https://lh3.googleusercontent.com/a/AAcHTteu2pMkksACEc4METIe-I8_m2MHGAn2-1hGpNeEjbPIDg=s96-c"
    }
    ```
- **Response:**
    - 200 - "message": "success"

### 3. Get User

- **Method:** GET
- **URL:** http://localhost:8000/v1/user/get/{email}
- **Response:**
    - 200 -
    ```json
    {
        "username": "niteshtyagi232000@gmail.com",
        "name": "Nitesh Tyagiii",
        "email": "niteshtyagi232000@gmail.com",
        "password": "niteshtyagi232000@gmail.com",
        "picture": "https://lh3.googleusercontent.com/a/AAcHTteu2pMkksAGEc4METIe-I8_m2MHGAn2-1hGpNeEjbPIDg=s96-c",
        "job": "Worker"
    }
    ```

### 4. Update User

- **Method:** PATCH
- **URL:** http://localhost:8000/v1/user/update
- **Request:**
    ```json
    {
        "username": "example",
        "name": "John Doe",
        "email": "example@gmail.com",
        "password": "example",
        "picture": "https://lh3.googleusercontent.com/a/AAcHTteu2pMkksAGEc4METIe-I8_m2MHGAn2-1hGpNeEjbPIDg=s96-c",
        "job": "Worker"
    }
    ```
- **Response:**
    - 200 - "message": "success"
    - 502 - "message": "no such user found"

### 5. Login User

- **Method:** POST
- **URL:** http://localhost:8000/v1/user/login
- **Request:**
    ```json
    {
        "username": "example",
        "password": "example"
    }
    ```
- **Response:**
    - 200 - "email": "example@gmail.com"
    - 502 - "message": "no such user found"
    - 502 - "message": "wrong password"
