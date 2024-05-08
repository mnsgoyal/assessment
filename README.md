## API

Below is a list of API endpoints with their respective input and output.

### Get Employee Details

#### Endpoint

```
GET
/employee/<employeeID>
```


`employeeID`: A intiger value, e.g. `1`

#### Output

```json
{
    "ID": 1,
    "Name": "Divij",
    "Position": "SE",
    "Salary": 500
}
```

### Get Employee Details

#### Endpoint

```
GET
/employee/<page>/<limit>
```


`page`: A intiger value, e.g. `1`
`limit`: A intiger value, e.g. `1`

#### Output

```json
{
    "PaginationDetails": {
        "Previous": false,
        "Next": true,
        "Limit": 2,
        "TotalPages": 3,
        "Page": 1
    },
    "EmpDetails": [
        {
            "ID": 1,
            "Name": "Robin",
            "Position": "Technical Architech",
            "Salary": 2000
        },
        {
            "ID": 2,
            "Name": "James",
            "Position": "Technical Manager",
            "Salary": 2300
        }
    ]
}
```

### Add New Employee Details

#### Endpoint

```
POST
/employee
```

#### Input

```json
{
    "Name": "Divij",
    "Position": "SSE",
    "Salary": 500
}
```

`Name`: Name of the employee, e.g. `Divij`   
`Position`: Position of the employee, e.g. `SSE`
`Salary`: Salary of the employee, e.g. `500`

### Get Stored Readings

#### Endpoint

```
DELETE
/employee/<employeeID>
```


`employeeID`: A intiger value, e.g. `1`

### Update Employee Details

#### Endpoint

```
PATCH
/employee
```

#### Input

```json
{
    "ID": 1,
    "Name": "Divij",
    "Position": "SSE",
    "Salary": 500
}
```

`ID`: ID of the employee, e.g. `1`   
`Name`: Name of the employee, e.g. `Divij`   
`Position`: Position of the employee, e.g. `SSE`
`Salary`: Salary of the employee, e.g. `500`