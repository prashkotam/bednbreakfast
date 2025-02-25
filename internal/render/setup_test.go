package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/prashkotam/bednbreakfast/internal/config"
	"github.com/prashkotam/bednbreakfast/internal/models"
)

var session *scs.SessionManager
var testApp config.Appconfig

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})

	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session
	
	os.Exit(m.Run())


}
