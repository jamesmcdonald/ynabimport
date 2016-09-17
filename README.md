Cook data from skandiabanken for import into YNAB.  Entirely based on
guesswork and reverse-engineering from failures, so ymmv.

To use this, follow these steps:

Go to Kontoutskrift and select the date range you're interested in. It
doesn't work very well with 'Siste 30 dager' because there are usually
uncleared junk transactions at the top that just say 'Varekjøp'. I usually do
from a couple of days before the last import until yesterday.

Click 'Last ned', 'CSV'
Select the defaults:
  Skilletegn: Semikolon
  Decimaltegn: Komma
Click 'Last ned CSV'

You need to run this script from a terminal. On Linux or MacOS, you should be
fine; on Windows you'll need to install cygwin or msys or something to get
perl and iconv.

Run this script on the CSV, something like:
```
  perl ~/Downloads/ynabimport.pl 8888888888_2016_07_12-2016_08_05.csv | iconv -f iso-8859-1 > ~/Desktop/import.csv
```
(the iconv bit is necessary because the file is in ISO-8859-1 format, and YNAB likes UTF-8)

Just drag the resulting import.csv from your desktop into YNAB! You will need
to mess around with the payee names. At the moment the support for payee name
matching in new YNAB is pretty weak.
