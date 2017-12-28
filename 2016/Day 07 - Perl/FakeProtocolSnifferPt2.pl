use strict;
use warnings;
use Data::Dumper;

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


my $fname = "realInput.txt";#"realInput.txt";
open(my $fh, "<$fname") or die "file $fname not found.\n$!";

my $counter = 0;
while (<$fh>){
	chomp;

	#get patterns out
	my @patternsOutside;
	my @patternsInside;
	while ($_ =~ /(\[?[^\]\[]+\]?)/g){
		
		my $codePart = $1;

		#extract patterns
		#negative look ahead to allow for overlapping matches as patterns do overlap
		#http://blogs.perl.org/users/jhannah_mutation_grid/2010/09/overlapping-regex-matches.html
		while ($codePart =~ /(?=(\w)(\w)\1)/g){
			

			if ($1 ne $2){ #2 diff chars in pattern!

				my $pattern = "$1$2$1";

				if ($codePart =~ /(\[|\])/){ #inside [] ?
					push @patternsInside, $pattern;
				}else{
					push @patternsOutside, $pattern;
				}
			}
		}
	}
	
	#print "found outside:\n";
	#print Dumper(@patternsOutside);
	#print "found inside: \n";
	#print Dumper(@patternsInside);
	
	#match patterns
	my %patternsInsideHash = map { $_ => 1 } @patternsInside;
	foreach my $outPattern (@patternsOutside){
		
		$outPattern =~ /(\w)(\w)\w/;
		my $reversePattern = "$2$1$2";
		
		if (exists($patternsInsideHash{$reversePattern})){
			print "$_\n";
			$counter++;
			last; #deal with it
		}
	}
}

print "found $counter patterns!\n";