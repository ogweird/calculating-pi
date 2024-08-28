<?php
function pi_leibniz($n) {
    if ($n == 0) {
        return $n;
    }

    $pi = 0;
    $sign = 1;

    for ($i = 1; $i <= ($n * 2); $i += 2) {
        $pi = $pi + $sign * (4 / $i);
        $sign = -$sign;
    }

    return $pi;
}

function pi_wallis($n) {
    if ($n == 0) {
        return $n;
    }
    
    $pi = 4.;

    for ($i = 3; $i <= ($n + 2); $i += 2) {
        $pi = $pi * (($i - 1) / $i) * (($i + 1) / $i);
    }

    return $pi;
}

function pi_nilakantha($n) {
    if ($n == 0) {
        return $n;
    }

    $pi = 3;
    $sign = 1;

    for ($i = 2; $i <= ($n * 2); $i += 2) {
        $pi = $pi + $sign * (4 / ($i * ($i + 1) * ($i + 2)));
        $sign = -$sign;
    }

    return $pi;
}
?>