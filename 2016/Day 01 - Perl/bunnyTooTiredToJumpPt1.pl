#Code pigglet loves dirty code
#
#    _._ _..._ .-',     _.._(`))
#   '-. `     '  /-._.-'    ',/
#      )         \            '.
#     / _    _    |             \
#    |  a    a    /              |
#    \   .-.                     ;  
#     '-('' ).-'       ,'       ;
#        '-;           |      .'
#           \           \    /
#           | 7  .__  _.-\   \
#           | |  |  ``/  /`  /
#          /,_|  |   /,_/   /
#             /,_/      '`-'


use strict;
use warnings;

my $path = 'R3, L5, R1, R2, L5, R2, R3, L2, L5, R5, L4, L3, R5, L1, R3, R4, R1, L3, R3, L2, L5, L2, R4, R5, R5, L4, L3, L3, R4, R4, R5, L5, L3, R2, R2, L3, L4, L5, R1, R3, L3, R2, L3, R5, L194, L2, L5, R2, R1, R1, L1, L5, L4, R4, R2, R2, L4, L1, R2, R53, R3, L5, R72, R2, L5, R3, L4, R187, L4, L5, L2, R1, R3, R5, L4, L4, R2, R5, L5, L4, L3, R5, L2, R1, R1, R4, L1, R2, L3, R5, L4, R2, L3, R1, L4, R4, L1, L2, R3, L1, L1, R4, R3, L4, R2, R5, L2, L3, L3, L1, R3, R5, R2, R3, R1, R2, L1, L4, L5, L2, R4, R5, L2, R4, R4, L3, R2, R1, L4, R3, L3, L4, L3, L1, R3, L2, R2, L4, L4, L5, R3, R5, R3, L2, R5, L2, L1, L5, L1, R2, R4, L5, R2, L4, L5, L4, L5, L2, L5, L4, R5, R3, R2, R2, L3, R3, L2, L5';
my @directionNames = ('north','east','south','west');
my $startDirection = 0;
my $startX = 0;
my $startY = 0;
my %visitedLocations;

sub GetNewDirection { #(my $currDirection, my $input)
	(my $currDirection, my $input) = @_;

	my $turnTo = substr $input, 0, 1;
	
	if ($turnTo eq 'R'){
		$currDirection += 1;
	} else {
		$currDirection -= 1;
	}

	return $currDirection % 4;
}

sub GetNewLocation { #($currX, $currY, $currDirection, $input)

	(my $currX, my $currY, my $currDirection, my $input) = @_;
	
	my $steps = substr $input, 1;
	
	if ($currDirection == 0){ #north
		$currY += $steps;
	}elsif ($currDirection == 1){ #east
		$currX += $steps;
	}elsif ($currDirection == 2){ #south
		$currY -= $steps;
	}elsif ($currDirection == 3){ #west
		$currX -= $steps;
	}

	return ($currX, $currY);
}

sub TraversePath { #($input, $startX, $startY, $startDirection)
	(my $input, my $currX, my $currY, my $currDirection) = @_;

	my @inputs = split(',', $input);
	foreach (@inputs){
		$_ =~ s/^\s+|\s+$//g; #remove leading and trailing whitespace

		$currDirection = GetNewDirection($currDirection, $_);
		(my $newX,my $newY) = GetNewLocation($currX,$currY,$currDirection,$_);

		$currX = $newX;
		$currY = $newY;
		print "[$_] -> $directionNames[$currDirection] to ($newX,$newY), he is " . (abs($newX) + abs($newY)) . " blocks aways.\n"

	}

	return ($currX, $currY, $currDirection);
}

(my $currX, my $currY, my $direction) = TraversePath($path, $startX, $startY, $startDirection);
print "bunny is at ($currX,$currY) heading $directionNames[0], He is " . (abs($currX) + abs($currY)) . " blocks away\n";