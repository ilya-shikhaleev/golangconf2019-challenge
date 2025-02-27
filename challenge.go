package challenge

// gophers is a list of GolangKazan GolangConf-2019 participants.
//
// Used in tests (see challenge_test.go).
var gophers = []gopher{
	// Participants (add yourself below):
	{
		name: "Alexander Kiryukhin",
		id:   "neonxp",
		post: "https://vk.com/feed?w=wall476865374_157",
	},

	// Testers. Do not need to contain a valid URL in post,
	// but should use `tester: true` and can't win a prize.
	{
		name:   "Iskander Sharipov",
		id:     "quasilyte",
		post:   "https://todo.add.a.post/123",
		tester: true,
	},
	{
		name:   "Oleg Kovalov",
		id:     "cristaloleg",
		post:   "https://prodam.garaz",
		tester: true,
	},
}

// gopher is a struct that describes GolangKazan GolangConf-2019 challenge participant.
type gopher struct {
	// name is participant full name (first + second).
	name string

	// id is a secondary identifier that is used if name alone can't
	// identify the participant. Optional, can be left empty.
	id string

	// post is a link to a media post about this challenge.
	//
	// We accept any kind of posts, here are some examples:
	//	- VK or Facebook post on a wall
	//	- Tweet
	//	- Instagram post
	//	- Etc, etc. (but read rules below beforehand)
	//
	// Rules:
	//	1. Post must include both #GolangKazan #GolangConf2019 hash tags.
	//	2. Post must survive until we stop the challenge and announce the winner.
	//	3. Account that does the announcement must have at least 10 followers/subscribers.
	post string

	// tester can submit invalid post string, but can't win a prize.
	tester bool
}

// key returns a hopefully unique participant key.
func (g gopher) key() string {
	if g.id == "" {
		return g.name
	}
	return g.name + "/" + g.id
}
