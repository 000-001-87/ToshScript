# ToshScript
ToshScript is a basic language that allows more advanced features that Scratch does not have. It also has package building and it is based on the Go Programming Language. It is licensed under the GNU General Public License v3.0.
## Compiling it yourself
### Requirements
Go 1.10> (this has been only compiled on Go 1.17 on Linux)
*nix-like Environmment 
Clone this repo using the following commands
### Compiling it

    git clone https://github.com/RAFhub6/ToshScript.git
    cd ToshScript
Before running ToshScript, you must compile this first. Run the following command:

    make build
You should see a "build" file appear. If not, something has went wrong. Please report any errors to the Issues tab. Now to run ToshScript, enter `make run` into the terminal. You should see the following output:

    ~$ make run
    ToshScript Demo
    If you see this, you have successfully compiled ToshScript! Congratulations!
This comes with a demo file (`index.ssi`). You can find the commands in the file and use them in your own script. To run a custom script, run `build yourscriptname.ssi`. There is also an `index.ssi-build` that stores variables for the script. To do this with your script, also create a file called `yourscriptname.ssi-build`.

