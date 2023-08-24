package main

type communicator interface{
	sendMessage(to string, message string) error
	receiveMessage(from string) (message string, err error)
}

type email struct{
}
func (e email)sendMessage(to string, message string) error{
	// TODO: real send email code goes here
	fmt.Println("message sent")
	return nil
}

type voiceMail struct{
	
}

type postalMail struct{
	
}

type pigeon struct{
	
}



func main() {



}
