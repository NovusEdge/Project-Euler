sub isTriplet { (@_[2])**2 == (@_[0])**2 + (@_[1])**2 }

sub check { 
    ( @_[0] < @_[1] < @_[2] and isTriplet(@_[0], @_[1], @_[2]) and @_[0]+@_[1]+@_[2] == 1000 )
}

sub main {
    for (my $i = 0; $i < 1000; $i++) {
        for (my $j = 0; $j < 1000; $j++) {
            $k = ($i**2 + $j**2) ** 0.5;
            if (check($i, $j, int($k))){
                return $i*$j*int($k);
            }
        }
    }
}

$start = time();
$ans = main();
$exc_time = time() - $start;

print "\nAnswer: $ans\n";
print "Time Taken: $exc_time seconds\n\n";