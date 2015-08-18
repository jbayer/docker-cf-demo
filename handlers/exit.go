package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jbayer/lattice-app/helpers"
)

type Exit struct {
	Time time.Time
}

func (p *Exit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index, _ := helpers.FetchIndex()

	w.WriteHeader(http.StatusOK)
	styledTemplate.Execute(w, Body{
		Body: fmt.Sprintf(`
<div class="goodbye">
<img src="https://fatmintech.files.wordpress.com/2014/06/c8c75-6a00e551c39e1c883401a3fd1ba08f970b-pi.png" width="50%" height="50%"></img>
</div>
<div class="my-index-goodbye">Shutting down instance</div>
<div class="index-goodbye">%d</div>
<div class="mid-color-goodbye"></div>
<div class="bottom-color-goodbye"></div>
    `, index, time.Since(p.Time)),
	})

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("shutting down")
		os.Exit(1)
	}()
}
