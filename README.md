oauth_fb
========

* Facebook authentiation written in Google Go
* Graph Api profile access


usage 

    import "github.com/42races/goauth"

    // Initialize Authenticate with facebook
    fba := oauth_fb.Init("your_app_id", "your_app_secret", "callback_url", options )
    
    // Url to be used to get app permission from facebook user should go here and accept app permissions
    fba.LoginUrl()
    
    // Authenticate
    // Facebook will return the auth_code in parameter to the callback url, now try to get AccessToken
    
    token, err := fba.Authenticate(auth_code)
    
    if err != nil { panic("Authentication failed") }
    
    // Once you have the token you can use Facebook Graph Api
    // Initialize Graph Api
    
    graph := oauth_fb.Graph{"your access token here"}
    
    // Now Get you profile details as follows
    
    profile, err := graph.GetObject("me")
    
    //the profile will have all the profile info as per the Profile structure
    
    
    
    
    
