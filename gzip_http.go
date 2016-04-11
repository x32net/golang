func handler(w http.ResponseWriter, r *http.Request) {
    response := (THIS IS MY RESPONSE STRUCT)
    w.Header().Set("Content-Type", "application/json")
    if settings.COMPRESS_RESPONSE { // this is a global flag I defined
        w.Header().Set("Content-Encoding", "gzip")
        gz := gzip.NewWriter(w)
        json.NewEncoder(gz).Encode(response)
        gz.Close()
    } else {
        json.NewEncoder(w).Encode(response)
    }
}
