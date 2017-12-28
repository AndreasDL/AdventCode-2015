use strict;
use warnings;

my $fname = "realInput.txt";
my $decompFile = "";

open(my $fh, "<$fname");
while(<$fh>){
	chomp;
	#print "$_ => \n";

	my $pos = 0;
	my $newLine = "";
	while ($_ =~ /\((\d+)x(\d+)\)/g){
		#readability++
		my $chars = $1;
		my $rep   = $2;
		my $startPat = $-[0]; #start pattern (perl build in magic)
		my $stopPat  = $+[0]; #stop pattern (perl build in magic)
		my $str      = substr $_, $stopPat, $chars;
	
		next if $pos >= $stopPat; #overlap => skip

		#add piece before patterns
		$newLine .= substr $_, $pos, ($startPat - $pos);
	
		#decomp pattern
		for (my $i =0; $i < $rep; $i++){
			$newLine .= $str;	
		}
		
		#move position
		$pos = $stopPat + $chars;

		#print "\t$_ =[$str x $rep]=> $newLine\n";
	}

	$newLine .= substr $_, $pos;

	#print "$_ => $newLine\n\n";
	$decompFile .= "$newLine";
}


print $decompFile, "\n\n";

print "found ", length $decompFile, " chars!\n";
