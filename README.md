# LoginProject_Backend

1.
GET http://localhost:8000/v1/user/check/{email}

Res:
200 - "message": "success"
502 - "message": "mongo: no documents in result"

2.
POST http://localhost:8000/v1/user/make
Req:
{
    "email" : "example@gmail.com",
    "name" : "John Doe",
    "picture" : "https://lh3.googleusercontent.com/a/AAcHTteu2pMkksACEc4METIe-I8_m2MHGAn2-1hGpNeEjbPIDg=s96-c"
}
Res:
200 - "message": "success"

3.
GET http://localhost:8000/v1/user/get/{email}
Res:
200 - 
{
    "username": "niteshtyagi232000@gmail.com",
    "name": "Nitesh Tyagiii",
    "email": "niteshtyagi232000@gmail.com",
    "password": "niteshtyagi232000@gmail.com",
    "picture": "https://lh3.googleusercontent.com/a/AAcHTteu2pMkksAGEc4METIe-I8_m2MHGAn2-1hGpNeEjbPIDg=s96-c",
    "job": "Worker"
}

4.
PATCH http://localhost:8000/v1/user/update
Req:
{
    "username": "example",
    "name": "John Doe",
    "email": "example@gmail.com",
    "password": "example",
    "picture": "https://lh3.googleusercontent.com/a/AAcHTteu2pMkksAGEc4METIe-I8_m2MHGAn2-1hGpNeEjbPIDg=s96-c",
    "job": "Worker"
}

Res:
200 - "message": "success"
502 - "message": "no such user found"

5.
POST http://localhost:8000/v1/user/login
Req:
{
    "username": "example",
    "password": "example"
}

Res:
200 - "email": "example@gmail.com"
502 - "message": "no such user found"
502 - "message": "wrong password"
