# DevOps API

This documentation is used to get a better understanding of the DevOps API to be used for the custom terraform provider and see how the CRUD operations work.

API is composed of 4 resources:

```
1. Engineer - an individual engineer
- id (unique numeric or alphanumeric identifier)
- name (string)
- valid email address(string)

1. Dev - collection of developer engineers
- id (unique numeric or alphanumeric identifier)
- name (string)
- engineers (list of engineer resources)

1. Ops - collections of operations engineers
- id (unique numeric or alphanumeric identifier)
- name (string)
- engineers (list of engineer resources)

1. DevOps - combination of a collection of developer and operations engineers
- id (unique numeric or alphanumeric identifier)
- Devs (list of Dev resources)
- Ops (list of Ops resources)
```

How to start the API:

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
