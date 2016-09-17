#!/usr/bin/perl -w

# ynabimport.pl
# James McDonald <james@jamesmcdonald.com>

use strict;
print "Date,Payee,Category,Memo,Outflow,Inflow\n";

# Field separator in the bank export
my $FS=";";
# Expect CR as well as LF
$/="\r\n";
# Separate print output with a comma
$,=',';

while(<>) {
    chomp;
    s/"//g;
    s/,/./g;
    my @F = split($FS);
    if (scalar(@F) > 0 && $F[0] =~ /(\d{4})-(\d{2})-(\d{2})/) {
        $F[0] =~ s@(\d{4})-(\d{2})-(\d{2})@$3/$2/$1@;
        # If we have a 'credit' value, stick it in the 'debit' column as a negative
        my $val = scalar(@F)==7 && $F[6] ne ''?'-'.$F[6]:$F[5];
        print $F[0], $F[4], "", "", $val . "\n";
    }
}
