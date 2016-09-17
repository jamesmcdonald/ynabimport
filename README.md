# ynabimport

Cook data from banks for import into YNAB.  Entirely based on
guesswork and reverse-engineering from failures, so ymmv.

## Running

You need to run this script from a terminal where you have perl
available. On Linux or MacOS, you should be fine; on Windows you'll need
to install cygwin or msys or something.

Run this script with something like:
```
  perl ynabimport.pl -b bank 8888888888_2016_07_12-2016_08_05.csv > ~/Desktop/import.csv
```

Where `8888888888_2016_07_12-2016_08_05.csv` is the file you downloaded
from the bank, and `bank` is the name of your bank. Currently supported
banks are 'skandiabanken' and 'dnb'.

Drag the resulting import.csv from your desktop into YNAB! You will need
to mess around with the payee names. At the moment the support for payee
name matching in new YNAB is pretty weak.

## Getting data

You need to get a 'CSV' or 'text' download from the relevant bank. The
important thing is downloading the data with semicolon-separated fields,
because that's what the script expects. This is the default in
Skandiabanken. Detailed instructions for Skandiabanken are below; for
other banks the procedure should be similar.

### Skandiabanken export

Go to Kontoutskrift and select the date range you're interested in. It
doesn't work very well with 'Siste 30 dager' because there are usually
uncleared junk transactions at the top that just say 'Varekjøp'. I usually do
from a couple of days before the last import until yesterday.

* Click 'Last ned', 'CSV'
* Select the defaults
  * Skilletegn: Semikolon
  * Decimaltegn: Komma
* Click 'Last ned CSV'

