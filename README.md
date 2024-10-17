# Go Example Project
This project is a Go http server

Project package details 
- [chi](https://github.com/go-chi/chi) is used for the router
- [nosurf](https://github.com/justinas/nosurf) is used for CSRF protection
- [scs](https://github.com/alexedwards/scs) is used for session management

# How to use this repo
Before using this repo, delete go.mod and go.sum. After doing this, run "go mod init your_module_name" in your project directory (ideally replace your_module_name with the link to your project's remote github repo eg. github.com/your_github_username/your_github_repo). Next, run "go tidy" in your project directory to add all the needed dependecies. Then, you might notice that there are errors in most of the go files. This is because the import statements in some of the go files refer to this github repo, as "go mod init" was run with the link to this github repo. As such, you must manually go into each of the files with a error and change "github.com/Goodmorningpeople/go_web_example/xxx" to whatever you entered as the argument to "go mod init". After this, your project should work as expected.
