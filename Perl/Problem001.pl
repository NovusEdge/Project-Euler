$start  = time();
$ans = 0;

for($i = 0; $i < 1000; $i += 1){
    if( $i%3==0 or $i%5==0 ){
        $ans += $i;
    }
} 

$exc_time = time() - $start;
print "\nAnswer: $ans\n";
print "Time Taken: $exc_time\n\n";