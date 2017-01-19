package main


import(
	"fmt"
	"net"
	"os"
)

// Error func
func CheckError (err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}



func main() {

	fmt.Println("Server running")
	// Prepare an address at any adress
	ServerAddr, err := net.ResolveUDPAddr("udp",":20014")
	CheckError(err)

	fmt.Printf("ServerAddr: ", ServerAddr)

	// Listen to selected port
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)

	defer ServerConn.Close()

	buf := make([]byte, 1024)


	for {
		n,addr,err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), "from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}


