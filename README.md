# gsendk

__gsendk__ is util sending a char-data to a foreground window from a CSV-file. For example it need for filling a big web forms.

Usage:
---
__gsendk <-h> <-n NNN> <-p NNN> file.csv__
##### Flags:
-h help\
-n \<seconds\> (wait this many seconds before sending text). The time used for switch a window for insert data foreground.\
-p \<milliseconds\> (wait this many milliseconds after each key sending)

##### Special words in input file meaning same keyboard button:
- TAB as Tab;
- ENTER as Enter;
- BSP as Backspace;
- ESC as Escape;
- UP as Arrow up;
- DOWN as Arrow down;
- LEFT as Arrow left;
- RIGHT as Arrow right;
- PAGEDOWN as PageDown;
- PAGEUP as PageUp;
- HOME as Home;
- END as End;
- P-\<milliseconds\> meaning pause. Example P-100 (pause 100 ms).

THANKS FOR ["github.com/micmonay/keybd_event"]("github.com/micmonay/keybd_event")
