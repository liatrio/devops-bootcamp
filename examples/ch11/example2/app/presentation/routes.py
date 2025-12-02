from flask import Blueprint, request, jsonify
from ..services.user_service import get_users, create_user

bp = Blueprint("users", __name__)

@bp.route("/users", methods=["GET"])
def list_users():
    return jsonify(get_users())

@bp.route("/users", methods=["POST"])
def add_user():
    payload = request.get_json() or {}
    try:
        user = create_user(payload)
    except ValueError as e:
        return jsonify({"error": str(e)}), 400
    return jsonify(user), 201
