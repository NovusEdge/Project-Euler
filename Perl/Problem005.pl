sub gcd {
	my ($x, $y) = @_;
	while ($x) { ($x, $y) = ($y % $x, $x) }
	$y
}
 
sub lcm {
	my ($x, $y) = @_;
	($x && $y) and $x / gcd($x, $y) * $y or 0
}

$start = time();
$ans = 1;
for (my $i = 1; $i < 21; $i++){
    $ans = lcm($ans, $i+1);
}

$exc_time = time() - $start;
print "Answer: $ans\n";
print "Time Taken: $exc_time seconds\n";