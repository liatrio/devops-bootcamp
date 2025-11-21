from app import create_app

app = create_app()

if __name__ == "__main__":
    # NOTE: debug=True is for development only. Disable in production.
    app.run(port=5001, debug=True)
