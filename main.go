package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {

	if len(os.Args) > 1 {

		if os.Args[1] == "get" {
			get_local_ip(os.Args[2])
			if len(os.Args) > 2 {
				receiver_setup(os.Args[2])
			} else {
				fmt.Println("Please specify a port number")
			}
		}

		if os.Args[1] == "give" {
			if len(os.Args) > 3 { // abomination incoming
				fmt.Println("Giving file")
				out, err := exec.Command("curl", "-X", "POST", "-F", "file=@"+os.Args[2], "http://"+os.Args[3]+"/upload").CombinedOutput()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf(string(out))
			} else {
				fmt.Println("Specify a file to send and the target device's address in the format: ip_address:port")
			}
		}

		if os.Args[1] == "pack" {
			if len(os.Args) > 2 {
				var pack_name string = os.Args[2] + ".rawr"
				if _, err := os.Stat(os.Args[2]); err == nil {
					pack_result := strings.ReplaceAll(pack_name, "/", "")
					fmt.Println("Packaging \"" + pack_result + "\"")
					exec.Command("zip", "-r", pack_result, os.Args[2]).Run()
					fmt.Println("Package created:", pack_result)
				} else {
					fmt.Println("File does not exist")
				}
			} else {
				fmt.Println("Specify a folder to pack")
			}
		}

		if os.Args[1] == "unpack" {
			if len(os.Args) > 2 {
				var unpack_name string = os.Args[2]
				if _, err := os.Stat(unpack_name); err == nil {
					fmt.Println("Unpacking \"" + os.Args[2] + "\"")
					exec.Command("unzip", unpack_name).Run()
					os.Remove(unpack_name)
					fmt.Println("Unpacked file:", unpack_name)
				} else {
					fmt.Println("File does not exist")
				}
			} else {
				fmt.Println("Specify a folder to unpack")
			}
		}

		if os.Args[1] == "help" || os.Args[1] == "--help" {
			help_message()
		}
		if os.Args[1] == "love" {
			fmt.Print("\n Love you Mum Mimiko, Auntie Metal and Mum² Suletta <3 \n\n")
		}
	} else {
		help_message()
	}
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) { // i just vibe coded thisss why am i stupid.. >m<
	// Limit the size of the memory to 10MB
	r.ParseMultipartForm(10 << 20) // 10 MB
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	// Create a file to write the uploaded content
	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully uploaded file: %s", handler.Filename)
}

func receiver_setup(server_port string) {
	http.HandleFunc("/upload", uploadFileHandler)
	fmt.Println("Server set up, ready to transmit!")

	// Start the server in a new goroutine
	go func() {
		if err := http.ListenAndServe(":"+server_port, nil); err != nil {
			log.Fatalf("ListenAndServe: %s", err)
		}
	}()

	fmt.Println("Press Enter to stop the server...")
	var input string
	fmt.Scanln(&input) // wait until user presses thingie
	fmt.Println("Stopping the server...")
}

func get_local_ip(server_port string) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, iface := range ifaces {
		// I have no ideaaaaaaa
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil {
					fmt.Println("Server Address:", ipNet.IP.String()+":"+server_port)
					return // Exit after finding the first local IP
				}
			}
		}
	}
}

func help_message() {
	fmt.Println()
	fmt.Println("rawr - Simple package manager")
	fmt.Println("Usage info: \n")
	fmt.Println("rawr get [port]: host an upload server for receiving packages")
	fmt.Println("rawr give [file] [ip_address:port]: give a package to an open host")
	fmt.Println("rawr install [package]: install a package")
	fmt.Println("rawr pack [folder]: create a rawr package")
	fmt.Println("rawr unpack [package]: unpackage a packaged rawr package")
	fmt.Println("rawr help/--help/(nothing): Display this message")
	fmt.Println("")
}
