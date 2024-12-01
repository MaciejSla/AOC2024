let raw = open data.txt | lines | split column "   " list1 list2
let list1 = $raw | select list1 | sort
let list2 = $raw | select list2 | sort

$list1 | merge $list2 | each {|line|
    let num1 = $line.list1 | into int
    let num2 = $line.list2 | into int
    $num1 - $num2 | math abs
} | math sum