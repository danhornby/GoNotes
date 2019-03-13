# GoNotes
Notes API built in GoLang

Getting Started:
This is a Note taking api that has simple GET, POST & DELETE Methods.

The application posts to port 8000 on localhost:
http://localhost:8000/

Format:

JSON
{
  "id": ""
  "Title": ""
  "Content": ""
}

The application has the following methods defined:

(GET) All Notes:

Accessible via:
hhtp://localhost:8000/
http://localhost:8000/Notes/

(GET) Single Note:

Accessible via:
http://localhost:8000/Notes/{<id>}
