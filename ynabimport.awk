#!/usr/bin/awk -f
# Cook data from skandiabanken for import into YNAB
# James McDonald <james@jamesmcdonald.com>

BEGIN {
    print "Date,Payee,Category,Memo,Outflow,Inflow";
    FS="\t";
    OFS=",";
}

$1 ~ /^"[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]"$/ {
    gsub("\"", "");
    gsub(",", ".");
    y=substr($1, 1, 4);
    m=substr($1, 6, 2);
    d=substr($1, 9, 2);
    print d "/" m "/" y, $5, "", "", $6, $7;
}
