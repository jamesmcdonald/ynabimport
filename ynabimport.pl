#!/usr/bin/perl

# ynabimport.pl
# James McDonald <james@jamesmcdonald.com>

use warnings;
use strict;
use Getopt::Std;

print "Date,Payee,Category,Memo,Outflow,Inflow\n";

# Expect CR as well as LF
$/="\r\n";
# Separate print output with a comma
$,=',';
# Field separator in the bank export
my $FS=";";

my $bank=\&skandiabanken;
our $opt_b;

getopts('b:');

if ($opt_b) {
	$opt_b =~ /^d/ && ($bank = \&dnb);
	$opt_b =~ /^s/ && ($bank = \&skandiabanken);
}

sub skandiabanken {
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
}

sub dnb {
	print STDERR "WARNING: DNB support is COMPLETELY untested\n";
	print STDERR "WARNING: Nothing prevents it from eating your cat\n";
	while(<>) {
		chomp;
		s/"//g;
		s/,/./g;
		my @F = split($FS);
		if (scalar(@F) > 0 && $F[0] =~ /(\d{2})\.(\d{2})\.(\d{4})/) {
			$F[0] =~ s@(\d{2})\.(\d{2})\.(\d{4})@$1/$2/$3@;
			my $val = scalar(@F)==5 && $F[4] ne ''?'-'.$F[4]:$F[3];
			print $F[0], $F[1], "", "", $val . "\n";
		}
	}
}

&$bank;
