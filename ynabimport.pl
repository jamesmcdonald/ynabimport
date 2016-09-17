#!/usr/bin/perl

# ynabimport.pl
# James McDonald <james@jamesmcdonald.com>

use warnings;
use strict;
use Getopt::Std;
use 5.010;
use Encode qw/from_to/;

# Accept CR as well as LF in input files
$/="\r\n";

my $parser=\&skandiabanken;
my $encoding='iso-8859-1';
our ($opt_b, $opt_i);

my %parsers = (
	'dnb' => \&dnb,
	'skandiabanken' => \&skandiabanken,
);

getopts('b:i:');

if ($opt_b) {
	my @p = grep /^$opt_b/, %parsers;
	if (!@p) {
		print STDERR "$0: unsupported bank '$opt_b'\n";
		exit 1;
	} elsif (@p > 1) {
		print STDERR "$0: multiple banks match '$opt_b': @p\n";
		exit 1;
	}
	$parser = $parsers{$p[0]};
}

if ($opt_i) {
	$encoding = $opt_i;
}

sub skandiabanken {
	my %result;
	s/"//g;
	s/,/./g;
	my @F = split(';');
	if (scalar(@F) > 0 && $F[0] =~ /(\d{4})-(\d{2})-(\d{2})/) {
		$result{'day'} = $3;
		$result{'month'} = $2;
		$result{'year'} = $1;
		$result{'value'} = scalar(@F)==7 && $F[6] ne ''?'-'.$F[6]:$F[5];
		$result{'description'} = $F[4];
	} else {
		return;
	}
	return %result;
}

sub dnb {
	state $warned = 0;
	if (!$warned) {
		print STDERR "WARNING: DNB support is COMPLETELY untested\n";
		print STDERR "WARNING: Nothing prevents it from eating your cat\n";
		$warned = 1;
	}
	my %result;
	s/"//g;
	s/,/./g;
	my @F = split(';');
	if (scalar(@F) > 0 && $F[0] =~ /(\d{2})\.(\d{2})\.(\d{4})/) {
		$result{'day'} = $1;
		$result{'month'} = $2;
		$result{'year'} = $3;
		$result{'value'} = scalar(@F)==5 && $F[4] ne ''?'-'.$F[4]:$F[3];
		$result{'description'} = $F[1];
	} else {
		return;
	}
	return %result;
}

print "Date,Payee,Category,Memo,Outflow,Inflow\n";

while(<>) {
	chomp;
	from_to($_, $encoding, 'utf-8');
	my %trans = &$parser($_);
	next if (!%trans);
	$,=',';
	$\="\n";
	print "$trans{'year'}-$trans{'month'}-$trans{'day'}", $trans{'description'}, "", "", $trans{'value'};
}
