package main

import "fmt"
import "os"
import "math/rand"
import "syscall"
import "time"

func print_message() {
	messages := []string{
		"Did you hear about the developer who read from /dev/random as root? They wanted a sudo-random number generator.",
		"The Unix permissions system was actually developed by a disgraced academic! They were doing great until someone accused them of sudo-science.",
		"If you run Vim as root, are you writing under a sudonym?",
		"How does Walter White escalate his permissions? Sudoephedrine.",
		"Remember to be root before using ::before, ::after or any other CSS sudo-elements.",
		"Be careful: if you type the wrong password too many times, you’ll be sentenced to chroot jail.",
		"What do you call a fake command-line application which acts as another user? Pseudo-sudo.",
		"While most people regard Peter Salus’s “A Quarter Century of UNIX” fondly, it has been accused in some quarters of being a sudo-history.",
		"🛡 Windows needs your permission to continue.",
		"How’d the amoeba get root permissions on Kubernetes? It used its sudo pod.",
	}
	message := messages[rand.Intn(len(messages))]
	fmt.Println(message)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	print_message()
	time.Sleep(1 * time.Second)
	err := syscall.Exec("/usr/bin/sudo", os.Args, os.Environ())
	if err != nil {
		fmt.Println(err)
	}
}
