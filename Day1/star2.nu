let raw = open data.txt | lines | split column "   " list1 list2
let frequency = $raw | get list2 | reduce --fold {} { |it, acc|
    try {
        $acc | update $"($it)" { $in + 1 }
    } catch {
        $acc | insert $it { 1 }
    }
}

$raw | get list1 | each { |num|
    try {
        ($frequency | get $num) * ($num | into int)
    } catch {
        0
    }
} | math sum
