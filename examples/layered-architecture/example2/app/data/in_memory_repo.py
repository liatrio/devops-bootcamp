USERS = [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]
NEXT_ID = 3

def get_all():
    return USERS.copy()

def add(user):
    global NEXT_ID
    u = {"id": NEXT_ID, "name": user["name"]}
    NEXT_ID += 1
    USERS.append(u)
    return u
