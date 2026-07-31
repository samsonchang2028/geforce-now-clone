from flask import Flask, render_template, jsonify, request  # type: ignore

app = Flask(__name__)


# items = [{"id": 1, "name": "Bob"}, {"id": 2, "name": "Sam"}]


# @app.route("/")
# def hello():
#     return "Hello World"


# @app.route("/hello-world")
# def helloWorld():
#     return "Chickem butt"


@app.route("/page")
def page():
    return render_template("index1.html", title="My Page")


# @app.route("/api/items", methods=["GET"])
# def get_items():
#     return jsonify(items)


# @app.route("/api/items/<int:item_id>", methods=["GET"])
# def get_item(item_id):
#     item = None
#     for i in items:
#         if i["id"] == i:
#             item = i
#             break
#     if not item:
#         return jsonify({"error": "item not found"}), 404
#     return jsonify(item)


# @app.route("/api/items", methods=["POST"])
# def create_item():
#     data = request.get_json()
#     if not data or "name" not in data:
#         return jsonify({"error": "Name is required"}), 400
#     new_item = {"id": len(items) + 1, "name": data["name"]}
#     items.append(new_item)
#     return jsonify(new_item), 201


items = []  # type: ignore


@app.route("/api/items", methods=["GET"])
def get_items():
    return jsonify(items)


@app.route("/api/item/<string:name_id>", methods=["GET"])
def get_name(name_id):
    item = None
    for name in items:
        if name == name_id:
            item = name
    if not item:
        return jsonify({"error": "item not found"}), 404
    return jsonify(item)


@app.route("/api/items", methods=["POST"])
def create_item():
    data = request.get_json()
    item_name = {"name": data.get("name")}
    items.append(item_name)
    return jsonify({"message": "Item added successfully"}), 201


if __name__ == "__main__":
    app.run(debug=True)
