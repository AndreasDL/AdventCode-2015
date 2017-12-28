use strict;
use warnings;

#the code below is so dirty, that this pigglet showed up!
#
#                         ,.
#                        (_|,.
#                       ,' /, )_______   _
#                    __j o``-'        `.'-)'
#                   (")                 \'
#                    `-j                |
#                      `-._(           /
#                         |_\  |--^.  /
#                        /_]'|_| /_)_/
#                           /_]'  /_]'


my $fname = "realInput.txt";
open(my $fh, "<$fname") or die "file $fname not found.\n$!";

my $counter = 0;
while (<$fh>){
	chomp;

	my $isValid = 1;
	my $foundPatternOutside = undef; #track if we found at least one pattern outside []
	
	while ($isValid and $_ =~ /(\[?[^\]\[]+\]?)/g){
		
		my $codePart = $1;

		my $hasPattern = ($codePart =~ /(\w)(\w)\2\1/ and $1 ne $2 );
		my $isInside = ($codePart =~ /(\[|\])/);

		if ($hasPattern){
			if ($isInside){
				$isValid = 0; #pattern inside [] => not valid
			}else{
				$foundPatternOutside = 1; # pattern outside => we got atleast one
			}
		}

		#print "$codePart => ";
		#print "P" if $hasPattern;
		#print "I" if $isInside;
		#print "V" if $isValid;
		#print "O" if $foundPatternOutside;
		#print "\n";
	}
	#print("$_ ==>> $isValid\n\n");

	$isValid = 0 if not $foundPatternOutside; #not pattern found outside [] => line is false

	$counter++ if $isValid;
}

print "found $counter lines!\n";