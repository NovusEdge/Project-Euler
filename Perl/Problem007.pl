sub isPrime {
    my $n = @_[0];
    if ($n == 0 or $n == 1) { return 0; }
    if ($n == 2){ return 1; }
    for (my $i = 2; $i <= int($n**0.5); $i++) {
        if ($n % $i == 0) { return 0; }
    }
    return 1;
}

$start = time();

$ans = 2; $i = 0;
while ($i != 10001){
    if (isPrime($ans)){ $i++; }
    $ans++;
}
$ans--; 

$exc_time = time() - $start;

print "\nAnswer: $ans\n";
print "Time Taken: $exc_time seconds\n\n";