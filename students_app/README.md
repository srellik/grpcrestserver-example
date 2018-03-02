This example is for [grpcrestserver](https://github.com/srellik/grpcrestserver), it uses bazel, protobuf and grpc.

```
git clone https://github.com/srellik/grpcrestserver-example && cd grpcrestserver-example

bazel run //students_app:server

You should see server running at :8888 port.

curl -v -d '{"user_id": "12345", "password": "pwd123"}' -X POST http://localhost:8888/login | jq

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkcyI6IntcIlVzZXJJZFwiOlwiMTIzNDVcIixcIlBhc3N3b3JkXCI6XCJcIixcIk1ldGFkYXRhXCI6bnVsbH0iLCJ0aW1lIjoxNTIwMDE3MjQyNjA4ODY3NzcyfQ.njIGonwRLiL6US1YI27rIF3V8WPOy8xBRg2_b0oW70o"
}

curl -v -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkcyI6IntcIlVzZXJJZFwiOlwiMTIzNDVcIixcIlBhc3N3b3JkXCI6XCJcIixcIk1ldGFkYXRhXCI6bnVsbH0iLCJ0aW1lIjoxNTIwMDE3MjQyNjA4ODY3NzcyfQ.njIGonwRLiL6US1YI27rIF3V8WPOy8xBRg2_b0oW70o" http://localhost:8888/v1/students | jq

{}

curl -v -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkcyI6IntcIlVzZXJJZFwiOlwiMTIzNDVcIixcIlBhc3N3b3JkXCI6XCJcIixcIk1ldGFkYXRhXCI6bnVsbH0iLCJ0aW1lIjoxNTIwMDE3MjQyNjA4ODY3NzcyfQ.njIGonwRLiL6US1YI27rIF3V8WPOy8xBRg2_b0oW70o" -d '{"first_name": "Abigail", "last_name": "Fisher", "age": 35, "mother_name": "Amanda", "father_name": "Steven", "gender": "FEMALE"}' -X POST http://localhost:8888/v1/students | jq

{
  "id": "7432234034536129345",
  "first_name": "Abigail",
  "last_name": "Fisher",
  "age": 35,
  "mother_name": "Amanda",
  "father_name": "Steven",
  "gender": "FEMALE"
}

 curl -v -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkcyI6IntcIlVzZXJJZFwiOlwiMTIzNDVcIixcIlBhc3N3b3JkXCI6XCJcIixcIk1ldGFkYXRhXCI6bnVsbH0iLCJ0aW1lIjoxNTIwMDE3MjQyNjA4ODY3NzcyfQ.njIGonwRLiL6US1YI27rIF3V8WPOy8xBRg2_b0oW70o" -d '{"first_name": "Christopher", "last_name": "Hemmings", "age": 35, "mother_name": "Sophie", "father_name": "Justin", "gender": "MALE"}' -X POST http://localhost:8888/v1/students | jq

curl -v -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkcyI6IntcIlVzZXJJZFwiOlwiMTIzNDVcIixcIlBhc3N3b3JkXCI6XCJcIixcIk1ldGFkYXRhXCI6bnVsbH0iLCJ0aW1lIjoxNTIwMDE3MjQyNjA4ODY3NzcyfQ.njIGonwRLiL6US1YI27rIF3V8WPOy8xBRg2_b0oW70o" http://localhost:8888/v1/students | jq

{
  "students": [
    {
      "id": "7432234034536129345",
      "first_name": "Abigail",
      "last_name": "Fisher",
      "age": 35,
      "mother_name": "Amanda",
      "father_name": "Steven",
      "gender": "FEMALE"
    },
    {
      "id": "6742714662793448032",
      "first_name": "Christopher",
      "last_name": "Hemmings",
      "age": 35,
      "mother_name": "Sophie",
      "father_name": "Justin",
      "gender": "MALE"
    }
  ]
}

Change the signature
curl -v -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkcyI6IntcIlVzZXJJZFwiOlwiMTIzNDVcIixcIlBhc3N3b3JkXCI6XCJcIixcIk1ldGFkYXRhXCI6bnVsbH0iLCJ0aW1lIjoxNTIwMDE3MjQyNjA4ODY3NzcyfQ.njIGonwRLiL6US1YI27rIF3V8WPOy8xBRg2_b0oW70" http://localhost:8888/v1/students | jq

{
  "error": "invalid token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVkcyI6IntcIlVzZXJJZFwiOlwiMTIzNDVcIixcIlBhc3N3b3JkXCI6XCJcIixcIk1ldGFkYXRhXCI6bnVsbH0iLCJ0aW1lIjoxNTIwMDE3MjQyNjA4ODY3NzcyfQ.njIGonwRLiL6US1YI27rIF3V8WPOy8xBRg2_b0oW70",
  "code": 16
}
```
