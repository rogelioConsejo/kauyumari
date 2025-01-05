# System vision
This is an LMS (Learning Management System) that is focused on building a marketing funnel, so initially you want
to be able to build a list of leads (free users) by offering some free content, and then you want to be able to
convert those leads into customers by offering some paid content.

## Email confirmation
Because this system is focused towards building a marketing funnel, the email is the most important part of the system.

## Drip feed content
Once a user has signed up, he can enroll in a course, and the course will have a drip feed content, so the user will
progressively receive the content of the course.

## Live classes
For the second iteration, after the MVP, we want to add live classes, so the teacher can schedule a live class, and
the students can access it through the platform, then the live class can be also referenced as a video recording, so
the students can access it later.

# MVP
To keep the scope small, we will focus on the following features:
- User registration
- User login
- User logout
- List available courses
- Enroll in a course
- Access course content
- Access course content in a drip feed way

# Next steps
## For the MVP
### MUST
- [ ] add user registry persistence + implementation
- [ ] email authentication method persistence implementation
- [ ] login Email confirmation (it should be similar to the email authentication method, but it should be a different
method, as it should be used to confirm the email address, not to authenticate the user, maybe we can extract some
common code)
- [x] login Access using default authentication method
- [ ] login Doorman
- [x] add default authentication method
- [ ] SendGrid integration implementation
- [ ] admin access to add, edit and delete courses
- [ ] list available courses
- [ ] enroll in a course
- [ ] access course content
- [ ] drip feed content

### SHOULD
- [ ] add content that can be purchased or checked out (instead of a course, it could be a book, a video, etc)
- [ ] admin access to edit email templates
- [ ] admin access to list users
- [ ] analytics (how many users are registered, how many are active, how many are inactive, how many are paying, etc)

### COULD
- [ ] add password authentication method (needs to be able to reset password, so this would need to be implemented 
and will need to extend the system, so it may not be as trivial as it seems)
- [ ] add sms/whatsapp authentication method
- [ ] add SSO authentication method
- [ ] a bunch of other things
- [ ] associate user with authentication method on the access component