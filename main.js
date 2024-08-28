const pi_leibniz = (n) => {
    if (n == 0) {
        return n
    } 

    let pi = 0
    let sign = 1

    for (let i = 1; i <= (n * 2); i += 2) {
        pi = pi + sign * (4 / i)
        sign = -sign
    }

    return pi
}

const pi_wallis = (n) => {
    if (n == 0) {
        return n
    } 

    let pi = 4.

    for (let i = 3; i <= (n + 2); i += 2) {
        pi = pi * ((i - 1) / 1) * ((i + 1) / 1) 
    }

    return pi
}

const pi_nilakantha = (n) => {
    if (n == 0) {
        return n
    }

    let pi = 3
    let sign = 1

    for (let i = 2; i <= (n * 2); i += 2) {
        pi = pi + sign * (4 / (i * (i + 1) * (i + 2)));
        sign = -sign;
    }

    return pi
}