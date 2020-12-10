sub T { @_[0] * ( @_[0]+1 ) / 2 }


sub Factors {
    my $n = @_[ 0 ]; my @factLis = [];

    for (my $i = 1; $i <= int($n**0.5); $i++) {
        if ($n % $i == 0){ push(@factLis, int($n/$i), $i) }
    }
    ( @factLis )
}

sub main() {
    my $i = 1;
    while (1) {
        if (scalar(Factors(T($i))-1) > 500){
            return T($i);
        }
        $i++;
    }
}

$start = time();
$ans = main();
$exc_time = time() - $start;

print "\nAnswer: $ans\n";
print "Time Taken: $exc_time seconds\n\n";