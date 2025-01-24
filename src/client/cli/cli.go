package cli

import (
	"fmt"
	"os"

	"github.com/gorilla/websocket"
	"github.com/hertzcodes/client/handlers"
	"github.com/manifoldco/promptui"
)

type session struct {
	Username string
	Room     string
}

func RunCLI() {
	for {
		// Display the main menu
		menu := []string{"Login", "Exit"}
		prompt := promptui.Select{
			Label: "Main Menu",
			Items: menu,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "Login":
			login()
		case "Exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}
}

func login() {
	// Prompt for username
	usernamePrompt := promptui.Prompt{
		Label: "Enter your username",
		HideEntered: true,
	}

	username, err := usernamePrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Prompt for password
	passwordPrompt := promptui.Prompt{
		Label: "Enter your password",
		Mask:  '*',
		HideEntered: true,
	}

	password, err := passwordPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Simulate a login process
	if handlers.Login(username, password) {
		fmt.Println("Login successful!")
		userMenu(session{Username: username})
	} else {
		fmt.Println("Invalid username or password.")
	}
}

func userMenu(session session) {
	for {
		menu := []string{"Join a room", "Logout"}
		prompt := promptui.Select{
			Label: "User  Menu",
			Items: menu,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "Join a room":
			room := promptui.Prompt{
				Label: "Enter room code",
				HideEntered: true,
			}

			r, _ := room.Run()
			session.Room = r
			conn := handlers.Connect(session.Username, session.Room)
			go input(conn)
			if output(conn) != nil {
				return
			}
		case "Logout":
			fmt.Println("Logging out...")
			return
		}
	}
}

func input(conn *websocket.Conn) {
	defer conn.Close()
	for {
		pr := promptui.Prompt{Label: "", HideEntered: true}
		ln, _ := pr.Run()
		if ln == "#leave" {
			return
		}
		err := conn.WriteMessage(websocket.TextMessage, []byte(ln))
		if err != nil {
			return
		}

	}
}

func output(conn *websocket.Conn) error {
	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", msg)
	}
}
