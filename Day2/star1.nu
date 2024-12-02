open data.txt | lines | reduce --fold 0 { |line, acc|
    let num_arr = $line | split words | each {into int}

    mut safe = 1
    mut prev_val = $num_arr | first
    mut prev_sign: string = ""

    for $num in ($num_arr | skip 1) {
        # Don't really like this sign check but whatever
        let sign = $"(($prev_val - $num) > 0)"
        if ($prev_sign | is-empty) {
            $prev_sign = $sign
        } else if (not ($prev_sign | str contains $sign)) {
            $safe = 0
            break
        }
        if (($prev_val - $num | math abs) in 1..3) {
            $prev_val = $num
        } else {
            $safe = 0
            break
        }
    }
    $safe + $acc
}