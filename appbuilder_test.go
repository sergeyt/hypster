package hypster

import (
	. "github.com/onsi/gomega"
	. "github.com/sergeyt/goblin"
	"testing"
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
}
