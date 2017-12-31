use strict;
use warnings;

my $fname = "input.txt";
open(my $fh, "<$fname") or die "file $fname not found.\n$!";

my $counter = 0;
while (<$fh>){
	chomp;

	next if /ab|cd|pq|xy/;
	next unless /(.)\1/;
	next unless (() = /[aeiou]/g) >= 3;

	$counter++
}


print "$counter\n"