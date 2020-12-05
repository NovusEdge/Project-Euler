use integer;
use List::Util;

$start = time();
$ans = 0;

my $NUM = $ARGV[0];
$NUM = 1 if ($NUM < 1);
my $count;
my @flags = ();
while ($NUM--) {
    $count = 0; 
    my @flags = (0 .. 2*10**6);
    for my $i (2 .. 2*10**6 ) {
        next unless defined $flags[$i];
        for (my $k=$i+$i; $k <= 2*10**6; $k+=$i) { undef $flags[$k]; }
        $ans += $i;
    }
}

$exc_time = time() - $start;

print "\nAnswer: $ans\n";
print "Time Taken: $exc_time seconds\n\n";