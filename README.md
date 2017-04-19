# ynabimport

Cook data from banks for import into YNAB.  Entirely based on
guesswork and reverse-engineering from failures, so ymmv.

[![Build Status](https://travis-ci.org/jamesmcdonald/ynabimport.svg?branch=master)](https://travis-ci.org/jamesmcdonald/ynabimport)

## Building

I have re-implemented this in Go. Check out the `perl` branch for the old perl
version. It's much easier to make self-contained multi-platform binaries with
Go. I intend to build some soon so you don't have to build it yourself.

To build from source, you need Go installed and your `$GOPATH` set up. Then it's as easy as:
```
  go get github.com/jamesmcdonald/ynabimport/...
```

## Running

I have designed the default behaviour in such a way that you can either use
`ynabimport` from the command line or by simply dropping files on it.

Run with something like:
```
  ynabimport 8888888888_2016_07_12-2016_08_05.csv
```

Where `8888888888_2016_07_12-2016_08_05.csv` is the file you downloaded
from the bank. By default, it expects the format Skandiabanken uses.

In a GUI, just drop your files on the `ynabimport` icon.

This will create a file (or files) in the same place as the original with the
same base name, but with the extension `.ynabimport.csv`. In our example, the
new file would be called `8888888888_2016_07_12-2016_08_05.ynabimport.csv`.

Drag the resulting CSV from your desktop into YNAB. Behold the magic! You will
almost certainly need to mess around with the payee names.

### But I don't use Skandiabanken!

I have added support for DnB as well, possibly with other banks to follow.
Changing format is slightly more complicated, but not terribly. From the
command line it's just:

```
  ynabimport -format dnb somefile.csv
```

To be able to drag-and-drop, you'll have to make a new shortcut, and set the
target to include the `format` option, for example `"C:\Program Files\Awesome
Programs\ynabimport.exe" -format dnb`.

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
