var images = []string{".png",".jpg",".jpeg",".gif",".ico"}
func addAvatar(w http.ResponseWriter, r *http.Request) {
        userName := getUserName(r)
        if err != nil || userName == ""{
                 http.Redirect(w, r, "/", 302)
        }
        file, header, err := r.FormFile("file")
        if err != nil {
                fmt.Fprintf(w, "Mime File not valid", 500)
                return
        }
        defer file.Close()
        for i, v := range images {
                matched, err := regexp.MatchString(v, header.Filename)
                if matched == true {
                        break
                }
                if matched == false && i == 4{
                        http.Error(w, "File not valid",500)
                        fmt.Println(matched, err,i,v)
                        return
                }
        }

        defer file.Close()
        out, err := os.Create("static/"+header.Filename)
        if err != nil {
                fmt.Fprintf(w, "Error",500)
                return
        }

        defer out.Close()
        _, err = db.Exec("UPDATE user SET avatarpath = ? WHERE username=?","/static/"+header.Filename, userName)
        if err != nil {
            http.Error(w, "Server error, unable to update your account.", 500)
            return
        }

        // write the content from POST to the file
        _, err = io.Copy(out, file)
        if err != nil {
                fmt.Fprintln(w, err)
        }
          http.Redirect(w, r, "/", 302)
        
}
