package main

func ErrorHelper(err error) {
	InfoLogger.Printf("%T\n %s\n %#v\n", err, err, err)
	WarningLogger.Printf("%T\n %s\n %#v\n", err, err, err)
	ErrorLogger.Fatal(err)
}
