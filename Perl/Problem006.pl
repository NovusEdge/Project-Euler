sub squareSum {
    my $n = @_[0];    
    ( ( $n*($n+1)/2 )**2 )
}

sub sumSquare {
    my $n = @_[0];
    ( $n * ($n+1) * (2*$n + 1) / 6 )
}

$start = time();
$ans = squareSum(100) - sumSquare(100);
$exc_time = time() - $start;
print "\nAnswer: $ans\n";
print "Time Taken: $exc_time seconds\n\n";