// URL Path To Array In Golang
// Here’s a short Go function that converts the URL path into an array;
// it’s useful if you want to implement REST style queries into a simple application.
func parseComponents(r *http.Request) []string {
    //Creates an App Engine Context - required to access App Engine specific services
    c := appengine.NewContext(r)
    //The URL that the user queried.
    path := r.URL.Path
    path = strings.TrimSpace(path)
    //Log the URL received:
    c.Infof("URL: ", path)
    //Cut off the leading and trailing forward slashes, if they exist.
    //This cuts off the leading forward slash.
    if strings.HasPrefix(path, "/") {
        path = path[1:]
    }
    //This cuts off the trailing forward slash.
    if strings.HasSuffix(path, "/") {
        cut_off_last_char_len := len(path) - 1
        path = path[:cut_off_last_char_len]
    }
    //We need to isolate the individual components of the path.
    components := strings.Split(path, "/")
    return components
}
