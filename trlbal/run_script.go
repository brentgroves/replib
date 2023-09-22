package trlbal

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

// type CancelFunc func()
// import (
//     "log"
//     "os/exec"
// )
// func main() {

//     cmd := exec.Command("firefox")

//     err := cmd.Run()

//     if err != nil {
//         log.Fatal(err)
//     }
// }

// type customOutput struct{}

//	func (c customOutput) Write(p []byte) (int, error) {
//		fmt.Println("received output: ", string(p))
//		return len(p), nil
//	}
//
//	func (c customOutput) cancel(){
//		fmt.Fprint(os.Stderr, "request cancelled\n")
//	}
//
// https://www.sohamkamani.com/golang/exec-shell-command/
// Hello returns a greeting for the named person.
func RunScript(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// create a new *Cmd instance
	// here we pass the command as the first argument and the arguments to pass to the command as the
	// remaining arguments in the function
	cmd := exec.Command("ls", "./")
	// an alternate way is: out, err := exec.Command("ls", "-l").Output()

	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
		log.Fatal(err)
	}
	// otherwise, print the output from running the command
	fmt.Println("Output: ", string(out))

	// Note that when we run exec, our application does not spawn a shell,
	// and runs the given command directly. This means that any
	// shell-based processing,
	// like glob patterns or expansions will not be done.

	// https://www.sohamkamani.com/golang/exec-shell-command/
	// If we tried executing ping using cmd.Output,
	// we wouldn’t get any output, since the Output method waits
	// for the command to execute, and the ping command executes
	// indefinitely.
	// Instead, we can create a custom Stdout attribute to read output
	// continuously:

	// cmd2 := exec.Command("ping", "google.com")
	// // pipe the commands output to the applications
	// // standard output
	// cmd2.Stdout = os.Stdout

	// // Runs the command and waits for completion
	// // but the output is instantly piped to Stdout
	// if err := cmd2.Run(); err != nil {
	// 	fmt.Println("could not run command: ", err)
	// }

	// fmt.Printf("Path Separator: '%c'\n", os.PathSeparator)

	// Instead of using os.Stdout, we can create our own
	// writer that implements the io.Writer interface.

	// cmd3 := exec.Command("ping", "google.com")

	// // pipe the commands output to the applications
	// // standard output
	// cmd3.Stdout = customOutput{}

	// // Run still runs the command and waits for completion
	// // but the output is instantly piped to Stdout
	// if err := cmd3.Run(); err != nil {
	// 	fmt.Println("could not run command: ", err)
	// }

	cmd4 := exec.Command("grep", "apple")

	// Create a new pipe, which gives us a reader/writer pair
	reader, writer := io.Pipe()
	// assign the reader to Stdin for the command
	cmd4.Stdin = reader
	// the output is printed to the console
	cmd4.Stdout = os.Stdout

	go func() {
		defer writer.Close()
		// the writer is connected to the reader via the pipe
		// so all data written here is passed on to the commands
		// standard input
		writer.Write([]byte("1. pear\n"))
		writer.Write([]byte("2. grapes\n"))
		writer.Write([]byte("3. apple\n"))
		writer.Write([]byte("4. banana\n"))
	}()

	if err := cmd4.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}

	// To stop these processes, we need to send a kill signal from
	// our application. We can do this by adding a context instance
	// to the command.

	// If the context gets cancelled, the command terminates as well.
	// https://www.sohamkamani.com/golang/context/
	ctx := context.Background()

	// The context now times out after 1 second
	// alternately, we can call `cancel()` to terminate immediately
	// var cancel context.CancelFunc
	// ctx, cancel = context.WithTimeout(ctx, 1*time.Second)
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	// defer cancel() // cancel when we are finished consuming integers
	cmd5 := exec.CommandContext(ctx, "sleep", "100")
	// cmd5 := exec.CommandContext(ctx, "/home/brent/src/reports/volume/go/runner/main/test.sh", "dev", "reports11", "30011", "1", "reports11", "30311", "reports")

	out5, err5 := cmd5.Output()
	if err5 != nil {
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Output: ", string(out5))

	// The *Cmd instance provides us with an input stream which we
	// can write into. Let’s use it to pass input to a grep child
	// process:

	// fmt.Printf("Path Separator: '%c'\n", os.PathSeparator)

	fmt.Printf("/home/brent/src/reports/volume/go/runner/main\n")
	//  ./test.sh dev reports11 30011 1 reports11 30311 reports
	// /home/brent/src/reports/volume/go/runner/main/test.sh dev reports11 30011 1 reports11 30311 reports
	cmd6 := exec.Command("/home/brent/src/reports/volume/go/runner/main/test.sh", "dev", "reports11", "30011", "1", "reports11", "30311", "reports")
	// cmd := exec.Command("firefox")
	// pipe the commands output to the applications
	// standard output
	cmd6.Stdout = os.Stdout

	err6 := cmd6.Run()

	if err6 != nil {
		log.Fatal(err6)
	}

	return message
}
