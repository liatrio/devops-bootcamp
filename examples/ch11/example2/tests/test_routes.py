from app import create_app
import json

def test_list_users():
    client = create_app().test_client()
    r = client.get("/users")
    assert r.status_code == 200
    assert isinstance(r.get_json(), list)

def test_create_user():
    client = create_app().test_client()
    r = client.post("/users", json={"name": "Diana"})
    assert r.status_code == 201
    assert r.get_json()["name"] == "Diana"
