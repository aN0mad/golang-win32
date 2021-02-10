package main

//sys	MessageBox(handle int, text string, caption string, typeBox int) (retValue int, err error) [failretval==0] = user32.MessageBoxW

//Don't need this one as we can pass a string to the function and go will auto-convert the type for us
/*
sys	MessageBox(handle int, text *uint16, caption *uint16, typeBox int) (retValue int, err error) [failretval==0] = user32.MessageBoxW
*/
