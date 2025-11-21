from flask import Flask, request, jsonify

app = Flask(__name__)

# In-memory "DB"
USERS = [
    {"id": 1, "name": "Alice"},
    {"id": 2, "name": "Bob"}
]
NEXT_ID = 3

@app.route("/users", methods=["GET"])
def list_users():
    return jsonify(USERS)

@app.route("/users", methods=["POST"])
def create_user():
    global NEXT_ID
    payload = request.get_json() or {}
    name = payload.get("name")
    if not name:
        return jsonify({"error": "name required"}), 400
    user = {"id": NEXT_ID, "name": name}
    NEXT_ID += 1
    USERS.append(user)
    return jsonify(user), 201

if __name__ == "__main__":
    app.run(port=5000, debug=True)
