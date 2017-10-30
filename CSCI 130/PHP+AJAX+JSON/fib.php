<html>
    <body>
    <?php

    function fib($n) {
        echo "<br>";
        $curr = 1;
        $prev = 1;
        
        for($i = 0; $i < 20; $i++) {
            if($i != 0 && $i != 1) {
                $ncurr = $curr + $prev;
                $prev = $curr;
                $curr = $ncurr;
            }
            echo $curr . "br";
        }
    }
    fib(20);
    ?>

    </body>
</html>