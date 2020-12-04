use List::Util qw(max);

sub isPrime {
    my $n = @_[0];
    if ($n == 0 or $n == 1) { return 0; }
    if ($n == 2){ return 1; }
    for (my $i = 2; $i <= int($n**0.5); $i++) {
        if ($n % $i == 0) { return 0; }
    }
    return 1;
}

sub Factors {
    my $n = @_[ 0 ]; 
    my @factLis = [ ];
    for (my $i = 1; $i <= int($n**0.5); $i++) {
        if ($n % $i == 0){ push(@factLis, int($n/$i), $i) }
    }
    return @factLis;
}

sub primeFactors {
    my $n = @_[0]; my @k = Factors($n);
    my @pfact = grep { isPrime($_) } @k;
    return @pfact
}

$start = time();
$ans = max(primeFactors(600851475143));
$exc_time = time() - $start;

print "\nAnswer: $ans\n";
print "Time Taken: $exc_time seconds\n\n";