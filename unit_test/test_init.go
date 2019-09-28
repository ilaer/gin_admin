package unit_test

import (
	"../app"
	_ "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"net/http"
	_ "net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"testing"
	_ "testing"
)

func TestInitApp(t *testing.T){
	r:= app.Initapp()
	w:=httptest.NewRecorder()
	req,_:=http.NewRequest(http.MethodGet,"/",nil)
	r.ServeHTTP(w,req)
	assert.Equal(t,300,w.Code)

}
