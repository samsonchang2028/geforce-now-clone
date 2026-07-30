from flask import Flask, render_template  # type: ignore

app = Flask(__name__)


@app.route("/")
def hello():
    return "Hello World"


@app.route("/hello-world")
def helloWorld():
    return "Chickem butt"


@app.route("/page")
def page():
    return render_template("index.html", title="My Page")


if __name__ == "__main__":
    app.run(debug=True)
