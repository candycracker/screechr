# screechr
A simple project

Welcome to Screehr

use following command to get access token (john has operator role and alice has admin role)

curl -X POST http://localhost:8080/login
   -H 'Content-Type: application/json'
   -d '{"username":"john","password":"123456"}'

or

curl -X POST http://localhost:8080/login
   -H 'Content-Type: application/json'
   -d '{"username":"alice","password":"123456"}'

expected response:
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOjEsInVpZCI6MTIzNDU2Nzg5MCwiaWF0IjoxNjYwMjY0Mzk2LCJpc3MiOiJqaWF3ZWkifQ.0XzXf9-gCw4eB4GOZ5IAJ113LxhywVGHXgtp3ey3pv4"}

-------------------------------------------------------------------------------------------------------------------------------------------
after receive access token then you might able to test following api

/get_profile:

curl -H "Authorization: Bearer $ACCESS_TOKEN" 
-X GET http://localhost:8080/get_profile 
-H 'Content-Type: application/json' 
-d '{"id":$id,"token":$token}'

example:

curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOjEsInVpZCI6MTIzNDU2Nzg5MCwiaWF0IjoxNjYwMjYzNTg4LCJpc3MiOiJqaWF3ZWkifQ.5cRy9DRBwCelhEzR9WSS3CFYVxdn_k6nlMviuvffzqQ" -X GET http://localhost:8080/get_profile -H 'Content-Type: application/json' -d '{"id":1234567890,"token":"JG3LDSkEzbQgnGcIU7o1P8p2FxuHUMg8"}'

-------------------------------------------------------------------------------------------------------------------------------------------
/update_profile:

curl -H "Authorization: Bearer $ACCESS_TOKEN" 
   -X GET http://localhost:8080/update_profile
   -H 'Content-Type: application/json'
   -d '{"id": $id,"user_name": $uid, "first_name": $name, "last_name": $name, "url": $url, "token": $token}'

example:

curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOjEsInVpZCI6MTIzNDU2Nzg5MCwiaWF0IjoxNjYwMjYzNTg4LCJpc3MiOiJqaWF3ZWkifQ.5cRy9DRBwCelhEzR9WSS3CFYVxdn_k6nlMviuvffzqQ" -X GET http://localhost:8080/update_profile -H 'Content-Type: application/json' -d '{"id":1234567890,"user_name": "AAA", "first_name": "BBB", "last_name": "CCC", "url": "*****.jpg","token":"JG3LDSkEzbQgnGcIU7o1P8p2FxuHUMg8"}'

-------------------------------------------------------------------------------------------------------------------------------------------
/update_profile_picture:

curl -H "Authorization: Bearer $ACCESS_TOKEN" 
   -X GET http://localhost:8080/update_profile_picture
   -H 'Content-Type: application/json'
   -d '{"id": $id, "url": $url}'

example:

curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOjEsInVpZCI6MTIzNDU2Nzg5MCwiaWF0IjoxNjYwMjYzNTg4LCJpc3MiOiJqaWF3ZWkifQ.5cRy9DRBwCelhEzR9WSS3CFYVxdn_k6nlMviuvffzqQ" -X GET http://localhost:8080/update_profile_picture -H 'Content-Type: application/json' -d '{"id":1234567890, "url": "*****.jpg","token":"JG3LDSkEzbQgnGcIU7o1P8p2FxuHUMg8"}'

-------------------------------------------------------------------------------------------------------------------------------------------
/get_screeches:

curl -H "Authorization: Bearer $ACCESS_TOKEN" 
   -X GET http://localhost:8080/get_screeches
   -H 'Content-Type: application/json'
   -d '{"order_by_ascend": $false, "user_id": $id}'

example 1:

curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOjEsInVpZCI6MTIzNDU2Nzg5MCwiaWF0IjoxNjYwMjYzNTg4LCJpc3MiOiJqaWF3ZWkifQ.5cRy9DRBwCelhEzR9WSS3CFYVxdn_k6nlMviuvffzqQ" -X GET http://localhost:8080/get_screeches -H 'Content-Type: application/json' -d '{"order_by_ascend":false}'

example 2: 
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOjEsInVpZCI6MTIzNDU2Nzg5MCwiaWF0IjoxNjYwMjYzNTg4LCJpc3MiOiJqaWF3ZWkifQ.5cRy9DRBwCelhEzR9WSS3CFYVxdn_k6nlMviuvffzqQ" -X GET http://localhost:8080/get_screeches -H 'Content-Type: application/json' -d '{"order_by_ascend":true, "user_id":1234567890}'
-------------------------------------------------------------------------------------------------------------------------------------------
/get_screech:

curl -H "Authorization: Bearer $ACCESS_TOKEN" 
   -X GET http://localhost:8080/get_screech
   -H 'Content-Type: application/json'
   -d '{"id":$id}'

example:
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOjEsInVpZCI6MTIzNDU2Nzg5MCwiaWF0IjoxNjYwMjYzNTg4LCJpc3MiOiJqaWF3ZWkifQ.5cRy9DRBwCelhEzR9WSS3CFYVxdn_k6nlMviuvffzqQ" -X GET http://localhost:8080/get_screech -H 'Content-Type: application/json' -d '{"id":1}'

-------------------------------------------------------------------------------------------------------------------------------------------
/create_screech:

curl -H "Authorization: Bearer $ACCESS_TOKEN" 
   -X GET http://localhost:8080/create_screech
   -H 'Content-Type: application/json'
   -d '{"user_id":$id, "content":$content}'

example:
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOjEsInVpZCI6MTIzNDU2Nzg5MCwiaWF0IjoxNjYwMjYzNTg4LCJpc3MiOiJqaWF3ZWkifQ.5cRy9DRBwCelhEzR9WSS3CFYVxdn_k6nlMviuvffzqQ" -X GET http://localhost:8080/create_screech -H 'Content-Type: application/json' -d '{"user_id":1234567890, "content":"PPPPPPPPPPPPPP"}'
-------------------------------------------------------------------------------------------------------------------------------------------
/update_screech:

curl -H "Authorization: Bearer $ACCESS_TOKEN" 
   -X GET http://localhost:8080/update_screech
   -H 'Content-Type: application/json'
   -d '{"id":$id, "content":$content}'

example:
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOjEsInVpZCI6MTIzNDU2Nzg5MCwiaWF0IjoxNjYwMjYzNTg4LCJpc3MiOiJqaWF3ZWkifQ.5cRy9DRBwCelhEzR9WSS3CFYVxdn_k6nlMviuvffzqQ" -X GET http://localhost:8080/get_screech -H 'Content-Type: application/json' -d '{"id":1, "content":"xxxxxxxxxxxx"}'
-------------------------------------------------------------------------------------------------------------------------------------------


for user who has admin role can also use following api to retrieve any profile

/admin/get_profile:

curl -H "Authorization: Bearer $ACCESS_TOKEN" 
   -X GET http://localhost:8080/admin/get_profile
   -H 'Content-Type: application/json'
   -d '{"id":$id}'

example:

curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFsaWNlIiwicGFzc3dvcmQiOiIxMjM0NTYiLCJyb2xlIjoyLCJ1aWQiOjk5ODc2NTQzMjEsImlhdCI6MTY2MDI2Njc2MywiaXNzIjoiamlhd2VpIn0.NYlo2U4FWUpTVUgrysKakOy8Bq2pJdmDk5M22ZTRr7k" -X GET http://localhost:8080/admin/get_profile -H 'Content-Type: application/json' -d '{"id":1234567890}'

