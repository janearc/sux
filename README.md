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
