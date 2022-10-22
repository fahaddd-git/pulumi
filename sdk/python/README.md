Pulumi Python SDK
The Pulumi Python SDK lets you write cloud programs in Python. It is not required, but recommended that you use a virtual environment to isolate dependencies of your projects and ensure reproducibility between machines.

Installation
Using pip:

$ pip install pulumi

Using a Poetry project:

$ poetry add pulumi
This SDK is meant for use with the Pulumi CLI. Visit Pulumi's Download & Install to install the CLI.

Existing Python projects can opt-in to using the built-in virtual environment support by setting the virtualenv option. To manually create a virtual environment and install dependencies, run the following commands in your project directory:

Windows
$ python -m venv venv
$ venv\Scripts\pip install -r requirements.txt

MacOS
$ python3 -m venv venv
$ venv/bin/pip install -r requirements.txt

Linux
$ python3 -m venv venv && venv/bin/pip install -r requirements.txt

Building and Testing
For anybody who wants to build from source, here is how you do it.

Make Targets
To build the SDK, simply run make from the root directory (where this README lives, at sdk/python/ from the repo's root). This will build the code, run tests, and install the package and its supporting artifacts.

At the moment, for local development, everything is installed into $HOME/.dev-pulumi. You will want this on your $PATH.

The tests will verify that everything works, but feel free to try running pulumi preview and/or pulumi up from the examples/minimal/ directory.