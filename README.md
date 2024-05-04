# Usage Guidelines

The database used is MySQL, and you can use Postman app to make a request to the API

This is the list of routes that you can use:

// User <br>
Register: http://localhost:9090/users/register (POST)<br>
Login: http://localhost:9090/users/login (POST)<br>
Update User: http://localhost:9090/users/:id (PUT)<br>
Delete User: http://localhost:9090/users/:id (DELETE)<br>
Logout: http://localhost:9090/users/logout (POST)<br>

// Photo<br>
Get Photos: http://localhost:9090/photos (GET)<br>
Create Photo: http://localhost:9090/photos (POST)<br>
Update Photo: http://localhost:9090/photos/:id (PUT)<br>
DELETE Photo: http://localhost:9090/photos/:id (DELETE)<br>

When using POST method, you can parse some data in JSON format. For example, you can write the data like this for registering a new user:<br>
{<br>
&nbsp;&nbsp;&nbsp;&nbsp;"username": "johndoe",<br>
&nbsp;&nbsp;&nbsp;&nbsp;"email": "johndoe@gmail.com",<br>
&nbsp;&nbsp;&nbsp;&nbsp;"password": "123456"<br>
}
