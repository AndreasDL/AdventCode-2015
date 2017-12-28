use strict;
use warnings;
use Data::Dumper;

my $fname = "realInput.txt";

open(my $fh, "<$fname");
while(<$fh>){
	chomp;

	print "$_";

	my $counter = 0;
	my @matches = ();

	$counter += length $_;

	#get matches
	while ($_ =~ /\((\d+)x(\d+)\)/g){
		my %match = (
			chars => $1,
			rep   => $2,
			startPat => $-[0], # start pattern (perl build in magic)
			stopPat  => $+[0], # stop pattern (perl build in magic)
			startPosStr => $+[0],
			stopPosStr  => $+[0] + $1,
			n => 1
		);

		push @matches, \%match;
	}

	for (my $i = 0 ; $i < scalar @matches; $i++){
		
		my %currMatch = %{$matches[$i]};

		$counter -= ($currMatch{stopPat} - $currMatch{startPat}) * $currMatch{n}; #remove pattern from str length
		$counter += $currMatch{chars} * ($currMatch{rep} - 1) * $currMatch{n}; #add repetitions
		
		#multiply repetitions on overlap
		#assumer patterns always completely overlap
		my $j = $i + 1;
		while ($j < scalar @matches and $matches[$j]{startPat} < $currMatch{stopPosStr}){

			#print "found overlap between [ $currMatch{startPosStr} - $currMatch{stopPosStr} ] ";
			#print "and [ $matches[$j]->{startPat} - $matches[$j]{stopPat} ]\n";

			$matches[$j]->{n} *= $currMatch{rep};
			$j++;
		}
	}

	#print Dumper(@matches);
	print " found $counter chars!\n\n";
}