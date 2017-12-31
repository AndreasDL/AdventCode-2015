use strict;
use warnings;

my $fname = "input.txt";
open(my $fh, "<$fname") or die "file $fname not found.\n$!";

my $counter = 0;
while (<$fh>){
	chomp;

	next unless /(.).\1/;
	next unless /(..).*\1/;

	$counter++
}


print "$counter\n"

# 110 too high !