# sux

"state sucks, man." so there's an interesting problem in trying to use tools like gpt
in that they lack any ability to understand state. this is, from one query to the next
the backend doesn't have any prior knowledge of queries you might have passed it.

for example, if i ask openai, "hey my name is jane and i am wondering if you have some
ideas for funny midle names for a person named jane." the back end returns something
like "well rhonda is a very funny middle name for a person named jane," if i then send
a subsequent query saying "what if i changed my name to aloysius everdander abercrombie,"
the backend doesn't remember our previous statement and can't reply with the obvious
answer, "that's long for jane, so i've been told."

the way we get around this issue is we have to continually dope the back end with
context or state. perhaps ironically this means taking the things you've already said,
and feeding them to the backend, and asking it to summarize, and then you use that
summary to add context to the next query and so on.

in order for this to be effective, there needs to be storage involved.

an open question for me is, _how do we determine whether to reach out to storage
to find answers to things the backend might not know, or to try to find answers
we don't have based upon storage whose contents we do not know much or anything
about?_

but i'll figure that out when i get there.

# how do i bake this cake though?

```go
import (
	"github.com/janearc/sux"
)

s := sux.NewSux()

// this is a unique identifier that refers to both:
//   - the contextualized session with the openai backend
//   - the storage backend that is used to provide context to the backend
sid := s.NewSession()

uuid := sid.ToUUID() // this is a uuid, but you probably don't need to use this

// ask the backend a question
response, err := s.Query(sid, "how many parsecs is the kessel run?")

// magical queries against the storage backend
things, err := s.StorageQuery(sid, "string to search for")
things, err := s.StorageQueryRegex(sid, "regex to search for")

// what is in a things?
for _, thing := range things {
    fmt.Fprintf("'%s' [id: %s]: %s\n", thing.Name, thing.ID, thing.DataUnMarshal)
	fmt.Fprintf("'%s' [id: %s]: %s\n", thing.Name, thing.ID, thing.Summary)
	fmt.Fprintf("'%s' [id: %s]: %v\n", thing.Name, thing.ID, thing.Metadata)
}

// I have a sid and I want to resume a conversation with the backend
_, err := s.SidValid(sid)
response, err := s.Query(sid, "why does javascript smell like a wet dog?") // my editor just autocompleted this don't be mad at me ok

// I have some stuff I want to actually store (but maybe don't do this because the magic is sux does this for you)
thing := sux.NewThing(
	sux.MarshalData(anObject), // please don't feed me json, just let me marshal this for you
)

// that is correct, you do not get to set your own name or summary, that is for the computer to do.
fmt.Fprintf("'%s' [id: %s]: %s\n", thing.Name, thing.ID, thing.Summary)

// if this seems okay with you, go ahead and commit that sucker
xact, err := thing.Commit()
fmt.Fprintf("'%s' [id: %s]: written (confirm: %s)\n", thing.Name, thing.ID, xact)
```