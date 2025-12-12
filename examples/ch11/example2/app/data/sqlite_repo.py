import sqlite3
from pathlib import Path

DB_PATH = Path(__file__).parent.parent.parent / "data.sqlite"

def _get_conn():
    conn = sqlite3.connect(str(DB_PATH))
    conn.row_factory = sqlite3.Row
    return conn

def _ensure_table():
    with _get_conn() as c:
        c.execute("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT NOT NULL)")

def get_all():
    _ensure_table()
    with _get_conn() as c:
        rows = c.execute("SELECT id, name FROM users").fetchall()
        return [dict(r) for r in rows]

def add(user):
    _ensure_table()
    with _get_conn() as c:
        cur = c.execute("INSERT INTO users (name) VALUES (?)", (user["name"],))
        return {"id": cur.lastrowid, "name": user["name"]}
