<?php

$start_time = microtime(true);

function fibonacci() {
    $a = 0;
    $b = 1;
    
    return function() use (&$a, &$b) {
        $result = $a;
        $c = $a + $b;
        $a = $b;
        $b = $c;
        return $result;
    };
}

$fib = fibonacci();
for ($i = 0; $i < 100000000; $i++) {
    $fib();
}

$end_time = microtime(true);
$execution_time = $end_time - $start_time;
echo "Execution time for PHP: " . $execution_time . " seconds\n";
