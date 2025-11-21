# keep business logic here (validation, business rules)
from ..data import get_all, add

def get_users():
    return get_all()

def create_user(payload):
    name = payload.get("name")
    if not name or not isinstance(name, str):
        raise ValueError("name required")
    return add({"name": name})
