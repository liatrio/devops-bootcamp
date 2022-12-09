# DevOps API

This documentation is used to get a better understanding of the DevOps API to be used for the custom terraform provider and see how the CRUD operations work.

API is composed of 4 resources:

```
1. Engineer - an individual engineer
    - id (unique numeric or alphanumeric identifier)
    - name (string)
    - valid email address(string)

2. Dev - collection of developer engineers
    - id (unique numeric or alphanumeric identifier)
    - name (string)
    - engineers (list of engineer resources)

3. Ops - collections of operations engineers
    - id (unique numeric or alphanumeric identifier)
    - name (string)
    - engineers (list of engineer resources)

4. DevOps - combination of a collection of developer and operations engineers
    - id (unique numeric or alphanumeric identifier)
    - Devs (list of Dev resources)
    - Ops (list of Ops resources)
```

## How to start the API:

Run it on locally:
- This will build, test, and run the API in the foreground.
```
make run
```

Running it as a docker container:
-  This will build, test the go code and then build the docker image then run it as a daemon process.
```
make docker-run
```

## How to use crud operations:

To make things a bit simpler we provided some scripts that go through CRUD operations for the resources.

These scripts are located in the `scripts` directory of this current directory.

All the scripts excute curl requests that the api recieves. 
Feel free to investigate and tinker with the scripts to have a better understanding client-side requests for the terrafrom custom provider.

Example call to creating new engineer resources from the script:

From current directory:
```
scripts/engineer.sh post 
```

From scripts directory:
```
./dev.sh post
```

This should send a POST request to the API that it will recieve and return json pertaining to the new engieer made:

Example response:
- This is how the response from the api will look like when you create a new engineer resource.
```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Fri, 09 Dec 2022 17:16:21 GMT
Content-Length: 68

{
    "name": "bob",
    "id": "D7SJA",
    "email": "bob@bob.com"
}
```

