use List::Util qw(max);

sub isPallindrome {
    my $n = @_[0]; my $k = reverse $n;
    return  $k == $n;
}

$start = time(); @palls = ();

for($i = 999; $i >= 100; $i--) {
    for($j = 999; $j >= 100; $j--) {
        if (isPallindrome($i*$j)) {
            push(@palls, $i*$j);
        }
    }
}

$ans = max(@palls);
$exc_time = time() - $start;
print "\nAnswer: $ans\n";
print "Time Taken: $exc_time seconds\n\n";