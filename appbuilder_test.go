package hypster

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/franela/go-supertest"
	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	g := Goblin(t)

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	describe := g.Describe
	it := g.It
	services := make(map[string]interface{})

	describe("hypster", func() {
		it("NewApp should return new instance of AppBuilder", func() {
			app := NewApp(services)
			Expect(app).ShouldNot(BeNil())
		})

		it("AppBuilder.Route should return RouteBuilder", func() {
			app := NewApp(services)
			rb := app.Route("/test")
			Expect(rb).ShouldNot(BeNil())
			rb2 := app.Route("/test")
			Expect(rb2).Should(Equal(rb))
		})
	})

	testApp(g)
}

const hi = "hi"

func sayhi(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(hi))
}

func echo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	b, _ := ioutil.ReadAll(req.Body)
	w.Write(b)
}

func testApp(g *G) {

	services := make(map[string]interface{})
	a := NewApp(services)
	a.Head("/", sayhi)
	a.Get("/", sayhi)
	a.Post("/", echo)
	a.Put("/", echo)
	a.Delete("/", echo)
	a.Options("/", sayhi)

	server := httptest.NewServer(a)
	defer server.Close()

	describe := g.Describe
	it := g.It

	describe("app", func() {
		it("HEAD should respond 200", func(done Done) {
			NewRequest(server.URL).
				Head("/").
				Expect(200, done)
		})

		it("GET should say hi with status 200", func(done Done) {
			NewRequest(server.URL).
				Get("/").
				Expect(200, hi, done)
		})

		it("POST should respond ok with status 200", func(done Done) {
			NewRequest(server.URL).
				Post("/").
				Send("ok").
				Expect(200, "ok", done)
		})

		it("PUT should respond ok with status 200", func(done Done) {
			NewRequest(server.URL).
				Put("/").
				Send("ok").
				Expect(200, "ok", done)
		})

		it("DELETE should respond ok with status 200", func(done Done) {
			NewRequest(server.URL).
				Delete("/").
				Send("ok").
				Expect(200, "ok", done)
		})

		it("OPTIONS should say hi with status 200", func(done Done) {
			NewRequest(server.URL).
				Options("/").
				Send(hi).
				Expect(200, hi, done)
		})
	})
}
