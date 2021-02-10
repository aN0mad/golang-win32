# Win32Testing
Some source code for playing with and learning to use win32 with golang.

## Folders & files
| Folders    | Status     | Description  |
|:----------:|:----------:|:------------|
| messageBox              | Done        | Mapping and creating the win32 function MessageBoxW call by hand|
| mkwinsyscall-MessageBox | Done        | Using mkwinsyscall to map and create the MessageBoxW function |
| processInject-shellcode | In progress | Using mkwinsyscall to create win32 calls to do process injection |
| processinject-dll       | TBD         | Using mkwinsyscall to create win32 calls to do reflective dll injection |

## Resources
1. [How To Call Windows APIs In Golang](https://anubissec.github.io/How-To-Call-Windows-APIs-In-Golang/ "Read 1st")
2. [GopherCon 2020: Justen Walker - Safety Not Guaranteed: Calling Windows APIs using Unsafe & Syscall](https://youtu.be/EsPcKkESYPA "Watch 2nd")
3. [Gophercon 2020 Winapi repository](https://github.com/justenwalker/gophercon-2020-winapi "Github repo from talk")
4. [Breaking all the rules using go to call windows api](https://medium.com/@justen.walker/breaking-all-the-rules-using-go-to-call-windows-api-2cbfd8c79724 "Older Justen Walker blogposts")