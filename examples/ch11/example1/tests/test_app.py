from app import app
import json

def test_list_users():
    client = app.test_client()
    r = client.get("/users")
    assert r.status_code == 200
    data = r.get_json()
    assert isinstance(data, list)

def test_create_user():
    client = app.test_client()
    r = client.post("/users", json={"name": "Charlie"})
    assert r.status_code == 201
    assert r.get_json()["name"] == "Charlie"
