package CLI

import (
	Facebook "HiddenEye/controllers/Facebook"
	Location "HiddenEye/controllers/Location"
	helpers "HiddenEye/helpers"
	Colors "HiddenEye/utils"
	"fmt"
)

func Hidden_Eye() {

	const Head = `

    ██░ ██  ██ ▓█████▄ ▓█████▄ ▓█████  ███▄    █    ▓█████ ██   ██ ▓█████ 
   ▓██░ ██ ▓██ ▒██▀ ██▌▒██▀ ██▌▓█   ▀  ██ ▀█   █    ▓█   ▀ ▒██  ██ ▓█   ▀ 
   ▒██▀▀██ ▒██ ░██   █▌░██   █▌▒███   ▓██  ▀█ ██▒   ▒███    ▒██ ██ ▒███   
   ░▓█ ░██ ░██ ░▓█▄   ▌░▓█▄   ▌▒▓█  ▄ ▓██▒  ▐▌██▒   ▒▓█  ▄  ░ ▐██▓▒▓█  ▄ 
   ░▓█▒░██▓ ██ ░▒████▓ ░▒████▓ ░▒████▒▒██░   ▓██░   ░▒████▒ ░ ██▒▓░░▒████▒
 ▒ ░░▒░▒░▓   ▒▒▓  ▒  ▒▒▓  ▒ ░░ ▒░ ░░ ▒░   ▒ ▒    ░░ ▒░ ░  ██▒▒▒ ░░ ▒░ ░
 ▒ ░▒░ ░ ▒ ░ ░ ▒  ▒  ░ ▒  ▒  ░ ░  ░░ ░░   ░ ▒░    ░ ░  ░▓██ ░▒░  ░ ░  ░
 ░  ░░ ░ ▒ ░ ░ ░  ░  ░ ░  ░    ░      ░   ░ ░       ░   ▒ ▒ ░░     ░   
 ░  ░  ░ ░     ░       ░       ░  ░         ░       ░  ░░ ░        ░  ░
				░       ░                                  ░ ░            
		  
   `

	fmt.Println(string(Colors.Red()), Head)

	fmt.Println(Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]", Colors.White(), "Modern Phishing Tool With Advanced Functionality", Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]")
	fmt.Println(Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]", Colors.White(), "  Phising , keylogger , Information Collector ", Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]")
	fmt.Println(Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]", Colors.White(), "\t  https://github.com/muhammadqazi", Colors.Red(), "\t\t[", Colors.White(), " * ", Colors.Red(), "]")
	fmt.Println(Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]", Colors.White(), "\t\t  By Muhammad Qazi", Colors.Red(), "\t\t\t[", Colors.White(), " * ", Colors.Red(), "]")

	fmt.Println(Colors.Red(), "\n\n--------------------------------------------------------------------------------------------------")
	fmt.Println(Colors.Red(), "\n\nPhising Modules:")

	fmt.Println(Colors.Red(), "\n\n["+Colors.White()+"1"+Colors.Red()+"]", Colors.White(), " Facebook",
		Colors.Red(), "\t["+Colors.White()+"2"+Colors.Red()+"]", Colors.White(), " Google",
		Colors.Red(), "\t\t["+Colors.White()+"3"+Colors.Red()+"]", Colors.White(), " LinkedIn",
	)

	fmt.Println(Colors.Red(), "\n["+Colors.White()+"4"+Colors.Red()+"]", Colors.White(), " Github\t",
		Colors.Red(), "\t["+Colors.White()+"5"+Colors.Red()+"]", Colors.White(), " StackOverflow",
		Colors.Red(), "\t["+Colors.White()+"6"+Colors.Red()+"]", Colors.White(), " WordPress",
	)

	fmt.Println(Colors.Red(), "\n["+Colors.White()+"7"+Colors.Red()+"]", Colors.White(), " Instagram\t",
		Colors.Red(), "\t["+Colors.White()+"8"+Colors.Red()+"]", Colors.White(), " Twitter",
		Colors.Red(), "\t\t["+Colors.White()+"9"+Colors.Red()+"]", Colors.White(), " SnapChat",
	)

	fmt.Println(Colors.Red(), "\n["+Colors.White()+"10"+Colors.Red()+"]", Colors.White(), " Yahoo\t",
		Colors.Red(), "\t["+Colors.White()+"11"+Colors.Red()+"]", Colors.White(), " Twitch",
		Colors.Red(), "\t\t["+Colors.White()+"12"+Colors.Red()+"]", Colors.White(), " Microsoft",
	)

	fmt.Println(Colors.Red(), "\n["+Colors.White()+"13"+Colors.Red()+"]", Colors.White(), " Steam\t",
		Colors.Red(), "\t["+Colors.White()+"14"+Colors.Red()+"]", Colors.White(), " iCloud",
		Colors.Red(), "\t\t["+Colors.White()+"15"+Colors.Red()+"]", Colors.White(), " Gitlab",
	)

	fmt.Println(Colors.Red(), "\n["+Colors.White()+"16"+Colors.Red()+"]", Colors.White(), " NetFlix\t",
		Colors.Red(), "\t["+Colors.White()+"17"+Colors.Red()+"]", Colors.White(), " Pinterest",
		Colors.Red(), "\t["+Colors.White()+"18"+Colors.Red()+"]", Colors.White(), " Origin",
	)

	fmt.Println(Colors.Red(), "\n["+Colors.White()+"19"+Colors.Red()+"]", Colors.White(), " ProtonMail",
		Colors.Red(), "\t["+Colors.White()+"20"+Colors.Red()+"]", Colors.White(), " Spotify",
		Colors.Red(), "\t["+Colors.White()+"21"+Colors.Red()+"]", Colors.White(), " Quora",
	)

	fmt.Println(Colors.Red(), "\n["+Colors.White()+"22"+Colors.Red()+"]", Colors.White(), " Adobe\t",
		Colors.Red(), "\t["+Colors.White()+"23"+Colors.Red()+"]", Colors.White(), " Cryptocurrency",
		Colors.Red(), "\t["+Colors.White()+"24"+Colors.Red()+"]", Colors.White(), " DeviantArt",
	)

	fmt.Println(Colors.Red(), "\n["+Colors.White()+"25"+Colors.Red()+"]", Colors.White(), " DropBox\t",
		Colors.Red(), "\t["+Colors.White()+"26"+Colors.Red()+"]", Colors.White(), " eBay",
		Colors.Red(), "\t\t["+Colors.White()+"27"+Colors.Red()+"]", Colors.White(), " PayPal",
	)

	fmt.Println(Colors.Red(), "\n["+Colors.White()+"28"+Colors.Red()+"]", Colors.White(), " Verizon\t",
		Colors.Red(), "\t["+Colors.White()+"29"+Colors.Red()+"]", Colors.White(), " Reddit",
		Colors.Red(), "\t\t["+Colors.White()+"30"+Colors.Red()+"]", Colors.White(), " PlayStation",
	)

	fmt.Println(Colors.Red(), "\n["+Colors.White()+"31"+Colors.Red()+"]", Colors.White(), " XBOX\t",
		Colors.Red(), "\t["+Colors.White()+"32"+Colors.Red()+"]", Colors.White(), " Custom",
	)

	fmt.Println(Colors.Red(), "\n\nAdditional Modules:")

	fmt.Println(Colors.Red(), "\n\n["+Colors.White()+"33"+Colors.Red()+"]", Colors.White(), " Get Target Location",
		Colors.Red(), "\n\n["+Colors.White()+"34"+Colors.Red()+"]", " Exit",
	)

	fmt.Print(Colors.Red(), "\n\n┌─[ "+Colors.Green()+"HIDDEN-EYE"+Colors.Blue()+"~"+Colors.White()+"@HOME"+Colors.Red()+" ]"+"\n└──╼ "+Colors.White())

	option := helpers.Input_User(1, 34)

	if option == 1 {
		Facebook.Facebook()
	} else if option == 33 {
		Location.Get_Location()
	}
}
