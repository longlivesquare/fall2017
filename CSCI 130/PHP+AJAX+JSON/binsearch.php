<html>
    <body>
        
    <?php 

    $arr = array();

    function randArr() {
        for($i=0; $i < 100; $i++) {
            array_push($arr, rand(0, 100));
        }
        array_multisort($arr);
    }

    function binSearch() {
        $find = $_POST["bs"];
        $min = 0;
        $max = 99;

        while($max > $min) {
            $test = floor(($max + $min)/2);
            if ($test == $find) {
                echo "<br>" . $find . " is in the array";
                return;
            }
        }
        echo "<br>" . $find . " is not in the array";
    }

    function printArr() {
        $arrlength = count($arr);
        echo "<br>["
        for($i = 0; $i < $arrlength; $i++) {
            echo $arr[$i] . ", ";
        }
        echo "]<br>";
    }
    randArr();
    binSearch();

    ?>


    </body>
</html>