##soauth

### A social authentiation plugin written in Google Go

It requires a webview(browser) to authenticate, currently **soauth** supports the following

* Authentication with facebook
* Authentication with foursquare
* Authenticaiton with github
* Facebook Graph Api
    * Profile details access
    * Connections - friend list

##usage

    // you may import only facebook or foursqure as per the requirement
    import (
        "fmt"
        "github.com/42races/soauth/facebook"
        "github.com/42races/soauth/foursqure"
    )

    // Initialize in the login page
    auth := facebook.Init("your_app_id", "your_app_secret", "callback_url", nil )
    // the last argument is for options, I will update it once the option usage is implemented

    // Url to be accessed in webview(Browser) to authenticate
    login_url := auth.LoginUrl()

    // Authenticate
    // Facebook will return the auth_code in parameter to the **callback url**, now try to get AccessToken
    // In callback action in your application

    auth := facebook.Init("your_app_id", "your_app_secret", "callback_url", nil )
    token, err := auth.Authenticate(auth_code)

    if err != nil {
        fmt.Println("Authentication failed")
        return err
    }

    fmt.Println("Auth token", token)
    // store the token in the database and use it for further api access

For foursquare and Github the procedure is same, just import package foursquare or github

## Facebook Graph Api Access
    // Once you have the token you can use Facebook Graph Api
    // Initialize Graph Api

    graph := oauth_fb.Graph{"your access token here"}

    // Now Get you profile details as follows

    profile, err := graph.GetObject("me")

    //the profile will have all the profile info as per the Profile structure defined in facebook pakage(facebook/facebook.go)

    // Friend list

    friends, err := graph.GetConnections("me/friends")


## Contributing

    Fork it
    Create your feature branch (git checkout -b my-new-feature)
    Commit your changes (git commit -am 'Added some feature')
    Push to the branch (git push origin my-new-feature)
    Create new Pull Request
